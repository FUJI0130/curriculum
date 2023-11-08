// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/FUJI0130/curriculum/src/core/domain/userdm (interfaces: UserRepository)

// Package mock_user is a generated GoMock package.
package mock_user

import (
	context "context"
	userdm "github.com/FUJI0130/curriculum/src/core/domain/userdm"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserRepository is a mock of UserRepository interface
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// FindByEmail mocks base method
func (m *MockUserRepository) FindByEmail(arg0 context.Context, arg1 string) (*userdm.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", arg0, arg1)
	ret0, _ := ret[0].(*userdm.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail
func (mr *MockUserRepositoryMockRecorder) FindByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindByEmail), arg0, arg1)
}

// FindByUserID mocks base method
func (m *MockUserRepository) FindByUserID(arg0 context.Context, arg1 string) (*userdm.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserID", arg0, arg1)
	ret0, _ := ret[0].(*userdm.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserID indicates an expected call of FindByUserID
func (mr *MockUserRepositoryMockRecorder) FindByUserID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserID", reflect.TypeOf((*MockUserRepository)(nil).FindByUserID), arg0, arg1)
}

// FindByUserName mocks base method
func (m *MockUserRepository) FindByUserName(arg0 context.Context, arg1 string) (*userdm.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserName", arg0, arg1)
	ret0, _ := ret[0].(*userdm.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserName indicates an expected call of FindByUserName
func (mr *MockUserRepositoryMockRecorder) FindByUserName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserName", reflect.TypeOf((*MockUserRepository)(nil).FindByUserName), arg0, arg1)
}

// FindByUserNames mocks base method
func (m *MockUserRepository) FindByUserNames(arg0 context.Context, arg1 []string) (map[string]*userdm.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserNames", arg0, arg1)
	ret0, _ := ret[0].(map[string]*userdm.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserNames indicates an expected call of FindByUserNames
func (mr *MockUserRepositoryMockRecorder) FindByUserNames(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserNames", reflect.TypeOf((*MockUserRepository)(nil).FindByUserNames), arg0, arg1)
}

// FindCareersByUserID mocks base method
func (m *MockUserRepository) FindCareersByUserID(arg0 context.Context, arg1 string) ([]userdm.Career, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCareersByUserID", arg0, arg1)
	ret0, _ := ret[0].([]userdm.Career)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCareersByUserID indicates an expected call of FindCareersByUserID
func (mr *MockUserRepositoryMockRecorder) FindCareersByUserID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCareersByUserID", reflect.TypeOf((*MockUserRepository)(nil).FindCareersByUserID), arg0, arg1)
}

// FindSkillsByUserID mocks base method
func (m *MockUserRepository) FindSkillsByUserID(arg0 context.Context, arg1 string) ([]userdm.Skill, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindSkillsByUserID", arg0, arg1)
	ret0, _ := ret[0].([]userdm.Skill)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSkillsByUserID indicates an expected call of FindSkillsByUserID
func (mr *MockUserRepositoryMockRecorder) FindSkillsByUserID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSkillsByUserID", reflect.TypeOf((*MockUserRepository)(nil).FindSkillsByUserID), arg0, arg1)
}

// Store mocks base method
func (m *MockUserRepository) Store(arg0 context.Context, arg1 *userdm.UserDomain) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store
func (mr *MockUserRepositoryMockRecorder) Store(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockUserRepository)(nil).Store), arg0, arg1)
}

// Update mocks base method
func (m *MockUserRepository) Update(arg0 context.Context, arg1 *userdm.UserDomain) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockUserRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepository)(nil).Update), arg0, arg1)
}
