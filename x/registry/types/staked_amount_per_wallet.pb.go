// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: registry/staked_amount_per_wallet.proto

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

type StakedAmountPerWallet struct {
	Index        string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Wallet       string `protobuf:"bytes,2,opt,name=wallet,proto3" json:"wallet,omitempty"`
	StakedAmount string `protobuf:"bytes,3,opt,name=stakedAmount,proto3" json:"stakedAmount,omitempty"`
}

func (m *StakedAmountPerWallet) Reset()         { *m = StakedAmountPerWallet{} }
func (m *StakedAmountPerWallet) String() string { return proto.CompactTextString(m) }
func (*StakedAmountPerWallet) ProtoMessage()    {}
func (*StakedAmountPerWallet) Descriptor() ([]byte, []int) {
	return fileDescriptor_3230b7fb836c91d8, []int{0}
}
func (m *StakedAmountPerWallet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StakedAmountPerWallet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StakedAmountPerWallet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StakedAmountPerWallet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakedAmountPerWallet.Merge(m, src)
}
func (m *StakedAmountPerWallet) XXX_Size() int {
	return m.Size()
}
func (m *StakedAmountPerWallet) XXX_DiscardUnknown() {
	xxx_messageInfo_StakedAmountPerWallet.DiscardUnknown(m)
}

var xxx_messageInfo_StakedAmountPerWallet proto.InternalMessageInfo

func (m *StakedAmountPerWallet) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *StakedAmountPerWallet) GetWallet() string {
	if m != nil {
		return m.Wallet
	}
	return ""
}

func (m *StakedAmountPerWallet) GetStakedAmount() string {
	if m != nil {
		return m.StakedAmount
	}
	return ""
}

func init() {
	proto.RegisterType((*StakedAmountPerWallet)(nil), "pricechain.cprc.registry.StakedAmountPerWallet")
}

func init() {
	proto.RegisterFile("registry/staked_amount_per_wallet.proto", fileDescriptor_3230b7fb836c91d8)
}

var fileDescriptor_3230b7fb836c91d8 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0x4a, 0x4d, 0xcf,
	0x2c, 0x2e, 0x29, 0xaa, 0xd4, 0x2f, 0x2e, 0x49, 0xcc, 0x4e, 0x4d, 0x89, 0x4f, 0xcc, 0xcd, 0x2f,
	0xcd, 0x2b, 0x89, 0x2f, 0x48, 0x2d, 0x8a, 0x2f, 0x4f, 0xcc, 0xc9, 0x49, 0x2d, 0xd1, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x92, 0x28, 0x28, 0xca, 0x4c, 0x4e, 0x4d, 0xce, 0x48, 0xcc, 0xcc, 0xd3,
	0x4b, 0x2e, 0x28, 0x4a, 0xd6, 0x83, 0x69, 0x54, 0xca, 0xe4, 0x12, 0x0d, 0x06, 0xeb, 0x75, 0x04,
	0x6b, 0x0d, 0x48, 0x2d, 0x0a, 0x07, 0x6b, 0x14, 0x12, 0xe1, 0x62, 0xcd, 0xcc, 0x4b, 0x49, 0xad,
	0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84, 0xc4, 0xb8, 0xd8, 0x20, 0x06, 0x4b,
	0x30, 0x81, 0x85, 0xa1, 0x3c, 0x21, 0x25, 0x2e, 0x9e, 0x62, 0x24, 0x63, 0x24, 0x98, 0xc1, 0xb2,
	0x28, 0x62, 0x4e, 0xae, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c,
	0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x9d,
	0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x1f, 0x00, 0x72, 0xa9, 0x33, 0xc8,
	0xa5, 0xfa, 0x20, 0x97, 0xea, 0x57, 0xe8, 0xc3, 0x3d, 0x59, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4,
	0x06, 0xf6, 0x92, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x74, 0x24, 0x6a, 0x4d, 0xfd, 0x00, 0x00,
	0x00,
}

func (m *StakedAmountPerWallet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StakedAmountPerWallet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StakedAmountPerWallet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StakedAmount) > 0 {
		i -= len(m.StakedAmount)
		copy(dAtA[i:], m.StakedAmount)
		i = encodeVarintStakedAmountPerWallet(dAtA, i, uint64(len(m.StakedAmount)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Wallet) > 0 {
		i -= len(m.Wallet)
		copy(dAtA[i:], m.Wallet)
		i = encodeVarintStakedAmountPerWallet(dAtA, i, uint64(len(m.Wallet)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintStakedAmountPerWallet(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStakedAmountPerWallet(dAtA []byte, offset int, v uint64) int {
	offset -= sovStakedAmountPerWallet(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StakedAmountPerWallet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovStakedAmountPerWallet(uint64(l))
	}
	l = len(m.Wallet)
	if l > 0 {
		n += 1 + l + sovStakedAmountPerWallet(uint64(l))
	}
	l = len(m.StakedAmount)
	if l > 0 {
		n += 1 + l + sovStakedAmountPerWallet(uint64(l))
	}
	return n
}

func sovStakedAmountPerWallet(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStakedAmountPerWallet(x uint64) (n int) {
	return sovStakedAmountPerWallet(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StakedAmountPerWallet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStakedAmountPerWallet
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
			return fmt.Errorf("proto: StakedAmountPerWallet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StakedAmountPerWallet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakedAmountPerWallet
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
				return ErrInvalidLengthStakedAmountPerWallet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStakedAmountPerWallet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Wallet", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakedAmountPerWallet
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
				return ErrInvalidLengthStakedAmountPerWallet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStakedAmountPerWallet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Wallet = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakedAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakedAmountPerWallet
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
				return ErrInvalidLengthStakedAmountPerWallet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStakedAmountPerWallet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakedAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStakedAmountPerWallet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStakedAmountPerWallet
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
func skipStakedAmountPerWallet(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStakedAmountPerWallet
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
					return 0, ErrIntOverflowStakedAmountPerWallet
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
					return 0, ErrIntOverflowStakedAmountPerWallet
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
				return 0, ErrInvalidLengthStakedAmountPerWallet
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStakedAmountPerWallet
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStakedAmountPerWallet
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStakedAmountPerWallet        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStakedAmountPerWallet          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStakedAmountPerWallet = fmt.Errorf("proto: unexpected end of group")
)
