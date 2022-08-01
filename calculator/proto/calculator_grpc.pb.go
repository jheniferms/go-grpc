// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.3
// source: calculator.proto

package proto

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

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	Primes(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (CalculatorService_PrimesClient, error)
	Avg(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AvgClient, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Primes(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (CalculatorService_PrimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[0], "/calculator.CalculatorService/Primes", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimesClient interface {
	Recv() (*PrimeResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimesClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimesClient) Recv() (*PrimeResponse, error) {
	m := new(PrimeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) Avg(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AvgClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[1], "/calculator.CalculatorService/Avg", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceAvgClient{stream}
	return x, nil
}

type CalculatorService_AvgClient interface {
	Send(*AvgRequest) error
	CloseAndRecv() (*AvgResponse, error)
	grpc.ClientStream
}

type calculatorServiceAvgClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceAvgClient) Send(m *AvgRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceAvgClient) CloseAndRecv() (*AvgResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AvgResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
// All implementations must embed UnimplementedCalculatorServiceServer
// for forward compatibility
type CalculatorServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	Primes(*PrimeRequest, CalculatorService_PrimesServer) error
	Avg(CalculatorService_AvgServer) error
	mustEmbedUnimplementedCalculatorServiceServer()
}

// UnimplementedCalculatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (UnimplementedCalculatorServiceServer) Sum(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedCalculatorServiceServer) Primes(*PrimeRequest, CalculatorService_PrimesServer) error {
	return status.Errorf(codes.Unimplemented, "method Primes not implemented")
}
func (UnimplementedCalculatorServiceServer) Avg(CalculatorService_AvgServer) error {
	return status.Errorf(codes.Unimplemented, "method Avg not implemented")
}
func (UnimplementedCalculatorServiceServer) mustEmbedUnimplementedCalculatorServiceServer() {}

// UnsafeCalculatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServiceServer will
// result in compilation errors.
type UnsafeCalculatorServiceServer interface {
	mustEmbedUnimplementedCalculatorServiceServer()
}

func RegisterCalculatorServiceServer(s grpc.ServiceRegistrar, srv CalculatorServiceServer) {
	s.RegisterService(&CalculatorService_ServiceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Primes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).Primes(m, &calculatorServicePrimesServer{stream})
}

type CalculatorService_PrimesServer interface {
	Send(*PrimeResponse) error
	grpc.ServerStream
}

type calculatorServicePrimesServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimesServer) Send(m *PrimeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_Avg_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).Avg(&calculatorServiceAvgServer{stream})
}

type CalculatorService_AvgServer interface {
	SendAndClose(*AvgResponse) error
	Recv() (*AvgRequest, error)
	grpc.ServerStream
}

type calculatorServiceAvgServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceAvgServer) SendAndClose(m *AvgResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceAvgServer) Recv() (*AvgRequest, error) {
	m := new(AvgRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorService_ServiceDesc is the grpc.ServiceDesc for CalculatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Primes",
			Handler:       _CalculatorService_Primes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Avg",
			Handler:       _CalculatorService_Avg_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "calculator.proto",
}
