// Code generated by mockery v2.38.0. DO NOT EDIT.

package iammocks

import (
	context "context"

	access "github.com/yandex-cloud/go-genproto/yandex/cloud/access"

	iam "github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"

	mock "github.com/stretchr/testify/mock"

	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

// ServiceAccountServiceServer is an autogenerated mock type for the ServiceAccountServiceServer type
type ServiceAccountServiceServer struct {
	mock.Mock
}

type ServiceAccountServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *ServiceAccountServiceServer) EXPECT() *ServiceAccountServiceServer_Expecter {
	return &ServiceAccountServiceServer_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *ServiceAccountServiceServer) Create(_a0 context.Context, _a1 *iam.CreateServiceAccountRequest) (*operation.Operation, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *operation.Operation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iam.CreateServiceAccountRequest) (*operation.Operation, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iam.CreateServiceAccountRequest) *operation.Operation); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operation.Operation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iam.CreateServiceAccountRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceAccountServiceServer_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ServiceAccountServiceServer_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *iam.CreateServiceAccountRequest
func (_e *ServiceAccountServiceServer_Expecter) Create(_a0 interface{}, _a1 interface{}) *ServiceAccountServiceServer_Create_Call {
	return &ServiceAccountServiceServer_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *ServiceAccountServiceServer_Create_Call) Run(run func(_a0 context.Context, _a1 *iam.CreateServiceAccountRequest)) *ServiceAccountServiceServer_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*iam.CreateServiceAccountRequest))
	})
	return _c
}

func (_c *ServiceAccountServiceServer_Create_Call) Return(_a0 *operation.Operation, _a1 error) *ServiceAccountServiceServer_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ServiceAccountServiceServer_Create_Call) RunAndReturn(run func(context.Context, *iam.CreateServiceAccountRequest) (*operation.Operation, error)) *ServiceAccountServiceServer_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *ServiceAccountServiceServer) Delete(_a0 context.Context, _a1 *iam.DeleteServiceAccountRequest) (*operation.Operation, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 *operation.Operation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iam.DeleteServiceAccountRequest) (*operation.Operation, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iam.DeleteServiceAccountRequest) *operation.Operation); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operation.Operation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iam.DeleteServiceAccountRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceAccountServiceServer_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type ServiceAccountServiceServer_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *iam.DeleteServiceAccountRequest
func (_e *ServiceAccountServiceServer_Expecter) Delete(_a0 interface{}, _a1 interface{}) *ServiceAccountServiceServer_Delete_Call {
	return &ServiceAccountServiceServer_Delete_Call{Call: _e.mock.On("Delete", _a0, _a1)}
}

func (_c *ServiceAccountServiceServer_Delete_Call) Run(run func(_a0 context.Context, _a1 *iam.DeleteServiceAccountRequest)) *ServiceAccountServiceServer_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*iam.DeleteServiceAccountRequest))
	})
	return _c
}

func (_c *ServiceAccountServiceServer_Delete_Call) Return(_a0 *operation.Operation, _a1 error) *ServiceAccountServiceServer_Delete_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ServiceAccountServiceServer_Delete_Call) RunAndReturn(run func(context.Context, *iam.DeleteServiceAccountRequest) (*operation.Operation, error)) *ServiceAccountServiceServer_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *ServiceAccountServiceServer) Get(_a0 context.Context, _a1 *iam.GetServiceAccountRequest) (*iam.ServiceAccount, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *iam.ServiceAccount
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iam.GetServiceAccountRequest) (*iam.ServiceAccount, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iam.GetServiceAccountRequest) *iam.ServiceAccount); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.ServiceAccount)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iam.GetServiceAccountRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceAccountServiceServer_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type ServiceAccountServiceServer_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *iam.GetServiceAccountRequest
func (_e *ServiceAccountServiceServer_Expecter) Get(_a0 interface{}, _a1 interface{}) *ServiceAccountServiceServer_Get_Call {
	return &ServiceAccountServiceServer_Get_Call{Call: _e.mock.On("Get", _a0, _a1)}
}

func (_c *ServiceAccountServiceServer_Get_Call) Run(run func(_a0 context.Context, _a1 *iam.GetServiceAccountRequest)) *ServiceAccountServiceServer_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*iam.GetServiceAccountRequest))
	})
	return _c
}

func (_c *ServiceAccountServiceServer_Get_Call) Return(_a0 *iam.ServiceAccount, _a1 error) *ServiceAccountServiceServer_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ServiceAccountServiceServer_Get_Call) RunAndReturn(run func(context.Context, *iam.GetServiceAccountRequest) (*iam.ServiceAccount, error)) *ServiceAccountServiceServer_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: _a0, _a1
func (_m *ServiceAccountServiceServer) List(_a0 context.Context, _a1 *iam.ListServiceAccountsRequest) (*iam.ListServiceAccountsResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *iam.ListServiceAccountsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iam.ListServiceAccountsRequest) (*iam.ListServiceAccountsResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iam.ListServiceAccountsRequest) *iam.ListServiceAccountsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.ListServiceAccountsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iam.ListServiceAccountsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceAccountServiceServer_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type ServiceAccountServiceServer_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *iam.ListServiceAccountsRequest
func (_e *ServiceAccountServiceServer_Expecter) List(_a0 interface{}, _a1 interface{}) *ServiceAccountServiceServer_List_Call {
	return &ServiceAccountServiceServer_List_Call{Call: _e.mock.On("List", _a0, _a1)}
}

