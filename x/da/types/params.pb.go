// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sunrise/da/params.proto

package types

import (
	bytes "bytes"
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params defines the parameters for the module.
type Params struct {
	VoteThreshold       cosmossdk_io_math.LegacyDec              `protobuf:"bytes,1,opt,name=vote_threshold,json=voteThreshold,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"vote_threshold"`
	SlashEpoch          uint64                                   `protobuf:"varint,2,opt,name=slash_epoch,json=slashEpoch,proto3" json:"slash_epoch,omitempty"`
	EpochMaxFault       uint64                                   `protobuf:"varint,3,opt,name=epoch_max_fault,json=epochMaxFault,proto3" json:"epoch_max_fault,omitempty"`
	SlashFraction       cosmossdk_io_math.LegacyDec              `protobuf:"bytes,4,opt,name=slash_fraction,json=slashFraction,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"slash_fraction"`
	ReplicationFactor   cosmossdk_io_math.LegacyDec              `protobuf:"bytes,5,opt,name=replication_factor,json=replicationFactor,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"replication_factor"`
	MinShardCount       uint64                                   `protobuf:"varint,6,opt,name=min_shard_count,json=minShardCount,proto3" json:"min_shard_count,omitempty"`
	MaxShardCount       uint64                                   `protobuf:"varint,7,opt,name=max_shard_count,json=maxShardCount,proto3" json:"max_shard_count,omitempty"`
	MaxShardSize        uint64                                   `protobuf:"varint,8,opt,name=max_shard_size,json=maxShardSize,proto3" json:"max_shard_size,omitempty"`
	ChallengePeriod     time.Duration                            `protobuf:"bytes,9,opt,name=challenge_period,json=challengePeriod,proto3,stdduration" json:"challenge_period"`
	ProofPeriod         time.Duration                            `protobuf:"bytes,10,opt,name=proof_period,json=proofPeriod,proto3,stdduration" json:"proof_period"`
	ChallengeCollateral github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,11,rep,name=challenge_collateral,json=challengeCollateral,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"challenge_collateral"`
	ZkpProvingKey       []byte                                   `protobuf:"bytes,12,opt,name=zkp_proving_key,json=zkpProvingKey,proto3" json:"zkp_proving_key,omitempty"`
	ZkpVerifyingKey     []byte                                   `protobuf:"bytes,13,opt,name=zkp_verifying_key,json=zkpVerifyingKey,proto3" json:"zkp_verifying_key,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_35d95cf6364b3a25, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetSlashEpoch() uint64 {
	if m != nil {
		return m.SlashEpoch
	}
	return 0
}

func (m *Params) GetEpochMaxFault() uint64 {
	if m != nil {
		return m.EpochMaxFault
	}
	return 0
}

func (m *Params) GetMinShardCount() uint64 {
	if m != nil {
		return m.MinShardCount
	}
	return 0
}

func (m *Params) GetMaxShardCount() uint64 {
	if m != nil {
		return m.MaxShardCount
	}
	return 0
}

func (m *Params) GetMaxShardSize() uint64 {
	if m != nil {
		return m.MaxShardSize
	}
	return 0
}

func (m *Params) GetChallengePeriod() time.Duration {
	if m != nil {
		return m.ChallengePeriod
	}
	return 0
}

func (m *Params) GetProofPeriod() time.Duration {
	if m != nil {
		return m.ProofPeriod
	}
	return 0
}

func (m *Params) GetChallengeCollateral() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.ChallengeCollateral
	}
	return nil
}

func (m *Params) GetZkpProvingKey() []byte {
	if m != nil {
		return m.ZkpProvingKey
	}
	return nil
}

func (m *Params) GetZkpVerifyingKey() []byte {
	if m != nil {
		return m.ZkpVerifyingKey
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "sunrise.da.Params")
}

func init() { proto.RegisterFile("sunrise/da/params.proto", fileDescriptor_35d95cf6364b3a25) }

