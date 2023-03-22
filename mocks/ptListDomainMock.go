package mocks

import (
	"reflect"
	"time"

	"github.com/golang/mock/gomock"
)

// MockPtListDomainInt is a mock of PtListDomainInt interface.
type MockPtListDomainInt struct {
	ctrl     *gomock.Controller
	recorder *MockPtListDomainIntMockRecorder
}

// MockPtListDomainIntMockRecorder is the mock recorder for MockPtListDomainInt.
type MockPtListDomainIntMockRecorder struct {
	mock *MockPtListDomainInt
}

// NewMockPtListDomainInt creates a new mock instance.
func NewMockPtListDomainInt(ctrl *gomock.Controller) *MockPtListDomainInt {
	mock := &MockPtListDomainInt{ctrl: ctrl}
	mock.recorder = &MockPtListDomainIntMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPtListDomainInt) EXPECT() *MockPtListDomainIntMockRecorder {
	return m.recorder
}

// AddPeriod mocks base method.
func (m *MockPtListDomainInt) AddPeriod(period string, t time.Time) time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPeriod", period, t)
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// AddPeriod indicates an expected call of AddPeriod.
func (mr *MockPtListDomainIntMockRecorder) AddPeriod(period, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPeriod", reflect.TypeOf((*MockPtListDomainInt)(nil).AddPeriod), period, t)
}

// GetInvocationTimestamp mocks base method.
func (m *MockPtListDomainInt) GetInvocationTimestamp(period string, t time.Time) time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInvocationTimestamp", period, t)
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetInvocationTimestamp indicates an expected call of GetInvocationTimestamp.
func (mr *MockPtListDomainIntMockRecorder) GetInvocationTimestamp(period, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInvocationTimestamp", reflect.TypeOf((*MockPtListDomainInt)(nil).GetInvocationTimestamp), period, t)
}
