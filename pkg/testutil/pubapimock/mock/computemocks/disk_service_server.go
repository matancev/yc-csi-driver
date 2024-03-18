// Code generated by mockery v2.38.0. DO NOT EDIT.

package computemocks

import (
	access "github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	compute "github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"

	context "context"

	mock "github.com/stretchr/testify/mock"

	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

// DiskServiceServer is an autogenerated mock type for the DiskServiceServer type
type DiskServiceServer struct {
	mock.Mock
}

type DiskServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *DiskServiceServer) EXPECT() *DiskServiceServer_Expecter {
	return &DiskServiceServer_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *DiskServiceServer) Create(_a0 context.Context, _a1 *compute.CreateDiskRequest) (*operation.Operation, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *operation.Operation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *compute.CreateDiskRequest) (*operation.Operation, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *compute.CreateDiskRequest) *operation.Operation); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operation.Operation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *compute.CreateDiskRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DiskServiceServer_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type DiskServiceServer_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *compute.CreateDiskRequest
func (_e *DiskServiceServer_Expecter) Create(_a0 interface{}, _a1 interface{}) *DiskServiceServer_Create_Call {
	return &DiskServiceServer_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *DiskServiceServer_Create_Call) Run(run func(_a0 context.Context, _a1 *compute.CreateDiskRequest)) *DiskServiceServer_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*compute.CreateDiskRequest))
	})
	return _c
}

func (_c *DiskServiceServer_Create_Call) Return(_a0 *operation.Operation, _a1 error) *DiskServiceServer_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DiskServiceServer_Create_Call) RunAndReturn(run func(context.Context, *compute.CreateDiskRequest) (*operation.Operation, error)) *DiskServiceServer_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *DiskServiceServer) Delete(_a0 context.Context, _a1 *compute.DeleteDiskRequest) (*operation.Operation, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 *operation.Operation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *compute.DeleteDiskRequest) (*operation.Operation, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *compute.DeleteDiskRequest) *operation.Operation); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operation.Operation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *compute.DeleteDiskRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DiskServiceServer_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type DiskServiceServer_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *compute.DeleteDiskRequest
func (_e *DiskServiceServer_Expecter) Delete(_a0 interface{}, _a1 interface{}) *DiskServiceServer_Delete_Call {
	return &DiskServiceServer_Delete_Call{Call: _e.mock.On("Delete", _a0, _a1)}
}

func (_c *DiskServiceServer_Delete_Call) Run(run func(_a0 context.Context, _a1 *compute.DeleteDiskRequest)) *DiskServiceServer_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*compute.DeleteDiskRequest))
	})
	return _c
}

func (_c *DiskServiceServer_Delete_Call) Return(_a0 *operation.Operation, _a1 error) *DiskServiceServer_Delete_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DiskServiceServer_Delete_Call) RunAndReturn(run func(context.Context, *compute.DeleteDiskRequest) (*operation.Operation, error)) *DiskServiceServer_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *DiskServiceServer) Get(_a0 context.Context, _a1 *compute.GetDiskRequest) (*compute.Disk, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *compute.Disk
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *compute.GetDiskRequest) (*compute.Disk, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *compute.GetDiskRequest) *compute.Disk); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.Disk)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *compute.GetDiskRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DiskServiceServer_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type DiskServiceServer_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *compute.GetDiskRequest
func (_e *DiskServiceServer_Expecter) Get(_a0 interface{}, _a1 interface{}) *DiskServiceServer_Get_Call {
	return &DiskServiceServer_Get_Call{Call: _e.mock.On("Get", _a0, _a1)}
}

func (_c *DiskServiceServer_Get_Call) Run(run func(_a0 context.Context, _a1 *compute.GetDiskRequest)) *DiskServiceServer_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*compute.GetDiskRequest))
	})
	return _c
}

func (_c *DiskServiceServer_Get_Call) Return(_a0 *compute.Disk, _a1 error) *DiskServiceServer_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DiskServiceServer_Get_Call) RunAndReturn(run func(context.Context, *compute.GetDiskRequest) (*compute.Disk, error)) *DiskServiceServer_Get_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: _a0, _a1
func (_m *DiskServiceServer) List(_a0 context.Context, _a1 *compute.ListDisksRequest) (*compute.ListDisksResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *compute.ListDisksResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *compute.ListDisksRequest) (*compute.ListDisksResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *compute.ListDisksRequest) *compute.ListDisksResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.ListDisksResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *compute.ListDisksRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DiskServiceServer_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type DiskServiceServer_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *compute.ListDisksRequest
func (_e *DiskServiceServer_Expecter) List(_a0 interface{}, _a1 interface{}) *DiskServiceServer_List_Call {
	return &DiskServiceServer_List_Call{Call: _e.mock.On("List", _a0, _a1)}
}

