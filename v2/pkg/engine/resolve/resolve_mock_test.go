// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/wundergraph/graphql-go-tools/v2/pkg/engine/resolve (interfaces: DataSource,BeforeFetchHook,AfterFetchHook)

// Package resolve is a generated GoMock package.
package resolve

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDataSource is a mock of DataSource interface.
type MockDataSource struct {
	ctrl     *gomock.Controller
	recorder *MockDataSourceMockRecorder
}

// MockDataSourceMockRecorder is the mock recorder for MockDataSource.
type MockDataSourceMockRecorder struct {
	mock *MockDataSource
}

// NewMockDataSource creates a new mock instance.
func NewMockDataSource(ctrl *gomock.Controller) *MockDataSource {
	mock := &MockDataSource{ctrl: ctrl}
	mock.recorder = &MockDataSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataSource) EXPECT() *MockDataSourceMockRecorder {
	return m.recorder
}

// Load mocks base method.
func (m *MockDataSource) Load(arg0 context.Context, arg1 []byte, arg2 io.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load.
func (mr *MockDataSourceMockRecorder) Load(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockDataSource)(nil).Load), arg0, arg1, arg2)
}

// MockBeforeFetchHook is a mock of BeforeFetchHook interface.
type MockBeforeFetchHook struct {
	ctrl     *gomock.Controller
	recorder *MockBeforeFetchHookMockRecorder
}

// MockBeforeFetchHookMockRecorder is the mock recorder for MockBeforeFetchHook.
type MockBeforeFetchHookMockRecorder struct {
	mock *MockBeforeFetchHook
}

// NewMockBeforeFetchHook creates a new mock instance.
func NewMockBeforeFetchHook(ctrl *gomock.Controller) *MockBeforeFetchHook {
	mock := &MockBeforeFetchHook{ctrl: ctrl}
	mock.recorder = &MockBeforeFetchHookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBeforeFetchHook) EXPECT() *MockBeforeFetchHookMockRecorder {
	return m.recorder
}

// OnBeforeFetch mocks base method.
func (m *MockBeforeFetchHook) OnBeforeFetch(arg0 HookContext, arg1 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnBeforeFetch", arg0, arg1)
}

// OnBeforeFetch indicates an expected call of OnBeforeFetch.
func (mr *MockBeforeFetchHookMockRecorder) OnBeforeFetch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnBeforeFetch", reflect.TypeOf((*MockBeforeFetchHook)(nil).OnBeforeFetch), arg0, arg1)
}

// MockAfterFetchHook is a mock of AfterFetchHook interface.
type MockAfterFetchHook struct {
	ctrl     *gomock.Controller
	recorder *MockAfterFetchHookMockRecorder
}

// MockAfterFetchHookMockRecorder is the mock recorder for MockAfterFetchHook.
type MockAfterFetchHookMockRecorder struct {
	mock *MockAfterFetchHook
}

// NewMockAfterFetchHook creates a new mock instance.
func NewMockAfterFetchHook(ctrl *gomock.Controller) *MockAfterFetchHook {
	mock := &MockAfterFetchHook{ctrl: ctrl}
	mock.recorder = &MockAfterFetchHookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAfterFetchHook) EXPECT() *MockAfterFetchHookMockRecorder {
	return m.recorder
}

// OnData mocks base method.
func (m *MockAfterFetchHook) OnData(arg0 HookContext, arg1 []byte, arg2 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnData", arg0, arg1, arg2)
}

// OnData indicates an expected call of OnData.
func (mr *MockAfterFetchHookMockRecorder) OnData(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnData", reflect.TypeOf((*MockAfterFetchHook)(nil).OnData), arg0, arg1, arg2)
}

// OnError mocks base method.
func (m *MockAfterFetchHook) OnError(arg0 HookContext, arg1 []byte, arg2 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnError", arg0, arg1, arg2)
}

// OnError indicates an expected call of OnError.
func (mr *MockAfterFetchHookMockRecorder) OnError(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnError", reflect.TypeOf((*MockAfterFetchHook)(nil).OnError), arg0, arg1, arg2)
}
