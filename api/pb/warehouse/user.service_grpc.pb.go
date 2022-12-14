// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: warehouse/user.service.proto

package warehouse

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	transaction "rabex/api/pb/warehouse/transaction"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserOperationsClient is the client API for UserOperations service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserOperationsClient interface {
	// Warehouse and Account
	CreateWarehouseAccount(ctx context.Context, in *UserAccountRequest, opts ...grpc.CallOption) (*UserAccountResponse, error)
	// Treasury Account
	CreateTreasuryAccount(ctx context.Context, in *UserAccountRequest, opts ...grpc.CallOption) (*UserAccountResponse, error)
	// GET LIMITS TO B SHOWN TO USER
	GetLimits(ctx context.Context, in *AssetLimitsRequest, opts ...grpc.CallOption) (*LimitsResponse, error)
	// users can see their warehouse balance
	// if they have multiple assets, they would see many records
	GetWarehouseBalance(ctx context.Context, in *BalancesRequest, opts ...grpc.CallOption) (*BalancesResponse, error)
	// RequestForDeposit requests for a new deposit
	RequestForDeposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*DepositResponse, error)
	// list user deposits
	SearchDeposits(ctx context.Context, in *SearchFiatDepositsRequest, opts ...grpc.CallOption) (*SearchFiatDepositsResponse, error)
	// see the details of a transaction
	// refer to Figma TOMAN-WITHDRAW_DETAILS -> Request details
	TransactionDetails(ctx context.Context, in *transaction.Request, opts ...grpc.CallOption) (*TransactionDetailsResponse, error)
	// RequestForWithdraw request for details
	RequestForWithdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*WithdrawResponse, error)
	SearchWithdraw(ctx context.Context, in *SearchFiatWithdrawalsRequest, opts ...grpc.CallOption) (*SearchFiatWithdrawalsResponse, error)
	WithdrawDetails(ctx context.Context, in *transaction.TransactionFilter, opts ...grpc.CallOption) (*TransactionDetailsResponse, error)
}

type userOperationsClient struct {
	cc grpc.ClientConnInterface
}

func NewUserOperationsClient(cc grpc.ClientConnInterface) UserOperationsClient {
	return &userOperationsClient{cc}
}

