package mocks

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockLoggerInt is a mock of LoggerInt interface.
type MockLoggerInt struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerIntMockRecorder
}

// MockLoggerIntMockRecorder is the mock recorder for MockLoggerInt.
type MockLoggerIntMockRecorder struct {
	mock *MockLoggerInt
}

// NewMockLoggerInt creates a new mock instance.
func NewMockLoggerInt(ctrl *gomock.Controller) *MockLoggerInt {
	mock := &MockLoggerInt{ctrl: ctrl}
	mock.recorder = &MockLoggerIntMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoggerInt) EXPECT() *MockLoggerIntMockRecorder {
	return m.recorder
}

// Debug mocks base method.
func (m *MockLoggerInt) Debug(args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debug", varargs...)
}

// Debug indicates an expected call of Debug.
func (mr *MockLoggerIntMockRecorder) Debug(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockLoggerInt)(nil).Debug), args...)
}

// Debugf mocks base method.
func (m *MockLoggerInt) Debugf(msg string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{msg}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockLoggerIntMockRecorder) Debugf(msg interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{msg}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockLoggerInt)(nil).Debugf), varargs...)
}

// Debugw mocks base method.
func (m *MockLoggerInt) Debugw(msg string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{msg}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugw", varargs...)
}

// Debugw indicates an expected call of Debugw.
func (mr *MockLoggerIntMockRecorder) Debugw(msg interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{msg}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugw", reflect.TypeOf((*MockLoggerInt)(nil).Debugw), varargs...)
}

// Error mocks base method.
func (m *MockLoggerInt) Error(args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MockLoggerIntMockRecorder) Error(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLoggerInt)(nil).Error), args...)
}

// Errorf mocks base method.
func (m *MockLoggerInt) Errorf(msg string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{msg}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockLoggerIntMockRecorder) Errorf(msg interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{msg}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockLoggerInt)(nil).Errorf), varargs...)
}

// Errorw mocks base method.
func (m *MockLoggerInt) Errorw(msg string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{msg}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorw", varargs...)
}

// Errorw indicates an expected call of Errorw.
func (mr *MockLoggerIntMockRecorder) Errorw(msg interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{msg}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorw", reflect.TypeOf((*MockLoggerInt)(nil).Errorw), varargs...)
}

// Fatal mocks base method.
func (m *MockLoggerInt) Fatal(args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatal", varargs...)
}

// Fatal indicates an expected call of Fatal.
func (mr *MockLoggerIntMockRecorder) Fatal(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatal", reflect.TypeOf((*MockLoggerInt)(nil).Fatal), args...)
}

// Fatalf mocks base method.
func (m *MockLoggerInt) Fatalf(msg string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{msg}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatalf", varargs...)
}

// Fatalf indicates an expected call of Fatalf.
func (mr *MockLoggerIntMockRecorder) Fatalf(msg interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{msg}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatalf", reflect.TypeOf((*MockLoggerInt)(nil).Fatalf), varargs...)
}

// Fatalw mocks base method.
func (m *MockLoggerInt) Fatalw(msg string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{msg}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatalw", varargs...)
}

// Fatalw indicates an expected call of Fatalw.
func (mr *MockLoggerIntMockRecorder) Fatalw(msg interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{msg}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatalw", reflect.TypeOf((*MockLoggerInt)(nil).Fatalw), varargs...)
}

// Info mocks base method.
func (m *MockLoggerInt) Info(args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info.
func (mr *MockLoggerIntMockRecorder) Info(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockLoggerInt)(nil).Info), args...)
}

// Infof mocks base method.
func (m *MockLoggerInt) Infof(msg string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{msg}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof.
func (mr *MockLoggerIntMockRecorder) Infof(msg interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{msg}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockLoggerInt)(nil).Infof), varargs...)
}

// Infow mocks base method.
func (m *MockLoggerInt) Infow(msg string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{msg}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infow", varargs...)
}

// Infow indicates an expected call of Infow.
func (mr *MockLoggerIntMockRecorder) Infow(msg interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{msg}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infow", reflect.TypeOf((*MockLoggerInt)(nil).Infow), varargs...)
}

// Initialize mocks base method.
func (m *MockLoggerInt) Initialize() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Initialize")
}

// Initialize indicates an expected call of Initialize.
func (mr *MockLoggerIntMockRecorder) Initialize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockLoggerInt)(nil).Initialize))
}

// Warn mocks base method.
func (m *MockLoggerInt) Warn(args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warn", varargs...)
}

// Warn indicates an expected call of Warn.
func (mr *MockLoggerIntMockRecorder) Warn(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warn", reflect.TypeOf((*MockLoggerInt)(nil).Warn), args...)
}

// Warnf mocks base method.
func (m *MockLoggerInt) Warnf(msg string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{msg}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnf", varargs...)
}

// Warnf indicates an expected call of Warnf.
func (mr *MockLoggerIntMockRecorder) Warnf(msg interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{msg}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnf", reflect.TypeOf((*MockLoggerInt)(nil).Warnf), varargs...)
}

// Warnw mocks base method.
func (m *MockLoggerInt) Warnw(msg string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{msg}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnw", varargs...)
}

// Warnw indicates an expected call of Warnw.
func (mr *MockLoggerIntMockRecorder) Warnw(msg interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{msg}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnw", reflect.TypeOf((*MockLoggerInt)(nil).Warnw), varargs...)
}
