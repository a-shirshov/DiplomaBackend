// Code generated by MockGen. DO NOT EDIT.
// Source: internal/microservices/user/usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	models "Diploma/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUsecase is a mock of Usecase interface.
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase.
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance.
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// GetFavouriteKudagoEventsIDs mocks base method.
func (m *MockUsecase) GetFavouriteKudagoEventsIDs(userID int) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavouriteKudagoEventsIDs", userID)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavouriteKudagoEventsIDs indicates an expected call of GetFavouriteKudagoEventsIDs.
func (mr *MockUsecaseMockRecorder) GetFavouriteKudagoEventsIDs(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavouriteKudagoEventsIDs", reflect.TypeOf((*MockUsecase)(nil).GetFavouriteKudagoEventsIDs), userID)
}

// GetUser mocks base method.
func (m *MockUsecase) GetUser(arg0 int) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUsecaseMockRecorder) GetUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUsecase)(nil).GetUser), arg0)
}

// UpdateUser mocks base method.
func (m *MockUsecase) UpdateUser(arg0 *models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUsecaseMockRecorder) UpdateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUsecase)(nil).UpdateUser), arg0)
}

// UpdateUserImage mocks base method.
func (m *MockUsecase) UpdateUserImage(userID int, imgUUID string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserImage", userID, imgUUID)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserImage indicates an expected call of UpdateUserImage.
func (mr *MockUsecaseMockRecorder) UpdateUserImage(userID, imgUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserImage", reflect.TypeOf((*MockUsecase)(nil).UpdateUserImage), userID, imgUUID)
}
