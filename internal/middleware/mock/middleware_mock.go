// Code generated by MockGen. DO NOT EDIT.
// Source: internal/middleware/middleware.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockMiddleware is a mock of Middleware interface.
type MockMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MockMiddlewareMockRecorder
}

// MockMiddlewareMockRecorder is the mock recorder for MockMiddleware.
type MockMiddlewareMockRecorder struct {
	mock *MockMiddleware
}

// NewMockMiddleware creates a new mock instance.
func NewMockMiddleware(ctrl *gomock.Controller) *MockMiddleware {
	mock := &MockMiddleware{ctrl: ctrl}
	mock.recorder = &MockMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMiddleware) EXPECT() *MockMiddlewareMockRecorder {
	return m.recorder
}

// CORSMiddleware mocks base method.
func (m *MockMiddleware) CORSMiddleware() gin.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CORSMiddleware")
	ret0, _ := ret[0].(gin.HandlerFunc)
	return ret0
}

// CORSMiddleware indicates an expected call of CORSMiddleware.
func (mr *MockMiddlewareMockRecorder) CORSMiddleware() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CORSMiddleware", reflect.TypeOf((*MockMiddleware)(nil).CORSMiddleware))
}

// MiddlewareValidateLoginUser mocks base method.
func (m *MockMiddleware) MiddlewareValidateLoginUser() gin.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MiddlewareValidateLoginUser")
	ret0, _ := ret[0].(gin.HandlerFunc)
	return ret0
}

// MiddlewareValidateLoginUser indicates an expected call of MiddlewareValidateLoginUser.
func (mr *MockMiddlewareMockRecorder) MiddlewareValidateLoginUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MiddlewareValidateLoginUser", reflect.TypeOf((*MockMiddleware)(nil).MiddlewareValidateLoginUser))
}

// MiddlewareValidateRedeemCode mocks base method.
func (m *MockMiddleware) MiddlewareValidateRedeemCode() gin.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MiddlewareValidateRedeemCode")
	ret0, _ := ret[0].(gin.HandlerFunc)
	return ret0
}

// MiddlewareValidateRedeemCode indicates an expected call of MiddlewareValidateRedeemCode.
func (mr *MockMiddlewareMockRecorder) MiddlewareValidateRedeemCode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MiddlewareValidateRedeemCode", reflect.TypeOf((*MockMiddleware)(nil).MiddlewareValidateRedeemCode))
}

// MiddlewareValidateUser mocks base method.
func (m *MockMiddleware) MiddlewareValidateUser() gin.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MiddlewareValidateUser")
	ret0, _ := ret[0].(gin.HandlerFunc)
	return ret0
}

// MiddlewareValidateUser indicates an expected call of MiddlewareValidateUser.
func (mr *MockMiddlewareMockRecorder) MiddlewareValidateUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MiddlewareValidateUser", reflect.TypeOf((*MockMiddleware)(nil).MiddlewareValidateUser))
}

// MiddlewareValidateUserEvent mocks base method.
func (m *MockMiddleware) MiddlewareValidateUserEvent() gin.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MiddlewareValidateUserEvent")
	ret0, _ := ret[0].(gin.HandlerFunc)
	return ret0
}

// MiddlewareValidateUserEvent indicates an expected call of MiddlewareValidateUserEvent.
func (mr *MockMiddlewareMockRecorder) MiddlewareValidateUserEvent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MiddlewareValidateUserEvent", reflect.TypeOf((*MockMiddleware)(nil).MiddlewareValidateUserEvent))
}

// MiddlewareValidateUserFormData mocks base method.
func (m *MockMiddleware) MiddlewareValidateUserFormData() gin.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MiddlewareValidateUserFormData")
	ret0, _ := ret[0].(gin.HandlerFunc)
	return ret0
}

// MiddlewareValidateUserFormData indicates an expected call of MiddlewareValidateUserFormData.
func (mr *MockMiddlewareMockRecorder) MiddlewareValidateUserFormData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MiddlewareValidateUserFormData", reflect.TypeOf((*MockMiddleware)(nil).MiddlewareValidateUserFormData))
}

// TokenAuthMiddleware mocks base method.
func (m *MockMiddleware) TokenAuthMiddleware() gin.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenAuthMiddleware")
	ret0, _ := ret[0].(gin.HandlerFunc)
	return ret0
}

// TokenAuthMiddleware indicates an expected call of TokenAuthMiddleware.
func (mr *MockMiddlewareMockRecorder) TokenAuthMiddleware() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenAuthMiddleware", reflect.TypeOf((*MockMiddleware)(nil).TokenAuthMiddleware))
}
