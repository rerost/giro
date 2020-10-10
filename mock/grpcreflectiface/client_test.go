// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rerost/giro/domain/grpcreflectiface (interfaces: Client)

// Package mock_grpcreflectiface is a generated GoMock package.
package mock_grpcreflectiface

import (
	gomock "github.com/golang/mock/gomock"
	desc "github.com/jhump/protoreflect/desc"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ListServices mocks base method
func (m *MockClient) ListServices() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServices")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices
func (mr *MockClientMockRecorder) ListServices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockClient)(nil).ListServices))
}

// ResolveMessage mocks base method
func (m *MockClient) ResolveMessage(arg0 string) (*desc.MessageDescriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveMessage", arg0)
	ret0, _ := ret[0].(*desc.MessageDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveMessage indicates an expected call of ResolveMessage
func (mr *MockClientMockRecorder) ResolveMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveMessage", reflect.TypeOf((*MockClient)(nil).ResolveMessage), arg0)
}

// ResolveService mocks base method
func (m *MockClient) ResolveService(arg0 string) (*desc.ServiceDescriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveService", arg0)
	ret0, _ := ret[0].(*desc.ServiceDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveService indicates an expected call of ResolveService
func (mr *MockClientMockRecorder) ResolveService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveService", reflect.TypeOf((*MockClient)(nil).ResolveService), arg0)
}