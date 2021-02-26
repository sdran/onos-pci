// Code generated by MockGen. DO NOT EDIT.
// Source: go/onos/e2t/e2/e2.pb.go

// Package e2 is a generated GoMock package.
package e2

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockE2TServiceClient is a mock of E2TServiceClient interface
type MockE2TServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockE2TServiceClientMockRecorder
}

// MockE2TServiceClientMockRecorder is the mock recorder for MockE2TServiceClient
type MockE2TServiceClientMockRecorder struct {
	mock *MockE2TServiceClient
}

// NewMockE2TServiceClient creates a new mock instance
func NewMockE2TServiceClient(ctrl *gomock.Controller) *MockE2TServiceClient {
	mock := &MockE2TServiceClient{ctrl: ctrl}
	mock.recorder = &MockE2TServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockE2TServiceClient) EXPECT() *MockE2TServiceClientMockRecorder {
	return m.recorder
}

// Stream mocks base method
func (m *MockE2TServiceClient) Stream(ctx context.Context, opts ...grpc.CallOption) (E2TService_StreamClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Stream", varargs...)
	ret0, _ := ret[0].(E2TService_StreamClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stream indicates an expected call of Stream
func (mr *MockE2TServiceClientMockRecorder) Stream(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stream", reflect.TypeOf((*MockE2TServiceClient)(nil).Stream), varargs...)
}

// MockE2TService_StreamClient is a mock of E2TService_StreamClient interface
type MockE2TService_StreamClient struct {
	ctrl     *gomock.Controller
	recorder *MockE2TService_StreamClientMockRecorder
}

// MockE2TService_StreamClientMockRecorder is the mock recorder for MockE2TService_StreamClient
type MockE2TService_StreamClientMockRecorder struct {
	mock *MockE2TService_StreamClient
}

// NewMockE2TService_StreamClient creates a new mock instance
func NewMockE2TService_StreamClient(ctrl *gomock.Controller) *MockE2TService_StreamClient {
	mock := &MockE2TService_StreamClient{ctrl: ctrl}
	mock.recorder = &MockE2TService_StreamClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockE2TService_StreamClient) EXPECT() *MockE2TService_StreamClientMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockE2TService_StreamClient) Send(arg0 *StreamRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockE2TService_StreamClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockE2TService_StreamClient)(nil).Send), arg0)
}

// Recv mocks base method
func (m *MockE2TService_StreamClient) Recv() (*StreamResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*StreamResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockE2TService_StreamClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockE2TService_StreamClient)(nil).Recv))
}

// Header mocks base method
func (m *MockE2TService_StreamClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockE2TService_StreamClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockE2TService_StreamClient)(nil).Header))
}

// Trailer mocks base method
func (m *MockE2TService_StreamClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockE2TService_StreamClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockE2TService_StreamClient)(nil).Trailer))
}

// CloseSend mocks base method
func (m *MockE2TService_StreamClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockE2TService_StreamClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockE2TService_StreamClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockE2TService_StreamClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockE2TService_StreamClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockE2TService_StreamClient)(nil).Context))
}

// SendMsg mocks base method
func (m_2 *MockE2TService_StreamClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockE2TService_StreamClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockE2TService_StreamClient)(nil).SendMsg), m)
}

// RecvMsg mocks base method
func (m_2 *MockE2TService_StreamClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockE2TService_StreamClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockE2TService_StreamClient)(nil).RecvMsg), m)
}

// MockE2TServiceServer is a mock of E2TServiceServer interface
type MockE2TServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockE2TServiceServerMockRecorder
}

// MockE2TServiceServerMockRecorder is the mock recorder for MockE2TServiceServer
type MockE2TServiceServerMockRecorder struct {
	mock *MockE2TServiceServer
}

// NewMockE2TServiceServer creates a new mock instance
func NewMockE2TServiceServer(ctrl *gomock.Controller) *MockE2TServiceServer {
	mock := &MockE2TServiceServer{ctrl: ctrl}
	mock.recorder = &MockE2TServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockE2TServiceServer) EXPECT() *MockE2TServiceServerMockRecorder {
	return m.recorder
}

// Stream mocks base method
func (m *MockE2TServiceServer) Stream(arg0 E2TService_StreamServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stream", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stream indicates an expected call of Stream
func (mr *MockE2TServiceServerMockRecorder) Stream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stream", reflect.TypeOf((*MockE2TServiceServer)(nil).Stream), arg0)
}

// MockE2TService_StreamServer is a mock of E2TService_StreamServer interface
type MockE2TService_StreamServer struct {
	ctrl     *gomock.Controller
	recorder *MockE2TService_StreamServerMockRecorder
}

// MockE2TService_StreamServerMockRecorder is the mock recorder for MockE2TService_StreamServer
type MockE2TService_StreamServerMockRecorder struct {
	mock *MockE2TService_StreamServer
}

// NewMockE2TService_StreamServer creates a new mock instance
func NewMockE2TService_StreamServer(ctrl *gomock.Controller) *MockE2TService_StreamServer {
	mock := &MockE2TService_StreamServer{ctrl: ctrl}
	mock.recorder = &MockE2TService_StreamServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockE2TService_StreamServer) EXPECT() *MockE2TService_StreamServerMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockE2TService_StreamServer) Send(arg0 *StreamResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockE2TService_StreamServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockE2TService_StreamServer)(nil).Send), arg0)
}

// Recv mocks base method
func (m *MockE2TService_StreamServer) Recv() (*StreamRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*StreamRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockE2TService_StreamServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockE2TService_StreamServer)(nil).Recv))
}

// SetHeader mocks base method
func (m *MockE2TService_StreamServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockE2TService_StreamServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockE2TService_StreamServer)(nil).SetHeader), arg0)
}

// SendHeader mocks base method
func (m *MockE2TService_StreamServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockE2TService_StreamServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockE2TService_StreamServer)(nil).SendHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockE2TService_StreamServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockE2TService_StreamServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockE2TService_StreamServer)(nil).SetTrailer), arg0)
}

// Context mocks base method
func (m *MockE2TService_StreamServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockE2TService_StreamServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockE2TService_StreamServer)(nil).Context))
}

// SendMsg mocks base method
func (m_2 *MockE2TService_StreamServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockE2TService_StreamServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockE2TService_StreamServer)(nil).SendMsg), m)
}

// RecvMsg mocks base method
func (m_2 *MockE2TService_StreamServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockE2TService_StreamServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockE2TService_StreamServer)(nil).RecvMsg), m)
}
