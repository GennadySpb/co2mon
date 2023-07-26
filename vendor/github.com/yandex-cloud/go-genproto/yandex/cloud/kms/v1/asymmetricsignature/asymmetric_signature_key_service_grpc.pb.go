// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/kms/v1/asymmetricsignature/asymmetric_signature_key_service.proto

package kms

import (
	context "context"
	access "github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AsymmetricSignatureKeyService_Create_FullMethodName               = "/yandex.cloud.kms.v1.asymmetricsignature.AsymmetricSignatureKeyService/Create"
	AsymmetricSignatureKeyService_Get_FullMethodName                  = "/yandex.cloud.kms.v1.asymmetricsignature.AsymmetricSignatureKeyService/Get"
	AsymmetricSignatureKeyService_List_FullMethodName                 = "/yandex.cloud.kms.v1.asymmetricsignature.AsymmetricSignatureKeyService/List"
	AsymmetricSignatureKeyService_Update_FullMethodName               = "/yandex.cloud.kms.v1.asymmetricsignature.AsymmetricSignatureKeyService/Update"
	AsymmetricSignatureKeyService_Delete_FullMethodName               = "/yandex.cloud.kms.v1.asymmetricsignature.AsymmetricSignatureKeyService/Delete"
	AsymmetricSignatureKeyService_ListOperations_FullMethodName       = "/yandex.cloud.kms.v1.asymmetricsignature.AsymmetricSignatureKeyService/ListOperations"
	AsymmetricSignatureKeyService_ListAccessBindings_FullMethodName   = "/yandex.cloud.kms.v1.asymmetricsignature.AsymmetricSignatureKeyService/ListAccessBindings"
	AsymmetricSignatureKeyService_SetAccessBindings_FullMethodName    = "/yandex.cloud.kms.v1.asymmetricsignature.AsymmetricSignatureKeyService/SetAccessBindings"
	AsymmetricSignatureKeyService_UpdateAccessBindings_FullMethodName = "/yandex.cloud.kms.v1.asymmetricsignature.AsymmetricSignatureKeyService/UpdateAccessBindings"
)

