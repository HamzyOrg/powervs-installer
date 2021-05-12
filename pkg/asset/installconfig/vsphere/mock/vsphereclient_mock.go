// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/installer/pkg/asset/installconfig/vsphere (interfaces: Finder)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	object "github.com/vmware/govmomi/object"
	reflect "reflect"
)

// MockFinder is a mock of Finder interface
type MockFinder struct {
	ctrl     *gomock.Controller
	recorder *MockFinderMockRecorder
}

// MockFinderMockRecorder is the mock recorder for MockFinder
type MockFinderMockRecorder struct {
	mock *MockFinder
}

// NewMockFinder creates a new mock instance
func NewMockFinder(ctrl *gomock.Controller) *MockFinder {
	mock := &MockFinder{ctrl: ctrl}
	mock.recorder = &MockFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFinder) EXPECT() *MockFinderMockRecorder {
	return m.recorder
}

// ClusterComputeResource mocks base method
func (m *MockFinder) ClusterComputeResource(arg0 context.Context, arg1 string) (*object.ClusterComputeResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterComputeResource", arg0, arg1)
	ret0, _ := ret[0].(*object.ClusterComputeResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClusterComputeResource indicates an expected call of ClusterComputeResource
func (mr *MockFinderMockRecorder) ClusterComputeResource(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterComputeResource", reflect.TypeOf((*MockFinder)(nil).ClusterComputeResource), arg0, arg1)
}

// ClusterComputeResourceList mocks base method
func (m *MockFinder) ClusterComputeResourceList(arg0 context.Context, arg1 string) ([]*object.ClusterComputeResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterComputeResourceList", arg0, arg1)
	ret0, _ := ret[0].([]*object.ClusterComputeResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClusterComputeResourceList indicates an expected call of ClusterComputeResourceList
func (mr *MockFinderMockRecorder) ClusterComputeResourceList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterComputeResourceList", reflect.TypeOf((*MockFinder)(nil).ClusterComputeResourceList), arg0, arg1)
}

// Datacenter mocks base method
func (m *MockFinder) Datacenter(arg0 context.Context, arg1 string) (*object.Datacenter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Datacenter", arg0, arg1)
	ret0, _ := ret[0].(*object.Datacenter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Datacenter indicates an expected call of Datacenter
func (mr *MockFinderMockRecorder) Datacenter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Datacenter", reflect.TypeOf((*MockFinder)(nil).Datacenter), arg0, arg1)
}

// DatacenterList mocks base method
func (m *MockFinder) DatacenterList(arg0 context.Context, arg1 string) ([]*object.Datacenter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatacenterList", arg0, arg1)
	ret0, _ := ret[0].([]*object.Datacenter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DatacenterList indicates an expected call of DatacenterList
func (mr *MockFinderMockRecorder) DatacenterList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatacenterList", reflect.TypeOf((*MockFinder)(nil).DatacenterList), arg0, arg1)
}

// DatastoreList mocks base method
func (m *MockFinder) DatastoreList(arg0 context.Context, arg1 string) ([]*object.Datastore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatastoreList", arg0, arg1)
	ret0, _ := ret[0].([]*object.Datastore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DatastoreList indicates an expected call of DatastoreList
func (mr *MockFinderMockRecorder) DatastoreList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatastoreList", reflect.TypeOf((*MockFinder)(nil).DatastoreList), arg0, arg1)
}

// Folder mocks base method
func (m *MockFinder) Folder(arg0 context.Context, arg1 string) (*object.Folder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Folder", arg0, arg1)
	ret0, _ := ret[0].(*object.Folder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Folder indicates an expected call of Folder
func (mr *MockFinderMockRecorder) Folder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Folder", reflect.TypeOf((*MockFinder)(nil).Folder), arg0, arg1)
}

// Network mocks base method
func (m *MockFinder) Network(arg0 context.Context, arg1 string) (object.NetworkReference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Network", arg0, arg1)
	ret0, _ := ret[0].(object.NetworkReference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Network indicates an expected call of Network
func (mr *MockFinderMockRecorder) Network(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Network", reflect.TypeOf((*MockFinder)(nil).Network), arg0, arg1)
}

// NetworkList mocks base method
func (m *MockFinder) NetworkList(arg0 context.Context, arg1 string) ([]object.NetworkReference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkList", arg0, arg1)
	ret0, _ := ret[0].([]object.NetworkReference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetworkList indicates an expected call of NetworkList
func (mr *MockFinderMockRecorder) NetworkList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkList", reflect.TypeOf((*MockFinder)(nil).NetworkList), arg0, arg1)
}
