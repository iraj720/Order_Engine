// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: market/services.proto

package market

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MarketClient is the client API for Market service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketClient interface {
	MarketPriceSanity(ctx context.Context, in *MarketPriceSanityRequest, opts ...grpc.CallOption) (*MarketPriceSanityResponse, error)
	// UpdateMarket provides access for admins to update a market settings
	UpdateMarket(ctx context.Context, in *UpdateMarketRequest, opts ...grpc.CallOption) (*UpdateMarketResponse, error)
	// ListMarkets provide access to list of markets by pagination
	ListMarkets(ctx context.Context, in *ListMarketsRequest, opts ...grpc.CallOption) (*ListMarketsResponse, error)
	// GetMarket fetches one pair/symbol by id
	GetMarket(ctx context.Context, in *GetMarketRequest, opts ...grpc.CallOption) (*GetMarketResponse, error)
	// GetOtcMarketStatus returns the market sell/buy status and BTCUSDT price
	GetOtcMarketStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetOtcMarketStatusResponse, error)
	// UpdateOtcMarketStatus provide access to update OtcMarket buy/sell status
	UpdateOtcMarketStatus(ctx context.Context, in *UpdateOtcMarketStatusRequest, opts ...grpc.CallOption) (*UpdateOtcMarketStatusResponse, error)
	// AddMarket provide access to add market and its settings
	AddMarket(ctx context.Context, in *AddMarketRequest, opts ...grpc.CallOption) (*AddMarketResponse, error)
	// GetMarketsPrice returns last price of requested markets
	GetMarketsPrice(ctx context.Context, in *GetMarketsPriceRequest, opts ...grpc.CallOption) (*GetMarketsPriceResponse, error)
	// GetMarketCharts returns KLine charts for selected market
	GetMarketCharts(ctx context.Context, in *GetMarketChartsRequest, opts ...grpc.CallOption) (*GetMarketChartsResponse, error)
	UpdateBaseSelector(ctx context.Context, in *UpdateBaseSelectorRequest, opts ...grpc.CallOption) (*MarketSelectors, error)
	UpdateQuoteSelector(ctx context.Context, in *UpdateQuoteSelectorRequest, opts ...grpc.CallOption) (*MarketSelectors, error)
	GetAllSelectors(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*MarketSelectorsGetResponse, error)
	UpdateSelector(ctx context.Context, in *UpdateSelectorRequest, opts ...grpc.CallOption) (*MarketSelectorsResponse, error)
	UpdateSubSelector(ctx context.Context, in *UpdateSubSelectorRequest, opts ...grpc.CallOption) (*MarketSelectorsResponse, error)
	Kline(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Market_KlineClient, error)
	MarketRefPrice(ctx context.Context, in *MarketRefPriceRequest, opts ...grpc.CallOption) (*MarketRefPriceResponse, error)
	TickerHealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TickerHealthCheckResponse, error)
}

type marketClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketClient(cc grpc.ClientConnInterface) MarketClient {
	return &marketClient{cc}
}

