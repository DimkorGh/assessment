package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewPtListDomain(t *testing.T) {
	t.Run("Test ptList domain constructor", func(t *testing.T) {
		got := NewPtListDomain()

		assert.Implements(t, new(PtListDomainInt), got)
	})
}

func TestAddPeriod(t *testing.T) {
	type args struct {
		period string
		t      time.Time
		tZone  *time.Location
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "Add one hour period",
			args: args{
				period: oneHourPeriod,
				t:      time.Date(2022, 1, 1, 21, 5, 6, 7, time.UTC),
				tZone:  time.UTC,
			},
			want: time.Date(2022, 1, 1, 22, 5, 6, 7, time.UTC),
		},
		{
			name: "Add one day period",
			args: args{
				period: oneDayPeriod,
				t:      time.Date(2022, 1, 1, 21, 5, 6, 7, time.UTC),
				tZone:  time.UTC,
			},
			want: time.Date(2022, 1, 2, 21, 5, 6, 7, time.UTC),
		},
		{
			name: "Add one month period and get last day of month timestamp",
			args: args{
				period: oneMonthPeriod,
				t:      time.Date(2022, 1, 1, 21, 5, 6, 7, time.UTC),
				tZone:  time.UTC,
			},
			want: time.Date(2022, 1, 31, 21, 0, 0, 0, time.UTC),
		},
		{
			name: "Add one year period and get last day of year timestamp",
			args: args{
				period: oneYearPeriod,
				t:      time.Date(2022, 1, 1, 21, 5, 6, 7, time.UTC),
				tZone:  time.UTC,
			},
			want: time.Date(2023, 12, 31, 21, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pld := NewPtListDomain()
			got := pld.AddPeriod(tt.args.period, tt.args.t)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetInvocationTimestamp(t *testing.T) {
	type args struct {
		period string
		t      time.Time
		tZone  *time.Location
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "Get invocation timestamp for one hour period",
			args: args{
				period: oneHourPeriod,
				t:      time.Date(2022, 1, 2, 3, 4, 5, 6, time.UTC),
				tZone:  time.UTC,
			},
			want: time.Date(2022, 1, 2, 4, 0, 0, 0, time.UTC),
		},
		{
			name: "Get invocation timestamp for one day period",
			args: args{
				period: oneDayPeriod,
				t:      time.Date(2022, 1, 2, 3, 4, 5, 6, time.UTC),
				tZone:  time.UTC,
			},
			want: time.Date(2022, 1, 2, 4, 0, 0, 0, time.UTC),
		},
		{
			name: "Get invocation timestamp for one month period",
			args: args{
				period: oneMonthPeriod,
				t:      time.Date(2022, 1, 2, 3, 4, 5, 6, time.UTC),
				tZone:  time.UTC,
			},
			want: time.Date(2022, 1, 31, 4, 0, 0, 0, time.UTC),
		},
		{
			name: "Get invocation timestamp for one year period",
			args: args{
				period: oneYearPeriod,
				t:      time.Date(2022, 1, 2, 3, 4, 5, 6, time.UTC),
				tZone:  time.UTC,
			},
			want: time.Date(2022, 12, 31, 4, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pld := NewPtListDomain()
			got := pld.GetInvocationTimestamp(tt.args.period, tt.args.t)

			assert.Equal(t, tt.want, got)
		})
	}
}
