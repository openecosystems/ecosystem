// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: platform/encryption/v2alpha/encryption.proto

package encryptionv2alphapb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	EncryptionService_Encrypt_FullMethodName = "/platform.encryption.v2alpha.EncryptionService/Encrypt"
	EncryptionService_Decrypt_FullMethodName = "/platform.encryption.v2alpha.EncryptionService/Decrypt"
)

// EncryptionServiceClient is the client API for EncryptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EncryptionServiceClient interface {
	Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error)
	Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error)
}

type encryptionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEncryptionServiceClient(cc grpc.ClientConnInterface) EncryptionServiceClient {
	return &encryptionServiceClient{cc}
}

func (c *encryptionServiceClient) Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EncryptResponse)
	err := c.cc.Invoke(ctx, EncryptionService_Encrypt_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encryptionServiceClient) Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DecryptResponse)
	err := c.cc.Invoke(ctx, EncryptionService_Decrypt_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EncryptionServiceServer is the server API for EncryptionService service.
// All implementations must embed UnimplementedEncryptionServiceServer
// for forward compatibility.
type EncryptionServiceServer interface {
	Encrypt(context.Context, *EncryptRequest) (*EncryptResponse, error)
	Decrypt(context.Context, *DecryptRequest) (*DecryptResponse, error)
	mustEmbedUnimplementedEncryptionServiceServer()
}

// UnimplementedEncryptionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEncryptionServiceServer struct{}

func (UnimplementedEncryptionServiceServer) Encrypt(context.Context, *EncryptRequest) (*EncryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encrypt not implemented")
}
func (UnimplementedEncryptionServiceServer) Decrypt(context.Context, *DecryptRequest) (*DecryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrypt not implemented")
}
func (UnimplementedEncryptionServiceServer) mustEmbedUnimplementedEncryptionServiceServer() {}
func (UnimplementedEncryptionServiceServer) testEmbeddedByValue()                           {}

// UnsafeEncryptionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EncryptionServiceServer will
// result in compilation errors.
type UnsafeEncryptionServiceServer interface {
	mustEmbedUnimplementedEncryptionServiceServer()
}

func RegisterEncryptionServiceServer(s grpc.ServiceRegistrar, srv EncryptionServiceServer) {
	// If the following call pancis, it indicates UnimplementedEncryptionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&EncryptionService_ServiceDesc, srv)
}

func _EncryptionService_Encrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncryptionServiceServer).Encrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EncryptionService_Encrypt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncryptionServiceServer).Encrypt(ctx, req.(*EncryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EncryptionService_Decrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncryptionServiceServer).Decrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EncryptionService_Decrypt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncryptionServiceServer).Decrypt(ctx, req.(*DecryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EncryptionService_ServiceDesc is the grpc.ServiceDesc for EncryptionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EncryptionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "platform.encryption.v2alpha.EncryptionService",
	HandlerType: (*EncryptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Encrypt",
			Handler:    _EncryptionService_Encrypt_Handler,
		},
		{
			MethodName: "Decrypt",
			Handler:    _EncryptionService_Decrypt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "platform/encryption/v2alpha/encryption.proto",
}
