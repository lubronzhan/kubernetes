// +build !providerless

/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mockdeploymentclient

import (
	context "context"
	reflect "reflect"

	resources "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2017-05-10/resources"
	gomock "github.com/golang/mock/gomock"
	retry "k8s.io/legacy-cloud-providers/azure/retry"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockInterface) Get(ctx context.Context, resourceGroupName, deploymentName string) (resources.DeploymentExtended, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, resourceGroupName, deploymentName)
	ret0, _ := ret[0].(resources.DeploymentExtended)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockInterfaceMockRecorder) Get(ctx, resourceGroupName, deploymentName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterface)(nil).Get), ctx, resourceGroupName, deploymentName)
}

// List mocks base method
func (m *MockInterface) List(ctx context.Context, resourceGroupName string) ([]resources.DeploymentExtended, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, resourceGroupName)
	ret0, _ := ret[0].([]resources.DeploymentExtended)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockInterfaceMockRecorder) List(ctx, resourceGroupName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockInterface)(nil).List), ctx, resourceGroupName)
}

// ExportTemplate mocks base method
func (m *MockInterface) ExportTemplate(ctx context.Context, resourceGroupName, deploymentName string) (resources.DeploymentExportResult, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportTemplate", ctx, resourceGroupName, deploymentName)
	ret0, _ := ret[0].(resources.DeploymentExportResult)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// ExportTemplate indicates an expected call of ExportTemplate
func (mr *MockInterfaceMockRecorder) ExportTemplate(ctx, resourceGroupName, deploymentName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportTemplate", reflect.TypeOf((*MockInterface)(nil).ExportTemplate), ctx, resourceGroupName, deploymentName)
}

// CreateOrUpdate mocks base method
func (m *MockInterface) CreateOrUpdate(ctx context.Context, resourceGroupName, managedClusterName string, parameters resources.Deployment, etag string) *retry.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", ctx, resourceGroupName, managedClusterName, parameters, etag)
	ret0, _ := ret[0].(*retry.Error)
	return ret0
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate
func (mr *MockInterfaceMockRecorder) CreateOrUpdate(ctx, resourceGroupName, managedClusterName, parameters, etag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockInterface)(nil).CreateOrUpdate), ctx, resourceGroupName, managedClusterName, parameters, etag)
}

// Delete mocks base method
func (m *MockInterface) Delete(ctx context.Context, resourceGroupName, deploymentName string) *retry.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, resourceGroupName, deploymentName)
	ret0, _ := ret[0].(*retry.Error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockInterfaceMockRecorder) Delete(ctx, resourceGroupName, deploymentName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockInterface)(nil).Delete), ctx, resourceGroupName, deploymentName)
}
