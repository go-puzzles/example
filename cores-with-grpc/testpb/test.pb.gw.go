// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: test.proto

/*
Package testpb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package testpb

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

func request_ExampleHelloService_SayHello_0(ctx context.Context, marshaler runtime.Marshaler, client ExampleHelloServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq HelloRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.SayHello(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ExampleHelloService_SayHello_0(ctx context.Context, marshaler runtime.Marshaler, server ExampleHelloServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq HelloRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.SayHello(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterExampleHelloServiceHandlerServer registers the http handlers for service ExampleHelloService to "mux".
// UnaryRPC     :call ExampleHelloServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterExampleHelloServiceHandlerFromEndpoint instead.
func RegisterExampleHelloServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server ExampleHelloServiceServer) error {

	mux.Handle("POST", pattern_ExampleHelloService_SayHello_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/testpb.ExampleHelloService/SayHello", runtime.WithHTTPPathPattern("/hello"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ExampleHelloService_SayHello_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ExampleHelloService_SayHello_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterExampleHelloServiceHandlerFromEndpoint is same as RegisterExampleHelloServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterExampleHelloServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
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

	return RegisterExampleHelloServiceHandler(ctx, mux, conn)
}

// RegisterExampleHelloServiceHandler registers the http handlers for service ExampleHelloService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterExampleHelloServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterExampleHelloServiceHandlerClient(ctx, mux, NewExampleHelloServiceClient(conn))
}

// RegisterExampleHelloServiceHandlerClient registers the http handlers for service ExampleHelloService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "ExampleHelloServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "ExampleHelloServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "ExampleHelloServiceClient" to call the correct interceptors.
func RegisterExampleHelloServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client ExampleHelloServiceClient) error {

	mux.Handle("POST", pattern_ExampleHelloService_SayHello_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/testpb.ExampleHelloService/SayHello", runtime.WithHTTPPathPattern("/hello"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ExampleHelloService_SayHello_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ExampleHelloService_SayHello_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_ExampleHelloService_SayHello_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"hello"}, ""))
)

var (
	forward_ExampleHelloService_SayHello_0 = runtime.ForwardResponseMessage
)
