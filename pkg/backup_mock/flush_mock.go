// Code generated by MockGen. DO NOT EDIT.
// Source: types.go

// Package backup_mock is a generated GoMock package.
package backup_mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFlush is a mock of IFlush interface.
type MockFlush struct {
	ctrl     *gomock.Controller
	recorder *MockFlushMockRecorder
}

// MockFlushMockRecorder is the mock recorder for MockFlush.
type MockFlushMockRecorder struct {
	mock *MockFlush
}

// NewMockFlush creates a new mock instance.
func NewMockFlush(ctrl *gomock.Controller) *MockFlush {
	mock := &MockFlush{ctrl: ctrl}
	mock.recorder = &MockFlushMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFlush) EXPECT() *MockFlushMockRecorder {
	return m.recorder
}

// Flush mocks base method.
func (m *MockFlush) Flush() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Flush")
}

// Flush indicates an expected call of Flush.
func (mr *MockFlushMockRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockFlush)(nil).Flush))
}
