package mocks

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockPtListServiceInt is a mock of PtListServiceInt interface.
type MockPtListServiceInt struct {
	ctrl     *gomock.Controller
	recorder *MockPtListServiceIntMockRecorder
}

// MockPtListServiceIntMockRecorder is the mock recorder for MockPtListServiceInt.
type MockPtListServiceIntMockRecorder struct {
	mock *MockPtListServiceInt
}

// NewMockPtListServiceInt creates a new mock instance.
func NewMockPtListServiceInt(ctrl *gomock.Controller) *MockPtListServiceInt {
	mock := &MockPtListServiceInt{ctrl: ctrl}
	mock.recorder = &MockPtListServiceIntMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPtListServiceInt) EXPECT() *MockPtListServiceIntMockRecorder {
	return m.recorder
}

// GetTimestampsList mocks base method.
func (m *MockPtListServiceInt) GetTimestampsList(tZone, period, startTime, endTime string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimestampsList", tZone, period, startTime, endTime)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTimestampsList indicates an expected call of GetTimestampsList.
func (mr *MockPtListServiceIntMockRecorder) GetTimestampsList(tZone, period, startTime, endTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimestampsList", reflect.TypeOf((*MockPtListServiceInt)(nil).GetTimestampsList), tZone, period, startTime, endTime)
}
