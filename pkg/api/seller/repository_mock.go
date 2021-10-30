// Code generated by MockGen. DO NOT EDIT.
// Source: repository_interface.go

// Package seller is a generated GoMock package.
package seller

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIRepository is a mock of IRepository interface.
type MockIRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryMockRecorder
}

// MockIRepositoryMockRecorder is the mock recorder for MockIRepository.
type MockIRepositoryMockRecorder struct {
	mock *MockIRepository
}

// NewMockIRepository creates a new mock instance.
func NewMockIRepository(ctrl *gomock.Controller) *MockIRepository {
	mock := &MockIRepository{ctrl: ctrl}
	mock.recorder = &MockIRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepository) EXPECT() *MockIRepositoryMockRecorder {
	return m.recorder
}

// FindByUUID mocks base method.
func (m *MockIRepository) FindByUUID(uuid string) (*Seller, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUUID", uuid)
	ret0, _ := ret[0].(*Seller)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUUID indicates an expected call of FindByUUID.
func (mr *MockIRepositoryMockRecorder) FindByUUID(uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUUID", reflect.TypeOf((*MockIRepository)(nil).FindByUUID), uuid)
}

// getTop mocks base method.
func (m *MockIRepository) getTop(size int) ([]*Seller, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getTop", size)
	ret0, _ := ret[0].([]*Seller)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getTop indicates an expected call of getTop.
func (mr *MockIRepositoryMockRecorder) getTop(size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getTop", reflect.TypeOf((*MockIRepository)(nil).getTop), size)
}

// list mocks base method.
func (m *MockIRepository) list() ([]*Seller, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "list")
	ret0, _ := ret[0].([]*Seller)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// list indicates an expected call of list.
func (mr *MockIRepositoryMockRecorder) list() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "list", reflect.TypeOf((*MockIRepository)(nil).list))
}
