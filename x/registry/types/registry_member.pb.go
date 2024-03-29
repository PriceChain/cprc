// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: registry/registry_member.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type RegistryMember struct {
	Id             uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Wallet         string   `protobuf:"bytes,2,opt,name=wallet,proto3" json:"wallet,omitempty"`
	RegistryId     string   `protobuf:"bytes,3,opt,name=registryId,proto3" json:"registryId,omitempty"`
	StakedAmount   string   `protobuf:"bytes,4,opt,name=stakedAmount,proto3" json:"stakedAmount,omitempty"`
	Address        string   `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Status         string   `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	PopCount       string   `protobuf:"bytes,7,opt,name=popCount,proto3" json:"popCount,omitempty"`
	Level          string   `protobuf:"bytes,8,opt,name=level,proto3" json:"level,omitempty"`
	Reputations    []string `protobuf:"bytes,9,rep,name=reputations,proto3" json:"reputations,omitempty"`
	Scores         []string `protobuf:"bytes,10,rep,name=scores,proto3" json:"scores,omitempty"`
	RewardSum      string   `protobuf:"bytes,11,opt,name=rewardSum,proto3" json:"rewardSum,omitempty"`
	RewardPaid     string   `protobuf:"bytes,12,opt,name=rewardPaid,proto3" json:"rewardPaid,omitempty"`
	RewardPaidDate string   `protobuf:"bytes,13,opt,name=rewardPaidDate,proto3" json:"rewardPaidDate,omitempty"`
	Reserved       string   `protobuf:"bytes,14,opt,name=reserved,proto3" json:"reserved,omitempty"`
}

func (m *RegistryMember) Reset()         { *m = RegistryMember{} }
func (m *RegistryMember) String() string { return proto.CompactTextString(m) }
func (*RegistryMember) ProtoMessage()    {}
func (*RegistryMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b8caeb548177525, []int{0}
}
func (m *RegistryMember) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegistryMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegistryMember.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegistryMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistryMember.Merge(m, src)
}
func (m *RegistryMember) XXX_Size() int {
	return m.Size()
}
func (m *RegistryMember) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistryMember.DiscardUnknown(m)
}

var xxx_messageInfo_RegistryMember proto.InternalMessageInfo

func (m *RegistryMember) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RegistryMember) GetWallet() string {
	if m != nil {
		return m.Wallet
	}
	return ""
}

func (m *RegistryMember) GetRegistryId() string {
	if m != nil {
		return m.RegistryId
	}
	return ""
}

func (m *RegistryMember) GetStakedAmount() string {
	if m != nil {
		return m.StakedAmount
	}
	return ""
}

func (m *RegistryMember) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RegistryMember) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *RegistryMember) GetPopCount() string {
	if m != nil {
		return m.PopCount
	}
	return ""
}

func (m *RegistryMember) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *RegistryMember) GetReputations() []string {
	if m != nil {
		return m.Reputations
	}
	return nil
}

func (m *RegistryMember) GetScores() []string {
	if m != nil {
		return m.Scores
	}
	return nil
}

func (m *RegistryMember) GetRewardSum() string {
	if m != nil {
		return m.RewardSum
	}
	return ""
}

func (m *RegistryMember) GetRewardPaid() string {
	if m != nil {
		return m.RewardPaid
	}
	return ""
}

func (m *RegistryMember) GetRewardPaidDate() string {
	if m != nil {
		return m.RewardPaidDate
	}
	return ""
}

func (m *RegistryMember) GetReserved() string {
	if m != nil {
		return m.Reserved
	}
	return ""
}

func init() {
	proto.RegisterType((*RegistryMember)(nil), "pricechain.cprc.registry.RegistryMember")
}

func init() { proto.RegisterFile("registry/registry_member.proto", fileDescriptor_3b8caeb548177525) }

