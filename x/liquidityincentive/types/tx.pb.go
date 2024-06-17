// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sunrise/liquidityincentive/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// MsgUpdateParams is the Msg/UpdateParams request type.
type MsgUpdateParams struct {
	// authority is the address that controls the module (defaults to x/gov unless
	// overwritten).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// NOTE: All parameters must be supplied.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *MsgUpdateParams) Reset()         { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParams) ProtoMessage()    {}
func (*MsgUpdateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a8cffcd57d82c0d, []int{0}
}
func (m *MsgUpdateParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParams.Merge(m, src)
}
func (m *MsgUpdateParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParams proto.InternalMessageInfo

func (m *MsgUpdateParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateParams) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// MsgUpdateParamsResponse defines the response structure for executing a
// MsgUpdateParams message.
type MsgUpdateParamsResponse struct {
}

func (m *MsgUpdateParamsResponse) Reset()         { *m = MsgUpdateParamsResponse{} }
func (m *MsgUpdateParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsResponse) ProtoMessage()    {}
func (*MsgUpdateParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a8cffcd57d82c0d, []int{1}
}
func (m *MsgUpdateParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsResponse.Merge(m, src)
}
func (m *MsgUpdateParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsResponse proto.InternalMessageInfo

type MsgVoteGauge struct {
	Sender  string       `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Weights []PoolWeight `protobuf:"bytes,2,rep,name=weights,proto3" json:"weights"`
}

func (m *MsgVoteGauge) Reset()         { *m = MsgVoteGauge{} }
func (m *MsgVoteGauge) String() string { return proto.CompactTextString(m) }
func (*MsgVoteGauge) ProtoMessage()    {}
func (*MsgVoteGauge) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a8cffcd57d82c0d, []int{2}
}
func (m *MsgVoteGauge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgVoteGauge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgVoteGauge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgVoteGauge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgVoteGauge.Merge(m, src)
}
func (m *MsgVoteGauge) XXX_Size() int {
	return m.Size()
}
func (m *MsgVoteGauge) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgVoteGauge.DiscardUnknown(m)
}

var xxx_messageInfo_MsgVoteGauge proto.InternalMessageInfo

func (m *MsgVoteGauge) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgVoteGauge) GetWeights() []PoolWeight {
	if m != nil {
		return m.Weights
	}
	return nil
}

type MsgVoteGaugeResponse struct {
}

func (m *MsgVoteGaugeResponse) Reset()         { *m = MsgVoteGaugeResponse{} }
func (m *MsgVoteGaugeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgVoteGaugeResponse) ProtoMessage()    {}
func (*MsgVoteGaugeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a8cffcd57d82c0d, []int{3}
}
func (m *MsgVoteGaugeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgVoteGaugeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgVoteGaugeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgVoteGaugeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgVoteGaugeResponse.Merge(m, src)
}
func (m *MsgVoteGaugeResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgVoteGaugeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgVoteGaugeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgVoteGaugeResponse proto.InternalMessageInfo

type MsgCollectVoteRewards struct {
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgCollectVoteRewards) Reset()         { *m = MsgCollectVoteRewards{} }
func (m *MsgCollectVoteRewards) String() string { return proto.CompactTextString(m) }
func (*MsgCollectVoteRewards) ProtoMessage()    {}
func (*MsgCollectVoteRewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a8cffcd57d82c0d, []int{4}
}
func (m *MsgCollectVoteRewards) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCollectVoteRewards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCollectVoteRewards.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCollectVoteRewards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCollectVoteRewards.Merge(m, src)
}
func (m *MsgCollectVoteRewards) XXX_Size() int {
	return m.Size()
}
func (m *MsgCollectVoteRewards) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCollectVoteRewards.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCollectVoteRewards proto.InternalMessageInfo

func (m *MsgCollectVoteRewards) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

type MsgCollectVoteRewardsResponse struct {
}

func (m *MsgCollectVoteRewardsResponse) Reset()         { *m = MsgCollectVoteRewardsResponse{} }
func (m *MsgCollectVoteRewardsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCollectVoteRewardsResponse) ProtoMessage()    {}
func (*MsgCollectVoteRewardsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a8cffcd57d82c0d, []int{5}
}
func (m *MsgCollectVoteRewardsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCollectVoteRewardsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCollectVoteRewardsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCollectVoteRewardsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCollectVoteRewardsResponse.Merge(m, src)
}
func (m *MsgCollectVoteRewardsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCollectVoteRewardsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCollectVoteRewardsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCollectVoteRewardsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUpdateParams)(nil), "sunrise.liquidityincentive.MsgUpdateParams")
	proto.RegisterType((*MsgUpdateParamsResponse)(nil), "sunrise.liquidityincentive.MsgUpdateParamsResponse")
	proto.RegisterType((*MsgVoteGauge)(nil), "sunrise.liquidityincentive.MsgVoteGauge")
	proto.RegisterType((*MsgVoteGaugeResponse)(nil), "sunrise.liquidityincentive.MsgVoteGaugeResponse")
	proto.RegisterType((*MsgCollectVoteRewards)(nil), "sunrise.liquidityincentive.MsgCollectVoteRewards")
	proto.RegisterType((*MsgCollectVoteRewardsResponse)(nil), "sunrise.liquidityincentive.MsgCollectVoteRewardsResponse")
}

func init() {
	proto.RegisterFile("sunrise/liquidityincentive/tx.proto", fileDescriptor_2a8cffcd57d82c0d)
}

var fileDescriptor_2a8cffcd57d82c0d = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xcd, 0xb6, 0x10, 0x94, 0x4d, 0x25, 0x84, 0x15, 0xda, 0xd4, 0x12, 0x6e, 0x65, 0xa4, 0x12,
	0x05, 0xb0, 0x69, 0x2a, 0x21, 0x11, 0x21, 0x21, 0x82, 0x80, 0x53, 0x24, 0x64, 0xbe, 0x24, 0x2e,
	0xc8, 0x8d, 0x57, 0x9b, 0x95, 0x6c, 0xaf, 0xd9, 0x59, 0xb7, 0x8d, 0xb8, 0xa0, 0x1e, 0x39, 0xf1,
	0x33, 0x38, 0xe6, 0xc0, 0x8f, 0xe8, 0xb1, 0x70, 0xe2, 0x84, 0x50, 0x72, 0xc8, 0x85, 0x1f, 0x81,
	0x6c, 0xaf, 0x13, 0x9a, 0xa6, 0x2e, 0xbd, 0x24, 0xde, 0x99, 0xf7, 0xde, 0xbc, 0x37, 0x5e, 0xe3,
	0x9b, 0x10, 0x87, 0x82, 0x01, 0xb1, 0x7d, 0xf6, 0x21, 0x66, 0x1e, 0x93, 0x03, 0x16, 0xf6, 0x48,
	0x28, 0xd9, 0x1e, 0xb1, 0xe5, 0x81, 0x15, 0x09, 0x2e, 0xb9, 0xa6, 0x2b, 0x90, 0x75, 0x1a, 0xa4,
	0x5f, 0x73, 0x03, 0x16, 0x72, 0x3b, 0xfd, 0xcd, 0xe0, 0xfa, 0x5a, 0x8f, 0x43, 0xc0, 0xc1, 0x0e,
	0x80, 0xda, 0x7b, 0xdb, 0xc9, 0x9f, 0x6a, 0xac, 0x67, 0x8d, 0xf7, 0xe9, 0xc9, 0xce, 0x0e, 0xaa,
	0x55, 0xa3, 0x9c, 0xf2, 0xac, 0x9e, 0x3c, 0xa9, 0xea, 0xad, 0x02, 0x77, 0x91, 0x2b, 0xdc, 0x20,
	0xa7, 0x6f, 0x15, 0x00, 0xa9, 0x1b, 0x53, 0x92, 0xe1, 0xcc, 0xef, 0x08, 0x5f, 0xed, 0x02, 0x7d,
	0x1d, 0x79, 0xae, 0x24, 0x2f, 0x52, 0x05, 0xed, 0x3e, 0xae, 0xb8, 0xb1, 0xec, 0x73, 0xc1, 0xe4,
	0xa0, 0x8e, 0x36, 0x51, 0xa3, 0xd2, 0xa9, 0xff, 0xf8, 0x76, 0xb7, 0xa6, 0xfc, 0x3d, 0xf6, 0x3c,
	0x41, 0x00, 0x5e, 0x4a, 0xc1, 0x42, 0xea, 0xcc, 0xa0, 0xda, 0x53, 0x5c, 0xce, 0x3c, 0xd4, 0x97,
	0x36, 0x51, 0xa3, 0xda, 0x32, 0xad, 0xb3, 0xd7, 0x64, 0x65, 0xb3, 0x3a, 0x95, 0xa3, 0x5f, 0x1b,
	0xa5, 0xaf, 0x93, 0x61, 0x13, 0x39, 0x8a, 0xdc, 0x7e, 0x74, 0x38, 0x19, 0x36, 0x67, 0xb2, 0x9f,
	0x27, 0xc3, 0xe6, 0x9d, 0x3c, 0xcd, 0xc1, 0xa2, 0x3c, 0x73, 0xfe, 0xcd, 0x75, 0xbc, 0x36, 0x57,
	0x72, 0x08, 0x44, 0x3c, 0x04, 0x62, 0x7e, 0xc4, 0x2b, 0x5d, 0xa0, 0x6f, 0xb8, 0x24, 0xcf, 0x93,
	0x25, 0x68, 0xab, 0xb8, 0x0c, 0x24, 0xf4, 0x88, 0xc8, 0x72, 0x3a, 0xea, 0xa4, 0x3d, 0xc3, 0x57,
	0xf6, 0x09, 0xa3, 0x7d, 0x99, 0x64, 0x59, 0x6e, 0x54, 0x5b, 0x5b, 0x85, 0x59, 0x38, 0xf7, 0xdf,
	0xa6, 0xf0, 0xce, 0xa5, 0x24, 0x8f, 0x93, 0x93, 0xdb, 0xd5, 0x24, 0x8b, 0x12, 0x35, 0x57, 0x71,
	0xed, 0xdf, 0xe1, 0x53, 0x53, 0x0f, 0xf1, 0xf5, 0x2e, 0xd0, 0x27, 0xdc, 0xf7, 0x49, 0x4f, 0x26,
	0x6d, 0x87, 0xec, 0xbb, 0xc2, 0x83, 0xb3, 0xdc, 0x9d, 0x54, 0xdd, 0xc0, 0x37, 0x16, 0xb2, 0x73,
	0xf9, 0xd6, 0x9f, 0x25, 0xbc, 0xdc, 0x05, 0xaa, 0x45, 0x78, 0xe5, 0xc4, 0x6b, 0xbe, 0x5d, 0x14,
	0x69, 0x6e, 0x81, 0xfa, 0xce, 0x05, 0xc0, 0xf9, 0x64, 0x8d, 0xe2, 0xca, 0x6c, 0xd5, 0x8d, 0x73,
	0x14, 0xa6, 0x48, 0xfd, 0xde, 0xff, 0x22, 0xa7, 0x83, 0x0e, 0x11, 0xd6, 0x16, 0xec, 0x6f, 0xfb,
	0x1c, 0xa1, 0xd3, 0x14, 0xfd, 0xc1, 0x85, 0x29, 0xb9, 0x09, 0xfd, 0xf2, 0xa7, 0xe4, 0x1a, 0x77,
	0x5e, 0x1d, 0x8d, 0x0c, 0x74, 0x3c, 0x32, 0xd0, 0xef, 0x91, 0x81, 0xbe, 0x8c, 0x8d, 0xd2, 0xf1,
	0xd8, 0x28, 0xfd, 0x1c, 0x1b, 0xa5, 0x77, 0x6d, 0xca, 0x64, 0x3f, 0xde, 0xb5, 0x7a, 0x3c, 0xb0,
	0xd5, 0x14, 0xdf, 0x1d, 0x10, 0x61, 0x17, 0xde, 0x6e, 0x39, 0x88, 0x08, 0xec, 0x96, 0xd3, 0xcf,
	0x75, 0xe7, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x2f, 0xbf, 0x38, 0x9f, 0x04, 0x00, 0x00,
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
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	VoteGauge(ctx context.Context, in *MsgVoteGauge, opts ...grpc.CallOption) (*MsgVoteGaugeResponse, error)
	CollectVoteRewards(ctx context.Context, in *MsgCollectVoteRewards, opts ...grpc.CallOption) (*MsgCollectVoteRewardsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/sunrise.liquidityincentive.Msg/UpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) VoteGauge(ctx context.Context, in *MsgVoteGauge, opts ...grpc.CallOption) (*MsgVoteGaugeResponse, error) {
	out := new(MsgVoteGaugeResponse)
	err := c.cc.Invoke(ctx, "/sunrise.liquidityincentive.Msg/VoteGauge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CollectVoteRewards(ctx context.Context, in *MsgCollectVoteRewards, opts ...grpc.CallOption) (*MsgCollectVoteRewardsResponse, error) {
	out := new(MsgCollectVoteRewardsResponse)
	err := c.cc.Invoke(ctx, "/sunrise.liquidityincentive.Msg/CollectVoteRewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	VoteGauge(context.Context, *MsgVoteGauge) (*MsgVoteGaugeResponse, error)
	CollectVoteRewards(context.Context, *MsgCollectVoteRewards) (*MsgCollectVoteRewardsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateParams(ctx context.Context, req *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (*UnimplementedMsgServer) VoteGauge(ctx context.Context, req *MsgVoteGauge) (*MsgVoteGaugeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteGauge not implemented")
}
func (*UnimplementedMsgServer) CollectVoteRewards(ctx context.Context, req *MsgCollectVoteRewards) (*MsgCollectVoteRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectVoteRewards not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sunrise.liquidityincentive.Msg/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_VoteGauge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgVoteGauge)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).VoteGauge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sunrise.liquidityincentive.Msg/VoteGauge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).VoteGauge(ctx, req.(*MsgVoteGauge))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CollectVoteRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCollectVoteRewards)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CollectVoteRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sunrise.liquidityincentive.Msg/CollectVoteRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CollectVoteRewards(ctx, req.(*MsgCollectVoteRewards))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sunrise.liquidityincentive.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "VoteGauge",
			Handler:    _Msg_VoteGauge_Handler,
		},
		{
			MethodName: "CollectVoteRewards",
			Handler:    _Msg_CollectVoteRewards_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sunrise/liquidityincentive/tx.proto",
}

func (m *MsgUpdateParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgVoteGauge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgVoteGauge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgVoteGauge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Weights) > 0 {
		for iNdEx := len(m.Weights) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Weights[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgVoteGaugeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgVoteGaugeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgVoteGaugeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgCollectVoteRewards) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCollectVoteRewards) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCollectVoteRewards) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCollectVoteRewardsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCollectVoteRewardsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCollectVoteRewardsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgUpdateParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgVoteGauge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Weights) > 0 {
		for _, e := range m.Weights {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgVoteGaugeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgCollectVoteRewards) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCollectVoteRewardsResponse) Size() (n int) {
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
func (m *MsgUpdateParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
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
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *MsgUpdateParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgVoteGauge) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgVoteGauge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgVoteGauge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weights", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Weights = append(m.Weights, PoolWeight{})
			if err := m.Weights[len(m.Weights)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *MsgVoteGaugeResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgVoteGaugeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgVoteGaugeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgCollectVoteRewards) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCollectVoteRewards: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCollectVoteRewards: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
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
func (m *MsgCollectVoteRewardsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCollectVoteRewardsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCollectVoteRewardsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
