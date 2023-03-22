package service

import (
	"time"

	"assessment/internal/periodic_task_list/domain"
)

const timestampFormat = "20060102T150405Z"

type PtListServiceInt interface {
	GetTimestampsList(tZone, period, startTime, endTime string) ([]string, error)
}

type PtListService struct {
	domain domain.PtListDomainInt
}

func NewPtListService(domain domain.PtListDomainInt) *PtListService {
	return &PtListService{
		domain: domain,
	}
}

func (pls *PtListService) GetTimestampsList(tZone, period, startTime, endTime string) ([]string, error) {
	var ptList []string

	timezone, err := time.LoadLocation(tZone)
	if err != nil {
		return ptList, &ConvertingInputDataError{errorMessage: err.Error()}
	}

	startTimestamp, err := time.Parse(timestampFormat, startTime)
	if err != nil {
		return ptList, &ConvertingInputDataError{errorMessage: err.Error()}
	}

	endTimestamp, err := time.Parse(timestampFormat, endTime)
	if err != nil {
		return ptList, &ConvertingInputDataError{errorMessage: err.Error()}
	}

	if endTimestamp.Before(startTimestamp) {
		return ptList, &ConvertingInputDataError{errorMessage: "Error endTimestamp should not be before startTimestamp"}
	}

	startTimestamp = startTimestamp.In(timezone)
	endTimestamp = endTimestamp.In(timezone)

	currentTimestamp := pls.domain.GetInvocationTimestamp(period, startTimestamp)

	for currentTimestamp.Before(endTimestamp) {
		ptList = append(ptList, currentTimestamp.UTC().Format(timestampFormat))

		currentTimestamp = pls.domain.AddPeriod(period, currentTimestamp)
	}

	return ptList, nil
}

type ConvertingInputDataError struct {
	errorMessage string
}

func (cid *ConvertingInputDataError) Error() string {
	return cid.errorMessage
}