var fileDescriptor_35d95cf6364b3a25 = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x6d, 0xd8, 0x28, 0x9b, 0xdb, 0x6e, 0x2c, 0x4c, 0x22, 0x0c, 0x29, 0xa9, 0x00, 0xa1, 0x6a,
	0x12, 0x31, 0x83, 0x1b, 0xc7, 0x6e, 0xf4, 0xc0, 0x00, 0x4d, 0x19, 0xe2, 0xc0, 0x25, 0x72, 0x1d,
	0x27, 0xb1, 0x9a, 0xc4, 0x91, 0xed, 0x54, 0x4d, 0x0f, 0xfc, 0x06, 0x8e, 0x1c, 0xb9, 0x70, 0xe1,
	0x97, 0xec, 0xb8, 0x23, 0xe2, 0xb0, 0xa1, 0xf5, 0xc2, 0xcf, 0x40, 0x76, 0x92, 0xae, 0xdc, 0xc6,
	0x29, 0xf1, 0xf3, 0x7b, 0xcf, 0xcf, 0xdf, 0xf7, 0x19, 0xdc, 0x17, 0x45, 0xc6, 0xa9, 0x20, 0x30,
	0x40, 0x30, 0x47, 0x1c, 0xa5, 0xc2, 0xcd, 0x39, 0x93, 0xcc, 0x04, 0xf5, 0x86, 0x1b, 0xa0, 0x3d,
	0x1b, 0x33, 0x91, 0x32, 0x01, 0xc7, 0x48, 0x10, 0x38, 0x3d, 0x18, 0x13, 0x89, 0x0e, 0x20, 0x66,
	0x34, 0xab, 0xb8, 0x7b, 0xbb, 0x11, 0x8b, 0x98, 0xfe, 0x85, 0xea, 0xaf, 0x46, 0xed, 0x88, 0xb1,
	0x28, 0x21, 0x50, 0xaf, 0xc6, 0x45, 0x08, 0x83, 0x82, 0x23, 0x49, 0x59, 0xad, 0x7a, 0xf4, 0xbd,
	0x0d, 0xda, 0x27, 0xfa, 0x48, 0xf3, 0x0d, 0xd8, 0x9a, 0x32, 0x49, 0x7c, 0x19, 0x73, 0x22, 0x62,
	0x96, 0x04, 0x96, 0xd1, 0x37, 0x06, 0x9b, 0xc3, 0xc7, 0x67, 0x17, 0x4e, 0xeb, 0xd7, 0x85, 0xf3,
	0xb0, 0x0a, 0x20, 0x82, 0x89, 0x4b, 0x19, 0x4c, 0x91, 0x8c, 0xdd, 0xb7, 0x24, 0x42, 0xb8, 0x3c,
	0x22, 0xd8, 0xeb, 0x29, 0xe9, 0x87, 0x46, 0x69, 0x3a, 0xa0, 0x23, 0x12, 0x24, 0x62, 0x9f, 0xe4,
	0x0c, 0xc7, 0xd6, 0xad, 0xbe, 0x31, 0x58, 0xf7, 0x80, 0x86, 0x5e, 0x2b, 0xc4, 0x7c, 0x0a, 0xb6,
	0xf5, 0x96, 0x9f, 0xa2, 0x99, 0x1f, 0xa2, 0x22, 0x91, 0xd6, 0x9a, 0x26, 0xf5, 0x34, 0xfc, 0x0e,
	0xcd, 0x46, 0x0a, 0x54, 0xa1, 0x2a, 0xa3, 0x90, 0x23, 0xac, 0x72, 0x5b, 0xeb, 0xff, 0x11, 0x4a,
	0x4b, 0x47, 0xb5, 0xd2, 0xf4, 0x80, 0xc9, 0x49, 0x9e, 0x50, 0xac, 0x0b, 0xe0, 0x87, 0x08, 0x4b,
	0xc6, 0xad, 0xdb, 0x37, 0xf7, 0xdb, 0x59, 0x91, 0x8f, 0xb4, 0x5a, 0xdd, 0x23, 0xa5, 0x99, 0x2f,
	0x62, 0xc4, 0x03, 0x1f, 0xb3, 0x22, 0x93, 0x56, 0xbb, 0xba, 0x47, 0x4a, 0xb3, 0x53, 0x85, 0x1e,
	0x2a, 0x50, 0xf3, 0xd0, 0xec, 0x1f, 0xde, 0x9d, 0x9a, 0x87, 0x66, 0x2b, 0xbc, 0x27, 0x60, 0xeb,
	0x9a, 0x27, 0xe8, 0x9c, 0x58, 0x1b, 0x9a, 0xd6, 0x6d, 0x68, 0xa7, 0x74, 0x4e, 0xcc, 0xf7, 0xe0,
	0x2e, 0x8e, 0x51, 0x92, 0x90, 0x2c, 0x22, 0x7e, 0x4e, 0x38, 0x65, 0x81, 0xb5, 0xd9, 0x37, 0x06,
	0x9d, 0x17, 0x0f, 0xdc, 0xaa, 0xe1, 0x6e, 0xd3, 0x70, 0xf7, 0xa8, 0x6e, 0xf8, 0x70, 0x43, 0x5d,
	0xf1, 0xeb, 0xa5, 0x63, 0x78, 0xdb, 0x4b, 0xf1, 0x89, 0xd6, 0x9a, 0x23, 0xd0, 0xcd, 0x39, 0x63,
	0x61, 0xe3, 0x05, 0x6e, 0xee, 0xd5, 0xd1, 0xc2, 0xda, 0xe7, 0x33, 0xd8, 0xbd, 0xce, 0x85, 0x59,
	0x92, 0x20, 0x49, 0x38, 0x4a, 0xac, 0x4e, 0x7f, 0x4d, 0xfb, 0x55, 0xc5, 0x75, 0xd5, 0x08, 0xbb,
	0xf5, 0x08, 0xbb, 0x87, 0x8c, 0x66, 0xc3, 0xe7, 0xca, 0xef, 0xc7, 0xa5, 0x33, 0x88, 0xa8, 0x8c,
	0x8b, 0xb1, 0x8b, 0x59, 0x0a, 0xeb, 0x79, 0xaf, 0x3e, 0xcf, 0x44, 0x30, 0x81, 0xb2, 0xcc, 0x89,
	0xd0, 0x02, 0xe1, 0xdd, 0x5b, 0x1e, 0x74, 0xb8, 0x3c, 0x47, 0x55, 0x79, 0x3e, 0xc9, 0xfd, 0x9c,
	0xb3, 0x29, 0xcd, 0x22, 0x7f, 0x42, 0x4a, 0xab, 0xdb, 0x37, 0x06, 0x5d, 0xaf, 0x37, 0x9f, 0xe4,
	0x27, 0x15, 0x7a, 0x4c, 0x4a, 0x73, 0x1f, 0xec, 0x28, 0xde, 0x94, 0x70, 0x1a, 0x96, 0x0d, 0xb3,
	0xa7, 0x99, 0xca, 0xe0, 0x63, 0x83, 0x1f, 0x93, 0xf2, 0xd5, 0xfa, 0x9f, 0x6f, 0x8e, 0x31, 0x3c,
	0x3a, 0xbb, 0xb2, 0x8d, 0xf3, 0x2b, 0xdb, 0xf8, 0x7d, 0x65, 0x1b, 0x5f, 0x16, 0x76, 0xeb, 0x7c,
	0x61, 0xb7, 0x7e, 0x2e, 0xec, 0xd6, 0xa7, 0xfd, 0x95, 0xc8, 0xf5, 0x73, 0x4d, 0x50, 0x49, 0x78,
	0xb3, 0x80, 0x33, 0xf5, 0xac, 0x75, 0xf4, 0x71, 0x5b, 0x57, 0xf2, 0xe5, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x90, 0x6c, 0xaf, 0xe5, 0xf1, 0x03, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.VoteThreshold.Equal(that1.VoteThreshold) {
		return false
	}
	if this.SlashEpoch != that1.SlashEpoch {
		return false
	}
	if this.EpochMaxFault != that1.EpochMaxFault {
		return false
	}
	if !this.SlashFraction.Equal(that1.SlashFraction) {
		return false
	}
	if !this.ReplicationFactor.Equal(that1.ReplicationFactor) {
		return false
	}
	if this.MinShardCount != that1.MinShardCount {
		return false
	}
	if this.MaxShardCount != that1.MaxShardCount {
		return false
	}
	if this.MaxShardSize != that1.MaxShardSize {
		return false
	}
	if this.ChallengePeriod != that1.ChallengePeriod {
		return false
	}
	if this.ProofPeriod != that1.ProofPeriod {
		return false
	}
	if len(this.ChallengeCollateral) != len(that1.ChallengeCollateral) {
		return false
	}
	for i := range this.ChallengeCollateral {
		if !this.ChallengeCollateral[i].Equal(&that1.ChallengeCollateral[i]) {
			return false
		}
	}
	if !bytes.Equal(this.ZkpProvingKey, that1.ZkpProvingKey) {
		return false
	}
	if !bytes.Equal(this.ZkpVerifyingKey, that1.ZkpVerifyingKey) {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ZkpVerifyingKey) > 0 {
		i -= len(m.ZkpVerifyingKey)
		copy(dAtA[i:], m.ZkpVerifyingKey)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ZkpVerifyingKey)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.ZkpProvingKey) > 0 {
		i -= len(m.ZkpProvingKey)
		copy(dAtA[i:], m.ZkpProvingKey)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ZkpProvingKey)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.ChallengeCollateral) > 0 {
		for iNdEx := len(m.ChallengeCollateral) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChallengeCollateral[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.ProofPeriod, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.ProofPeriod):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x52
	n2, err2 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.ChallengePeriod, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.ChallengePeriod):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintParams(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x4a
	if m.MaxShardSize != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxShardSize))
		i--
		dAtA[i] = 0x40
	}
	if m.MaxShardCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxShardCount))
		i--
		dAtA[i] = 0x38
	}
	if m.MinShardCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinShardCount))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.ReplicationFactor.Size()
		i -= size
		if _, err := m.ReplicationFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.SlashFraction.Size()
		i -= size
		if _, err := m.SlashFraction.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.EpochMaxFault != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.EpochMaxFault))
		i--
		dAtA[i] = 0x18
	}
	if m.SlashEpoch != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SlashEpoch))
		i--
		dAtA[i] = 0x10
	}
	{
		size := m.VoteThreshold.Size()
		i -= size
		if _, err := m.VoteThreshold.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.VoteThreshold.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.SlashEpoch != 0 {
		n += 1 + sovParams(uint64(m.SlashEpoch))
	}
	if m.EpochMaxFault != 0 {
		n += 1 + sovParams(uint64(m.EpochMaxFault))
	}
	l = m.SlashFraction.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.ReplicationFactor.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.MinShardCount != 0 {
		n += 1 + sovParams(uint64(m.MinShardCount))
	}
	if m.MaxShardCount != 0 {
		n += 1 + sovParams(uint64(m.MaxShardCount))
	}
	if m.MaxShardSize != 0 {
		n += 1 + sovParams(uint64(m.MaxShardSize))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.ChallengePeriod)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.ProofPeriod)
	n += 1 + l + sovParams(uint64(l))
	if len(m.ChallengeCollateral) > 0 {
		for _, e := range m.ChallengeCollateral {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = len(m.ZkpProvingKey)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.ZkpVerifyingKey)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteThreshold", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VoteThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashEpoch", wireType)
			}
			m.SlashEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SlashEpoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochMaxFault", wireType)
			}
			m.EpochMaxFault = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochMaxFault |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFraction", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashFraction.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplicationFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReplicationFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinShardCount", wireType)
			}
			m.MinShardCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinShardCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxShardCount", wireType)
			}
			m.MaxShardCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxShardCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxShardSize", wireType)
			}
			m.MaxShardSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxShardSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengePeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.ChallengePeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.ProofPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChallengeCollateral", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChallengeCollateral = append(m.ChallengeCollateral, types.Coin{})
			if err := m.ChallengeCollateral[len(m.ChallengeCollateral)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZkpProvingKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ZkpProvingKey = append(m.ZkpProvingKey[:0], dAtA[iNdEx:postIndex]...)
			if m.ZkpProvingKey == nil {
				m.ZkpProvingKey = []byte{}
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZkpVerifyingKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ZkpVerifyingKey = append(m.ZkpVerifyingKey[:0], dAtA[iNdEx:postIndex]...)
			if m.ZkpVerifyingKey == nil {
				m.ZkpVerifyingKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
