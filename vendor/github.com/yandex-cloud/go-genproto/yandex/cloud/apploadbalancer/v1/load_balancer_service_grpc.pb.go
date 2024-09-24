// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/apploadbalancer/v1/load_balancer_service.proto

package apploadbalancer

import (
	context "context"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LoadBalancerService_Get_FullMethodName             = "/yandex.cloud.apploadbalancer.v1.LoadBalancerService/Get"
	LoadBalancerService_List_FullMethodName            = "/yandex.cloud.apploadbalancer.v1.LoadBalancerService/List"
	LoadBalancerService_Create_FullMethodName          = "/yandex.cloud.apploadbalancer.v1.LoadBalancerService/Create"
	LoadBalancerService_Update_FullMethodName          = "/yandex.cloud.apploadbalancer.v1.LoadBalancerService/Update"
	LoadBalancerService_Delete_FullMethodName          = "/yandex.cloud.apploadbalancer.v1.LoadBalancerService/Delete"
	LoadBalancerService_Start_FullMethodName           = "/yandex.cloud.apploadbalancer.v1.LoadBalancerService/Start"
	LoadBalancerService_Stop_FullMethodName            = "/yandex.cloud.apploadbalancer.v1.LoadBalancerService/Stop"
	LoadBalancerService_AddListener_FullMethodName     = "/yandex.cloud.apploadbalancer.v1.LoadBalancerService/AddListener"
	LoadBalancerService_RemoveListener_FullMethodName  = "/yandex.cloud.apploadbalancer.v1.LoadBalancerService/RemoveListener"
	LoadBalancerService_UpdateListener_FullMethodName  = "/yandex.cloud.apploadbalancer.v1.LoadBalancerService/UpdateListener"
	LoadBalancerService_AddSniMatch_FullMethodName     = "/yandex.cloud.apploadbalancer.v1.LoadBalancerService/AddSniMatch"
	LoadBalancerService_UpdateSniMatch_FullMethodName  = "/yandex.cloud.apploadbalancer.v1.LoadBalancerService/UpdateSniMatch"
	LoadBalancerService_RemoveSniMatch_FullMethodName  = "/yandex.cloud.apploadbalancer.v1.LoadBalancerService/RemoveSniMatch"
	LoadBalancerService_GetTargetStates_FullMethodName = "/yandex.cloud.apploadbalancer.v1.LoadBalancerService/GetTargetStates"
	LoadBalancerService_ListOperations_FullMethodName  = "/yandex.cloud.apploadbalancer.v1.LoadBalancerService/ListOperations"
)

