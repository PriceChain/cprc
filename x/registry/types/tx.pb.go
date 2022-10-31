// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: registry/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgCreateRegistry struct {
	Creator               string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Name                  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	StakeAmount           string `protobuf:"bytes,3,opt,name=stakeAmount,proto3" json:"stakeAmount,omitempty"`
	Quorum                string `protobuf:"bytes,4,opt,name=quorum,proto3" json:"quorum,omitempty"`
	ConsensusExpiringTime string `protobuf:"bytes,5,opt,name=consensusExpiringTime,proto3" json:"consensusExpiringTime,omitempty"`
}

func (m *MsgCreateRegistry) Reset()         { *m = MsgCreateRegistry{} }
func (m *MsgCreateRegistry) String() string { return proto.CompactTextString(m) }
func (*MsgCreateRegistry) ProtoMessage()    {}
func (*MsgCreateRegistry) Descriptor() ([]byte, []int) {
	return fileDescriptor_33996d822bb1b5d4, []int{0}
}
func (m *MsgCreateRegistry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateRegistry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateRegistry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateRegistry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateRegistry.Merge(m, src)
}
func (m *MsgCreateRegistry) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateRegistry) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateRegistry.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateRegistry proto.InternalMessageInfo

func (m *MsgCreateRegistry) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateRegistry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgCreateRegistry) GetStakeAmount() string {
	if m != nil {
		return m.StakeAmount
	}
	return ""
}

func (m *MsgCreateRegistry) GetQuorum() string {
	if m != nil {
		return m.Quorum
	}
	return ""
}

func (m *MsgCreateRegistry) GetConsensusExpiringTime() string {
	if m != nil {
		return m.ConsensusExpiringTime
	}
	return ""
}

type MsgCreateRegistryResponse struct {
}

func (m *MsgCreateRegistryResponse) Reset()         { *m = MsgCreateRegistryResponse{} }
func (m *MsgCreateRegistryResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateRegistryResponse) ProtoMessage()    {}
func (*MsgCreateRegistryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_33996d822bb1b5d4, []int{1}
}
func (m *MsgCreateRegistryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateRegistryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateRegistryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateRegistryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateRegistryResponse.Merge(m, src)
}
func (m *MsgCreateRegistryResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateRegistryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateRegistryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateRegistryResponse proto.InternalMessageInfo

type MsgJoinRegistryCoOperator struct {
	Creator     string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	RegistryId  string `protobuf:"bytes,2,opt,name=registryId,proto3" json:"registryId,omitempty"`
	StakeAmount string `protobuf:"bytes,3,opt,name=stakeAmount,proto3" json:"stakeAmount,omitempty"`
}

func (m *MsgJoinRegistryCoOperator) Reset()         { *m = MsgJoinRegistryCoOperator{} }
func (m *MsgJoinRegistryCoOperator) String() string { return proto.CompactTextString(m) }
func (*MsgJoinRegistryCoOperator) ProtoMessage()    {}
func (*MsgJoinRegistryCoOperator) Descriptor() ([]byte, []int) {
	return fileDescriptor_33996d822bb1b5d4, []int{2}
}
func (m *MsgJoinRegistryCoOperator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgJoinRegistryCoOperator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgJoinRegistryCoOperator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgJoinRegistryCoOperator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgJoinRegistryCoOperator.Merge(m, src)
}
func (m *MsgJoinRegistryCoOperator) XXX_Size() int {
	return m.Size()
}
func (m *MsgJoinRegistryCoOperator) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgJoinRegistryCoOperator.DiscardUnknown(m)
}

var xxx_messageInfo_MsgJoinRegistryCoOperator proto.InternalMessageInfo

func (m *MsgJoinRegistryCoOperator) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgJoinRegistryCoOperator) GetRegistryId() string {
	if m != nil {
		return m.RegistryId
	}
	return ""
}

func (m *MsgJoinRegistryCoOperator) GetStakeAmount() string {
	if m != nil {
		return m.StakeAmount
	}
	return ""
}

type MsgJoinRegistryCoOperatorResponse struct {
}

func (m *MsgJoinRegistryCoOperatorResponse) Reset()         { *m = MsgJoinRegistryCoOperatorResponse{} }
func (m *MsgJoinRegistryCoOperatorResponse) String() string { return proto.CompactTextString(m) }
func (*MsgJoinRegistryCoOperatorResponse) ProtoMessage()    {}
func (*MsgJoinRegistryCoOperatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_33996d822bb1b5d4, []int{3}
}
func (m *MsgJoinRegistryCoOperatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgJoinRegistryCoOperatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgJoinRegistryCoOperatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgJoinRegistryCoOperatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgJoinRegistryCoOperatorResponse.Merge(m, src)
}
func (m *MsgJoinRegistryCoOperatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgJoinRegistryCoOperatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgJoinRegistryCoOperatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgJoinRegistryCoOperatorResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateRegistry)(nil), "pricechain.rd_net.registry.MsgCreateRegistry")
	proto.RegisterType((*MsgCreateRegistryResponse)(nil), "pricechain.rd_net.registry.MsgCreateRegistryResponse")
	proto.RegisterType((*MsgJoinRegistryCoOperator)(nil), "pricechain.rd_net.registry.MsgJoinRegistryCoOperator")
	proto.RegisterType((*MsgJoinRegistryCoOperatorResponse)(nil), "pricechain.rd_net.registry.MsgJoinRegistryCoOperatorResponse")
}

