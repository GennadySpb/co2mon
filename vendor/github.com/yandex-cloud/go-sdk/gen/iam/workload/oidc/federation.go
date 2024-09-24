// Code generated by sdkgen. DO NOT EDIT.

// nolint
package oidc

import (
	"context"

	"google.golang.org/grpc"

	oidc "github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1/workload/oidc"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// FederationServiceClient is a oidc.FederationServiceClient with
// lazy GRPC connection initialization.
type FederationServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements oidc.FederationServiceClient
func (c *FederationServiceClient) Create(ctx context.Context, in *oidc.CreateFederationRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return oidc.NewFederationServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements oidc.FederationServiceClient
func (c *FederationServiceClient) Delete(ctx context.Context, in *oidc.DeleteFederationRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return oidc.NewFederationServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements oidc.FederationServiceClient
func (c *FederationServiceClient) Get(ctx context.Context, in *oidc.GetFederationRequest, opts ...grpc.CallOption) (*oidc.Federation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return oidc.NewFederationServiceClient(conn).Get(ctx, in, opts...)
}

// List implements oidc.FederationServiceClient
func (c *FederationServiceClient) List(ctx context.Context, in *oidc.ListFederationsRequest, opts ...grpc.CallOption) (*oidc.ListFederationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return oidc.NewFederationServiceClient(conn).List(ctx, in, opts...)
}

type FederationIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *FederationServiceClient
	request *oidc.ListFederationsRequest

	items []*oidc.Federation
}

func (c *FederationServiceClient) FederationIterator(ctx context.Context, req *oidc.ListFederationsRequest, opts ...grpc.CallOption) *FederationIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &FederationIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *FederationIterator) Next() bool {
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

	it.items = response.Federations
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *FederationIterator) Take(size int64) ([]*oidc.Federation, error) {
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

	var result []*oidc.Federation

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *FederationIterator) TakeAll() ([]*oidc.Federation, error) {
	return it.Take(0)
}

func (it *FederationIterator) Value() *oidc.Federation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *FederationIterator) Error() error {
	return it.err
}

// Update implements oidc.FederationServiceClient
func (c *FederationServiceClient) Update(ctx context.Context, in *oidc.UpdateFederationRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return oidc.NewFederationServiceClient(conn).Update(ctx, in, opts...)
}