func (_c *DiskServiceServer_List_Call) Run(run func(_a0 context.Context, _a1 *compute.ListDisksRequest)) *DiskServiceServer_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*compute.ListDisksRequest))
	})
	return _c
}

func (_c *DiskServiceServer_List_Call) Return(_a0 *compute.ListDisksResponse, _a1 error) *DiskServiceServer_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DiskServiceServer_List_Call) RunAndReturn(run func(context.Context, *compute.ListDisksRequest) (*compute.ListDisksResponse, error)) *DiskServiceServer_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAccessBindings provides a mock function with given fields: _a0, _a1
func (_m *DiskServiceServer) ListAccessBindings(_a0 context.Context, _a1 *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
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

// DiskServiceServer_ListAccessBindings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAccessBindings'
type DiskServiceServer_ListAccessBindings_Call struct {
	*mock.Call
}

// ListAccessBindings is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *access.ListAccessBindingsRequest
func (_e *DiskServiceServer_Expecter) ListAccessBindings(_a0 interface{}, _a1 interface{}) *DiskServiceServer_ListAccessBindings_Call {
	return &DiskServiceServer_ListAccessBindings_Call{Call: _e.mock.On("ListAccessBindings", _a0, _a1)}
}

func (_c *DiskServiceServer_ListAccessBindings_Call) Run(run func(_a0 context.Context, _a1 *access.ListAccessBindingsRequest)) *DiskServiceServer_ListAccessBindings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*access.ListAccessBindingsRequest))
	})
	return _c
}

func (_c *DiskServiceServer_ListAccessBindings_Call) Return(_a0 *access.ListAccessBindingsResponse, _a1 error) *DiskServiceServer_ListAccessBindings_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DiskServiceServer_ListAccessBindings_Call) RunAndReturn(run func(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error)) *DiskServiceServer_ListAccessBindings_Call {
	_c.Call.Return(run)
	return _c
}

// ListOperations provides a mock function with given fields: _a0, _a1
func (_m *DiskServiceServer) ListOperations(_a0 context.Context, _a1 *compute.ListDiskOperationsRequest) (*compute.ListDiskOperationsResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for ListOperations")
	}

	var r0 *compute.ListDiskOperationsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *compute.ListDiskOperationsRequest) (*compute.ListDiskOperationsResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *compute.ListDiskOperationsRequest) *compute.ListDiskOperationsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.ListDiskOperationsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *compute.ListDiskOperationsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DiskServiceServer_ListOperations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListOperations'
type DiskServiceServer_ListOperations_Call struct {
	*mock.Call
}

// ListOperations is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *compute.ListDiskOperationsRequest
func (_e *DiskServiceServer_Expecter) ListOperations(_a0 interface{}, _a1 interface{}) *DiskServiceServer_ListOperations_Call {
	return &DiskServiceServer_ListOperations_Call{Call: _e.mock.On("ListOperations", _a0, _a1)}
}

func (_c *DiskServiceServer_ListOperations_Call) Run(run func(_a0 context.Context, _a1 *compute.ListDiskOperationsRequest)) *DiskServiceServer_ListOperations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*compute.ListDiskOperationsRequest))
	})
	return _c
}

func (_c *DiskServiceServer_ListOperations_Call) Return(_a0 *compute.ListDiskOperationsResponse, _a1 error) *DiskServiceServer_ListOperations_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DiskServiceServer_ListOperations_Call) RunAndReturn(run func(context.Context, *compute.ListDiskOperationsRequest) (*compute.ListDiskOperationsResponse, error)) *DiskServiceServer_ListOperations_Call {
	_c.Call.Return(run)
	return _c
}

// ListSnapshotSchedules provides a mock function with given fields: _a0, _a1
func (_m *DiskServiceServer) ListSnapshotSchedules(_a0 context.Context, _a1 *compute.ListDiskSnapshotSchedulesRequest) (*compute.ListDiskSnapshotSchedulesResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for ListSnapshotSchedules")
	}

	var r0 *compute.ListDiskSnapshotSchedulesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *compute.ListDiskSnapshotSchedulesRequest) (*compute.ListDiskSnapshotSchedulesResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *compute.ListDiskSnapshotSchedulesRequest) *compute.ListDiskSnapshotSchedulesResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compute.ListDiskSnapshotSchedulesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *compute.ListDiskSnapshotSchedulesRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DiskServiceServer_ListSnapshotSchedules_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSnapshotSchedules'
