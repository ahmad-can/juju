// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/provider/oci/common (interfaces: OCIVirtualNetworkingClient)

package testing

import (
	context "context"
	reflect "reflect"

	core "github.com/oracle/oci-go-sdk/v65/core"
	gomock "go.uber.org/mock/gomock"
)

// MockOCIVirtualNetworkingClient is a mock of OCIVirtualNetworkingClient interface.
type MockOCIVirtualNetworkingClient struct {
	ctrl     *gomock.Controller
	recorder *MockOCIVirtualNetworkingClientMockRecorder
}

// MockOCIVirtualNetworkingClientMockRecorder is the mock recorder for MockOCIVirtualNetworkingClient.
type MockOCIVirtualNetworkingClientMockRecorder struct {
	mock *MockOCIVirtualNetworkingClient
}

// NewMockOCIVirtualNetworkingClient creates a new mock instance.
func NewMockOCIVirtualNetworkingClient(ctrl *gomock.Controller) *MockOCIVirtualNetworkingClient {
	mock := &MockOCIVirtualNetworkingClient{ctrl: ctrl}
	mock.recorder = &MockOCIVirtualNetworkingClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOCIVirtualNetworkingClient) EXPECT() *MockOCIVirtualNetworkingClientMockRecorder {
	return m.recorder
}

