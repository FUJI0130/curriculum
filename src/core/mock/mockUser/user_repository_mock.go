// Code generated by MockGen. DO NOT EDIT.
// Source: src/core/domain/userdm/user_repository.go

// Package mockUser is a generated GoMock package.
package mockUser

import (
	context "context"
	reflect "reflect"

	userdm "github.com/FUJI0130/curriculum/src/core/domain/userdm"
	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// FindByName mocks base method.
func (m *MockUserRepository) FindByName(ctx context.Context, name string) (*userdm.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", ctx, name)
	ret0, _ := ret[0].(*userdm.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *MockUserRepositoryMockRecorder) FindByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockUserRepository)(nil).FindByName), ctx, name)
}

// FindByNames mocks base method.
func (m *MockUserRepository) FindByNames(ctx context.Context, names []string) (map[string]*userdm.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByNames", ctx, names)
	ret0, _ := ret[0].(map[string]*userdm.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByNames indicates an expected call of FindByNames.
func (mr *MockUserRepositoryMockRecorder) FindByNames(ctx, names interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByNames", reflect.TypeOf((*MockUserRepository)(nil).FindByNames), ctx, names)
}

// Store mocks base method.
func (m *MockUserRepository) Store(ctx context.Context, user *userdm.User, skills []*userdm.Skill, careers []*userdm.Career) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", ctx, user, skills, careers)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockUserRepositoryMockRecorder) Store(ctx, user, skills, careers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockUserRepository)(nil).Store), ctx, user, skills, careers)
}