func (_c *ServiceAccountServiceServer_List_Call) Run(run func(_a0 context.Context, _a1 *iam.ListServiceAccountsRequest)) *ServiceAccountServiceServer_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*iam.ListServiceAccountsRequest))
	})
	return _c
}

func (_c *ServiceAccountServiceServer_List_Call) Return(_a0 *iam.ListServiceAccountsResponse, _a1 error) *ServiceAccountServiceServer_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ServiceAccountServiceServer_List_Call) RunAndReturn(run func(context.Context, *iam.ListServiceAccountsRequest) (*iam.ListServiceAccountsResponse, error)) *ServiceAccountServiceServer_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAccessBindings provides a mock function with given fields: _a0, _a1
func (_m *ServiceAccountServiceServer) ListAccessBindings(_a0 context.Context, _a1 *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for ListAccessBindings")
	}

	var r0 *access.ListAccessBindingsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *access.ListAccessBindingsRequest) *access.ListAccessBindingsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*access.ListAccessBindingsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *access.ListAccessBindingsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceAccountServiceServer_ListAccessBindings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAccessBindings'
type ServiceAccountServiceServer_ListAccessBindings_Call struct {
	*mock.Call
}

// ListAccessBindings is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *access.ListAccessBindingsRequest
func (_e *ServiceAccountServiceServer_Expecter) ListAccessBindings(_a0 interface{}, _a1 interface{}) *ServiceAccountServiceServer_ListAccessBindings_Call {
	return &ServiceAccountServiceServer_ListAccessBindings_Call{Call: _e.mock.On("ListAccessBindings", _a0, _a1)}
}

func (_c *ServiceAccountServiceServer_ListAccessBindings_Call) Run(run func(_a0 context.Context, _a1 *access.ListAccessBindingsRequest)) *ServiceAccountServiceServer_ListAccessBindings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*access.ListAccessBindingsRequest))
	})
	return _c
}

func (_c *ServiceAccountServiceServer_ListAccessBindings_Call) Return(_a0 *access.ListAccessBindingsResponse, _a1 error) *ServiceAccountServiceServer_ListAccessBindings_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ServiceAccountServiceServer_ListAccessBindings_Call) RunAndReturn(run func(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error)) *ServiceAccountServiceServer_ListAccessBindings_Call {
	_c.Call.Return(run)
	return _c
}

// ListOperations provides a mock function with given fields: _a0, _a1
func (_m *ServiceAccountServiceServer) ListOperations(_a0 context.Context, _a1 *iam.ListServiceAccountOperationsRequest) (*iam.ListServiceAccountOperationsResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for ListOperations")
	}

	var r0 *iam.ListServiceAccountOperationsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iam.ListServiceAccountOperationsRequest) (*iam.ListServiceAccountOperationsResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iam.ListServiceAccountOperationsRequest) *iam.ListServiceAccountOperationsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.ListServiceAccountOperationsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iam.ListServiceAccountOperationsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceAccountServiceServer_ListOperations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListOperations'
type ServiceAccountServiceServer_ListOperations_Call struct {
	*mock.Call
}

// ListOperations is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *iam.ListServiceAccountOperationsRequest
func (_e *ServiceAccountServiceServer_Expecter) ListOperations(_a0 interface{}, _a1 interface{}) *ServiceAccountServiceServer_ListOperations_Call {
	return &ServiceAccountServiceServer_ListOperations_Call{Call: _e.mock.On("ListOperations", _a0, _a1)}
}

func (_c *ServiceAccountServiceServer_ListOperations_Call) Run(run func(_a0 context.Context, _a1 *iam.ListServiceAccountOperationsRequest)) *ServiceAccountServiceServer_ListOperations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*iam.ListServiceAccountOperationsRequest))
	})
	return _c
}

func (_c *ServiceAccountServiceServer_ListOperations_Call) Return(_a0 *iam.ListServiceAccountOperationsResponse, _a1 error) *ServiceAccountServiceServer_ListOperations_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ServiceAccountServiceServer_ListOperations_Call) RunAndReturn(run func(context.Context, *iam.ListServiceAccountOperationsRequest) (*iam.ListServiceAccountOperationsResponse, error)) *ServiceAccountServiceServer_ListOperations_Call {
	_c.Call.Return(run)
	return _c
}