func (c *marketClient) MarketPriceSanity(ctx context.Context, in *MarketPriceSanityRequest, opts ...grpc.CallOption) (*MarketPriceSanityResponse, error) {
	out := new(MarketPriceSanityResponse)
	err := c.cc.Invoke(ctx, "/rabex.api.market.Market/MarketPriceSanity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) UpdateMarket(ctx context.Context, in *UpdateMarketRequest, opts ...grpc.CallOption) (*UpdateMarketResponse, error) {
	out := new(UpdateMarketResponse)
	err := c.cc.Invoke(ctx, "/rabex.api.market.Market/UpdateMarket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) ListMarkets(ctx context.Context, in *ListMarketsRequest, opts ...grpc.CallOption) (*ListMarketsResponse, error) {
	out := new(ListMarketsResponse)
	err := c.cc.Invoke(ctx, "/rabex.api.market.Market/ListMarkets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) GetMarket(ctx context.Context, in *GetMarketRequest, opts ...grpc.CallOption) (*GetMarketResponse, error) {
	out := new(GetMarketResponse)
	err := c.cc.Invoke(ctx, "/rabex.api.market.Market/GetMarket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) GetOtcMarketStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetOtcMarketStatusResponse, error) {
	out := new(GetOtcMarketStatusResponse)
	err := c.cc.Invoke(ctx, "/rabex.api.market.Market/GetOtcMarketStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) UpdateOtcMarketStatus(ctx context.Context, in *UpdateOtcMarketStatusRequest, opts ...grpc.CallOption) (*UpdateOtcMarketStatusResponse, error) {
	out := new(UpdateOtcMarketStatusResponse)
	err := c.cc.Invoke(ctx, "/rabex.api.market.Market/UpdateOtcMarketStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) AddMarket(ctx context.Context, in *AddMarketRequest, opts ...grpc.CallOption) (*AddMarketResponse, error) {
	out := new(AddMarketResponse)
	err := c.cc.Invoke(ctx, "/rabex.api.market.Market/AddMarket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) GetMarketsPrice(ctx context.Context, in *GetMarketsPriceRequest, opts ...grpc.CallOption) (*GetMarketsPriceResponse, error) {
	out := new(GetMarketsPriceResponse)
	err := c.cc.Invoke(ctx, "/rabex.api.market.Market/GetMarketsPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) GetMarketCharts(ctx context.Context, in *GetMarketChartsRequest, opts ...grpc.CallOption) (*GetMarketChartsResponse, error) {
	out := new(GetMarketChartsResponse)
	err := c.cc.Invoke(ctx, "/rabex.api.market.Market/GetMarketCharts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) UpdateBaseSelector(ctx context.Context, in *UpdateBaseSelectorRequest, opts ...grpc.CallOption) (*MarketSelectors, error) {
	out := new(MarketSelectors)
	err := c.cc.Invoke(ctx, "/rabex.api.market.Market/UpdateBaseSelector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) UpdateQuoteSelector(ctx context.Context, in *UpdateQuoteSelectorRequest, opts ...grpc.CallOption) (*MarketSelectors, error) {
	out := new(MarketSelectors)
	err := c.cc.Invoke(ctx, "/rabex.api.market.Market/UpdateQuoteSelector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) GetAllSelectors(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*MarketSelectorsGetResponse, error) {
	out := new(MarketSelectorsGetResponse)
	err := c.cc.Invoke(ctx, "/rabex.api.market.Market/GetAllSelectors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) UpdateSelector(ctx context.Context, in *UpdateSelectorRequest, opts ...grpc.CallOption) (*MarketSelectorsResponse, error) {
	out := new(MarketSelectorsResponse)
	err := c.cc.Invoke(ctx, "/rabex.api.market.Market/UpdateSelector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) UpdateSubSelector(ctx context.Context, in *UpdateSubSelectorRequest, opts ...grpc.CallOption) (*MarketSelectorsResponse, error) {
	out := new(MarketSelectorsResponse)
	err := c.cc.Invoke(ctx, "/rabex.api.market.Market/UpdateSubSelector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) Kline(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Market_KlineClient, error) {
	stream, err := c.cc.NewStream(ctx, &Market_ServiceDesc.Streams[0], "/rabex.api.market.Market/Kline", opts...)
	if err != nil {
		return nil, err
	}
	x := &marketKlineClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Market_KlineClient interface {
	Recv() (*MarketKlineStream, error)
	grpc.ClientStream
}

type marketKlineClient struct {
	grpc.ClientStream
}

func (x *marketKlineClient) Recv() (*MarketKlineStream, error) {
	m := new(MarketKlineStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *marketClient) MarketRefPrice(ctx context.Context, in *MarketRefPriceRequest, opts ...grpc.CallOption) (*MarketRefPriceResponse, error) {
	out := new(MarketRefPriceResponse)
	err := c.cc.Invoke(ctx, "/rabex.api.market.Market/MarketRefPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) TickerHealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TickerHealthCheckResponse, error) {
	out := new(TickerHealthCheckResponse)
	err := c.cc.Invoke(ctx, "/rabex.api.market.Market/TickerHealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketServer is the server API for Market service.
// All implementations must embed UnimplementedMarketServer
// for forward compatibility
type MarketServer interface {
	MarketPriceSanity(context.Context, *MarketPriceSanityRequest) (*MarketPriceSanityResponse, error)
	// UpdateMarket provides access for admins to update a market settings
	UpdateMarket(context.Context, *UpdateMarketRequest) (*UpdateMarketResponse, error)
	// ListMarkets provide access to list of markets by pagination
	ListMarkets(context.Context, *ListMarketsRequest) (*ListMarketsResponse, error)
	// GetMarket fetches one pair/symbol by id
	GetMarket(context.Context, *GetMarketRequest) (*GetMarketResponse, error)
	// GetOtcMarketStatus returns the market sell/buy status and BTCUSDT price
	GetOtcMarketStatus(context.Context, *emptypb.Empty) (*GetOtcMarketStatusResponse, error)
	// UpdateOtcMarketStatus provide access to update OtcMarket buy/sell status
	UpdateOtcMarketStatus(context.Context, *UpdateOtcMarketStatusRequest) (*UpdateOtcMarketStatusResponse, error)
	// AddMarket provide access to add market and its settings
	AddMarket(context.Context, *AddMarketRequest) (*AddMarketResponse, error)
	// GetMarketsPrice returns last price of requested markets
	GetMarketsPrice(context.Context, *GetMarketsPriceRequest) (*GetMarketsPriceResponse, error)
	// GetMarketCharts returns KLine charts for selected market
	GetMarketCharts(context.Context, *GetMarketChartsRequest) (*GetMarketChartsResponse, error)
	UpdateBaseSelector(context.Context, *UpdateBaseSelectorRequest) (*MarketSelectors, error)
	UpdateQuoteSelector(context.Context, *UpdateQuoteSelectorRequest) (*MarketSelectors, error)
	GetAllSelectors(context.Context, *emptypb.Empty) (*MarketSelectorsGetResponse, error)
	UpdateSelector(context.Context, *UpdateSelectorRequest) (*MarketSelectorsResponse, error)
	UpdateSubSelector(context.Context, *UpdateSubSelectorRequest) (*MarketSelectorsResponse, error)
	Kline(*emptypb.Empty, Market_KlineServer) error
	MarketRefPrice(context.Context, *MarketRefPriceRequest) (*MarketRefPriceResponse, error)
	TickerHealthCheck(context.Context, *emptypb.Empty) (*TickerHealthCheckResponse, error)
	mustEmbedUnimplementedMarketServer()
}

// UnimplementedMarketServer must be embedded to have forward compatible implementations.
type UnimplementedMarketServer struct {
}

func (UnimplementedMarketServer) MarketPriceSanity(context.Context, *MarketPriceSanityRequest) (*MarketPriceSanityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarketPriceSanity not implemented")
}
func (UnimplementedMarketServer) UpdateMarket(context.Context, *UpdateMarketRequest) (*UpdateMarketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMarket not implemented")
}
func (UnimplementedMarketServer) ListMarkets(context.Context, *ListMarketsRequest) (*ListMarketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMarkets not implemented")
}
func (UnimplementedMarketServer) GetMarket(context.Context, *GetMarketRequest) (*GetMarketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMarket not implemented")
}
func (UnimplementedMarketServer) GetOtcMarketStatus(context.Context, *emptypb.Empty) (*GetOtcMarketStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOtcMarketStatus not implemented")
}
func (UnimplementedMarketServer) UpdateOtcMarketStatus(context.Context, *UpdateOtcMarketStatusRequest) (*UpdateOtcMarketStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOtcMarketStatus not implemented")
}
func (UnimplementedMarketServer) AddMarket(context.Context, *AddMarketRequest) (*AddMarketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMarket not implemented")
}
func (UnimplementedMarketServer) GetMarketsPrice(context.Context, *GetMarketsPriceRequest) (*GetMarketsPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMarketsPrice not implemented")
}
func (UnimplementedMarketServer) GetMarketCharts(context.Context, *GetMarketChartsRequest) (*GetMarketChartsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMarketCharts not implemented")
}
func (UnimplementedMarketServer) UpdateBaseSelector(context.Context, *UpdateBaseSelectorRequest) (*MarketSelectors, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBaseSelector not implemented")
}
func (UnimplementedMarketServer) UpdateQuoteSelector(context.Context, *UpdateQuoteSelectorRequest) (*MarketSelectors, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuoteSelector not implemented")
}
func (UnimplementedMarketServer) GetAllSelectors(context.Context, *emptypb.Empty) (*MarketSelectorsGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllSelectors not implemented")
}
func (UnimplementedMarketServer) UpdateSelector(context.Context, *UpdateSelectorRequest) (*MarketSelectorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSelector not implemented")
}
func (UnimplementedMarketServer) UpdateSubSelector(context.Context, *UpdateSubSelectorRequest) (*MarketSelectorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSubSelector not implemented")
}
func (UnimplementedMarketServer) Kline(*emptypb.Empty, Market_KlineServer) error {
	return status.Errorf(codes.Unimplemented, "method Kline not implemented")
}
func (UnimplementedMarketServer) MarketRefPrice(context.Context, *MarketRefPriceRequest) (*MarketRefPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarketRefPrice not implemented")
}
func (UnimplementedMarketServer) TickerHealthCheck(context.Context, *emptypb.Empty) (*TickerHealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TickerHealthCheck not implemented")
}
func (UnimplementedMarketServer) mustEmbedUnimplementedMarketServer() {}

// UnsafeMarketServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketServer will
// result in compilation errors.
type UnsafeMarketServer interface {
	mustEmbedUnimplementedMarketServer()
}

func RegisterMarketServer(s grpc.ServiceRegistrar, srv MarketServer) {
	s.RegisterService(&Market_ServiceDesc, srv)
}

func _Market_MarketPriceSanity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketPriceSanityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).MarketPriceSanity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.market.Market/MarketPriceSanity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).MarketPriceSanity(ctx, req.(*MarketPriceSanityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_UpdateMarket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMarketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).UpdateMarket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.market.Market/UpdateMarket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).UpdateMarket(ctx, req.(*UpdateMarketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_ListMarkets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMarketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).ListMarkets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.market.Market/ListMarkets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).ListMarkets(ctx, req.(*ListMarketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_GetMarket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMarketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetMarket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.market.Market/GetMarket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetMarket(ctx, req.(*GetMarketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_GetOtcMarketStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetOtcMarketStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.market.Market/GetOtcMarketStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetOtcMarketStatus(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_UpdateOtcMarketStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOtcMarketStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).UpdateOtcMarketStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.market.Market/UpdateOtcMarketStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).UpdateOtcMarketStatus(ctx, req.(*UpdateOtcMarketStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_AddMarket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMarketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).AddMarket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.market.Market/AddMarket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).AddMarket(ctx, req.(*AddMarketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_GetMarketsPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMarketsPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetMarketsPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.market.Market/GetMarketsPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetMarketsPrice(ctx, req.(*GetMarketsPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_GetMarketCharts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMarketChartsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetMarketCharts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.market.Market/GetMarketCharts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetMarketCharts(ctx, req.(*GetMarketChartsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_UpdateBaseSelector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBaseSelectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).UpdateBaseSelector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.market.Market/UpdateBaseSelector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).UpdateBaseSelector(ctx, req.(*UpdateBaseSelectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_UpdateQuoteSelector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQuoteSelectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).UpdateQuoteSelector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.market.Market/UpdateQuoteSelector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).UpdateQuoteSelector(ctx, req.(*UpdateQuoteSelectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_GetAllSelectors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetAllSelectors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.market.Market/GetAllSelectors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetAllSelectors(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_UpdateSelector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSelectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).UpdateSelector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.market.Market/UpdateSelector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).UpdateSelector(ctx, req.(*UpdateSelectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_UpdateSubSelector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSubSelectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).UpdateSubSelector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.market.Market/UpdateSubSelector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).UpdateSubSelector(ctx, req.(*UpdateSubSelectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_Kline_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MarketServer).Kline(m, &marketKlineServer{stream})
}

type Market_KlineServer interface {
	Send(*MarketKlineStream) error
	grpc.ServerStream
}

type marketKlineServer struct {
	grpc.ServerStream
}

func (x *marketKlineServer) Send(m *MarketKlineStream) error {
	return x.ServerStream.SendMsg(m)
}

func _Market_MarketRefPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketRefPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).MarketRefPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.market.Market/MarketRefPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).MarketRefPrice(ctx, req.(*MarketRefPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_TickerHealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).TickerHealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rabex.api.market.Market/TickerHealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).TickerHealthCheck(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Market_ServiceDesc is the grpc.ServiceDesc for Market service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Market_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rabex.api.market.Market",
	HandlerType: (*MarketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MarketPriceSanity",
			Handler:    _Market_MarketPriceSanity_Handler,
		},
		{
			MethodName: "UpdateMarket",
			Handler:    _Market_UpdateMarket_Handler,
		},
		{
			MethodName: "ListMarkets",
			Handler:    _Market_ListMarkets_Handler,
		},
		{
			MethodName: "GetMarket",
			Handler:    _Market_GetMarket_Handler,
		},
		{
			MethodName: "GetOtcMarketStatus",
			Handler:    _Market_GetOtcMarketStatus_Handler,
		},
		{
			MethodName: "UpdateOtcMarketStatus",
			Handler:    _Market_UpdateOtcMarketStatus_Handler,
		},
		{
			MethodName: "AddMarket",
			Handler:    _Market_AddMarket_Handler,
		},
		{
			MethodName: "GetMarketsPrice",
			Handler:    _Market_GetMarketsPrice_Handler,
		},
		{
			MethodName: "GetMarketCharts",
			Handler:    _Market_GetMarketCharts_Handler,
		},
		{
			MethodName: "UpdateBaseSelector",
			Handler:    _Market_UpdateBaseSelector_Handler,
		},
		{
			MethodName: "UpdateQuoteSelector",
			Handler:    _Market_UpdateQuoteSelector_Handler,
		},
		{
			MethodName: "GetAllSelectors",
			Handler:    _Market_GetAllSelectors_Handler,
		},
		{
			MethodName: "UpdateSelector",
			Handler:    _Market_UpdateSelector_Handler,
		},
		{
			MethodName: "UpdateSubSelector",
			Handler:    _Market_UpdateSubSelector_Handler,
		},
		{
			MethodName: "MarketRefPrice",
			Handler:    _Market_MarketRefPrice_Handler,
		},
		{
			MethodName: "TickerHealthCheck",
			Handler:    _Market_TickerHealthCheck_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Kline",
			Handler:       _Market_Kline_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "market/services.proto",
}
