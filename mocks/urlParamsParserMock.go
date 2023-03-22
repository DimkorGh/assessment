package mocks

import (
	"net/http"
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockUrlParamsParserInt is a mock of UrlParamsParserInt interface.
type MockUrlParamsParserInt struct {
	ctrl     *gomock.Controller
	recorder *MockUrlParamsParserIntMockRecorder
}

// MockUrlParamsParserIntMockRecorder is the mock recorder for MockUrlParamsParserInt.
type MockUrlParamsParserIntMockRecorder struct {
	mock *MockUrlParamsParserInt
}

// NewMockUrlParamsParserInt creates a new mock instance.
func NewMockUrlParamsParserInt(ctrl *gomock.Controller) *MockUrlParamsParserInt {
	mock := &MockUrlParamsParserInt{ctrl: ctrl}
	mock.recorder = &MockUrlParamsParserIntMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUrlParamsParserInt) EXPECT() *MockUrlParamsParserIntMockRecorder {
	return m.recorder
}

// ParseUrlParams mocks base method.
func (m *MockUrlParamsParserInt) ParseUrlParams(r *http.Request, data interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseUrlParams", r, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// ParseUrlParams indicates an expected call of ParseUrlParams.
func (mr *MockUrlParamsParserIntMockRecorder) ParseUrlParams(r, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseUrlParams", reflect.TypeOf((*MockUrlParamsParserInt)(nil).ParseUrlParams), r, data)
}
