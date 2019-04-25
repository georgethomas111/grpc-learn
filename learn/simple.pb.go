// Code generated by protoc-gen-go. DO NOT EDIT.
// source: simple.proto

package learn

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GossipRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GossipRequest) Reset()         { *m = GossipRequest{} }
func (m *GossipRequest) String() string { return proto.CompactTextString(m) }
func (*GossipRequest) ProtoMessage()    {}
func (*GossipRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ffd045dd4d042c1, []int{0}
}

func (m *GossipRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GossipRequest.Unmarshal(m, b)
}
func (m *GossipRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GossipRequest.Marshal(b, m, deterministic)
}
func (m *GossipRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GossipRequest.Merge(m, src)
}
func (m *GossipRequest) XXX_Size() int {
	return xxx_messageInfo_GossipRequest.Size(m)
}
func (m *GossipRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GossipRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GossipRequest proto.InternalMessageInfo

func (m *GossipRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// The response message containing the greetings
type Messages struct {
	Info                 string   `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Messages) Reset()         { *m = Messages{} }
func (m *Messages) String() string { return proto.CompactTextString(m) }
func (*Messages) ProtoMessage()    {}
func (*Messages) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ffd045dd4d042c1, []int{1}
}

func (m *Messages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Messages.Unmarshal(m, b)
}
func (m *Messages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Messages.Marshal(b, m, deterministic)
}
func (m *Messages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Messages.Merge(m, src)
}
func (m *Messages) XXX_Size() int {
	return xxx_messageInfo_Messages.Size(m)
}
func (m *Messages) XXX_DiscardUnknown() {
	xxx_messageInfo_Messages.DiscardUnknown(m)
}

var xxx_messageInfo_Messages proto.InternalMessageInfo

func (m *Messages) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func init() {
	proto.RegisterType((*GossipRequest)(nil), "learn.GossipRequest")
	proto.RegisterType((*Messages)(nil), "learn.Messages")
}

func init() { proto.RegisterFile("simple.proto", fileDescriptor_5ffd045dd4d042c1) }

var fileDescriptor_5ffd045dd4d042c1 = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xce, 0xcc, 0x2d,
	0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0x49, 0x4d, 0x2c, 0xca, 0x53,
	0x92, 0xe7, 0xe2, 0x75, 0xcf, 0x2f, 0x2e, 0xce, 0x2c, 0x08, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e,
	0x11, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c,
	0x51, 0x92, 0xe3, 0xe2, 0xf0, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x2d, 0x16, 0x12, 0xe2, 0x62,
	0xc9, 0xcc, 0x4b, 0xcb, 0x87, 0xca, 0x82, 0xd9, 0x46, 0x56, 0x5c, 0xac, 0x3e, 0x20, 0x93, 0x84,
	0x0c, 0xb9, 0xd8, 0x20, 0x26, 0x09, 0x89, 0xe8, 0x81, 0xcd, 0xd6, 0x43, 0x31, 0x58, 0x8a, 0x1f,
	0x2a, 0x0a, 0x33, 0xcd, 0x80, 0x31, 0x89, 0x0d, 0xec, 0x14, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x1e, 0x98, 0x49, 0xbc, 0x9a, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LearnClient is the client API for Learn service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LearnClient interface {
	Gossip(ctx context.Context, in *GossipRequest, opts ...grpc.CallOption) (Learn_GossipClient, error)
}

type learnClient struct {
	cc *grpc.ClientConn
}

func NewLearnClient(cc *grpc.ClientConn) LearnClient {
	return &learnClient{cc}
}

func (c *learnClient) Gossip(ctx context.Context, in *GossipRequest, opts ...grpc.CallOption) (Learn_GossipClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Learn_serviceDesc.Streams[0], "/learn.Learn/Gossip", opts...)
	if err != nil {
		return nil, err
	}
	x := &learnGossipClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Learn_GossipClient interface {
	Recv() (*Messages, error)
	grpc.ClientStream
}

type learnGossipClient struct {
	grpc.ClientStream
}

func (x *learnGossipClient) Recv() (*Messages, error) {
	m := new(Messages)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LearnServer is the server API for Learn service.
type LearnServer interface {
	Gossip(*GossipRequest, Learn_GossipServer) error
}

// UnimplementedLearnServer can be embedded to have forward compatible implementations.
type UnimplementedLearnServer struct {
}

func (*UnimplementedLearnServer) Gossip(req *GossipRequest, srv Learn_GossipServer) error {
	return status.Errorf(codes.Unimplemented, "method Gossip not implemented")
}

func RegisterLearnServer(s *grpc.Server, srv LearnServer) {
	s.RegisterService(&_Learn_serviceDesc, srv)
}

func _Learn_Gossip_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GossipRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LearnServer).Gossip(m, &learnGossipServer{stream})
}

type Learn_GossipServer interface {
	Send(*Messages) error
	grpc.ServerStream
}

type learnGossipServer struct {
	grpc.ServerStream
}

func (x *learnGossipServer) Send(m *Messages) error {
	return x.ServerStream.SendMsg(m)
}

var _Learn_serviceDesc = grpc.ServiceDesc{
	ServiceName: "learn.Learn",
	HandlerType: (*LearnServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Gossip",
			Handler:       _Learn_Gossip_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "simple.proto",
}