type DiskServiceServer_ListSnapshotSchedules_Call struct {
	*mock.Call
}

// ListSnapshotSchedules is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *compute.ListDiskSnapshotSchedulesRequest
func (_e *DiskServiceServer_Expecter) ListSnapshotSchedules(_a0 interface{}, _a1 interface{}) *DiskServiceServer_ListSnapshotSchedules_Call {
	return &DiskServiceServer_ListSnapshotSchedules_Call{Call: _e.mock.On("ListSnapshotSchedules", _a0, _a1)}
}

func (_c *DiskServiceServer_ListSnapshotSchedules_Call) Run(run func(_a0 context.Context, _a1 *compute.ListDiskSnapshotSchedulesRequest)) *DiskServiceServer_ListSnapshotSchedules_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*compute.ListDiskSnapshotSchedulesRequest))
	})
	return _c
}

func (_c *DiskServiceServer_ListSnapshotSchedules_Call) Return(_a0 *compute.ListDiskSnapshotSchedulesResponse, _a1 error) *DiskServiceServer_ListSnapshotSchedules_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DiskServiceServer_ListSnapshotSchedules_Call) RunAndReturn(run func(context.Context, *compute.ListDiskSnapshotSchedulesRequest) (*compute.ListDiskSnapshotSchedulesResponse, error)) *DiskServiceServer_ListSnapshotSchedules_Call {
	_c.Call.Return(run)
	return _c
}

// Move provides a mock function with given fields: _a0, _a1
func (_m *DiskServiceServer) Move(_a0 context.Context, _a1 *compute.MoveDiskRequest) (*operation.Operation, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Move")
	}

	var r0 *operation.Operation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *compute.MoveDiskRequest) (*operation.Operation, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *compute.MoveDiskRequest) *operation.Operation); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operation.Operation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *compute.MoveDiskRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DiskServiceServer_Move_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Move'
type DiskServiceServer_Move_Call struct {
	*mock.Call
}

// Move is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *compute.MoveDiskRequest
func (_e *DiskServiceServer_Expecter) Move(_a0 interface{}, _a1 interface{}) *DiskServiceServer_Move_Call {
	return &DiskServiceServer_Move_Call{Call: _e.mock.On("Move", _a0, _a1)}
}

func (_c *DiskServiceServer_Move_Call) Run(run func(_a0 context.Context, _a1 *compute.MoveDiskRequest)) *DiskServiceServer_Move_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*compute.MoveDiskRequest))
	})
	return _c
}

func (_c *DiskServiceServer_Move_Call) Return(_a0 *operation.Operation, _a1 error) *DiskServiceServer_Move_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DiskServiceServer_Move_Call) RunAndReturn(run func(context.Context, *compute.MoveDiskRequest) (*operation.Operation, error)) *DiskServiceServer_Move_Call {
	_c.Call.Return(run)
	return _c
}

// Relocate provides a mock function with given fields: _a0, _a1
func (_m *DiskServiceServer) Relocate(_a0 context.Context, _a1 *compute.RelocateDiskRequest) (*operation.Operation, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Relocate")
	}

	var r0 *operation.Operation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *compute.RelocateDiskRequest) (*operation.Operation, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *compute.RelocateDiskRequest) *operation.Operation); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operation.Operation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *compute.RelocateDiskRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DiskServiceServer_Relocate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Relocate'
type DiskServiceServer_Relocate_Call struct {
	*mock.Call
}

// Relocate is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *compute.RelocateDiskRequest
func (_e *DiskServiceServer_Expecter) Relocate(_a0 interface{}, _a1 interface{}) *DiskServiceServer_Relocate_Call {
	return &DiskServiceServer_Relocate_Call{Call: _e.mock.On("Relocate", _a0, _a1)}
}

func (_c *DiskServiceServer_Relocate_Call) Run(run func(_a0 context.Context, _a1 *compute.RelocateDiskRequest)) *DiskServiceServer_Relocate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*compute.RelocateDiskRequest))
	})
	return _c
}

