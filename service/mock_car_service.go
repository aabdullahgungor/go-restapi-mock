// Code generated by MockGen. DO NOT EDIT.
// Source: car_service_interface.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	model "github.com/aabdullahgungor/go-restapi-mock/model"
	gomock "github.com/golang/mock/gomock"
)

// MockICarService is a mock of ICarService interface.
type MockICarService struct {
	ctrl     *gomock.Controller
	recorder *MockICarServiceMockRecorder
}

// MockICarServiceMockRecorder is the mock recorder for MockICarService.
type MockICarServiceMockRecorder struct {
	mock *MockICarService
}

// NewMockICarService creates a new mock instance.
func NewMockICarService(ctrl *gomock.Controller) *MockICarService {
	mock := &MockICarService{ctrl: ctrl}
	mock.recorder = &MockICarServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICarService) EXPECT() *MockICarServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockICarService) Create(car *model.Car) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", car)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockICarServiceMockRecorder) Create(car interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockICarService)(nil).Create), car)
}

// Delete mocks base method.
func (m *MockICarService) Delete(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockICarServiceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockICarService)(nil).Delete), id)
}

// Edit mocks base method.
func (m *MockICarService) Edit(car *model.Car) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Edit", car)
	ret0, _ := ret[0].(error)
	return ret0
}

// Edit indicates an expected call of Edit.
func (mr *MockICarServiceMockRecorder) Edit(car interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edit", reflect.TypeOf((*MockICarService)(nil).Edit), car)
}

// GetAll mocks base method.
func (m *MockICarService) GetAll() ([]model.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]model.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockICarServiceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockICarService)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockICarService) GetById(id string) (model.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(model.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockICarServiceMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockICarService)(nil).GetById), id)
}
