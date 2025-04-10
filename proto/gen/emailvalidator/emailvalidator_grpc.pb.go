// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/proto/emailvalidator/emailvalidator.proto

package emailvalidator

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
	EmailValidator_Validate_FullMethodName = "/emailvalidator.EmailValidator/Validate"
)

// EmailValidatorClient is the client API for EmailValidator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmailValidatorClient interface {
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
}

type emailValidatorClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailValidatorClient(cc grpc.ClientConnInterface) EmailValidatorClient {
	return &emailValidatorClient{cc}
}

func (c *emailValidatorClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidateResponse)
	err := c.cc.Invoke(ctx, EmailValidator_Validate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailValidatorServer is the server API for EmailValidator service.
// All implementations must embed UnimplementedEmailValidatorServer
// for forward compatibility.
type EmailValidatorServer interface {
	Validate(context.Context, *ValidateRequest) (*ValidateResponse, error)
	mustEmbedUnimplementedEmailValidatorServer()
}

// UnimplementedEmailValidatorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEmailValidatorServer struct{}

func (UnimplementedEmailValidatorServer) Validate(context.Context, *ValidateRequest) (*ValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedEmailValidatorServer) mustEmbedUnimplementedEmailValidatorServer() {}
func (UnimplementedEmailValidatorServer) testEmbeddedByValue()                        {}

// UnsafeEmailValidatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmailValidatorServer will
// result in compilation errors.
type UnsafeEmailValidatorServer interface {
	mustEmbedUnimplementedEmailValidatorServer()
}

func RegisterEmailValidatorServer(s grpc.ServiceRegistrar, srv EmailValidatorServer) {
	// If the following call pancis, it indicates UnimplementedEmailValidatorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&EmailValidator_ServiceDesc, srv)
}

func _EmailValidator_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailValidatorServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailValidator_Validate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailValidatorServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EmailValidator_ServiceDesc is the grpc.ServiceDesc for EmailValidator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmailValidator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "emailvalidator.EmailValidator",
	HandlerType: (*EmailValidatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Validate",
			Handler:    _EmailValidator_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/proto/emailvalidator/emailvalidator.proto",
}
