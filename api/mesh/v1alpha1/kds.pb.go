// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mesh/v1alpha1/kds.proto

package v1alpha1

import (
	context "context"
	fmt "fmt"
	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type KumaResource struct {
	Meta                 *KumaResource_Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Spec                 *any.Any           `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *KumaResource) Reset()         { *m = KumaResource{} }
func (m *KumaResource) String() string { return proto.CompactTextString(m) }
func (*KumaResource) ProtoMessage()    {}
func (*KumaResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c4a288324484b61, []int{0}
}

func (m *KumaResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KumaResource.Unmarshal(m, b)
}
func (m *KumaResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KumaResource.Marshal(b, m, deterministic)
}
func (m *KumaResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KumaResource.Merge(m, src)
}
func (m *KumaResource) XXX_Size() int {
	return xxx_messageInfo_KumaResource.Size(m)
}
func (m *KumaResource) XXX_DiscardUnknown() {
	xxx_messageInfo_KumaResource.DiscardUnknown(m)
}

var xxx_messageInfo_KumaResource proto.InternalMessageInfo

func (m *KumaResource) GetMeta() *KumaResource_Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *KumaResource) GetSpec() *any.Any {
	if m != nil {
		return m.Spec
	}
	return nil
}

type KumaResource_Meta struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Mesh                 string               `protobuf:"bytes,2,opt,name=mesh,proto3" json:"mesh,omitempty"`
	CreationTime         *timestamp.Timestamp `protobuf:"bytes,3,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty"`
	ModificationTime     *timestamp.Timestamp `protobuf:"bytes,4,opt,name=modification_time,json=modificationTime,proto3" json:"modification_time,omitempty"`
	Version              string               `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *KumaResource_Meta) Reset()         { *m = KumaResource_Meta{} }
func (m *KumaResource_Meta) String() string { return proto.CompactTextString(m) }
func (*KumaResource_Meta) ProtoMessage()    {}
func (*KumaResource_Meta) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c4a288324484b61, []int{0, 0}
}

func (m *KumaResource_Meta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KumaResource_Meta.Unmarshal(m, b)
}
func (m *KumaResource_Meta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KumaResource_Meta.Marshal(b, m, deterministic)
}
func (m *KumaResource_Meta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KumaResource_Meta.Merge(m, src)
}
func (m *KumaResource_Meta) XXX_Size() int {
	return xxx_messageInfo_KumaResource_Meta.Size(m)
}
func (m *KumaResource_Meta) XXX_DiscardUnknown() {
	xxx_messageInfo_KumaResource_Meta.DiscardUnknown(m)
}

var xxx_messageInfo_KumaResource_Meta proto.InternalMessageInfo

func (m *KumaResource_Meta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KumaResource_Meta) GetMesh() string {
	if m != nil {
		return m.Mesh
	}
	return ""
}

func (m *KumaResource_Meta) GetCreationTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *KumaResource_Meta) GetModificationTime() *timestamp.Timestamp {
	if m != nil {
		return m.ModificationTime
	}
	return nil
}

func (m *KumaResource_Meta) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*KumaResource)(nil), "kuma.mesh.v1alpha1.KumaResource")
	proto.RegisterType((*KumaResource_Meta)(nil), "kuma.mesh.v1alpha1.KumaResource.Meta")
}

func init() { proto.RegisterFile("mesh/v1alpha1/kds.proto", fileDescriptor_5c4a288324484b61) }

