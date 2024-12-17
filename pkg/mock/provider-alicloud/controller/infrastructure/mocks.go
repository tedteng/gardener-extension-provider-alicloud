// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener-extension-provider-alicloud/pkg/controller/infrastructure (interfaces: TerraformChartOps)
//
// Generated by this command:
//
//	mockgen -package=infrastructure -destination=mocks.go github.com/gardener/gardener-extension-provider-alicloud/pkg/controller/infrastructure TerraformChartOps
//

// Package infrastructure is a generated GoMock package.
package infrastructure

import (
	reflect "reflect"

	client "github.com/gardener/gardener-extension-provider-alicloud/pkg/alicloud/client"
	v1alpha1 "github.com/gardener/gardener-extension-provider-alicloud/pkg/apis/alicloud/v1alpha1"
	infrastructure "github.com/gardener/gardener-extension-provider-alicloud/pkg/controller/infrastructure"
	v1alpha10 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	gomock "go.uber.org/mock/gomock"
)

// MockTerraformChartOps is a mock of TerraformChartOps interface.
type MockTerraformChartOps struct {
	ctrl     *gomock.Controller
	recorder *MockTerraformChartOpsMockRecorder
	isgomock struct{}
}

// MockTerraformChartOpsMockRecorder is the mock recorder for MockTerraformChartOps.
type MockTerraformChartOpsMockRecorder struct {
	mock *MockTerraformChartOps
}

// NewMockTerraformChartOps creates a new mock instance.
func NewMockTerraformChartOps(ctrl *gomock.Controller) *MockTerraformChartOps {
	mock := &MockTerraformChartOps{ctrl: ctrl}
	mock.recorder = &MockTerraformChartOpsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTerraformChartOps) EXPECT() *MockTerraformChartOpsMockRecorder {
	return m.recorder
}

// ComputeChartValues mocks base method.
func (m *MockTerraformChartOps) ComputeChartValues(infra *v1alpha10.Infrastructure, config *v1alpha1.InfrastructureConfig, podCIDR *string, values *infrastructure.InitializerValues) map[string]any {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputeChartValues", infra, config, podCIDR, values)
	ret0, _ := ret[0].(map[string]any)
	return ret0
}

// ComputeChartValues indicates an expected call of ComputeChartValues.
func (mr *MockTerraformChartOpsMockRecorder) ComputeChartValues(infra, config, podCIDR, values any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeChartValues", reflect.TypeOf((*MockTerraformChartOps)(nil).ComputeChartValues), infra, config, podCIDR, values)
}

// ComputeCreateVPCInitializerValues mocks base method.
func (m *MockTerraformChartOps) ComputeCreateVPCInitializerValues(config *v1alpha1.InfrastructureConfig, internetChargeType string) *infrastructure.InitializerValues {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputeCreateVPCInitializerValues", config, internetChargeType)
	ret0, _ := ret[0].(*infrastructure.InitializerValues)
	return ret0
}

// ComputeCreateVPCInitializerValues indicates an expected call of ComputeCreateVPCInitializerValues.
func (mr *MockTerraformChartOpsMockRecorder) ComputeCreateVPCInitializerValues(config, internetChargeType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeCreateVPCInitializerValues", reflect.TypeOf((*MockTerraformChartOps)(nil).ComputeCreateVPCInitializerValues), config, internetChargeType)
}

// ComputeUseVPCInitializerValues mocks base method.
func (m *MockTerraformChartOps) ComputeUseVPCInitializerValues(config *v1alpha1.InfrastructureConfig, info *client.VPCInfo) *infrastructure.InitializerValues {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputeUseVPCInitializerValues", config, info)
	ret0, _ := ret[0].(*infrastructure.InitializerValues)
	return ret0
}

// ComputeUseVPCInitializerValues indicates an expected call of ComputeUseVPCInitializerValues.
func (mr *MockTerraformChartOpsMockRecorder) ComputeUseVPCInitializerValues(config, info any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeUseVPCInitializerValues", reflect.TypeOf((*MockTerraformChartOps)(nil).ComputeUseVPCInitializerValues), config, info)
}