func (_c *DiskServiceServer_Relocate_Call) Return(_a0 *operation.Operation, _a1 error) *DiskServiceServer_Relocate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DiskServiceServer_Relocate_Call) RunAndReturn(run func(context.Context, *compute.RelocateDiskRequest) (*operation.Operation, error)) *DiskServiceServer_Relocate_Call {
	_c.Call.Return(run)
	return _c
}

// SetAccessBindings provides a mock function with given fields: _a0, _a1
func (_m *DiskServiceServer) SetAccessBindings(_a0 context.Context, _a1 *access.SetAccessBindingsRequest) (*operation.Operation, error) {
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

// DiskServiceServer_SetAccessBindings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetAccessBindings'
type DiskServiceServer_SetAccessBindings_Call struct {
	*mock.Call
}

// SetAccessBindings is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *access.SetAccessBindingsRequest
func (_e *DiskServiceServer_Expecter) SetAccessBindings(_a0 interface{}, _a1 interface{}) *DiskServiceServer_SetAccessBindings_Call {
	return &DiskServiceServer_SetAccessBindings_Call{Call: _e.mock.On("SetAccessBindings", _a0, _a1)}
}

func (_c *DiskServiceServer_SetAccessBindings_Call) Run(run func(_a0 context.Context, _a1 *access.SetAccessBindingsRequest)) *DiskServiceServer_SetAccessBindings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*access.SetAccessBindingsRequest))
	})
	return _c
}

func (_c *DiskServiceServer_SetAccessBindings_Call) Return(_a0 *operation.Operation, _a1 error) *DiskServiceServer_SetAccessBindings_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DiskServiceServer_SetAccessBindings_Call) RunAndReturn(run func(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error)) *DiskServiceServer_SetAccessBindings_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *DiskServiceServer) Update(_a0 context.Context, _a1 *compute.UpdateDiskRequest) (*operation.Operation, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *operation.Operation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *compute.UpdateDiskRequest) (*operation.Operation, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *compute.UpdateDiskRequest) *operation.Operation); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operation.Operation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *compute.UpdateDiskRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DiskServiceServer_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type DiskServiceServer_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *compute.UpdateDiskRequest
func (_e *DiskServiceServer_Expecter) Update(_a0 interface{}, _a1 interface{}) *DiskServiceServer_Update_Call {
	return &DiskServiceServer_Update_Call{Call: _e.mock.On("Update", _a0, _a1)}
}

func (_c *DiskServiceServer_Update_Call) Run(run func(_a0 context.Context, _a1 *compute.UpdateDiskRequest)) *DiskServiceServer_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*compute.UpdateDiskRequest))
	})
	return _c
}

func (_c *DiskServiceServer_Update_Call) Return(_a0 *operation.Operation, _a1 error) *DiskServiceServer_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DiskServiceServer_Update_Call) RunAndReturn(run func(context.Context, *compute.UpdateDiskRequest) (*operation.Operation, error)) *DiskServiceServer_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateAccessBindings provides a mock function with given fields: _a0, _a1
func (_m *DiskServiceServer) UpdateAccessBindings(_a0 context.Context, _a1 *access.UpdateAccessBindingsRequest) (*operation.Operation, error) {
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

// DiskServiceServer_UpdateAccessBindings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateAccessBindings'
type DiskServiceServer_UpdateAccessBindings_Call struct {
	*mock.Call
}

// UpdateAccessBindings is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *access.UpdateAccessBindingsRequest
func (_e *DiskServiceServer_Expecter) UpdateAccessBindings(_a0 interface{}, _a1 interface{}) *DiskServiceServer_UpdateAccessBindings_Call {
	return &DiskServiceServer_UpdateAccessBindings_Call{Call: _e.mock.On("UpdateAccessBindings", _a0, _a1)}
}

func (_c *DiskServiceServer_UpdateAccessBindings_Call) Run(run func(_a0 context.Context, _a1 *access.UpdateAccessBindingsRequest)) *DiskServiceServer_UpdateAccessBindings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*access.UpdateAccessBindingsRequest))
	})
	return _c
}

func (_c *DiskServiceServer_UpdateAccessBindings_Call) Return(_a0 *operation.Operation, _a1 error) *DiskServiceServer_UpdateAccessBindings_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DiskServiceServer_UpdateAccessBindings_Call) RunAndReturn(run func(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error)) *DiskServiceServer_UpdateAccessBindings_Call {
	_c.Call.Return(run)
	return _c
}

// NewDiskServiceServer creates a new instance of DiskServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDiskServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *DiskServiceServer {
	mock := &DiskServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}