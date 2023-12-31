// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/FUJI0130/curriculum/src/core/domain/userdm (interfaces: ExistByNameDomainService)

// Package mockExistByNameDomainService is a generated GoMock package.
package mockExistByNameDomainService

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockExistByNameDomainService is a mock of ExistByNameDomainService interface.
type MockExistByNameDomainService struct {
	ctrl     *gomock.Controller
	recorder *MockExistByNameDomainServiceMockRecorder
}

// MockExistByNameDomainServiceMockRecorder is the mock recorder for MockExistByNameDomainService.
type MockExistByNameDomainServiceMockRecorder struct {
	mock *MockExistByNameDomainService
}

// NewMockExistByNameDomainService creates a new mock instance.
func NewMockExistByNameDomainService(ctrl *gomock.Controller) *MockExistByNameDomainService {
	mock := &MockExistByNameDomainService{ctrl: ctrl}
	mock.recorder = &MockExistByNameDomainServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExistByNameDomainService) EXPECT() *MockExistByNameDomainServiceMockRecorder {
	return m.recorder
}

// Exec mocks base method.
func (m *MockExistByNameDomainService) Exec(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exec", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockExistByNameDomainServiceMockRecorder) Exec(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockExistByNameDomainService)(nil).Exec), arg0, arg1)
}
