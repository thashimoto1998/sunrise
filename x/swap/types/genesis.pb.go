// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sunrise/swap/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the swap module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params                  Params                   `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	IncomingInFlightPackets []IncomingInFlightPacket `protobuf:"bytes,2,rep,name=incoming_in_flight_packets,json=incomingInFlightPackets,proto3" json:"incoming_in_flight_packets"`
	OutgoingInFlightPackets []OutgoingInFlightPacket `protobuf:"bytes,3,rep,name=outgoing_in_flight_packets,json=outgoingInFlightPackets,proto3" json:"outgoing_in_flight_packets"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fd15ebb865a4f63, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetIncomingInFlightPackets() []IncomingInFlightPacket {
	if m != nil {
		return m.IncomingInFlightPackets
	}
	return nil
}

func (m *GenesisState) GetOutgoingInFlightPackets() []OutgoingInFlightPacket {
	if m != nil {
		return m.OutgoingInFlightPackets
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "sunrise.swap.GenesisState")
}

func init() { proto.RegisterFile("sunrise/swap/genesis.proto", fileDescriptor_1fd15ebb865a4f63) }

var fileDescriptor_1fd15ebb865a4f63 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x40, 0x93, 0x16, 0x75, 0x48, 0x3b, 0x45, 0x95, 0x28, 0x19, 0x4c, 0x05, 0x0c, 0x1d, 0x90,
	0x2d, 0x85, 0x1b, 0x74, 0x28, 0xea, 0x44, 0x05, 0x1b, 0x4b, 0xe4, 0x46, 0xc6, 0xb5, 0x68, 0xfc,
	0xad, 0xd8, 0x11, 0xf4, 0x16, 0x0c, 0x1c, 0xaa, 0x63, 0x47, 0x26, 0x84, 0x92, 0x8b, 0xa0, 0x3a,
	0x46, 0x22, 0x95, 0xb7, 0x44, 0xef, 0xf9, 0xbf, 0xaf, 0x1f, 0x25, 0xba, 0x92, 0xa5, 0xd0, 0x8c,
	0xe8, 0x37, 0xaa, 0x08, 0x67, 0x92, 0x69, 0xa1, 0xb1, 0x2a, 0xc1, 0x40, 0x3c, 0x72, 0x0c, 0x1f,
	0x59, 0x32, 0xe6, 0xc0, 0xc1, 0x02, 0x72, 0xfc, 0x6a, 0x9d, 0xe4, 0xba, 0xf3, 0x5e, 0xc8, 0xec,
	0x65, 0x2b, 0xf8, 0xc6, 0x64, 0x8a, 0xe6, 0xaf, 0xcc, 0x38, 0xe9, 0xa2, 0x23, 0x29, 0x5a, 0xd2,
	0xc2, 0x35, 0xae, 0x3e, 0x7b, 0xd1, 0xe8, 0xbe, 0xad, 0x3e, 0x19, 0x6a, 0x58, 0x9c, 0x46, 0x83,
	0x56, 0x98, 0x84, 0xd3, 0x70, 0x36, 0x4c, 0xc7, 0xf8, 0xff, 0x16, 0x78, 0x65, 0xd9, 0xfc, 0x6c,
	0xff, 0x7d, 0x19, 0x3c, 0x3a, 0x33, 0xe6, 0x51, 0x22, 0x64, 0x0e, 0x85, 0x90, 0x3c, 0x3b, 0x5d,
	0x41, 0x4f, 0x7a, 0xd3, 0xfe, 0x6c, 0x98, 0xde, 0x74, 0xe7, 0x2c, 0x9d, 0xbf, 0x94, 0x0b, 0x6b,
	0xaf, 0xac, 0xec, 0xe6, 0x9e, 0x0b, 0x2f, 0xb5, 0x21, 0xa8, 0x0c, 0x07, 0x7f, 0xa8, 0xef, 0x0b,
	0x3d, 0x38, 0xdf, 0x1f, 0x02, 0x2f, 0xd5, 0xf3, 0xc5, 0xbe, 0x46, 0xe1, 0xa1, 0x46, 0xe1, 0x4f,
	0x8d, 0xc2, 0x8f, 0x06, 0x05, 0x87, 0x06, 0x05, 0x5f, 0x0d, 0x0a, 0x9e, 0x6f, 0xb9, 0x30, 0x9b,
	0x6a, 0x8d, 0x73, 0x28, 0x88, 0x0b, 0x6d, 0xe9, 0x8e, 0x95, 0x7f, 0x3f, 0xe4, 0xbd, 0xbd, 0xb2,
	0xd9, 0x29, 0xa6, 0xd7, 0x03, 0x7b, 0xe5, 0xbb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9f, 0xf3,
	0x96, 0xa4, 0xe7, 0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OutgoingInFlightPackets) > 0 {
		for iNdEx := len(m.OutgoingInFlightPackets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OutgoingInFlightPackets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.IncomingInFlightPackets) > 0 {
		for iNdEx := len(m.IncomingInFlightPackets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IncomingInFlightPackets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.IncomingInFlightPackets) > 0 {
		for _, e := range m.IncomingInFlightPackets {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.OutgoingInFlightPackets) > 0 {
		for _, e := range m.OutgoingInFlightPackets {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncomingInFlightPackets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IncomingInFlightPackets = append(m.IncomingInFlightPackets, IncomingInFlightPacket{})
			if err := m.IncomingInFlightPackets[len(m.IncomingInFlightPackets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutgoingInFlightPackets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutgoingInFlightPackets = append(m.OutgoingInFlightPackets, OutgoingInFlightPacket{})
			if err := m.OutgoingInFlightPackets[len(m.OutgoingInFlightPackets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
