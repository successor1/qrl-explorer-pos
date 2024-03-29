// Code generated by MockGen. DO NOT EDIT.
// Source: foo.go

// Package mock_rpc is a generated GoMock package.
package mock_rpc

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMyHTTPClient is a mock of MyHTTPClient interface.
type MockMyHTTPClient struct {
	ctrl     *gomock.Controller
	recorder *MockMyHTTPClientMockRecorder
}

// MockMyHTTPClientMockRecorder is the mock recorder for MockMyHTTPClient.
type MockMyHTTPClientMockRecorder struct {
	mock *MockMyHTTPClient
}

// NewMockMyHTTPClient creates a new mock instance.
func NewMockMyHTTPClient(ctrl *gomock.Controller) *MockMyHTTPClient {
	mock := &MockMyHTTPClient{ctrl: ctrl}
	mock.recorder = &MockMyHTTPClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMyHTTPClient) EXPECT() *MockMyHTTPClientMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockMyHTTPClient) Do(req *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", req)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockMyHTTPClientMockRecorder) Do(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockMyHTTPClient)(nil).Do), req)
}