var fileDescriptor_3b8caeb548177525 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xcf, 0x6a, 0xe3, 0x30,
	0x10, 0xc6, 0xe3, 0xfc, 0xb7, 0x92, 0xf5, 0x41, 0x2c, 0xcb, 0xb0, 0x2c, 0xc2, 0xe4, 0xb0, 0x04,
	0x0a, 0xf6, 0xa1, 0x4f, 0xd0, 0xa6, 0x3d, 0xf4, 0x50, 0x08, 0xee, 0xad, 0x97, 0xa2, 0x58, 0x43,
	0x22, 0x6a, 0xc7, 0x46, 0x92, 0x93, 0xe6, 0x2d, 0xfa, 0x00, 0x7d, 0xa0, 0x1e, 0x73, 0xec, 0xb1,
	0x24, 0x2f, 0x52, 0x2c, 0xc7, 0x49, 0xda, 0x9b, 0x7e, 0xdf, 0x37, 0x9a, 0xe1, 0x63, 0x86, 0x30,
	0x85, 0x73, 0xa9, 0x8d, 0xda, 0x84, 0xf5, 0xe3, 0x29, 0xc5, 0x74, 0x86, 0x2a, 0xc8, 0x55, 0x66,
	0x32, 0x0a, 0xb9, 0x92, 0x31, 0xc6, 0x0b, 0x2e, 0x97, 0x41, 0x9c, 0xab, 0x38, 0xa8, 0xcb, 0x46,
	0x6f, 0x2d, 0xe2, 0x45, 0x07, 0xb8, 0xb7, 0x5f, 0xa8, 0x47, 0x9a, 0x52, 0x80, 0xe3, 0x3b, 0xe3,
	0x76, 0xd4, 0x94, 0x82, 0xfe, 0x21, 0xdd, 0x35, 0x4f, 0x12, 0x34, 0xd0, 0xf4, 0x9d, 0xb1, 0x1b,
	0x1d, 0x88, 0x32, 0x42, 0xea, 0x36, 0x77, 0x02, 0x5a, 0xd6, 0x3b, 0x53, 0xe8, 0x88, 0x0c, 0xb5,
	0xe1, 0xcf, 0x28, 0xae, 0xd2, 0xac, 0x58, 0x1a, 0x68, 0xdb, 0x8a, 0x6f, 0x1a, 0x05, 0xd2, 0xe3,
	0x42, 0x28, 0xd4, 0x1a, 0x3a, 0xd6, 0xae, 0xb1, 0x9c, 0xaa, 0x0d, 0x37, 0x85, 0x86, 0x6e, 0x35,
	0xb5, 0x22, 0xfa, 0x97, 0xf4, 0xf3, 0x2c, 0x9f, 0xd8, 0x8e, 0x3d, 0xeb, 0x1c, 0x99, 0xfe, 0x26,
	0x9d, 0x04, 0x57, 0x98, 0x40, 0xdf, 0x1a, 0x15, 0x50, 0x9f, 0x0c, 0x14, 0xe6, 0x85, 0xe1, 0x46,
	0x66, 0x4b, 0x0d, 0xae, 0xdf, 0x1a, 0xbb, 0xd1, 0xb9, 0x64, 0x67, 0xc5, 0x99, 0x42, 0x0d, 0xc4,
	0x9a, 0x07, 0xa2, 0xff, 0x88, 0xab, 0x70, 0xcd, 0x95, 0x78, 0x28, 0x52, 0x18, 0xd8, 0x9e, 0x27,
	0xa1, 0xca, 0x5f, 0xc2, 0x94, 0x4b, 0x01, 0xc3, 0x3a, 0x7f, 0xad, 0xd0, 0xff, 0xc4, 0x3b, 0xd1,
	0x0d, 0x37, 0x08, 0xbf, 0x6c, 0xcd, 0x0f, 0xb5, 0x4c, 0xa4, 0x50, 0xa3, 0x5a, 0xa1, 0x00, 0xaf,
	0x4a, 0x54, 0xf3, 0xf5, 0xed, 0xfb, 0x8e, 0x39, 0xdb, 0x1d, 0x73, 0x3e, 0x77, 0xcc, 0x79, 0xdd,
	0xb3, 0xc6, 0x76, 0xcf, 0x1a, 0x1f, 0x7b, 0xd6, 0x78, 0xbc, 0x98, 0x4b, 0xb3, 0x28, 0x66, 0x41,
	0x9c, 0xa5, 0xe1, 0xb4, 0xdc, 0xee, 0xa4, 0xdc, 0x6e, 0x58, 0x6e, 0x37, 0x7c, 0x39, 0x9e, 0x41,
	0x68, 0x36, 0x39, 0xea, 0x59, 0xd7, 0x9e, 0xc1, 0xe5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x56,
	0x09, 0x11, 0x15, 0x28, 0x02, 0x00, 0x00,
}