func (c *userOperationsClient) CreateWarehouseAccount(ctx context.Context, in *UserAccountRequest, opts ...grpc.CallOption) (*UserAccountResponse, error) {
	out := new(UserAccountResponse)
	err := c.cc.Invoke(ctx, "/rabex.api.warehouse.UserOperations/CreateWarehouseAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOperationsClient) CreateTreasuryAccount(ctx context.Context, in *UserAccountRequest, opts ...grpc.CallOption) (*UserAccountResponse, error) {
	out := new(UserAccountResponse)
	err := c.cc.Invoke(ctx, "/rabex.api.warehouse.UserOperations/CreateTreasuryAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOperationsClient) GetLimits(ctx context.Context, in *AssetLimitsRequest, opts ...grpc.CallOption) (*LimitsResponse, error) {
	out := new(LimitsResponse)
	err := c.cc.Invoke(ctx, "/rabex.api.warehouse.UserOperations/GetLimits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOperationsClient) GetWarehouseBalance(ctx context.Context, in *BalancesRequest, opts ...grpc.CallOption) (*BalancesResponse, error) {
	out := new(BalancesResponse)
	err := c.cc.Invoke(ctx, "/rabex.api.warehouse.UserOperations/GetWarehouseBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOperationsClient) RequestForDeposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*DepositResponse, error) {
	out := new(DepositResponse)
	err := c.cc.Invoke(ctx, "/rabex.api.warehouse.UserOperations/RequestForDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOperationsClient) SearchDeposits(ctx context.Context, in *SearchFiatDepositsRequest, opts ...grpc.CallOption) (*SearchFiatDepositsResponse, error) {
	out := new(SearchFiatDepositsResponse)
	err := c.cc.Invoke(ctx, "/rabex.api.warehouse.UserOperations/SearchDeposits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOperationsClient) TransactionDetails(ctx context.Context, in *transaction.Request, opts ...grpc.CallOption) (*TransactionDetailsResponse, error) {
	out := new(TransactionDetailsResponse)
	err := c.cc.Invoke(ctx, "/rabex.api.warehouse.UserOperations/TransactionDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOperationsClient) RequestForWithdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*WithdrawResponse, error) {
	out := new(WithdrawResponse)
	err := c.cc.Invoke(ctx, "/rabex.api.warehouse.UserOperations/RequestForWithdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOperationsClient) SearchWithdraw(ctx context.Context, in *SearchFiatWithdrawalsRequest, opts ...grpc.CallOption) (*SearchFiatWithdrawalsResponse, error) {
	out := new(SearchFiatWithdrawalsResponse)
	err := c.cc.Invoke(ctx, "/rabex.api.warehouse.UserOperations/SearchWithdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOperationsClient) WithdrawDetails(ctx context.Context, in *transaction.TransactionFilter, opts ...grpc.CallOption) (*TransactionDetailsResponse, error) {
	out := new(TransactionDetailsResponse)
	err := c.cc.Invoke(ctx, "/rabex.api.warehouse.UserOperations/WithdrawDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserOperationsServer is the server API for UserOperations service.
// All implementations must embed UnimplementedUserOperationsServer
// for forward compatibility
type UserOperationsServer interface {
	// Warehouse and Account
	CreateWarehouseAccount(context.Context, *UserAccountRequest) (*UserAccountResponse, error)
	// Treasury Account
	CreateTreasuryAccount(context.Context, *UserAccountRequest) (*UserAccountResponse, error)
	// GET LIMITS TO B SHOWN TO USER
	GetLimits(context.Context, *AssetLimitsRequest) (*LimitsResponse, error)
	// users can see their warehouse balance
	// if they have multiple assets, they would see many records
	GetWarehouseBalance(context.Context, *BalancesRequest) (*BalancesResponse, error)
	// RequestForDeposit requests for a new deposit
	RequestForDeposit(context.Context, *DepositRequest) (*DepositResponse, error)
	// list user deposits
	SearchDeposits(context.Context, *SearchFiatDepositsRequest) (*SearchFiatDepositsResponse, error)
	// see the details of a transaction
	// refer to Figma TOMAN-WITHDRAW_DETAILS -> Request details
	TransactionDetails(context.Context, *transaction.Request) (*TransactionDetailsResponse, error)
	// RequestForWithdraw request for details
	RequestForWithdraw(context.Context, *WithdrawRequest) (*WithdrawResponse, error)
	SearchWithdraw(context.Context, *SearchFiatWithdrawalsRequest) (*SearchFiatWithdrawalsResponse, error)
	WithdrawDetails(context.Context, *transaction.TransactionFilter) (*TransactionDetailsResponse, error)
	mustEmbedUnimplementedUserOperationsServer()
}

// UnimplementedUserOperationsServer must be embedded to have forward compatible implementations.
type UnimplementedUserOperationsServer struct {
}

func (UnimplementedUserOperationsServer) CreateWarehouseAccount(context.Context, *UserAccountRequest) (*UserAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWarehouseAccount not implemented")
}
func (UnimplementedUserOperationsServer) CreateTreasuryAccount(context.Context, *UserAccountRequest) (*UserAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTreasuryAccount not implemented")
}
func (UnimplementedUserOperationsServer) GetLimits(context.Context, *AssetLimitsRequest) (*LimitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLimits not implemented")
}
func (UnimplementedUserOperationsServer) GetWarehouseBalance(context.Context, *BalancesRequest) (*BalancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWarehouseBalance not implemented")
}
func (UnimplementedUserOperationsServer) RequestForDeposit(context.Context, *DepositRequest) (*DepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestForDeposit not implemented")
}
func (UnimplementedUserOperationsServer) SearchDeposits(context.Context, *SearchFiatDepositsRequest) (*SearchFiatDepositsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchDeposits not implemented")
}
func (UnimplementedUserOperationsServer) TransactionDetails(context.Context, *transaction.Request) (*TransactionDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransactionDetails not implemented")
}
func (UnimplementedUserOperationsServer) RequestForWithdraw(context.Context, *WithdrawRequest) (*WithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestForWithdraw not implemented")
}
func (UnimplementedUserOperationsServer) SearchWithdraw(context.Context, *SearchFiatWithdrawalsRequest) (*SearchFiatWithdrawalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchWithdraw not implemented")
}
func (UnimplementedUserOperationsServer) WithdrawDetails(context.Context, *transaction.TransactionFilter) (*TransactionDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawDetails not implemented")
}
func (UnimplementedUserOperationsServer) mustEmbedUnimplementedUserOperationsServer() {}

// UnsafeUserOperationsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserOperationsServer will
// result in compilation errors.
type UnsafeUserOperationsServer interface {
	mustEmbedUnimplementedUserOperationsServer()
}

func RegisterUserOperationsServer(s grpc.ServiceRegistrar, srv UserOperationsServer) {
	s.RegisterService(&UserOperations_ServiceDesc, srv)
}

func _UserOperations_CreateWarehouseAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserOperationsServer).CreateWarehouseAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.warehouse.UserOperations/CreateWarehouseAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserOperationsServer).CreateWarehouseAccount(ctx, req.(*UserAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserOperations_CreateTreasuryAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserOperationsServer).CreateTreasuryAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.warehouse.UserOperations/CreateTreasuryAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserOperationsServer).CreateTreasuryAccount(ctx, req.(*UserAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserOperations_GetLimits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetLimitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserOperationsServer).GetLimits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.warehouse.UserOperations/GetLimits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserOperationsServer).GetLimits(ctx, req.(*AssetLimitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserOperations_GetWarehouseBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserOperationsServer).GetWarehouseBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.warehouse.UserOperations/GetWarehouseBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserOperationsServer).GetWarehouseBalance(ctx, req.(*BalancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserOperations_RequestForDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserOperationsServer).RequestForDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.warehouse.UserOperations/RequestForDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserOperationsServer).RequestForDeposit(ctx, req.(*DepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserOperations_SearchDeposits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchFiatDepositsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserOperationsServer).SearchDeposits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.warehouse.UserOperations/SearchDeposits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserOperationsServer).SearchDeposits(ctx, req.(*SearchFiatDepositsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserOperations_TransactionDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(transaction.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserOperationsServer).TransactionDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.warehouse.UserOperations/TransactionDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserOperationsServer).TransactionDetails(ctx, req.(*transaction.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserOperations_RequestForWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserOperationsServer).RequestForWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.warehouse.UserOperations/RequestForWithdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserOperationsServer).RequestForWithdraw(ctx, req.(*WithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserOperations_SearchWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchFiatWithdrawalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserOperationsServer).SearchWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.warehouse.UserOperations/SearchWithdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserOperationsServer).SearchWithdraw(ctx, req.(*SearchFiatWithdrawalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserOperations_WithdrawDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(transaction.TransactionFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserOperationsServer).WithdrawDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.warehouse.UserOperations/WithdrawDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserOperationsServer).WithdrawDetails(ctx, req.(*transaction.TransactionFilter))
	}
	return interceptor(ctx, in, info, handler)
}

// UserOperations_ServiceDesc is the grpc.ServiceDesc for UserOperations service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserOperations_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rabex.api.warehouse.UserOperations",
	HandlerType: (*UserOperationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWarehouseAccount",
			Handler:    _UserOperations_CreateWarehouseAccount_Handler,
		},
		{
			MethodName: "CreateTreasuryAccount",
			Handler:    _UserOperations_CreateTreasuryAccount_Handler,
		},
		{
			MethodName: "GetLimits",
			Handler:    _UserOperations_GetLimits_Handler,
		},
		{
			MethodName: "GetWarehouseBalance",
			Handler:    _UserOperations_GetWarehouseBalance_Handler,
		},
		{
			MethodName: "RequestForDeposit",
			Handler:    _UserOperations_RequestForDeposit_Handler,
		},
		{
			MethodName: "SearchDeposits",
			Handler:    _UserOperations_SearchDeposits_Handler,
		},
		{
			MethodName: "TransactionDetails",
			Handler:    _UserOperations_TransactionDetails_Handler,
		},
		{
			MethodName: "RequestForWithdraw",
			Handler:    _UserOperations_RequestForWithdraw_Handler,
		},
		{
			MethodName: "SearchWithdraw",
			Handler:    _UserOperations_SearchWithdraw_Handler,
		},
		{
			MethodName: "WithdrawDetails",
			Handler:    _UserOperations_WithdrawDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "warehouse/user.service.proto",
}
