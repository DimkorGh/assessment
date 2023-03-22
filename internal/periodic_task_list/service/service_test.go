package service

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"assessment/internal/periodic_task_list/domain"
	"assessment/mocks"
)

func TestNewPtListService(t *testing.T) {
	t.Run("Test ptList service constructor", func(t *testing.T) {
		ptListDomain := domain.NewPtListDomain()
		got := NewPtListService(ptListDomain)

		assert.Implements(t, new(PtListServiceInt), got)
	})

}

func TestGetTimestampsListReturnsError(t *testing.T) {
	type args struct {
		tZone     string
		period    string
		startTime string
		endTime   string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Invalid timezone",
			args: args{
				tZone:     "Europe/Patra",
				period:    "1y",
				startTime: "20180214T204603Z",
				endTime:   "20211115T123456Z",
			},
			want:    []string(nil),
			wantErr: true,
		},
		{
			name: "Invalid start timestamp",
			args: args{
				tZone:     "Europe/Athens",
				period:    "1y",
				startTime: "1234",
				endTime:   "20211115T123456Z",
			},
			want:    []string(nil),
			wantErr: true,
		},
		{
			name: "Invalid end timestamp",
			args: args{
				tZone:     "Europe/Athens",
				period:    "1y",
				startTime: "20180214T204603Z",
				endTime:   "1234",
			},
			want:    []string(nil),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			ptListDomainMock := mocks.NewMockPtListDomainInt(ctrl)

			pls := NewPtListService(ptListDomainMock)
			got, err := pls.GetTimestampsList(tt.args.tZone, tt.args.period, tt.args.startTime, tt.args.endTime)

			assert.Error(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetTimestampsListReturnsPtList(t *testing.T) {
	type args struct {
		tZone     string
		period    string
		startTime string
		endTime   string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Get valid ptList",
			args: args{
				tZone:     "Europe/Athens",
				period:    "1y",
				startTime: "20180214T204603Z",
				endTime:   "20210214T204603Z",
			},
			want:    []string{"20180214T204603Z"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			ptListDomainMock := mocks.NewMockPtListDomainInt(ctrl)
			ptListDomainMock.
				EXPECT().
				GetInvocationTimestamp(gomock.Any(), gomock.Any()).
				Return(time.Date(2018, 2, 14, 20, 46, 3, 0, time.UTC))
			ptListDomainMock.
				EXPECT().
				AddPeriod(gomock.Any(), gomock.Any()).
				Return(time.Date(2022, 2, 14, 20, 46, 3, 0, time.UTC))

			pls := NewPtListService(ptListDomainMock)
			got, err := pls.GetTimestampsList(tt.args.tZone, tt.args.period, tt.args.startTime, tt.args.endTime)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
