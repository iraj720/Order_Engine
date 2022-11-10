// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: accounting/crypto/services.proto

package crypto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CryptoAPIClient is the client API for CryptoAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CryptoAPIClient interface {
	CryptoDeposit(ctx context.Context, in *CryptoDepositRequest, opts ...grpc.CallOption) (*CryptoDepositResponse, error)
	CryptoWithdraw(ctx context.Context, in *CryptoWithdrawRequest, opts ...grpc.CallOption) (*CryptoWithdrawResponse, error)
	// CryptoCorrectiveDeposit must be used to correct a finalized withdraw
	CryptoCorrectiveDeposit(ctx context.Context, in *CryptoCorrectiveDepositRequest, opts ...grpc.CallOption) (*CryptoCorrectiveDepositResponse, error)
}

type cryptoAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewCryptoAPIClient(cc grpc.ClientConnInterface) CryptoAPIClient {
	return &cryptoAPIClient{cc}
}

func (c *cryptoAPIClient) CryptoDeposit(ctx context.Context, in *CryptoDepositRequest, opts ...grpc.CallOption) (*CryptoDepositResponse, error) {
	out := new(CryptoDepositResponse)
	err := c.cc.Invoke(ctx, "/crypto.CryptoAPI/CryptoDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoAPIClient) CryptoWithdraw(ctx context.Context, in *CryptoWithdrawRequest, opts ...grpc.CallOption) (*CryptoWithdrawResponse, error) {
	out := new(CryptoWithdrawResponse)
	err := c.cc.Invoke(ctx, "/crypto.CryptoAPI/CryptoWithdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoAPIClient) CryptoCorrectiveDeposit(ctx context.Context, in *CryptoCorrectiveDepositRequest, opts ...grpc.CallOption) (*CryptoCorrectiveDepositResponse, error) {
	out := new(CryptoCorrectiveDepositResponse)
	err := c.cc.Invoke(ctx, "/crypto.CryptoAPI/CryptoCorrectiveDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CryptoAPIServer is the server API for CryptoAPI service.
// All implementations must embed UnimplementedCryptoAPIServer
// for forward compatibility
type CryptoAPIServer interface {
	CryptoDeposit(context.Context, *CryptoDepositRequest) (*CryptoDepositResponse, error)
	CryptoWithdraw(context.Context, *CryptoWithdrawRequest) (*CryptoWithdrawResponse, error)
	// CryptoCorrectiveDeposit must be used to correct a finalized withdraw
	CryptoCorrectiveDeposit(context.Context, *CryptoCorrectiveDepositRequest) (*CryptoCorrectiveDepositResponse, error)
	mustEmbedUnimplementedCryptoAPIServer()
}

// UnimplementedCryptoAPIServer must be embedded to have forward compatible implementations.
type UnimplementedCryptoAPIServer struct {
}

func (UnimplementedCryptoAPIServer) CryptoDeposit(context.Context, *CryptoDepositRequest) (*CryptoDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CryptoDeposit not implemented")
}
func (UnimplementedCryptoAPIServer) CryptoWithdraw(context.Context, *CryptoWithdrawRequest) (*CryptoWithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CryptoWithdraw not implemented")
}
func (UnimplementedCryptoAPIServer) CryptoCorrectiveDeposit(context.Context, *CryptoCorrectiveDepositRequest) (*CryptoCorrectiveDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CryptoCorrectiveDeposit not implemented")
}
func (UnimplementedCryptoAPIServer) mustEmbedUnimplementedCryptoAPIServer() {}

// UnsafeCryptoAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CryptoAPIServer will
// result in compilation errors.
type UnsafeCryptoAPIServer interface {
	mustEmbedUnimplementedCryptoAPIServer()
}

func RegisterCryptoAPIServer(s grpc.ServiceRegistrar, srv CryptoAPIServer) {
	s.RegisterService(&CryptoAPI_ServiceDesc, srv)
}

func _CryptoAPI_CryptoDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CryptoDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoAPIServer).CryptoDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crypto.CryptoAPI/CryptoDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoAPIServer).CryptoDeposit(ctx, req.(*CryptoDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoAPI_CryptoWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CryptoWithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoAPIServer).CryptoWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crypto.CryptoAPI/CryptoWithdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoAPIServer).CryptoWithdraw(ctx, req.(*CryptoWithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoAPI_CryptoCorrectiveDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CryptoCorrectiveDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoAPIServer).CryptoCorrectiveDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crypto.CryptoAPI/CryptoCorrectiveDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoAPIServer).CryptoCorrectiveDeposit(ctx, req.(*CryptoCorrectiveDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CryptoAPI_ServiceDesc is the grpc.ServiceDesc for CryptoAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CryptoAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "crypto.CryptoAPI",
	HandlerType: (*CryptoAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CryptoDeposit",
			Handler:    _CryptoAPI_CryptoDeposit_Handler,
		},
		{
			MethodName: "CryptoWithdraw",
			Handler:    _CryptoAPI_CryptoWithdraw_Handler,
		},
		{
			MethodName: "CryptoCorrectiveDeposit",
			Handler:    _CryptoAPI_CryptoCorrectiveDeposit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accounting/crypto/services.proto",
}