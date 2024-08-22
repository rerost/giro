// Code generated by MockGen. DO NOT EDIT.
// Source: ./client.go
//
// Generated by this command:
//
//	mockgen -source=./client.go -destination=../../mock/grpcreflectiface/client_test.go -package=mockgrpcreflectiface
//

// Package mockgrpcreflectiface is a generated GoMock package.
package mockgrpcreflectiface

import (
	reflect "reflect"

	desc "github.com/jhump/protoreflect/desc"
	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ListServices mocks base method.
func (m *MockClient) ListServices() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServices")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices.
func (mr *MockClientMockRecorder) ListServices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockClient)(nil).ListServices))
}

// ResolveMessage mocks base method.
func (m *MockClient) ResolveMessage(messageName string) (*desc.MessageDescriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveMessage", messageName)
	ret0, _ := ret[0].(*desc.MessageDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveMessage indicates an expected call of ResolveMessage.
func (mr *MockClientMockRecorder) ResolveMessage(messageName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveMessage", reflect.TypeOf((*MockClient)(nil).ResolveMessage), messageName)
}

// ResolveService mocks base method.
func (m *MockClient) ResolveService(serviceName string) (*desc.ServiceDescriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveService", serviceName)
	ret0, _ := ret[0].(*desc.ServiceDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveService indicates an expected call of ResolveService.
func (mr *MockClientMockRecorder) ResolveService(serviceName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveService", reflect.TypeOf((*MockClient)(nil).ResolveService), serviceName)
}
