// Code generated by MockGen. DO NOT EDIT.
// Source: src/core/domain/tagdm/tag_repository.go

// Package mockTag is a generated GoMock package.
package mockTag

import (
	context "context"
	reflect "reflect"

	tagdm "github.com/FUJI0130/curriculum/src/core/domain/tagdm"
	gomock "github.com/golang/mock/gomock"
)

// MockTagRepository is a mock of TagRepository interface.
type MockTagRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTagRepositoryMockRecorder
}

// MockTagRepositoryMockRecorder is the mock recorder for MockTagRepository.
type MockTagRepositoryMockRecorder struct {
	mock *MockTagRepository
}

// NewMockTagRepository creates a new mock instance.
func NewMockTagRepository(ctrl *gomock.Controller) *MockTagRepository {
	mock := &MockTagRepository{ctrl: ctrl}
	mock.recorder = &MockTagRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTagRepository) EXPECT() *MockTagRepositoryMockRecorder {
	return m.recorder
}

// CreateNewTag mocks base method.
func (m *MockTagRepository) CreateNewTag(ctx context.Context, name string) (*tagdm.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewTag", ctx, name)
	ret0, _ := ret[0].(*tagdm.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewTag indicates an expected call of CreateNewTag.
func (mr *MockTagRepositoryMockRecorder) CreateNewTag(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewTag", reflect.TypeOf((*MockTagRepository)(nil).CreateNewTag), ctx, name)
}

// FindByID mocks base method.
func (m *MockTagRepository) FindByID(ctx context.Context, id string) (*tagdm.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(*tagdm.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockTagRepositoryMockRecorder) FindByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockTagRepository)(nil).FindByID), ctx, id)
}

// FindByName mocks base method.
func (m *MockTagRepository) FindByName(ctx context.Context, name string) (*tagdm.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", ctx, name)
	ret0, _ := ret[0].(*tagdm.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *MockTagRepositoryMockRecorder) FindByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockTagRepository)(nil).FindByName), ctx, name)
}

// Store mocks base method.
func (m *MockTagRepository) Store(ctx context.Context, tag *tagdm.Tag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", ctx, tag)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockTagRepositoryMockRecorder) Store(ctx, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockTagRepository)(nil).Store), ctx, tag)
}
