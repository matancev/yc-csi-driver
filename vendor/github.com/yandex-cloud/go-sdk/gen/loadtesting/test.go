// Code generated by sdkgen. DO NOT EDIT.

// nolint
package api

import (
	"context"

	"google.golang.org/grpc"

	api "github.com/yandex-cloud/go-genproto/yandex/cloud/loadtesting/api/v1"
	test "github.com/yandex-cloud/go-genproto/yandex/cloud/loadtesting/api/v1/test"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// TestServiceClient is a api.TestServiceClient with
// lazy GRPC connection initialization.
type TestServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements api.TestServiceClient
func (c *TestServiceClient) Create(ctx context.Context, in *api.CreateTestRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return api.NewTestServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements api.TestServiceClient
func (c *TestServiceClient) Delete(ctx context.Context, in *api.DeleteTestRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return api.NewTestServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements api.TestServiceClient
func (c *TestServiceClient) Get(ctx context.Context, in *api.GetTestRequest, opts ...grpc.CallOption) (*test.Test, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return api.NewTestServiceClient(conn).Get(ctx, in, opts...)
}

// List implements api.TestServiceClient
func (c *TestServiceClient) List(ctx context.Context, in *api.ListTestsRequest, opts ...grpc.CallOption) (*api.ListTestsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return api.NewTestServiceClient(conn).List(ctx, in, opts...)
}

type TestIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *TestServiceClient
	request *api.ListTestsRequest

	items []*test.Test
}

func (c *TestServiceClient) TestIterator(ctx context.Context, req *api.ListTestsRequest, opts ...grpc.CallOption) *TestIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &TestIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *TestIterator) Next() bool {
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

	it.items = response.Tests
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *TestIterator) Take(size int64) ([]*test.Test, error) {
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

	var result []*test.Test

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *TestIterator) TakeAll() ([]*test.Test, error) {
	return it.Take(0)
}

func (it *TestIterator) Value() *test.Test {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *TestIterator) Error() error {
	return it.err
}

// Stop implements api.TestServiceClient
func (c *TestServiceClient) Stop(ctx context.Context, in *api.StopTestRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return api.NewTestServiceClient(conn).Stop(ctx, in, opts...)
}