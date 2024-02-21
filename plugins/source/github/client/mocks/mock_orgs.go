// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/github/client (interfaces: OrganizationsService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v59/github"
)

// MockOrganizationsService is a mock of OrganizationsService interface.
type MockOrganizationsService struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsServiceMockRecorder
}

// MockOrganizationsServiceMockRecorder is the mock recorder for MockOrganizationsService.
type MockOrganizationsServiceMockRecorder struct {
	mock *MockOrganizationsService
}

// NewMockOrganizationsService creates a new mock instance.
func NewMockOrganizationsService(ctrl *gomock.Controller) *MockOrganizationsService {
	mock := &MockOrganizationsService{ctrl: ctrl}
	mock.recorder = &MockOrganizationsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsService) EXPECT() *MockOrganizationsServiceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockOrganizationsService) Get(arg0 context.Context, arg1 string) (*github.Organization, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*github.Organization)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockOrganizationsServiceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOrganizationsService)(nil).Get), arg0, arg1)
}

// GetHookDelivery mocks base method.
func (m *MockOrganizationsService) GetHookDelivery(arg0 context.Context, arg1 string, arg2, arg3 int64) (*github.HookDelivery, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHookDelivery", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*github.HookDelivery)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetHookDelivery indicates an expected call of GetHookDelivery.
func (mr *MockOrganizationsServiceMockRecorder) GetHookDelivery(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHookDelivery", reflect.TypeOf((*MockOrganizationsService)(nil).GetHookDelivery), arg0, arg1, arg2, arg3)
}

// GetOrgMembership mocks base method.
func (m *MockOrganizationsService) GetOrgMembership(arg0 context.Context, arg1, arg2 string) (*github.Membership, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgMembership", arg0, arg1, arg2)
	ret0, _ := ret[0].(*github.Membership)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOrgMembership indicates an expected call of GetOrgMembership.
func (mr *MockOrganizationsServiceMockRecorder) GetOrgMembership(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgMembership", reflect.TypeOf((*MockOrganizationsService)(nil).GetOrgMembership), arg0, arg1, arg2)
}

// ListHookDeliveries mocks base method.
func (m *MockOrganizationsService) ListHookDeliveries(arg0 context.Context, arg1 string, arg2 int64, arg3 *github.ListCursorOptions) ([]*github.HookDelivery, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListHookDeliveries", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*github.HookDelivery)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListHookDeliveries indicates an expected call of ListHookDeliveries.
func (mr *MockOrganizationsServiceMockRecorder) ListHookDeliveries(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHookDeliveries", reflect.TypeOf((*MockOrganizationsService)(nil).ListHookDeliveries), arg0, arg1, arg2, arg3)
}

// ListHooks mocks base method.
func (m *MockOrganizationsService) ListHooks(arg0 context.Context, arg1 string, arg2 *github.ListOptions) ([]*github.Hook, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListHooks", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*github.Hook)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListHooks indicates an expected call of ListHooks.
func (mr *MockOrganizationsServiceMockRecorder) ListHooks(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHooks", reflect.TypeOf((*MockOrganizationsService)(nil).ListHooks), arg0, arg1, arg2)
}

// ListInstallations mocks base method.
func (m *MockOrganizationsService) ListInstallations(arg0 context.Context, arg1 string, arg2 *github.ListOptions) (*github.OrganizationInstallations, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInstallations", arg0, arg1, arg2)
	ret0, _ := ret[0].(*github.OrganizationInstallations)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListInstallations indicates an expected call of ListInstallations.
func (mr *MockOrganizationsServiceMockRecorder) ListInstallations(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInstallations", reflect.TypeOf((*MockOrganizationsService)(nil).ListInstallations), arg0, arg1, arg2)
}

// ListMembers mocks base method.
func (m *MockOrganizationsService) ListMembers(arg0 context.Context, arg1 string, arg2 *github.ListMembersOptions) ([]*github.User, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMembers", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*github.User)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListMembers indicates an expected call of ListMembers.
func (mr *MockOrganizationsServiceMockRecorder) ListMembers(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMembers", reflect.TypeOf((*MockOrganizationsService)(nil).ListMembers), arg0, arg1, arg2)
}
