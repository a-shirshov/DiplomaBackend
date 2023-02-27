// Code generated by MockGen. DO NOT EDIT.
// Source: internal/microservices/event/usecase.go

// Package mock is a generated GoMock package.
package mock

import (
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

// CheckKudaGoFavourite mocks base method.
func (m *MockUsecase) CheckKudaGoFavourite(userID, eventID int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckKudaGoFavourite", userID, eventID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckKudaGoFavourite indicates an expected call of CheckKudaGoFavourite.
func (mr *MockUsecaseMockRecorder) CheckKudaGoFavourite(userID, eventID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckKudaGoFavourite", reflect.TypeOf((*MockUsecase)(nil).CheckKudaGoFavourite), userID, eventID)
}

// CheckKudaGoMeeting mocks base method.
func (m *MockUsecase) CheckKudaGoMeeting(userID, eventID int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckKudaGoMeeting", userID, eventID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckKudaGoMeeting indicates an expected call of CheckKudaGoMeeting.
func (mr *MockUsecaseMockRecorder) CheckKudaGoMeeting(userID, eventID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckKudaGoMeeting", reflect.TypeOf((*MockUsecase)(nil).CheckKudaGoMeeting), userID, eventID)
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

// GetPeopleCountWithEventCreatedIfNecessary mocks base method.
func (m *MockUsecase) GetPeopleCountWithEventCreatedIfNecessary(eventID int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeopleCountWithEventCreatedIfNecessary", eventID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPeopleCountWithEventCreatedIfNecessary indicates an expected call of GetPeopleCountWithEventCreatedIfNecessary.
func (mr *MockUsecaseMockRecorder) GetPeopleCountWithEventCreatedIfNecessary(eventID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeopleCountWithEventCreatedIfNecessary", reflect.TypeOf((*MockUsecase)(nil).GetPeopleCountWithEventCreatedIfNecessary), eventID)
}

// SwitchEventFavourite mocks base method.
func (m *MockUsecase) SwitchEventFavourite(userID, eventID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SwitchEventFavourite", userID, eventID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SwitchEventFavourite indicates an expected call of SwitchEventFavourite.
func (mr *MockUsecaseMockRecorder) SwitchEventFavourite(userID, eventID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SwitchEventFavourite", reflect.TypeOf((*MockUsecase)(nil).SwitchEventFavourite), userID, eventID)
}

// SwitchEventMeeting mocks base method.
func (m *MockUsecase) SwitchEventMeeting(userID, eventID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SwitchEventMeeting", userID, eventID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SwitchEventMeeting indicates an expected call of SwitchEventMeeting.
func (mr *MockUsecaseMockRecorder) SwitchEventMeeting(userID, eventID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SwitchEventMeeting", reflect.TypeOf((*MockUsecase)(nil).SwitchEventMeeting), userID, eventID)
}