// CreateInternetGateway mocks base method.
func (m *MockOCIVirtualNetworkingClient) CreateInternetGateway(arg0 context.Context, arg1 core.CreateInternetGatewayRequest) (core.CreateInternetGatewayResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInternetGateway", arg0, arg1)
	ret0, _ := ret[0].(core.CreateInternetGatewayResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInternetGateway indicates an expected call of CreateInternetGateway.
func (mr *MockOCIVirtualNetworkingClientMockRecorder) CreateInternetGateway(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInternetGateway", reflect.TypeOf((*MockOCIVirtualNetworkingClient)(nil).CreateInternetGateway), arg0, arg1)
}

// CreateRouteTable mocks base method.
func (m *MockOCIVirtualNetworkingClient) CreateRouteTable(arg0 context.Context, arg1 core.CreateRouteTableRequest) (core.CreateRouteTableResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRouteTable", arg0, arg1)
	ret0, _ := ret[0].(core.CreateRouteTableResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRouteTable indicates an expected call of CreateRouteTable.
func (mr *MockOCIVirtualNetworkingClientMockRecorder) CreateRouteTable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRouteTable", reflect.TypeOf((*MockOCIVirtualNetworkingClient)(nil).CreateRouteTable), arg0, arg1)
}

// CreateSecurityList mocks base method.
func (m *MockOCIVirtualNetworkingClient) CreateSecurityList(arg0 context.Context, arg1 core.CreateSecurityListRequest) (core.CreateSecurityListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecurityList", arg0, arg1)
	ret0, _ := ret[0].(core.CreateSecurityListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecurityList indicates an expected call of CreateSecurityList.
func (mr *MockOCIVirtualNetworkingClientMockRecorder) CreateSecurityList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecurityList", reflect.TypeOf((*MockOCIVirtualNetworkingClient)(nil).CreateSecurityList), arg0, arg1)
}

// CreateSubnet mocks base method.
func (m *MockOCIVirtualNetworkingClient) CreateSubnet(arg0 context.Context, arg1 core.CreateSubnetRequest) (core.CreateSubnetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubnet", arg0, arg1)
	ret0, _ := ret[0].(core.CreateSubnetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSubnet indicates an expected call of CreateSubnet.
func (mr *MockOCIVirtualNetworkingClientMockRecorder) CreateSubnet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubnet", reflect.TypeOf((*MockOCIVirtualNetworkingClient)(nil).CreateSubnet), arg0, arg1)
}

// CreateVcn mocks base method.
func (m *MockOCIVirtualNetworkingClient) CreateVcn(arg0 context.Context, arg1 core.CreateVcnRequest) (core.CreateVcnResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVcn", arg0, arg1)
	ret0, _ := ret[0].(core.CreateVcnResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVcn indicates an expected call of CreateVcn.
func (mr *MockOCIVirtualNetworkingClientMockRecorder) CreateVcn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVcn", reflect.TypeOf((*MockOCIVirtualNetworkingClient)(nil).CreateVcn), arg0, arg1)
}

// DeleteInternetGateway mocks base method.
func (m *MockOCIVirtualNetworkingClient) DeleteInternetGateway(arg0 context.Context, arg1 core.DeleteInternetGatewayRequest) (core.DeleteInternetGatewayResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInternetGateway", arg0, arg1)
	ret0, _ := ret[0].(core.DeleteInternetGatewayResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteInternetGateway indicates an expected call of DeleteInternetGateway.
func (mr *MockOCIVirtualNetworkingClientMockRecorder) DeleteInternetGateway(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInternetGateway", reflect.TypeOf((*MockOCIVirtualNetworkingClient)(nil).DeleteInternetGateway), arg0, arg1)
}

// DeleteRouteTable mocks base method.
func (m *MockOCIVirtualNetworkingClient) DeleteRouteTable(arg0 context.Context, arg1 core.DeleteRouteTableRequest) (core.DeleteRouteTableResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRouteTable", arg0, arg1)
	ret0, _ := ret[0].(core.DeleteRouteTableResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRouteTable indicates an expected call of DeleteRouteTable.
func (mr *MockOCIVirtualNetworkingClientMockRecorder) DeleteRouteTable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRouteTable", reflect.TypeOf((*MockOCIVirtualNetworkingClient)(nil).DeleteRouteTable), arg0, arg1)
}

// DeleteSecurityList mocks base method.
func (m *MockOCIVirtualNetworkingClient) DeleteSecurityList(arg0 context.Context, arg1 core.DeleteSecurityListRequest) (core.DeleteSecurityListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecurityList", arg0, arg1)
	ret0, _ := ret[0].(core.DeleteSecurityListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSecurityList indicates an expected call of DeleteSecurityList.
func (mr *MockOCIVirtualNetworkingClientMockRecorder) DeleteSecurityList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecurityList", reflect.TypeOf((*MockOCIVirtualNetworkingClient)(nil).DeleteSecurityList), arg0, arg1)
}

// DeleteSubnet mocks base method.
func (m *MockOCIVirtualNetworkingClient) DeleteSubnet(arg0 context.Context, arg1 core.DeleteSubnetRequest) (core.DeleteSubnetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSubnet", arg0, arg1)
	ret0, _ := ret[0].(core.DeleteSubnetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSubnet indicates an expected call of DeleteSubnet.
func (mr *MockOCIVirtualNetworkingClientMockRecorder) DeleteSubnet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSubnet", reflect.TypeOf((*MockOCIVirtualNetworkingClient)(nil).DeleteSubnet), arg0, arg1)
}

// DeleteVcn mocks base method.
func (m *MockOCIVirtualNetworkingClient) DeleteVcn(arg0 context.Context, arg1 core.DeleteVcnRequest) (core.DeleteVcnResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVcn", arg0, arg1)
	ret0, _ := ret[0].(core.DeleteVcnResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteVcn indicates an expected call of DeleteVcn.
func (mr *MockOCIVirtualNetworkingClientMockRecorder) DeleteVcn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVcn", reflect.TypeOf((*MockOCIVirtualNetworkingClient)(nil).DeleteVcn), arg0, arg1)
}

// GetInternetGateway mocks base method.
func (m *MockOCIVirtualNetworkingClient) GetInternetGateway(arg0 context.Context, arg1 core.GetInternetGatewayRequest) (core.GetInternetGatewayResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInternetGateway", arg0, arg1)
	ret0, _ := ret[0].(core.GetInternetGatewayResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInternetGateway indicates an expected call of GetInternetGateway.
func (mr *MockOCIVirtualNetworkingClientMockRecorder) GetInternetGateway(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInternetGateway", reflect.TypeOf((*MockOCIVirtualNetworkingClient)(nil).GetInternetGateway), arg0, arg1)
}

// GetRouteTable mocks base method.
func (m *MockOCIVirtualNetworkingClient) GetRouteTable(arg0 context.Context, arg1 core.GetRouteTableRequest) (core.GetRouteTableResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRouteTable", arg0, arg1)
	ret0, _ := ret[0].(core.GetRouteTableResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRouteTable indicates an expected call of GetRouteTable.
func (mr *MockOCIVirtualNetworkingClientMockRecorder) GetRouteTable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRouteTable", reflect.TypeOf((*MockOCIVirtualNetworkingClient)(nil).GetRouteTable), arg0, arg1)
}

// GetSecurityList mocks base method.
func (m *MockOCIVirtualNetworkingClient) GetSecurityList(arg0 context.Context, arg1 core.GetSecurityListRequest) (core.GetSecurityListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecurityList", arg0, arg1)
	ret0, _ := ret[0].(core.GetSecurityListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecurityList indicates an expected call of GetSecurityList.
func (mr *MockOCIVirtualNetworkingClientMockRecorder) GetSecurityList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecurityList", reflect.TypeOf((*MockOCIVirtualNetworkingClient)(nil).GetSecurityList), arg0, arg1)
}

// GetSubnet mocks base method.
func (m *MockOCIVirtualNetworkingClient) GetSubnet(arg0 context.Context, arg1 core.GetSubnetRequest) (core.GetSubnetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubnet", arg0, arg1)
	ret0, _ := ret[0].(core.GetSubnetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubnet indicates an expected call of GetSubnet.
func (mr *MockOCIVirtualNetworkingClientMockRecorder) GetSubnet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubnet", reflect.TypeOf((*MockOCIVirtualNetworkingClient)(nil).GetSubnet), arg0, arg1)
}

// GetVcn mocks base method.
func (m *MockOCIVirtualNetworkingClient) GetVcn(arg0 context.Context, arg1 core.GetVcnRequest) (core.GetVcnResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVcn", arg0, arg1)
	ret0, _ := ret[0].(core.GetVcnResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVcn indicates an expected call of GetVcn.
func (mr *MockOCIVirtualNetworkingClientMockRecorder) GetVcn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVcn", reflect.TypeOf((*MockOCIVirtualNetworkingClient)(nil).GetVcn), arg0, arg1)
}

// GetVnic mocks base method.
func (m *MockOCIVirtualNetworkingClient) GetVnic(arg0 context.Context, arg1 core.GetVnicRequest) (core.GetVnicResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVnic", arg0, arg1)
	ret0, _ := ret[0].(core.GetVnicResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVnic indicates an expected call of GetVnic.
func (mr *MockOCIVirtualNetworkingClientMockRecorder) GetVnic(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVnic", reflect.TypeOf((*MockOCIVirtualNetworkingClient)(nil).GetVnic), arg0, arg1)
}

// ListInternetGateways mocks base method.
func (m *MockOCIVirtualNetworkingClient) ListInternetGateways(arg0 context.Context, arg1 core.ListInternetGatewaysRequest) (core.ListInternetGatewaysResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInternetGateways", arg0, arg1)
	ret0, _ := ret[0].(core.ListInternetGatewaysResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInternetGateways indicates an expected call of ListInternetGateways.
func (mr *MockOCIVirtualNetworkingClientMockRecorder) ListInternetGateways(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInternetGateways", reflect.TypeOf((*MockOCIVirtualNetworkingClient)(nil).ListInternetGateways), arg0, arg1)
}

// ListRouteTables mocks base method.
func (m *MockOCIVirtualNetworkingClient) ListRouteTables(arg0 context.Context, arg1 core.ListRouteTablesRequest) (core.ListRouteTablesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRouteTables", arg0, arg1)
	ret0, _ := ret[0].(core.ListRouteTablesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRouteTables indicates an expected call of ListRouteTables.
func (mr *MockOCIVirtualNetworkingClientMockRecorder) ListRouteTables(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRouteTables", reflect.TypeOf((*MockOCIVirtualNetworkingClient)(nil).ListRouteTables), arg0, arg1)
}

// ListSecurityLists mocks base method.
func (m *MockOCIVirtualNetworkingClient) ListSecurityLists(arg0 context.Context, arg1 core.ListSecurityListsRequest) (core.ListSecurityListsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecurityLists", arg0, arg1)
	ret0, _ := ret[0].(core.ListSecurityListsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecurityLists indicates an expected call of ListSecurityLists.
func (mr *MockOCIVirtualNetworkingClientMockRecorder) ListSecurityLists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecurityLists", reflect.TypeOf((*MockOCIVirtualNetworkingClient)(nil).ListSecurityLists), arg0, arg1)
}

// ListSubnets mocks base method.
func (m *MockOCIVirtualNetworkingClient) ListSubnets(arg0 context.Context, arg1 core.ListSubnetsRequest) (core.ListSubnetsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSubnets", arg0, arg1)
	ret0, _ := ret[0].(core.ListSubnetsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSubnets indicates an expected call of ListSubnets.
func (mr *MockOCIVirtualNetworkingClientMockRecorder) ListSubnets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubnets", reflect.TypeOf((*MockOCIVirtualNetworkingClient)(nil).ListSubnets), arg0, arg1)
}

// ListVcns mocks base method.
func (m *MockOCIVirtualNetworkingClient) ListVcns(arg0 context.Context, arg1 core.ListVcnsRequest) (core.ListVcnsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVcns", arg0, arg1)
	ret0, _ := ret[0].(core.ListVcnsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVcns indicates an expected call of ListVcns.
func (mr *MockOCIVirtualNetworkingClientMockRecorder) ListVcns(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVcns", reflect.TypeOf((*MockOCIVirtualNetworkingClient)(nil).ListVcns), arg0, arg1)
}