var fileDescriptor_5c4a288324484b61 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xcb, 0x4e, 0xf2, 0x40,
	0x14, 0x4e, 0xf9, 0xfb, 0x5f, 0x98, 0x1f, 0x13, 0x1d, 0x49, 0xac, 0x8d, 0x11, 0xe3, 0x25, 0x61,
	0x35, 0x95, 0xba, 0x72, 0x65, 0x34, 0x44, 0x17, 0xc6, 0x4d, 0xd1, 0x0d, 0x1b, 0x33, 0x0c, 0x07,
	0x98, 0xc0, 0x74, 0x6a, 0x67, 0xda, 0x84, 0x97, 0xf2, 0x39, 0x5c, 0xfb, 0x44, 0xa6, 0x07, 0x6a,
	0x2a, 0x44, 0xdd, 0xb8, 0x3b, 0x3d, 0xdf, 0xed, 0xf4, 0x6b, 0xc9, 0x8e, 0x02, 0x33, 0x09, 0xf2,
	0x0e, 0x9f, 0x25, 0x13, 0xde, 0x09, 0xa6, 0x43, 0xc3, 0x92, 0x54, 0x5b, 0x4d, 0xe9, 0x34, 0x53,
	0x9c, 0x15, 0x28, 0x2b, 0x51, 0xbf, 0x35, 0xd6, 0x7a, 0x3c, 0x83, 0x00, 0x19, 0x83, 0x6c, 0x14,
	0x58, 0xa9, 0xc0, 0x58, 0xae, 0x92, 0x85, 0xc8, 0xdf, 0x83, 0x38, 0xd7, 0xf3, 0x80, 0x27, 0x32,
	0xc8, 0xc3, 0x60, 0x28, 0x8d, 0xd0, 0x39, 0xa4, 0xf3, 0x25, 0xba, 0xbb, 0x2a, 0xe7, 0xf1, 0x12,
	0x3a, 0x7c, 0xa9, 0x91, 0xc6, 0x6d, 0xa6, 0x78, 0x04, 0x46, 0x67, 0xa9, 0x00, 0x7a, 0x4e, 0x5c,
	0x05, 0x96, 0x7b, 0xce, 0x81, 0xd3, 0xfe, 0x1f, 0x9e, 0xb0, 0xf5, 0x6b, 0x58, 0x95, 0xcf, 0xee,
	0xc0, 0xf2, 0x08, 0x25, 0xb4, 0x4d, 0x5c, 0x93, 0x80, 0xf0, 0x6a, 0x28, 0x6d, 0xb2, 0x45, 0x2a,
	0x2b, 0x53, 0xd9, 0x65, 0x3c, 0x8f, 0x90, 0xe1, 0xbf, 0x3a, 0xc4, 0x2d, 0x84, 0x94, 0x12, 0x37,
	0xe6, 0x0a, 0x30, 0xad, 0x1e, 0xe1, 0x5c, 0xec, 0x8a, 0x3c, 0xb4, 0xa9, 0x47, 0x38, 0xd3, 0x0b,
	0xb2, 0x21, 0x52, 0xe0, 0x56, 0xea, 0xf8, 0xb1, 0x78, 0x77, 0xef, 0x17, 0x66, 0xf8, 0x6b, 0x19,
	0xf7, 0x65, 0x31, 0x51, 0xa3, 0x14, 0x14, 0x2b, 0x7a, 0x43, 0xb6, 0x94, 0x1e, 0xca, 0x91, 0x14,
	0x15, 0x13, 0xf7, 0x5b, 0x93, 0xcd, 0xaa, 0x08, 0x8d, 0x3c, 0xf2, 0x37, 0x87, 0xd4, 0x48, 0x1d,
	0x7b, 0xbf, 0xf1, 0xc0, 0xf2, 0x31, 0x7c, 0xae, 0x91, 0x66, 0x51, 0x4d, 0xb7, 0x6c, 0xbf, 0x07,
	0x69, 0x2e, 0x05, 0x50, 0x41, 0x68, 0x17, 0x66, 0x96, 0x57, 0x7b, 0x33, 0xf4, 0x88, 0xe1, 0x37,
	0x63, 0x3c, 0x91, 0x2c, 0x0f, 0x19, 0x32, 0xde, 0xa5, 0x11, 0x3c, 0x65, 0x60, 0xac, 0x7f, 0xfc,
	0x35, 0xc9, 0x24, 0x3a, 0x36, 0xd0, 0x76, 0x4e, 0x1d, 0xda, 0x27, 0xdb, 0x3d, 0x9b, 0x02, 0x57,
	0x1f, 0x53, 0xf6, 0x57, 0x0c, 0x56, 0x03, 0x5a, 0x9f, 0xe2, 0x15, 0xef, 0x07, 0x42, 0xaf, 0xc1,
	0x8a, 0xc9, 0xcf, 0x5a, 0x5f, 0x91, 0xfe, 0xbf, 0xf2, 0x9f, 0x1a, 0xfc, 0xc1, 0xf2, 0xcf, 0xde,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x0c, 0x25, 0xe8, 0x17, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KumaDiscoveryServiceClient is the client API for KumaDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KumaDiscoveryServiceClient interface {
	DeltaKumaResources(ctx context.Context, opts ...grpc.CallOption) (KumaDiscoveryService_DeltaKumaResourcesClient, error)
	StreamKumaResources(ctx context.Context, opts ...grpc.CallOption) (KumaDiscoveryService_StreamKumaResourcesClient, error)
	FetchKumaResources(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error)
}

type kumaDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewKumaDiscoveryServiceClient(cc *grpc.ClientConn) KumaDiscoveryServiceClient {
	return &kumaDiscoveryServiceClient{cc}
}

func (c *kumaDiscoveryServiceClient) DeltaKumaResources(ctx context.Context, opts ...grpc.CallOption) (KumaDiscoveryService_DeltaKumaResourcesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KumaDiscoveryService_serviceDesc.Streams[0], "/kuma.mesh.v1alpha1.KumaDiscoveryService/DeltaKumaResources", opts...)
	if err != nil {
		return nil, err
	}
	x := &kumaDiscoveryServiceDeltaKumaResourcesClient{stream}
	return x, nil
}