func (m *RegistryMember) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegistryMember) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegistryMember) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Reserved) > 0 {
		i -= len(m.Reserved)
		copy(dAtA[i:], m.Reserved)
		i = encodeVarintRegistryMember(dAtA, i, uint64(len(m.Reserved)))
		i--
		dAtA[i] = 0x72
	}
	if len(m.RewardPaidDate) > 0 {
		i -= len(m.RewardPaidDate)
		copy(dAtA[i:], m.RewardPaidDate)
		i = encodeVarintRegistryMember(dAtA, i, uint64(len(m.RewardPaidDate)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.RewardPaid) > 0 {
		i -= len(m.RewardPaid)
		copy(dAtA[i:], m.RewardPaid)
		i = encodeVarintRegistryMember(dAtA, i, uint64(len(m.RewardPaid)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.RewardSum) > 0 {
		i -= len(m.RewardSum)
		copy(dAtA[i:], m.RewardSum)
		i = encodeVarintRegistryMember(dAtA, i, uint64(len(m.RewardSum)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.Scores) > 0 {
		for iNdEx := len(m.Scores) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Scores[iNdEx])
			copy(dAtA[i:], m.Scores[iNdEx])
			i = encodeVarintRegistryMember(dAtA, i, uint64(len(m.Scores[iNdEx])))
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.Reputations) > 0 {
		for iNdEx := len(m.Reputations) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Reputations[iNdEx])
			copy(dAtA[i:], m.Reputations[iNdEx])
			i = encodeVarintRegistryMember(dAtA, i, uint64(len(m.Reputations[iNdEx])))
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.Level) > 0 {
		i -= len(m.Level)
		copy(dAtA[i:], m.Level)
		i = encodeVarintRegistryMember(dAtA, i, uint64(len(m.Level)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.PopCount) > 0 {
		i -= len(m.PopCount)
		copy(dAtA[i:], m.PopCount)
		i = encodeVarintRegistryMember(dAtA, i, uint64(len(m.PopCount)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintRegistryMember(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintRegistryMember(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.StakedAmount) > 0 {
		i -= len(m.StakedAmount)
		copy(dAtA[i:], m.StakedAmount)
		i = encodeVarintRegistryMember(dAtA, i, uint64(len(m.StakedAmount)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RegistryId) > 0 {
		i -= len(m.RegistryId)
		copy(dAtA[i:], m.RegistryId)
		i = encodeVarintRegistryMember(dAtA, i, uint64(len(m.RegistryId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Wallet) > 0 {
		i -= len(m.Wallet)
		copy(dAtA[i:], m.Wallet)
		i = encodeVarintRegistryMember(dAtA, i, uint64(len(m.Wallet)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintRegistryMember(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRegistryMember(dAtA []byte, offset int, v uint64) int {
	offset -= sovRegistryMember(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RegistryMember) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovRegistryMember(uint64(m.Id))
	}
	l = len(m.Wallet)
	if l > 0 {
		n += 1 + l + sovRegistryMember(uint64(l))
	}
	l = len(m.RegistryId)
	if l > 0 {
		n += 1 + l + sovRegistryMember(uint64(l))
	}
	l = len(m.StakedAmount)
	if l > 0 {
		n += 1 + l + sovRegistryMember(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovRegistryMember(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovRegistryMember(uint64(l))
	}
	l = len(m.PopCount)
	if l > 0 {
		n += 1 + l + sovRegistryMember(uint64(l))
	}
	l = len(m.Level)
	if l > 0 {
		n += 1 + l + sovRegistryMember(uint64(l))
	}
	if len(m.Reputations) > 0 {
		for _, s := range m.Reputations {
			l = len(s)
			n += 1 + l + sovRegistryMember(uint64(l))
		}
	}
	if len(m.Scores) > 0 {
		for _, s := range m.Scores {
			l = len(s)
			n += 1 + l + sovRegistryMember(uint64(l))
		}
	}
	l = len(m.RewardSum)
	if l > 0 {
		n += 1 + l + sovRegistryMember(uint64(l))
	}
	l = len(m.RewardPaid)
	if l > 0 {
		n += 1 + l + sovRegistryMember(uint64(l))
	}
	l = len(m.RewardPaidDate)
	if l > 0 {
		n += 1 + l + sovRegistryMember(uint64(l))
	}
	l = len(m.Reserved)
	if l > 0 {
		n += 1 + l + sovRegistryMember(uint64(l))
	}
	return n
}

func sovRegistryMember(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRegistryMember(x uint64) (n int) {
	return sovRegistryMember(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RegistryMember) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRegistryMember
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
			return fmt.Errorf("proto: RegistryMember: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegistryMember: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistryMember
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Wallet", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistryMember
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
				return ErrInvalidLengthRegistryMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegistryMember
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Wallet = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegistryId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistryMember
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
				return ErrInvalidLengthRegistryMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegistryMember
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RegistryId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakedAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistryMember
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
				return ErrInvalidLengthRegistryMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegistryMember
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakedAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistryMember
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
				return ErrInvalidLengthRegistryMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegistryMember
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistryMember
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
				return ErrInvalidLengthRegistryMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegistryMember
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PopCount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistryMember
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
				return ErrInvalidLengthRegistryMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegistryMember
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PopCount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistryMember
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
				return ErrInvalidLengthRegistryMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegistryMember
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Level = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reputations", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistryMember
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
				return ErrInvalidLengthRegistryMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegistryMember
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reputations = append(m.Reputations, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scores", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistryMember
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
				return ErrInvalidLengthRegistryMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegistryMember
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Scores = append(m.Scores, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardSum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistryMember
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
				return ErrInvalidLengthRegistryMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegistryMember
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardSum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardPaid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistryMember
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
				return ErrInvalidLengthRegistryMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegistryMember
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardPaid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardPaidDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistryMember
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
				return ErrInvalidLengthRegistryMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegistryMember
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardPaidDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reserved", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistryMember
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
				return ErrInvalidLengthRegistryMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegistryMember
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reserved = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRegistryMember(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRegistryMember
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
func skipRegistryMember(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRegistryMember
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
					return 0, ErrIntOverflowRegistryMember
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
					return 0, ErrIntOverflowRegistryMember
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
				return 0, ErrInvalidLengthRegistryMember
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRegistryMember
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRegistryMember
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRegistryMember        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRegistryMember          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRegistryMember = fmt.Errorf("proto: unexpected end of group")
)
