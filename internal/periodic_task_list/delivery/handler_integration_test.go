package delivery

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/spf13/viper"

	"assessment/internal/core/config"
	"assessment/internal/core/logging"
	"assessment/internal/periodic_task_list/domain"
	"assessment/internal/periodic_task_list/service"
	"assessment/internal/utils/parser"
	"assessment/internal/utils/validators"
)

func TestGetPtList(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		wantStatusCode int
		wantBody       string
	}{
		{
			name:           "Add one hour period",
			url:            "?period=1h&tz=Europe/Athens&t1=20210714T204603Z&t2=20210715T123456Z",
			wantStatusCode: iris.StatusOK,
			wantBody:       `["20210714T210000Z","20210714T220000Z","20210714T230000Z","20210715T000000Z","20210715T010000Z","20210715T020000Z","20210715T030000Z","20210715T040000Z","20210715T050000Z","20210715T060000Z","20210715T070000Z","20210715T080000Z","20210715T090000Z","20210715T100000Z","20210715T110000Z","20210715T120000Z"]` + "\n",
		},
		{
			name:           "Add one day period",
			url:            "?period=1d&tz=Europe/Athens&t1=20211010T204603Z&t2=20211115T123456Z",
			wantStatusCode: iris.StatusOK,
			wantBody:       `["20211010T210000Z","20211011T210000Z","20211012T210000Z","20211013T210000Z","20211014T210000Z","20211015T210000Z","20211016T210000Z","20211017T210000Z","20211018T210000Z","20211019T210000Z","20211020T210000Z","20211021T210000Z","20211022T210000Z","20211023T210000Z","20211024T210000Z","20211025T210000Z","20211026T210000Z","20211027T210000Z","20211028T210000Z","20211029T210000Z","20211030T210000Z","20211031T220000Z","20211101T220000Z","20211102T220000Z","20211103T220000Z","20211104T220000Z","20211105T220000Z","20211106T220000Z","20211107T220000Z","20211108T220000Z","20211109T220000Z","20211110T220000Z","20211111T220000Z","20211112T220000Z","20211113T220000Z","20211114T220000Z"]` + "\n",
		},
		{
			name:           "Add one month period",
			url:            "?period=1mo&tz=Europe/Athens&t1=20210214T204603Z&t2=20211115T123456Z",
			wantStatusCode: iris.StatusOK,
			wantBody:       `["20210228T210000Z","20210331T200000Z","20210430T200000Z","20210531T200000Z","20210630T200000Z","20210731T200000Z","20210831T200000Z","20210930T200000Z","20211031T210000Z"]` + "\n",
		},
		{
			name:           "Add one year period",
			url:            "?period=1y&tz=Europe/Athens&t1=20180214T204603Z&t2=20211115T123456Z",
			wantStatusCode: iris.StatusOK,
			wantBody:       `["20181231T210000Z","20191231T210000Z","20201231T210000Z"]` + "\n",
		},
		{
			name:           "End date before start date",
			url:            "?period=1d&tz=Europe/Athens&t2=20210714T204603Z&t1=20210715T123456Z",
			wantStatusCode: iris.StatusBadRequest,
			wantBody:       `{"status":"error","desc":"Error endTimestamp should not be before startTimestamp"}` + "\n",
		},
		{
			name:           "Invalid period input to url params",
			url:            "?period=1w&tz=Europe/Athens&t1=20210714T204603Z&t2=20210715T123456Z",
			wantStatusCode: iris.StatusBadRequest,
			wantBody:       `{"status":"error","desc":"Key: 'getPtListRequest.Period' Error:Field validation for 'Period' failed on the 'oneof' tag"}` + "\n",
		},
		{
			name:           "Invalid timezone input to url params",
			url:            "?period=1y&tz=Europe/Patra&t1=20180214T204603Z&t2=20211115T123456Z",
			wantStatusCode: iris.StatusBadRequest,
			wantBody:       `{"status":"error","desc":"Key: 'getPtListRequest.Tz' Error:Field validation for 'Tz' failed on the 'validateTimezone' tag"}` + "\n",
		},
		{
			name:           "Invalid T1 date timestamp format input to url params",
			url:            "?period=1y&tz=Europe/Athens&t1=20180214&t2=20211115T123456Z",
			wantStatusCode: iris.StatusBadRequest,
			wantBody:       `{"status":"error","desc":"Key: 'getPtListRequest.T1' Error:Field validation for 'T1' failed on the 'validateTimestampFormat' tag"}` + "\n",
		},
		{
			name:           "Invalid T2 date timestamp format input to url params",
			url:            "?period=1y&tz=Europe/Athens&t1=20180214T204603Z&t2=20211115",
			wantStatusCode: iris.StatusBadRequest,
			wantBody:       `{"status":"error","desc":"Key: 'getPtListRequest.T2' Error:Field validation for 'T2' failed on the 'validateTimestampFormat' tag"}` + "\n",
		},
		{
			name:           "Missing period from url params",
			url:            "?tz=Europe/Athens&t1=20180214T204603Z&t2=20211115T123456Z",
			wantStatusCode: iris.StatusBadRequest,
			wantBody:       `{"status":"error","desc":"Key: 'getPtListRequest.Period' Error:Field validation for 'Period' failed on the 'required' tag"}` + "\n",
		},
		{
			name:           "Missing timezone from url params",
			url:            "?period=1y&t1=20180214T204603Z&t2=20211115T123456Z",
			wantStatusCode: iris.StatusBadRequest,
			wantBody:       `{"status":"error","desc":"Key: 'getPtListRequest.Tz' Error:Field validation for 'Tz' failed on the 'required' tag"}` + "\n",
		},
		{
			name:           "Missing t1 field from url params",
			url:            "?period=1y&tz=Europe/Athens&t2=20211115T123456Z",
			wantStatusCode: iris.StatusBadRequest,
			wantBody:       `{"status":"error","desc":"Key: 'getPtListRequest.T1' Error:Field validation for 'T1' failed on the 'required' tag"}` + "\n",
		},
		{
			name:           "Missing t2 field from url params",
			url:            "?period=1y&tz=Europe/Athens&t1=20180214T204603Z",
			wantStatusCode: iris.StatusBadRequest,
			wantBody:       `{"status":"error","desc":"Key: 'getPtListRequest.T2' Error:Field validation for 'T2' failed on the 'required' tag"}` + "\n",
		},
	}

	for _, tt := range tests {
		if testing.Short() {
			t.Skip("this is integration test so skip it while running unit tests")
		}

		t.Run("Valid run without errors", func(t *testing.T) {
			ptListHandler := initializeHandler()

			app := iris.New()
			app.Get("/ptlist", ptListHandler.GetPtList)

			e := httptest.New(t, app)
			e.GET("/ptlist").
				WithURL(tt.url).
				Expect().
				Status(tt.wantStatusCode).
				Body().
				IsEqual(tt.wantBody)
		})
	}
}

func initializeHandler() *PtListHandler {
	vpr := viper.New()
	cfg := config.NewConfig(vpr)

	logger := logging.NewLogger(cfg)
	logger.Initialize()

	structValidator := validators.NewStructValidator(validator.New())
	urlParamsParser := parser.NewUrlParamsParser(structValidator)

	ptListDomain := domain.NewPtListDomain()
	ptListService := service.NewPtListService(ptListDomain)

	return NewPtListHandler(
		urlParamsParser,
		ptListService,
		logger,
	)
}
