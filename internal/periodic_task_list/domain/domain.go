package domain

import (
	"time"
)

const (
	oneHourPeriod  = "1h"
	oneDayPeriod   = "1d"
	oneMonthPeriod = "1mo"
	oneYearPeriod  = "1y"
)

type PtListDomainInt interface {
	AddPeriod(period string, t time.Time) time.Time
	GetInvocationTimestamp(period string, t time.Time) time.Time
}

type PtListDomain struct {
}

func NewPtListDomain() *PtListDomain {
	return &PtListDomain{}
}

func (pld *PtListDomain) AddPeriod(period string, t time.Time) time.Time {

	switch period {
	case oneHourPeriod:
		t = t.Add(time.Hour)
	case oneDayPeriod:
		t = t.AddDate(0, 0, 1)
	case oneMonthPeriod:
		t = t.AddDate(0, 1, -1)
		t = pld.getMonthLastDay(t)
	case oneYearPeriod:
		t = t.AddDate(1, 0, 0)
		t = pld.getYearLastDay(t)
	}

	return t
}

func (pld *PtListDomain) GetInvocationTimestamp(period string, t time.Time) time.Time {
	t = t.Add(time.Hour)
	t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())

	switch period {
	case oneMonthPeriod:
		return pld.getMonthLastDay(t)
	case oneYearPeriod:
		return pld.getYearLastDay(t)
	}

	return t
}

func (pld *PtListDomain) getMonthLastDay(t time.Time) time.Time {
	nextMonthFirstDay := time.Date(t.Year(), t.Month()+1, 1, t.Hour(), 0, 0, 0, t.Location())
	lastMonthLastDay := nextMonthFirstDay.AddDate(0, 0, -1)

	return lastMonthLastDay
}

func (pld *PtListDomain) getYearLastDay(t time.Time) time.Time {
	return time.Date(t.Year(), 12, 31, t.Hour(), 0, 0, 0, t.Location())
}
