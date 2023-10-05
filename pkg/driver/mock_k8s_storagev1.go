// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/storage/v1 (interfaces: VolumeAttachmentInterface,StorageV1Interface)

// Package driver is a generated GoMock package.
package driver

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/storage/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/storage/v1"
	v12 "k8s.io/client-go/kubernetes/typed/storage/v1"
	rest "k8s.io/client-go/rest"
)

// MockVolumeAttachmentInterface is a mock of VolumeAttachmentInterface interface.
type MockVolumeAttachmentInterface struct {
	ctrl     *gomock.Controller
	recorder *MockVolumeAttachmentInterfaceMockRecorder
}

// MockVolumeAttachmentInterfaceMockRecorder is the mock recorder for MockVolumeAttachmentInterface.
type MockVolumeAttachmentInterfaceMockRecorder struct {
	mock *MockVolumeAttachmentInterface
}

// NewMockVolumeAttachmentInterface creates a new mock instance.
func NewMockVolumeAttachmentInterface(ctrl *gomock.Controller) *MockVolumeAttachmentInterface {
	mock := &MockVolumeAttachmentInterface{ctrl: ctrl}
	mock.recorder = &MockVolumeAttachmentInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVolumeAttachmentInterface) EXPECT() *MockVolumeAttachmentInterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockVolumeAttachmentInterface) Apply(arg0 context.Context, arg1 *v11.VolumeAttachmentApplyConfiguration, arg2 v10.ApplyOptions) (*v1.VolumeAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.VolumeAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockVolumeAttachmentInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockVolumeAttachmentInterface)(nil).Apply), arg0, arg1, arg2)
}

// ApplyStatus mocks base method.
func (m *MockVolumeAttachmentInterface) ApplyStatus(arg0 context.Context, arg1 *v11.VolumeAttachmentApplyConfiguration, arg2 v10.ApplyOptions) (*v1.VolumeAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.VolumeAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplyStatus indicates an expected call of ApplyStatus.
func (mr *MockVolumeAttachmentInterfaceMockRecorder) ApplyStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyStatus", reflect.TypeOf((*MockVolumeAttachmentInterface)(nil).ApplyStatus), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockVolumeAttachmentInterface) Create(arg0 context.Context, arg1 *v1.VolumeAttachment, arg2 v10.CreateOptions) (*v1.VolumeAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.VolumeAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockVolumeAttachmentInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVolumeAttachmentInterface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockVolumeAttachmentInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockVolumeAttachmentInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVolumeAttachmentInterface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockVolumeAttachmentInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockVolumeAttachmentInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockVolumeAttachmentInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockVolumeAttachmentInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.VolumeAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.VolumeAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockVolumeAttachmentInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVolumeAttachmentInterface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockVolumeAttachmentInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.VolumeAttachmentList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.VolumeAttachmentList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockVolumeAttachmentInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVolumeAttachmentInterface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockVolumeAttachmentInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.VolumeAttachment, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.VolumeAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockVolumeAttachmentInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockVolumeAttachmentInterface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockVolumeAttachmentInterface) Update(arg0 context.Context, arg1 *v1.VolumeAttachment, arg2 v10.UpdateOptions) (*v1.VolumeAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.VolumeAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockVolumeAttachmentInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockVolumeAttachmentInterface)(nil).Update), arg0, arg1, arg2)
}

// UpdateStatus mocks base method.
func (m *MockVolumeAttachmentInterface) UpdateStatus(arg0 context.Context, arg1 *v1.VolumeAttachment, arg2 v10.UpdateOptions) (*v1.VolumeAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.VolumeAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockVolumeAttachmentInterfaceMockRecorder) UpdateStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockVolumeAttachmentInterface)(nil).UpdateStatus), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockVolumeAttachmentInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockVolumeAttachmentInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockVolumeAttachmentInterface)(nil).Watch), arg0, arg1)
}

// MockStorageV1Interface is a mock of StorageV1Interface interface.
type MockStorageV1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockStorageV1InterfaceMockRecorder
}

// MockStorageV1InterfaceMockRecorder is the mock recorder for MockStorageV1Interface.
type MockStorageV1InterfaceMockRecorder struct {
	mock *MockStorageV1Interface
}

// NewMockStorageV1Interface creates a new mock instance.
func NewMockStorageV1Interface(ctrl *gomock.Controller) *MockStorageV1Interface {
	mock := &MockStorageV1Interface{ctrl: ctrl}
	mock.recorder = &MockStorageV1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageV1Interface) EXPECT() *MockStorageV1InterfaceMockRecorder {
	return m.recorder
}

// CSIDrivers mocks base method.
func (m *MockStorageV1Interface) CSIDrivers() v12.CSIDriverInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CSIDrivers")
	ret0, _ := ret[0].(v12.CSIDriverInterface)
	return ret0
}

// CSIDrivers indicates an expected call of CSIDrivers.
func (mr *MockStorageV1InterfaceMockRecorder) CSIDrivers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CSIDrivers", reflect.TypeOf((*MockStorageV1Interface)(nil).CSIDrivers))
}

// CSINodes mocks base method.
func (m *MockStorageV1Interface) CSINodes() v12.CSINodeInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CSINodes")
	ret0, _ := ret[0].(v12.CSINodeInterface)
	return ret0
}

// CSINodes indicates an expected call of CSINodes.
func (mr *MockStorageV1InterfaceMockRecorder) CSINodes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CSINodes", reflect.TypeOf((*MockStorageV1Interface)(nil).CSINodes))
}

// CSIStorageCapacities mocks base method.
func (m *MockStorageV1Interface) CSIStorageCapacities(arg0 string) v12.CSIStorageCapacityInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CSIStorageCapacities", arg0)
	ret0, _ := ret[0].(v12.CSIStorageCapacityInterface)
	return ret0
}

// CSIStorageCapacities indicates an expected call of CSIStorageCapacities.
func (mr *MockStorageV1InterfaceMockRecorder) CSIStorageCapacities(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CSIStorageCapacities", reflect.TypeOf((*MockStorageV1Interface)(nil).CSIStorageCapacities), arg0)
}

// RESTClient mocks base method.
func (m *MockStorageV1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient.
func (mr *MockStorageV1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockStorageV1Interface)(nil).RESTClient))
}

// StorageClasses mocks base method.
func (m *MockStorageV1Interface) StorageClasses() v12.StorageClassInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageClasses")
	ret0, _ := ret[0].(v12.StorageClassInterface)
	return ret0
}

// StorageClasses indicates an expected call of StorageClasses.
func (mr *MockStorageV1InterfaceMockRecorder) StorageClasses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageClasses", reflect.TypeOf((*MockStorageV1Interface)(nil).StorageClasses))
}

// VolumeAttachments mocks base method.
func (m *MockStorageV1Interface) VolumeAttachments() v12.VolumeAttachmentInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeAttachments")
	ret0, _ := ret[0].(v12.VolumeAttachmentInterface)
	return ret0
}

// VolumeAttachments indicates an expected call of VolumeAttachments.
func (mr *MockStorageV1InterfaceMockRecorder) VolumeAttachments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeAttachments", reflect.TypeOf((*MockStorageV1Interface)(nil).VolumeAttachments))
}