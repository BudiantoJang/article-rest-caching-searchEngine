// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/port/validation.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockValidation is a mock of Validation interface.
type MockValidation struct {
	ctrl     *gomock.Controller
	recorder *MockValidationMockRecorder
}

// MockValidationMockRecorder is the mock recorder for MockValidation.
type MockValidationMockRecorder struct {
	mock *MockValidation
}

// NewMockValidation creates a new mock instance.
func NewMockValidation(ctrl *gomock.Controller) *MockValidation {
	mock := &MockValidation{ctrl: ctrl}
	mock.recorder = &MockValidationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidation) EXPECT() *MockValidationMockRecorder {
	return m.recorder
}

// ValidateRequest mocks base method.
func (m *MockValidation) ValidateRequest(ctx context.Context, req interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateRequest", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateRequest indicates an expected call of ValidateRequest.
func (mr *MockValidationMockRecorder) ValidateRequest(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateRequest", reflect.TypeOf((*MockValidation)(nil).ValidateRequest), ctx, req)
}
