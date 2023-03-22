package delivery

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/kataras/iris/v12/x/errors"
	"github.com/stretchr/testify/assert"

	"assessment/internal/core/logging"
	"assessment/internal/periodic_task_list/domain"
	"assessment/internal/periodic_task_list/service"
	"assessment/internal/utils/parser"
	"assessment/internal/utils/validators"
	"assessment/mocks"
)

func TestNewPtListHandler(t *testing.T) {
	t.Run("Test ptList handler constructor", func(t *testing.T) {
		logger := logging.NewLogger(nil)
		structValidator := validators.NewStructValidator(validator.New())
		urlParamsParser := parser.NewUrlParamsParser(structValidator)

		ptListDomain := domain.NewPtListDomain()
		ptListService := service.NewPtListService(ptListDomain)
		got := NewPtListHandler(urlParamsParser, ptListService, logger)

		assert.Implements(t, new(PtListHandlerInt), got)
	})
}

func TestGetPtListWithoutErrors(t *testing.T) {
	t.Run("Valid run without errors", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		loggerMock := mocks.NewMockLoggerInt(ctrl)

		urlParamsParser := mocks.NewMockUrlParamsParserInt(ctrl)
		urlParamsParser.
			EXPECT().
			ParseUrlParams(gomock.Any(), gomock.Any()).
			Return(nil)

		ptListServiceMock := mocks.NewMockPtListServiceInt(ctrl)
		ptListServiceMock.
			EXPECT().
			GetTimestampsList(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]string{"20210228T220000Z", "20210331T210000Z"}, nil)

		ptListHandler := NewPtListHandler(
			urlParamsParser,
			ptListServiceMock,
			loggerMock,
		)

		app := iris.New()
		app.Get("/ptlist", ptListHandler.GetPtList)

		e := httptest.New(t, app)
		e.GET("/ptlist").
			WithURL("?period=1y&tz=Europe/Athens&t1=20210214T204603Z&t2=20211115T123456Z").
			Expect().
			Status(iris.StatusOK).
			Body().
			IsEqual(`["20210228T220000Z","20210331T210000Z"]` + "\n")
	})
}

func TestGetPtListWithUrlParamParsingError(t *testing.T) {
	t.Run("Error while parsing url params", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		loggerMock := mocks.NewMockLoggerInt(ctrl)

		urlParamsParser := mocks.NewMockUrlParamsParserInt(ctrl)
		urlParamsParser.
			EXPECT().
			ParseUrlParams(gomock.Any(), gomock.Any()).
			Return(&parser.UrlParamParserError{ErrorMessage: "parsing error"})

		ptListServiceMock := mocks.NewMockPtListServiceInt(ctrl)

		ptListHandler := NewPtListHandler(
			urlParamsParser,
			ptListServiceMock,
			loggerMock,
		)

		app := iris.New()
		app.Get("/ptlist", ptListHandler.GetPtList)

		e := httptest.New(t, app)
		e.GET("/ptlist").
			WithURL("?period=1y&tz=Europe/Athens&t2=20211115T123456Z").
			Expect().
			Status(iris.StatusBadRequest).
			Body().
			IsEqual(`{"status":"error","desc":"Error while parsing url params : parsing error"}` + "\n")
	})
}

func TestGetPtListWithInternalErrorWillBeLogged(t *testing.T) {
	t.Run("Valid run without errors", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		loggerMock := mocks.NewMockLoggerInt(ctrl)
		loggerMock.
			EXPECT().
			Error(gomock.Any())

		urlParamsParser := mocks.NewMockUrlParamsParserInt(ctrl)
		urlParamsParser.
			EXPECT().
			ParseUrlParams(gomock.Any(), gomock.Any()).
			Return(nil)

		ptListServiceMock := mocks.NewMockPtListServiceInt(ctrl)
		ptListServiceMock.
			EXPECT().
			GetTimestampsList(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return([]string{}, errors.New("internal error"))

		ptListHandler := NewPtListHandler(
			urlParamsParser,
			ptListServiceMock,
			loggerMock,
		)

		app := iris.New()
		app.Get("/ptlist", ptListHandler.GetPtList)

		e := httptest.New(t, app)
		e.GET("/ptlist").
			WithURL("?period=1y&tz=Europe/Athens&t1=20210214T204603Z&t2=20211115T123456Z").
			Expect().
			Status(iris.StatusInternalServerError).
			Body().
			IsEqual(`{"status":"error","desc":"internal error"}` + "\n")
	})
}