// AsymmetricSignatureKeyServiceClient is the client API for AsymmetricSignatureKeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AsymmetricSignatureKeyServiceClient interface {
	// control plane
	// Creates an asymmetric KMS key in the specified folder.
	Create(ctx context.Context, in *CreateAsymmetricSignatureKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Returns the specified asymmetric KMS key.
	//
	//	To get the list of available asymmetric KMS keys, make a [SymmetricKeyService.List] request.
	Get(ctx context.Context, in *GetAsymmetricSignatureKeyRequest, opts ...grpc.CallOption) (*AsymmetricSignatureKey, error)
	// Returns the list of asymmetric KMS keys in the specified folder.
	List(ctx context.Context, in *ListAsymmetricSignatureKeysRequest, opts ...grpc.CallOption) (*ListAsymmetricSignatureKeysResponse, error)
	// Updates the specified asymmetric KMS key.
	Update(ctx context.Context, in *UpdateAsymmetricSignatureKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified asymmetric KMS key. This action also automatically schedules
	// the destruction of all of the key's versions in 72 hours.
	//
	// The key and its versions appear absent in [AsymmetricSignatureKeyService.Get] and [AsymmetricSignatureKeyService.List]
	// requests, but can be restored within 72 hours with a request to tech support.
	Delete(ctx context.Context, in *DeleteAsymmetricSignatureKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists operations for the specified asymmetric KMS key.
	ListOperations(ctx context.Context, in *ListAsymmetricSignatureKeyOperationsRequest, opts ...grpc.CallOption) (*ListAsymmetricSignatureKeyOperationsResponse, error)
	// Lists existing access bindings for the specified key.
	ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the key.
	SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates access bindings for the specified key.
	UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type asymmetricSignatureKeyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAsymmetricSignatureKeyServiceClient(cc grpc.ClientConnInterface) AsymmetricSignatureKeyServiceClient {
	return &asymmetricSignatureKeyServiceClient{cc}
}

func (c *asymmetricSignatureKeyServiceClient) Create(ctx context.Context, in *CreateAsymmetricSignatureKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, AsymmetricSignatureKeyService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asymmetricSignatureKeyServiceClient) Get(ctx context.Context, in *GetAsymmetricSignatureKeyRequest, opts ...grpc.CallOption) (*AsymmetricSignatureKey, error) {
	out := new(AsymmetricSignatureKey)
	err := c.cc.Invoke(ctx, AsymmetricSignatureKeyService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asymmetricSignatureKeyServiceClient) List(ctx context.Context, in *ListAsymmetricSignatureKeysRequest, opts ...grpc.CallOption) (*ListAsymmetricSignatureKeysResponse, error) {
	out := new(ListAsymmetricSignatureKeysResponse)
	err := c.cc.Invoke(ctx, AsymmetricSignatureKeyService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asymmetricSignatureKeyServiceClient) Update(ctx context.Context, in *UpdateAsymmetricSignatureKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, AsymmetricSignatureKeyService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asymmetricSignatureKeyServiceClient) Delete(ctx context.Context, in *DeleteAsymmetricSignatureKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, AsymmetricSignatureKeyService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asymmetricSignatureKeyServiceClient) ListOperations(ctx context.Context, in *ListAsymmetricSignatureKeyOperationsRequest, opts ...grpc.CallOption) (*ListAsymmetricSignatureKeyOperationsResponse, error) {
	out := new(ListAsymmetricSignatureKeyOperationsResponse)
	err := c.cc.Invoke(ctx, AsymmetricSignatureKeyService_ListOperations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asymmetricSignatureKeyServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	out := new(access.ListAccessBindingsResponse)
	err := c.cc.Invoke(ctx, AsymmetricSignatureKeyService_ListAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asymmetricSignatureKeyServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, AsymmetricSignatureKeyService_SetAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asymmetricSignatureKeyServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, AsymmetricSignatureKeyService_UpdateAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AsymmetricSignatureKeyServiceServer is the server API for AsymmetricSignatureKeyService service.
// All implementations should embed UnimplementedAsymmetricSignatureKeyServiceServer
// for forward compatibility
type AsymmetricSignatureKeyServiceServer interface {
	// control plane
	// Creates an asymmetric KMS key in the specified folder.
	Create(context.Context, *CreateAsymmetricSignatureKeyRequest) (*operation.Operation, error)
	// Returns the specified asymmetric KMS key.
	//
	//	To get the list of available asymmetric KMS keys, make a [SymmetricKeyService.List] request.
	Get(context.Context, *GetAsymmetricSignatureKeyRequest) (*AsymmetricSignatureKey, error)
	// Returns the list of asymmetric KMS keys in the specified folder.
	List(context.Context, *ListAsymmetricSignatureKeysRequest) (*ListAsymmetricSignatureKeysResponse, error)
	// Updates the specified asymmetric KMS key.
	Update(context.Context, *UpdateAsymmetricSignatureKeyRequest) (*operation.Operation, error)
	// Deletes the specified asymmetric KMS key. This action also automatically schedules
	// the destruction of all of the key's versions in 72 hours.
	//
	// The key and its versions appear absent in [AsymmetricSignatureKeyService.Get] and [AsymmetricSignatureKeyService.List]
	// requests, but can be restored within 72 hours with a request to tech support.
	Delete(context.Context, *DeleteAsymmetricSignatureKeyRequest) (*operation.Operation, error)
	// Lists operations for the specified asymmetric KMS key.
	ListOperations(context.Context, *ListAsymmetricSignatureKeyOperationsRequest) (*ListAsymmetricSignatureKeyOperationsResponse, error)
	// Lists existing access bindings for the specified key.
	ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the key.
	SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error)
	// Updates access bindings for the specified key.
	UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error)
}

// UnimplementedAsymmetricSignatureKeyServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAsymmetricSignatureKeyServiceServer struct {
}

func (UnimplementedAsymmetricSignatureKeyServiceServer) Create(context.Context, *CreateAsymmetricSignatureKeyRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAsymmetricSignatureKeyServiceServer) Get(context.Context, *GetAsymmetricSignatureKeyRequest) (*AsymmetricSignatureKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAsymmetricSignatureKeyServiceServer) List(context.Context, *ListAsymmetricSignatureKeysRequest) (*ListAsymmetricSignatureKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAsymmetricSignatureKeyServiceServer) Update(context.Context, *UpdateAsymmetricSignatureKeyRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAsymmetricSignatureKeyServiceServer) Delete(context.Context, *DeleteAsymmetricSignatureKeyRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAsymmetricSignatureKeyServiceServer) ListOperations(context.Context, *ListAsymmetricSignatureKeyOperationsRequest) (*ListAsymmetricSignatureKeyOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedAsymmetricSignatureKeyServiceServer) ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccessBindings not implemented")
}
func (UnimplementedAsymmetricSignatureKeyServiceServer) SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccessBindings not implemented")
}
func (UnimplementedAsymmetricSignatureKeyServiceServer) UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccessBindings not implemented")
}

// UnsafeAsymmetricSignatureKeyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AsymmetricSignatureKeyServiceServer will
// result in compilation errors.
type UnsafeAsymmetricSignatureKeyServiceServer interface {
	mustEmbedUnimplementedAsymmetricSignatureKeyServiceServer()
}

func RegisterAsymmetricSignatureKeyServiceServer(s grpc.ServiceRegistrar, srv AsymmetricSignatureKeyServiceServer) {
	s.RegisterService(&AsymmetricSignatureKeyService_ServiceDesc, srv)
}

func _AsymmetricSignatureKeyService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAsymmetricSignatureKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsymmetricSignatureKeyServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsymmetricSignatureKeyService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsymmetricSignatureKeyServiceServer).Create(ctx, req.(*CreateAsymmetricSignatureKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsymmetricSignatureKeyService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAsymmetricSignatureKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsymmetricSignatureKeyServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsymmetricSignatureKeyService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsymmetricSignatureKeyServiceServer).Get(ctx, req.(*GetAsymmetricSignatureKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsymmetricSignatureKeyService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAsymmetricSignatureKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsymmetricSignatureKeyServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsymmetricSignatureKeyService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsymmetricSignatureKeyServiceServer).List(ctx, req.(*ListAsymmetricSignatureKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsymmetricSignatureKeyService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAsymmetricSignatureKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsymmetricSignatureKeyServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsymmetricSignatureKeyService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsymmetricSignatureKeyServiceServer).Update(ctx, req.(*UpdateAsymmetricSignatureKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsymmetricSignatureKeyService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAsymmetricSignatureKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsymmetricSignatureKeyServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsymmetricSignatureKeyService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsymmetricSignatureKeyServiceServer).Delete(ctx, req.(*DeleteAsymmetricSignatureKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsymmetricSignatureKeyService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAsymmetricSignatureKeyOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsymmetricSignatureKeyServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsymmetricSignatureKeyService_ListOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsymmetricSignatureKeyServiceServer).ListOperations(ctx, req.(*ListAsymmetricSignatureKeyOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsymmetricSignatureKeyService_ListAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.ListAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsymmetricSignatureKeyServiceServer).ListAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsymmetricSignatureKeyService_ListAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsymmetricSignatureKeyServiceServer).ListAccessBindings(ctx, req.(*access.ListAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsymmetricSignatureKeyService_SetAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.SetAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsymmetricSignatureKeyServiceServer).SetAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsymmetricSignatureKeyService_SetAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsymmetricSignatureKeyServiceServer).SetAccessBindings(ctx, req.(*access.SetAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsymmetricSignatureKeyService_UpdateAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.UpdateAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsymmetricSignatureKeyServiceServer).UpdateAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AsymmetricSignatureKeyService_UpdateAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsymmetricSignatureKeyServiceServer).UpdateAccessBindings(ctx, req.(*access.UpdateAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AsymmetricSignatureKeyService_ServiceDesc is the grpc.ServiceDesc for AsymmetricSignatureKeyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AsymmetricSignatureKeyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.kms.v1.asymmetricsignature.AsymmetricSignatureKeyService",
	HandlerType: (*AsymmetricSignatureKeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AsymmetricSignatureKeyService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AsymmetricSignatureKeyService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _AsymmetricSignatureKeyService_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AsymmetricSignatureKeyService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AsymmetricSignatureKeyService_Delete_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _AsymmetricSignatureKeyService_ListOperations_Handler,
		},
		{
			MethodName: "ListAccessBindings",
			Handler:    _AsymmetricSignatureKeyService_ListAccessBindings_Handler,
		},
		{
			MethodName: "SetAccessBindings",
			Handler:    _AsymmetricSignatureKeyService_SetAccessBindings_Handler,
		},
		{
			MethodName: "UpdateAccessBindings",
			Handler:    _AsymmetricSignatureKeyService_UpdateAccessBindings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/kms/v1/asymmetricsignature/asymmetric_signature_key_service.proto",
}