// LoadBalancerServiceClient is the client API for LoadBalancerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing application load balancers.
type LoadBalancerServiceClient interface {
	// Returns the specified application load balancer.
	//
	// To get the list of all available application load balancers, make a [List] request.
	Get(ctx context.Context, in *GetLoadBalancerRequest, opts ...grpc.CallOption) (*LoadBalancer, error)
	// Lists application load balancers in the specified folder.
	List(ctx context.Context, in *ListLoadBalancersRequest, opts ...grpc.CallOption) (*ListLoadBalancersResponse, error)
	// Creates an application load balancer in the specified folder.
	Create(ctx context.Context, in *CreateLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified application load balancer.
	Update(ctx context.Context, in *UpdateLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified application load balancer.
	Delete(ctx context.Context, in *DeleteLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Starts the specified application load balancer.
	Start(ctx context.Context, in *StartLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Stops the specified application load balancer.
	Stop(ctx context.Context, in *StopLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Adds a listener to the specified application load balancer.
	AddListener(ctx context.Context, in *AddListenerRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified listener.
	RemoveListener(ctx context.Context, in *RemoveListenerRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified listener of the specified application load balancer.
	UpdateListener(ctx context.Context, in *UpdateListenerRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Adds a SNI handler to the specified listener.
	//
	// This request does not allow to add [TlsListener.default_handler]. Make an [UpdateListener] request instead.
	AddSniMatch(ctx context.Context, in *AddSniMatchRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified SNI handler of the specified listener.
	//
	// This request does not allow to update [TlsListener.default_handler]. Make an [UpdateListener] request instead.
	UpdateSniMatch(ctx context.Context, in *UpdateSniMatchRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified SNI handler.
	//
	// This request does not allow to delete [TlsListener.default_handler].
	RemoveSniMatch(ctx context.Context, in *RemoveSniMatchRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Returns the statuses of all targets of the specified backend group in all their availability zones.
	GetTargetStates(ctx context.Context, in *GetTargetStatesRequest, opts ...grpc.CallOption) (*GetTargetStatesResponse, error)
	// Lists operations for the specified application load balancer.
	ListOperations(ctx context.Context, in *ListLoadBalancerOperationsRequest, opts ...grpc.CallOption) (*ListLoadBalancerOperationsResponse, error)
}

type loadBalancerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoadBalancerServiceClient(cc grpc.ClientConnInterface) LoadBalancerServiceClient {
	return &loadBalancerServiceClient{cc}
}

func (c *loadBalancerServiceClient) Get(ctx context.Context, in *GetLoadBalancerRequest, opts ...grpc.CallOption) (*LoadBalancer, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoadBalancer)
	err := c.cc.Invoke(ctx, LoadBalancerService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancerServiceClient) List(ctx context.Context, in *ListLoadBalancersRequest, opts ...grpc.CallOption) (*ListLoadBalancersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListLoadBalancersResponse)
	err := c.cc.Invoke(ctx, LoadBalancerService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancerServiceClient) Create(ctx context.Context, in *CreateLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, LoadBalancerService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancerServiceClient) Update(ctx context.Context, in *UpdateLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, LoadBalancerService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancerServiceClient) Delete(ctx context.Context, in *DeleteLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, LoadBalancerService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancerServiceClient) Start(ctx context.Context, in *StartLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, LoadBalancerService_Start_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancerServiceClient) Stop(ctx context.Context, in *StopLoadBalancerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, LoadBalancerService_Stop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancerServiceClient) AddListener(ctx context.Context, in *AddListenerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, LoadBalancerService_AddListener_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancerServiceClient) RemoveListener(ctx context.Context, in *RemoveListenerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, LoadBalancerService_RemoveListener_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancerServiceClient) UpdateListener(ctx context.Context, in *UpdateListenerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, LoadBalancerService_UpdateListener_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancerServiceClient) AddSniMatch(ctx context.Context, in *AddSniMatchRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, LoadBalancerService_AddSniMatch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancerServiceClient) UpdateSniMatch(ctx context.Context, in *UpdateSniMatchRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, LoadBalancerService_UpdateSniMatch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancerServiceClient) RemoveSniMatch(ctx context.Context, in *RemoveSniMatchRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, LoadBalancerService_RemoveSniMatch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancerServiceClient) GetTargetStates(ctx context.Context, in *GetTargetStatesRequest, opts ...grpc.CallOption) (*GetTargetStatesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTargetStatesResponse)
	err := c.cc.Invoke(ctx, LoadBalancerService_GetTargetStates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancerServiceClient) ListOperations(ctx context.Context, in *ListLoadBalancerOperationsRequest, opts ...grpc.CallOption) (*ListLoadBalancerOperationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListLoadBalancerOperationsResponse)
	err := c.cc.Invoke(ctx, LoadBalancerService_ListOperations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoadBalancerServiceServer is the server API for LoadBalancerService service.
// All implementations should embed UnimplementedLoadBalancerServiceServer
// for forward compatibility.
//
// A set of methods for managing application load balancers.
type LoadBalancerServiceServer interface {
	// Returns the specified application load balancer.
	//
	// To get the list of all available application load balancers, make a [List] request.
	Get(context.Context, *GetLoadBalancerRequest) (*LoadBalancer, error)
	// Lists application load balancers in the specified folder.
	List(context.Context, *ListLoadBalancersRequest) (*ListLoadBalancersResponse, error)
	// Creates an application load balancer in the specified folder.
	Create(context.Context, *CreateLoadBalancerRequest) (*operation.Operation, error)
	// Updates the specified application load balancer.
	Update(context.Context, *UpdateLoadBalancerRequest) (*operation.Operation, error)
	// Deletes the specified application load balancer.
	Delete(context.Context, *DeleteLoadBalancerRequest) (*operation.Operation, error)
	// Starts the specified application load balancer.
	Start(context.Context, *StartLoadBalancerRequest) (*operation.Operation, error)
	// Stops the specified application load balancer.
	Stop(context.Context, *StopLoadBalancerRequest) (*operation.Operation, error)
	// Adds a listener to the specified application load balancer.
	AddListener(context.Context, *AddListenerRequest) (*operation.Operation, error)
	// Deletes the specified listener.
	RemoveListener(context.Context, *RemoveListenerRequest) (*operation.Operation, error)
	// Updates the specified listener of the specified application load balancer.
	UpdateListener(context.Context, *UpdateListenerRequest) (*operation.Operation, error)
	// Adds a SNI handler to the specified listener.
	//
	// This request does not allow to add [TlsListener.default_handler]. Make an [UpdateListener] request instead.
	AddSniMatch(context.Context, *AddSniMatchRequest) (*operation.Operation, error)
	// Updates the specified SNI handler of the specified listener.
	//
	// This request does not allow to update [TlsListener.default_handler]. Make an [UpdateListener] request instead.
	UpdateSniMatch(context.Context, *UpdateSniMatchRequest) (*operation.Operation, error)
	// Deletes the specified SNI handler.
	//
	// This request does not allow to delete [TlsListener.default_handler].
	RemoveSniMatch(context.Context, *RemoveSniMatchRequest) (*operation.Operation, error)
	// Returns the statuses of all targets of the specified backend group in all their availability zones.
	GetTargetStates(context.Context, *GetTargetStatesRequest) (*GetTargetStatesResponse, error)
	// Lists operations for the specified application load balancer.
	ListOperations(context.Context, *ListLoadBalancerOperationsRequest) (*ListLoadBalancerOperationsResponse, error)
}

// UnimplementedLoadBalancerServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLoadBalancerServiceServer struct{}

func (UnimplementedLoadBalancerServiceServer) Get(context.Context, *GetLoadBalancerRequest) (*LoadBalancer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedLoadBalancerServiceServer) List(context.Context, *ListLoadBalancersRequest) (*ListLoadBalancersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedLoadBalancerServiceServer) Create(context.Context, *CreateLoadBalancerRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedLoadBalancerServiceServer) Update(context.Context, *UpdateLoadBalancerRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedLoadBalancerServiceServer) Delete(context.Context, *DeleteLoadBalancerRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedLoadBalancerServiceServer) Start(context.Context, *StartLoadBalancerRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedLoadBalancerServiceServer) Stop(context.Context, *StopLoadBalancerRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedLoadBalancerServiceServer) AddListener(context.Context, *AddListenerRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddListener not implemented")
}
func (UnimplementedLoadBalancerServiceServer) RemoveListener(context.Context, *RemoveListenerRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveListener not implemented")
}
func (UnimplementedLoadBalancerServiceServer) UpdateListener(context.Context, *UpdateListenerRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateListener not implemented")
}
func (UnimplementedLoadBalancerServiceServer) AddSniMatch(context.Context, *AddSniMatchRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSniMatch not implemented")
}
func (UnimplementedLoadBalancerServiceServer) UpdateSniMatch(context.Context, *UpdateSniMatchRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSniMatch not implemented")
}
func (UnimplementedLoadBalancerServiceServer) RemoveSniMatch(context.Context, *RemoveSniMatchRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSniMatch not implemented")
}
func (UnimplementedLoadBalancerServiceServer) GetTargetStates(context.Context, *GetTargetStatesRequest) (*GetTargetStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTargetStates not implemented")
}
func (UnimplementedLoadBalancerServiceServer) ListOperations(context.Context, *ListLoadBalancerOperationsRequest) (*ListLoadBalancerOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedLoadBalancerServiceServer) testEmbeddedByValue() {}

// UnsafeLoadBalancerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoadBalancerServiceServer will
// result in compilation errors.
type UnsafeLoadBalancerServiceServer interface {
	mustEmbedUnimplementedLoadBalancerServiceServer()
}

func RegisterLoadBalancerServiceServer(s grpc.ServiceRegistrar, srv LoadBalancerServiceServer) {
	// If the following call pancis, it indicates UnimplementedLoadBalancerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LoadBalancerService_ServiceDesc, srv)
}

func _LoadBalancerService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLoadBalancerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancerServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoadBalancerService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancerServiceServer).Get(ctx, req.(*GetLoadBalancerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancerService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLoadBalancersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancerServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoadBalancerService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancerServiceServer).List(ctx, req.(*ListLoadBalancersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancerService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLoadBalancerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancerServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoadBalancerService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancerServiceServer).Create(ctx, req.(*CreateLoadBalancerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancerService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLoadBalancerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancerServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoadBalancerService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancerServiceServer).Update(ctx, req.(*UpdateLoadBalancerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancerService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLoadBalancerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancerServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoadBalancerService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancerServiceServer).Delete(ctx, req.(*DeleteLoadBalancerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancerService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartLoadBalancerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancerServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoadBalancerService_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancerServiceServer).Start(ctx, req.(*StartLoadBalancerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancerService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopLoadBalancerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancerServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoadBalancerService_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancerServiceServer).Stop(ctx, req.(*StopLoadBalancerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancerService_AddListener_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddListenerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancerServiceServer).AddListener(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoadBalancerService_AddListener_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancerServiceServer).AddListener(ctx, req.(*AddListenerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancerService_RemoveListener_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveListenerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancerServiceServer).RemoveListener(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoadBalancerService_RemoveListener_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancerServiceServer).RemoveListener(ctx, req.(*RemoveListenerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancerService_UpdateListener_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateListenerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancerServiceServer).UpdateListener(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoadBalancerService_UpdateListener_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancerServiceServer).UpdateListener(ctx, req.(*UpdateListenerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancerService_AddSniMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSniMatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancerServiceServer).AddSniMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoadBalancerService_AddSniMatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancerServiceServer).AddSniMatch(ctx, req.(*AddSniMatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancerService_UpdateSniMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSniMatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancerServiceServer).UpdateSniMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoadBalancerService_UpdateSniMatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancerServiceServer).UpdateSniMatch(ctx, req.(*UpdateSniMatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancerService_RemoveSniMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveSniMatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancerServiceServer).RemoveSniMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoadBalancerService_RemoveSniMatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancerServiceServer).RemoveSniMatch(ctx, req.(*RemoveSniMatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancerService_GetTargetStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTargetStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancerServiceServer).GetTargetStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoadBalancerService_GetTargetStates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancerServiceServer).GetTargetStates(ctx, req.(*GetTargetStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancerService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLoadBalancerOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancerServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoadBalancerService_ListOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancerServiceServer).ListOperations(ctx, req.(*ListLoadBalancerOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LoadBalancerService_ServiceDesc is the grpc.ServiceDesc for LoadBalancerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoadBalancerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.apploadbalancer.v1.LoadBalancerService",
	HandlerType: (*LoadBalancerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _LoadBalancerService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _LoadBalancerService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _LoadBalancerService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _LoadBalancerService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _LoadBalancerService_Delete_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _LoadBalancerService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _LoadBalancerService_Stop_Handler,
		},
		{
			MethodName: "AddListener",
			Handler:    _LoadBalancerService_AddListener_Handler,
		},
		{
			MethodName: "RemoveListener",
			Handler:    _LoadBalancerService_RemoveListener_Handler,
		},
		{
			MethodName: "UpdateListener",
			Handler:    _LoadBalancerService_UpdateListener_Handler,
		},
		{
			MethodName: "AddSniMatch",
			Handler:    _LoadBalancerService_AddSniMatch_Handler,
		},
		{
			MethodName: "UpdateSniMatch",
			Handler:    _LoadBalancerService_UpdateSniMatch_Handler,
		},
		{
			MethodName: "RemoveSniMatch",
			Handler:    _LoadBalancerService_RemoveSniMatch_Handler,
		},
		{
			MethodName: "GetTargetStates",
			Handler:    _LoadBalancerService_GetTargetStates_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _LoadBalancerService_ListOperations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/apploadbalancer/v1/load_balancer_service.proto",
}
