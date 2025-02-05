// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	models "github.com/Shemistan/Lesson_6/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockIService is a mock of IService interface.
type MockIService struct {
	ctrl     *gomock.Controller
	recorder *MockIServiceMockRecorder
}

// MockIServiceMockRecorder is the mock recorder for MockIService.
type MockIServiceMockRecorder struct {
	mock *MockIService
}

// NewMockIService creates a new mock instance.
func NewMockIService(ctrl *gomock.Controller) *MockIService {
	mock := &MockIService{ctrl: ctrl}
	mock.recorder = &MockIServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIService) EXPECT() *MockIServiceMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockIService) Add(user *models.User) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", user)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockIServiceMockRecorder) Add(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockIService)(nil).Add), user)
}

// DeleteUser mocks base method.
func (m *MockIService) DeleteUser(userId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockIServiceMockRecorder) DeleteUser(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockIService)(nil).DeleteUser), userId)
}

// GetStatistics mocks base method.
func (m *MockIService) GetStatistics() *models.Statistic {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatistics")
	ret0, _ := ret[0].(*models.Statistic)
	return ret0
}

// GetStatistics indicates an expected call of GetStatistics.
func (mr *MockIServiceMockRecorder) GetStatistics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatistics", reflect.TypeOf((*MockIService)(nil).GetStatistics))
}

// GetUser mocks base method.
func (m *MockIService) GetUser(userId int64) (*models.GetUserData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", userId)
	ret0, _ := ret[0].(*models.GetUserData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockIServiceMockRecorder) GetUser(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockIService)(nil).GetUser), userId)
}

// GetUserByLogin mocks base method.
func (m *MockIService) GetUserByLogin(login, password string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByLogin", login, password)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByLogin indicates an expected call of GetUserByLogin.
func (mr *MockIServiceMockRecorder) GetUserByLogin(login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByLogin", reflect.TypeOf((*MockIService)(nil).GetUserByLogin), login, password)
}

// GetUsers mocks base method.
func (m *MockIService) GetUsers() ([]*models.GetUserData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers")
	ret0, _ := ret[0].([]*models.GetUserData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockIServiceMockRecorder) GetUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockIService)(nil).GetUsers))
}

// UpdateUser mocks base method.
func (m *MockIService) UpdateUser(userId int64, user *models.UpdateUserData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", userId, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockIServiceMockRecorder) UpdateUser(userId, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockIService)(nil).UpdateUser), userId, user)
}