// SetAccessBindings provides a mock function with given fields: _a0, _a1
func (_m *ServiceAccountServiceServer) SetAccessBindings(_a0 context.Context, _a1 *access.SetAccessBindingsRequest) (*operation.Operation, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for SetAccessBindings")
	}

	var r0 *operation.Operation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *access.SetAccessBindingsRequest) *operation.Operation); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operation.Operation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *access.SetAccessBindingsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceAccountServiceServer_SetAccessBindings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetAccessBindings'
type ServiceAccountServiceServer_SetAccessBindings_Call struct {
	*mock.Call
}

// SetAccessBindings is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *access.SetAccessBindingsRequest
func (_e *ServiceAccountServiceServer_Expecter) SetAccessBindings(_a0 interface{}, _a1 interface{}) *ServiceAccountServiceServer_SetAccessBindings_Call {
	return &ServiceAccountServiceServer_SetAccessBindings_Call{Call: _e.mock.On("SetAccessBindings", _a0, _a1)}
}

func (_c *ServiceAccountServiceServer_SetAccessBindings_Call) Run(run func(_a0 context.Context, _a1 *access.SetAccessBindingsRequest)) *ServiceAccountServiceServer_SetAccessBindings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*access.SetAccessBindingsRequest))
	})
	return _c
}

func (_c *ServiceAccountServiceServer_SetAccessBindings_Call) Return(_a0 *operation.Operation, _a1 error) *ServiceAccountServiceServer_SetAccessBindings_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ServiceAccountServiceServer_SetAccessBindings_Call) RunAndReturn(run func(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error)) *ServiceAccountServiceServer_SetAccessBindings_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *ServiceAccountServiceServer) Update(_a0 context.Context, _a1 *iam.UpdateServiceAccountRequest) (*operation.Operation, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *operation.Operation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iam.UpdateServiceAccountRequest) (*operation.Operation, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iam.UpdateServiceAccountRequest) *operation.Operation); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operation.Operation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iam.UpdateServiceAccountRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceAccountServiceServer_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type ServiceAccountServiceServer_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *iam.UpdateServiceAccountRequest
func (_e *ServiceAccountServiceServer_Expecter) Update(_a0 interface{}, _a1 interface{}) *ServiceAccountServiceServer_Update_Call {
	return &ServiceAccountServiceServer_Update_Call{Call: _e.mock.On("Update", _a0, _a1)}
}

func (_c *ServiceAccountServiceServer_Update_Call) Run(run func(_a0 context.Context, _a1 *iam.UpdateServiceAccountRequest)) *ServiceAccountServiceServer_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*iam.UpdateServiceAccountRequest))
	})
	return _c
}

func (_c *ServiceAccountServiceServer_Update_Call) Return(_a0 *operation.Operation, _a1 error) *ServiceAccountServiceServer_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ServiceAccountServiceServer_Update_Call) RunAndReturn(run func(context.Context, *iam.UpdateServiceAccountRequest) (*operation.Operation, error)) *ServiceAccountServiceServer_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateAccessBindings provides a mock function with given fields: _a0, _a1
func (_m *ServiceAccountServiceServer) UpdateAccessBindings(_a0 context.Context, _a1 *access.UpdateAccessBindingsRequest) (*operation.Operation, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAccessBindings")
	}

	var r0 *operation.Operation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *access.UpdateAccessBindingsRequest) *operation.Operation); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operation.Operation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *access.UpdateAccessBindingsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceAccountServiceServer_UpdateAccessBindings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateAccessBindings'
type ServiceAccountServiceServer_UpdateAccessBindings_Call struct {
	*mock.Call
}

// UpdateAccessBindings is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *access.UpdateAccessBindingsRequest
func (_e *ServiceAccountServiceServer_Expecter) UpdateAccessBindings(_a0 interface{}, _a1 interface{}) *ServiceAccountServiceServer_UpdateAccessBindings_Call {
	return &ServiceAccountServiceServer_UpdateAccessBindings_Call{Call: _e.mock.On("UpdateAccessBindings", _a0, _a1)}
}

func (_c *ServiceAccountServiceServer_UpdateAccessBindings_Call) Run(run func(_a0 context.Context, _a1 *access.UpdateAccessBindingsRequest)) *ServiceAccountServiceServer_UpdateAccessBindings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*access.UpdateAccessBindingsRequest))
	})
	return _c
}

func (_c *ServiceAccountServiceServer_UpdateAccessBindings_Call) Return(_a0 *operation.Operation, _a1 error) *ServiceAccountServiceServer_UpdateAccessBindings_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ServiceAccountServiceServer_UpdateAccessBindings_Call) RunAndReturn(run func(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error)) *ServiceAccountServiceServer_UpdateAccessBindings_Call {
	_c.Call.Return(run)
	return _c
}

// NewServiceAccountServiceServer creates a new instance of ServiceAccountServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceAccountServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceAccountServiceServer {
	mock := &ServiceAccountServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}