func init() { proto.RegisterFile("registry/tx.proto", fileDescriptor_33996d822bb1b5d4) }

var fileDescriptor_33996d822bb1b5d4 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4f, 0x4b, 0xc3, 0x30,
	0x1c, 0x5d, 0xb6, 0x39, 0xf1, 0x27, 0x08, 0x0b, 0x38, 0xea, 0x84, 0x30, 0xeb, 0xc5, 0xcb, 0x5a,
	0xf0, 0xcf, 0xd1, 0x83, 0x0e, 0x11, 0x85, 0xa1, 0x0c, 0x4f, 0x5e, 0xa4, 0xeb, 0x42, 0x17, 0xa4,
	0x49, 0x4d, 0x52, 0xdd, 0xbe, 0xc4, 0xf0, 0xab, 0xf8, 0x2d, 0x3c, 0xee, 0xe8, 0x51, 0xb6, 0x2f,
	0x22, 0x8d, 0x0b, 0x0e, 0xb7, 0x32, 0xf4, 0xd6, 0xbc, 0xf7, 0xd2, 0xf7, 0xde, 0x2f, 0x3f, 0xa8,
	0x4a, 0x1a, 0x31, 0xa5, 0xe5, 0xd0, 0xd7, 0x03, 0x2f, 0x91, 0x42, 0x0b, 0x5c, 0x4f, 0x24, 0x0b,
	0x69, 0xd8, 0x0f, 0x18, 0xf7, 0x64, 0xef, 0x81, 0x53, 0xed, 0x59, 0x91, 0xfb, 0x86, 0xa0, 0xda,
	0x56, 0x51, 0x4b, 0xd2, 0x40, 0xd3, 0xce, 0x0c, 0xc5, 0x0e, 0xac, 0x87, 0x19, 0x22, 0xa4, 0x83,
	0x1a, 0xe8, 0x60, 0xa3, 0x63, 0x8f, 0x18, 0x43, 0x99, 0x07, 0x31, 0x75, 0x8a, 0x06, 0x36, 0xdf,
	0xb8, 0x01, 0x9b, 0x4a, 0x07, 0x8f, 0xf4, 0x2c, 0x16, 0x29, 0xd7, 0x4e, 0xc9, 0x50, 0xf3, 0x10,
	0xae, 0x41, 0xe5, 0x29, 0x15, 0x32, 0x8d, 0x9d, 0xb2, 0x21, 0x67, 0x27, 0x7c, 0x0c, 0xdb, 0xa1,
	0xe0, 0x8a, 0x72, 0x95, 0xaa, 0x8b, 0x41, 0xc2, 0x24, 0xe3, 0xd1, 0x1d, 0x8b, 0xa9, 0xb3, 0x66,
	0x64, 0xcb, 0x49, 0x77, 0x17, 0x76, 0x16, 0x22, 0x77, 0xa8, 0x4a, 0x32, 0xb1, 0xfb, 0x62, 0xc8,
	0x6b, 0xc1, 0xb8, 0xa5, 0x5a, 0xe2, 0x26, 0xa1, 0xd2, 0xa4, 0xcf, 0xef, 0x45, 0x00, 0xec, 0x4c,
	0xae, 0x7a, 0xb3, 0x76, 0x73, 0xc8, 0xea, 0x8e, 0xee, 0x3e, 0xec, 0xe5, 0x1a, 0xdb, 0x74, 0x87,
	0xa3, 0x22, 0x94, 0xda, 0x2a, 0xc2, 0xcf, 0xb0, 0xf5, 0x6b, 0xe4, 0x4d, 0x2f, 0xff, 0x95, 0xbc,
	0x85, 0xba, 0xf5, 0x93, 0x3f, 0xc9, 0xad, 0x3f, 0x1e, 0x21, 0xa8, 0xe5, 0xcc, 0x66, 0xd5, 0x1f,
	0x97, 0x5f, 0xab, 0x9f, 0xfe, 0xeb, 0x9a, 0x0d, 0x74, 0x7e, 0xf9, 0x3e, 0x21, 0x68, 0x3c, 0x21,
	0xe8, 0x73, 0x42, 0xd0, 0xeb, 0x94, 0x14, 0xc6, 0x53, 0x52, 0xf8, 0x98, 0x92, 0xc2, 0x7d, 0x33,
	0x62, 0xba, 0x9f, 0x76, 0xbd, 0x50, 0xc4, 0xfe, 0x6d, 0x66, 0xd1, 0xca, 0x2c, 0xfc, 0x6f, 0x0b,
	0x7f, 0xe0, 0xff, 0xec, 0xf9, 0x30, 0xa1, 0xaa, 0x5b, 0x31, 0xbb, 0x7e, 0xf4, 0x15, 0x00, 0x00,
	0xff, 0xff, 0x70, 0xc2, 0x6d, 0x19, 0x00, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateRegistry(ctx context.Context, in *MsgCreateRegistry, opts ...grpc.CallOption) (*MsgCreateRegistryResponse, error)
	JoinRegistryCoOperator(ctx context.Context, in *MsgJoinRegistryCoOperator, opts ...grpc.CallOption) (*MsgJoinRegistryCoOperatorResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateRegistry(ctx context.Context, in *MsgCreateRegistry, opts ...grpc.CallOption) (*MsgCreateRegistryResponse, error) {
	out := new(MsgCreateRegistryResponse)
	err := c.cc.Invoke(ctx, "/pricechain.rd_net.registry.Msg/CreateRegistry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) JoinRegistryCoOperator(ctx context.Context, in *MsgJoinRegistryCoOperator, opts ...grpc.CallOption) (*MsgJoinRegistryCoOperatorResponse, error) {
	out := new(MsgJoinRegistryCoOperatorResponse)
	err := c.cc.Invoke(ctx, "/pricechain.rd_net.registry.Msg/JoinRegistryCoOperator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateRegistry(context.Context, *MsgCreateRegistry) (*MsgCreateRegistryResponse, error)
	JoinRegistryCoOperator(context.Context, *MsgJoinRegistryCoOperator) (*MsgJoinRegistryCoOperatorResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateRegistry(ctx context.Context, req *MsgCreateRegistry) (*MsgCreateRegistryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRegistry not implemented")
}
func (*UnimplementedMsgServer) JoinRegistryCoOperator(ctx context.Context, req *MsgJoinRegistryCoOperator) (*MsgJoinRegistryCoOperatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinRegistryCoOperator not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateRegistry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateRegistry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateRegistry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pricechain.rd_net.registry.Msg/CreateRegistry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateRegistry(ctx, req.(*MsgCreateRegistry))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_JoinRegistryCoOperator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgJoinRegistryCoOperator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).JoinRegistryCoOperator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pricechain.rd_net.registry.Msg/JoinRegistryCoOperator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).JoinRegistryCoOperator(ctx, req.(*MsgJoinRegistryCoOperator))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pricechain.rd_net.registry.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRegistry",
			Handler:    _Msg_CreateRegistry_Handler,
		},
		{
			MethodName: "JoinRegistryCoOperator",
			Handler:    _Msg_JoinRegistryCoOperator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry/tx.proto",
}

func (m *MsgCreateRegistry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateRegistry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateRegistry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ConsensusExpiringTime) > 0 {
		i -= len(m.ConsensusExpiringTime)
		copy(dAtA[i:], m.ConsensusExpiringTime)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ConsensusExpiringTime)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Quorum) > 0 {
		i -= len(m.Quorum)
		copy(dAtA[i:], m.Quorum)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Quorum)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.StakeAmount) > 0 {
		i -= len(m.StakeAmount)
		copy(dAtA[i:], m.StakeAmount)
		i = encodeVarintTx(dAtA, i, uint64(len(m.StakeAmount)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateRegistryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateRegistryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateRegistryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgJoinRegistryCoOperator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgJoinRegistryCoOperator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgJoinRegistryCoOperator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StakeAmount) > 0 {
		i -= len(m.StakeAmount)
		copy(dAtA[i:], m.StakeAmount)
		i = encodeVarintTx(dAtA, i, uint64(len(m.StakeAmount)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RegistryId) > 0 {
		i -= len(m.RegistryId)
		copy(dAtA[i:], m.RegistryId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RegistryId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgJoinRegistryCoOperatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgJoinRegistryCoOperatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgJoinRegistryCoOperatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateRegistry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.StakeAmount)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Quorum)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ConsensusExpiringTime)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateRegistryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgJoinRegistryCoOperator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.RegistryId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.StakeAmount)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgJoinRegistryCoOperatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateRegistry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateRegistry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateRegistry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakeAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakeAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quorum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Quorum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusExpiringTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsensusExpiringTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateRegistryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateRegistryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateRegistryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgJoinRegistryCoOperator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgJoinRegistryCoOperator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgJoinRegistryCoOperator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegistryId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RegistryId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakeAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakeAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgJoinRegistryCoOperatorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgJoinRegistryCoOperatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgJoinRegistryCoOperatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
