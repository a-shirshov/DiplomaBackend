// Code generated by MockGen. DO NOT EDIT.
// Source: internal/microservices/event/usecase.go

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

// GetEvent mocks base method.
func (m *MockUsecase) GetEvent(eventId int) (*models.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEvent", eventId)
	ret0, _ := ret[0].(*models.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvent indicates an expected call of GetEvent.
func (mr *MockUsecaseMockRecorder) GetEvent(eventId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvent", reflect.TypeOf((*MockUsecase)(nil).GetEvent), eventId)
}

// GetEvents mocks base method.
func (m *MockUsecase) GetEvents(page int) ([]*models.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEvents", page)
	ret0, _ := ret[0].([]*models.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvents indicates an expected call of GetEvents.
func (mr *MockUsecaseMockRecorder) GetEvents(page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvents", reflect.TypeOf((*MockUsecase)(nil).GetEvents), page)
}

// GetPeopleCountAndCheckMeeting mocks base method.
func (m *MockUsecase) GetPeopleCountAndCheckMeeting(userID, eventID int) (int, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeopleCountAndCheckMeeting", userID, eventID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPeopleCountAndCheckMeeting indicates an expected call of GetPeopleCountAndCheckMeeting.
func (mr *MockUsecaseMockRecorder) GetPeopleCountAndCheckMeeting(userID, eventID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeopleCountAndCheckMeeting", reflect.TypeOf((*MockUsecase)(nil).GetPeopleCountAndCheckMeeting), userID, eventID)
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
