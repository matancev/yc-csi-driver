// Code generated by sdkgen. DO NOT EDIT.

// nolint
package resourcemanager

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	resourcemanager "github.com/yandex-cloud/go-genproto/yandex/cloud/resourcemanager/v1"
)

//revive:disable

// FolderServiceClient is a resourcemanager.FolderServiceClient with
// lazy GRPC connection initialization.
type FolderServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements resourcemanager.FolderServiceClient
func (c *FolderServiceClient) Create(ctx context.Context, in *resourcemanager.CreateFolderRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return resourcemanager.NewFolderServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements resourcemanager.FolderServiceClient
func (c *FolderServiceClient) Delete(ctx context.Context, in *resourcemanager.DeleteFolderRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return resourcemanager.NewFolderServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements resourcemanager.FolderServiceClient
func (c *FolderServiceClient) Get(ctx context.Context, in *resourcemanager.GetFolderRequest, opts ...grpc.CallOption) (*resourcemanager.Folder, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return resourcemanager.NewFolderServiceClient(conn).Get(ctx, in, opts...)
}

// List implements resourcemanager.FolderServiceClient
func (c *FolderServiceClient) List(ctx context.Context, in *resourcemanager.ListFoldersRequest, opts ...grpc.CallOption) (*resourcemanager.ListFoldersResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return resourcemanager.NewFolderServiceClient(conn).List(ctx, in, opts...)
}

type FolderIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *FolderServiceClient
	request *resourcemanager.ListFoldersRequest

	items []*resourcemanager.Folder
}

func (c *FolderServiceClient) FolderIterator(ctx context.Context, req *resourcemanager.ListFoldersRequest, opts ...grpc.CallOption) *FolderIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &FolderIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *FolderIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Folders
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *FolderIterator) Take(size int64) ([]*resourcemanager.Folder, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*resourcemanager.Folder

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *FolderIterator) TakeAll() ([]*resourcemanager.Folder, error) {
	return it.Take(0)
}

func (it *FolderIterator) Value() *resourcemanager.Folder {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *FolderIterator) Error() error {
	return it.err
}

// ListAccessBindings implements resourcemanager.FolderServiceClient
func (c *FolderServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return resourcemanager.NewFolderServiceClient(conn).ListAccessBindings(ctx, in, opts...)
}

type FolderAccessBindingsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *FolderServiceClient
	request *access.ListAccessBindingsRequest

	items []*access.AccessBinding
}

func (c *FolderServiceClient) FolderAccessBindingsIterator(ctx context.Context, req *access.ListAccessBindingsRequest, opts ...grpc.CallOption) *FolderAccessBindingsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &FolderAccessBindingsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *FolderAccessBindingsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.ListAccessBindings(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.AccessBindings
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *FolderAccessBindingsIterator) Take(size int64) ([]*access.AccessBinding, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*access.AccessBinding

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *FolderAccessBindingsIterator) TakeAll() ([]*access.AccessBinding, error) {
	return it.Take(0)
}

func (it *FolderAccessBindingsIterator) Value() *access.AccessBinding {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *FolderAccessBindingsIterator) Error() error {
	return it.err
}

// ListOperations implements resourcemanager.FolderServiceClient
func (c *FolderServiceClient) ListOperations(ctx context.Context, in *resourcemanager.ListFolderOperationsRequest, opts ...grpc.CallOption) (*resourcemanager.ListFolderOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return resourcemanager.NewFolderServiceClient(conn).ListOperations(ctx, in, opts...)
}

type FolderOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *FolderServiceClient
	request *resourcemanager.ListFolderOperationsRequest

	items []*operation.Operation
}

func (c *FolderServiceClient) FolderOperationsIterator(ctx context.Context, req *resourcemanager.ListFolderOperationsRequest, opts ...grpc.CallOption) *FolderOperationsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &FolderOperationsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *FolderOperationsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.ListOperations(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Operations
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *FolderOperationsIterator) Take(size int64) ([]*operation.Operation, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*operation.Operation

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *FolderOperationsIterator) TakeAll() ([]*operation.Operation, error) {
	return it.Take(0)
}

func (it *FolderOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *FolderOperationsIterator) Error() error {
	return it.err
}

// SetAccessBindings implements resourcemanager.FolderServiceClient
func (c *FolderServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return resourcemanager.NewFolderServiceClient(conn).SetAccessBindings(ctx, in, opts...)
}

// Update implements resourcemanager.FolderServiceClient
func (c *FolderServiceClient) Update(ctx context.Context, in *resourcemanager.UpdateFolderRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return resourcemanager.NewFolderServiceClient(conn).Update(ctx, in, opts...)
}

// UpdateAccessBindings implements resourcemanager.FolderServiceClient
func (c *FolderServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return resourcemanager.NewFolderServiceClient(conn).UpdateAccessBindings(ctx, in, opts...)
}