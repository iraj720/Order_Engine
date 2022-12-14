// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: accounting/treasury-statement/services.proto

/*
Package tStatement is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package tStatement

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

var (
	filter_TreasuryStatementAPI_GetBalanceAll_0 = &utilities.DoubleArray{Encoding: map[string]int{"treasury_id": 0}, Base: []int{1, 1, 0}, Check: []int{0, 1, 2}}
)

func request_TreasuryStatementAPI_GetBalanceAll_0(ctx context.Context, marshaler runtime.Marshaler, client TreasuryStatementAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetBalanceAllRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["treasury_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "treasury_id")
	}

	protoReq.TreasuryId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "treasury_id", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_TreasuryStatementAPI_GetBalanceAll_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetBalanceAll(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_TreasuryStatementAPI_GetBalanceAll_0(ctx context.Context, marshaler runtime.Marshaler, server TreasuryStatementAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetBalanceAllRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["treasury_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "treasury_id")
	}

	protoReq.TreasuryId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "treasury_id", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_TreasuryStatementAPI_GetBalanceAll_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetBalanceAll(ctx, &protoReq)
	return msg, metadata, err

}

func request_TreasuryStatementAPI_GetBalanceByAsset_0(ctx context.Context, marshaler runtime.Marshaler, client TreasuryStatementAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetBalanceByAssetRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["treasury_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "treasury_id")
	}

	protoReq.TreasuryId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "treasury_id", err)
	}

	val, ok = pathParams["asset_network"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "asset_network")
	}

	protoReq.AssetNetwork, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "asset_network", err)
	}

	msg, err := client.GetBalanceByAsset(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_TreasuryStatementAPI_GetBalanceByAsset_0(ctx context.Context, marshaler runtime.Marshaler, server TreasuryStatementAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetBalanceByAssetRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["treasury_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "treasury_id")
	}

	protoReq.TreasuryId, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "treasury_id", err)
	}

	val, ok = pathParams["asset_network"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "asset_network")
	}

	protoReq.AssetNetwork, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "asset_network", err)
	}

	msg, err := server.GetBalanceByAsset(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterTreasuryStatementAPIHandlerServer registers the http handlers for service TreasuryStatementAPI to "mux".
// UnaryRPC     :call TreasuryStatementAPIServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterTreasuryStatementAPIHandlerFromEndpoint instead.
func RegisterTreasuryStatementAPIHandlerServer(ctx context.Context, mux *runtime.ServeMux, server TreasuryStatementAPIServer) error {

	mux.Handle("GET", pattern_TreasuryStatementAPI_GetBalanceAll_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/tStatement.TreasuryStatementAPI/GetBalanceAll", runtime.WithHTTPPathPattern("/api/v1/accounting/treasury/statement/{treasury_id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_TreasuryStatementAPI_GetBalanceAll_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_TreasuryStatementAPI_GetBalanceAll_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_TreasuryStatementAPI_GetBalanceByAsset_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/tStatement.TreasuryStatementAPI/GetBalanceByAsset", runtime.WithHTTPPathPattern("/api/v1/accounting/treasury/statement/{treasury_id}/{asset_network}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_TreasuryStatementAPI_GetBalanceByAsset_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_TreasuryStatementAPI_GetBalanceByAsset_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterTreasuryStatementAPIHandlerFromEndpoint is same as RegisterTreasuryStatementAPIHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterTreasuryStatementAPIHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterTreasuryStatementAPIHandler(ctx, mux, conn)
}

// RegisterTreasuryStatementAPIHandler registers the http handlers for service TreasuryStatementAPI to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterTreasuryStatementAPIHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterTreasuryStatementAPIHandlerClient(ctx, mux, NewTreasuryStatementAPIClient(conn))
}

// RegisterTreasuryStatementAPIHandlerClient registers the http handlers for service TreasuryStatementAPI
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "TreasuryStatementAPIClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "TreasuryStatementAPIClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "TreasuryStatementAPIClient" to call the correct interceptors.
func RegisterTreasuryStatementAPIHandlerClient(ctx context.Context, mux *runtime.ServeMux, client TreasuryStatementAPIClient) error {

	mux.Handle("GET", pattern_TreasuryStatementAPI_GetBalanceAll_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/tStatement.TreasuryStatementAPI/GetBalanceAll", runtime.WithHTTPPathPattern("/api/v1/accounting/treasury/statement/{treasury_id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_TreasuryStatementAPI_GetBalanceAll_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_TreasuryStatementAPI_GetBalanceAll_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_TreasuryStatementAPI_GetBalanceByAsset_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/tStatement.TreasuryStatementAPI/GetBalanceByAsset", runtime.WithHTTPPathPattern("/api/v1/accounting/treasury/statement/{treasury_id}/{asset_network}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_TreasuryStatementAPI_GetBalanceByAsset_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_TreasuryStatementAPI_GetBalanceByAsset_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_TreasuryStatementAPI_GetBalanceAll_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4, 1, 0, 4, 1, 5, 5}, []string{"api", "v1", "accounting", "treasury", "statement", "treasury_id"}, ""))

	pattern_TreasuryStatementAPI_GetBalanceByAsset_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4, 1, 0, 4, 1, 5, 5, 1, 0, 4, 1, 5, 6}, []string{"api", "v1", "accounting", "treasury", "statement", "treasury_id", "asset_network"}, ""))
)

var (
	forward_TreasuryStatementAPI_GetBalanceAll_0 = runtime.ForwardResponseMessage

	forward_TreasuryStatementAPI_GetBalanceByAsset_0 = runtime.ForwardResponseMessage
)
