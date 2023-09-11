// Code generated by MockGen. DO NOT EDIT.
// Source: libraryService/internal/storage (interfaces: IStorage)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	entity "libraryService/internal/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIStorage is a mock of IStorage interface.
type MockIStorage struct {
	ctrl     *gomock.Controller
	recorder *MockIStorageMockRecorder
}

// MockIStorageMockRecorder is the mock recorder for MockIStorage.
type MockIStorageMockRecorder struct {
	mock *MockIStorage
}

// NewMockIStorage creates a new mock instance.
func NewMockIStorage(ctrl *gomock.Controller) *MockIStorage {
	mock := &MockIStorage{ctrl: ctrl}
	mock.recorder = &MockIStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStorage) EXPECT() *MockIStorageMockRecorder {
	return m.recorder
}

// CheckIfAuthorExists mocks base method.
func (m *MockIStorage) CheckIfAuthorExists(arg0 context.Context, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIfAuthorExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckIfAuthorExists indicates an expected call of CheckIfAuthorExists.
func (mr *MockIStorageMockRecorder) CheckIfAuthorExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIfAuthorExists", reflect.TypeOf((*MockIStorage)(nil).CheckIfAuthorExists), arg0, arg1)
}

// CheckIfBookExists mocks base method.
func (m *MockIStorage) CheckIfBookExists(arg0 context.Context, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIfBookExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckIfBookExists indicates an expected call of CheckIfBookExists.
func (mr *MockIStorageMockRecorder) CheckIfBookExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIfBookExists", reflect.TypeOf((*MockIStorage)(nil).CheckIfBookExists), arg0, arg1)
}

// GetAuthorsByBookTitle mocks base method.
func (m *MockIStorage) GetAuthorsByBookTitle(arg0 context.Context, arg1 string) []entity.Author {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorsByBookTitle", arg0, arg1)
	ret0, _ := ret[0].([]entity.Author)
	return ret0
}

// GetAuthorsByBookTitle indicates an expected call of GetAuthorsByBookTitle.
func (mr *MockIStorageMockRecorder) GetAuthorsByBookTitle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorsByBookTitle", reflect.TypeOf((*MockIStorage)(nil).GetAuthorsByBookTitle), arg0, arg1)
}

// GetBooksByAuthorName mocks base method.
func (m *MockIStorage) GetBooksByAuthorName(arg0 context.Context, arg1 string) []entity.Book {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBooksByAuthorName", arg0, arg1)
	ret0, _ := ret[0].([]entity.Book)
	return ret0
}

// GetBooksByAuthorName indicates an expected call of GetBooksByAuthorName.
func (mr *MockIStorageMockRecorder) GetBooksByAuthorName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBooksByAuthorName", reflect.TypeOf((*MockIStorage)(nil).GetBooksByAuthorName), arg0, arg1)
}