type KumaDiscoveryService_DeltaKumaResourcesClient interface {
	Send(*v2.DeltaDiscoveryRequest) error
	Recv() (*v2.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type kumaDiscoveryServiceDeltaKumaResourcesClient struct {
	grpc.ClientStream
}

func (x *kumaDiscoveryServiceDeltaKumaResourcesClient) Send(m *v2.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *kumaDiscoveryServiceDeltaKumaResourcesClient) Recv() (*v2.DeltaDiscoveryResponse, error) {
	m := new(v2.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kumaDiscoveryServiceClient) StreamKumaResources(ctx context.Context, opts ...grpc.CallOption) (KumaDiscoveryService_StreamKumaResourcesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KumaDiscoveryService_serviceDesc.Streams[1], "/kuma.mesh.v1alpha1.KumaDiscoveryService/StreamKumaResources", opts...)
	if err != nil {
		return nil, err
	}
	x := &kumaDiscoveryServiceStreamKumaResourcesClient{stream}
	return x, nil
}

type KumaDiscoveryService_StreamKumaResourcesClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type kumaDiscoveryServiceStreamKumaResourcesClient struct {
	grpc.ClientStream
}

func (x *kumaDiscoveryServiceStreamKumaResourcesClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *kumaDiscoveryServiceStreamKumaResourcesClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kumaDiscoveryServiceClient) FetchKumaResources(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error) {
	out := new(v2.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/kuma.mesh.v1alpha1.KumaDiscoveryService/FetchKumaResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KumaDiscoveryServiceServer is the server API for KumaDiscoveryService service.
type KumaDiscoveryServiceServer interface {
	DeltaKumaResources(KumaDiscoveryService_DeltaKumaResourcesServer) error
	StreamKumaResources(KumaDiscoveryService_StreamKumaResourcesServer) error
	FetchKumaResources(context.Context, *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error)
}

// UnimplementedKumaDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKumaDiscoveryServiceServer struct {
}

func (*UnimplementedKumaDiscoveryServiceServer) DeltaKumaResources(srv KumaDiscoveryService_DeltaKumaResourcesServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaKumaResources not implemented")
}
func (*UnimplementedKumaDiscoveryServiceServer) StreamKumaResources(srv KumaDiscoveryService_StreamKumaResourcesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamKumaResources not implemented")
}
func (*UnimplementedKumaDiscoveryServiceServer) FetchKumaResources(ctx context.Context, req *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchKumaResources not implemented")
}

func RegisterKumaDiscoveryServiceServer(s *grpc.Server, srv KumaDiscoveryServiceServer) {
	s.RegisterService(&_KumaDiscoveryService_serviceDesc, srv)
}

func _KumaDiscoveryService_DeltaKumaResources_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KumaDiscoveryServiceServer).DeltaKumaResources(&kumaDiscoveryServiceDeltaKumaResourcesServer{stream})
}

type KumaDiscoveryService_DeltaKumaResourcesServer interface {
	Send(*v2.DeltaDiscoveryResponse) error
	Recv() (*v2.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type kumaDiscoveryServiceDeltaKumaResourcesServer struct {
	grpc.ServerStream
}

func (x *kumaDiscoveryServiceDeltaKumaResourcesServer) Send(m *v2.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *kumaDiscoveryServiceDeltaKumaResourcesServer) Recv() (*v2.DeltaDiscoveryRequest, error) {
	m := new(v2.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _KumaDiscoveryService_StreamKumaResources_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KumaDiscoveryServiceServer).StreamKumaResources(&kumaDiscoveryServiceStreamKumaResourcesServer{stream})
}

type KumaDiscoveryService_StreamKumaResourcesServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type kumaDiscoveryServiceStreamKumaResourcesServer struct {
	grpc.ServerStream
}

func (x *kumaDiscoveryServiceStreamKumaResourcesServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *kumaDiscoveryServiceStreamKumaResourcesServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _KumaDiscoveryService_FetchKumaResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v2.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KumaDiscoveryServiceServer).FetchKumaResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kuma.mesh.v1alpha1.KumaDiscoveryService/FetchKumaResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KumaDiscoveryServiceServer).FetchKumaResources(ctx, req.(*v2.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KumaDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kuma.mesh.v1alpha1.KumaDiscoveryService",
	HandlerType: (*KumaDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchKumaResources",
			Handler:    _KumaDiscoveryService_FetchKumaResources_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DeltaKumaResources",
			Handler:       _KumaDiscoveryService_DeltaKumaResources_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamKumaResources",
			Handler:       _KumaDiscoveryService_StreamKumaResources_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "mesh/v1alpha1/kds.proto",
}
