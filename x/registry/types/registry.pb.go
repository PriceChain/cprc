// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: registry/registry.proto

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

type Registry struct {
	Id                    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	StakedAmount          string `protobuf:"bytes,3,opt,name=stakedAmount,proto3" json:"stakedAmount,omitempty"`
	Status                string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Price                 string `protobuf:"bytes,5,opt,name=price,proto3" json:"price,omitempty"`
	Quorum                string `protobuf:"bytes,6,opt,name=quorum,proto3" json:"quorum,omitempty"`
	ConsensusExpiringTime string `protobuf:"bytes,7,opt,name=consensusExpiringTime,proto3" json:"consensusExpiringTime,omitempty"`
	ActiveMembers         string `protobuf:"bytes,8,opt,name=activeMembers,proto3" json:"activeMembers,omitempty"`
	ProdInfo              string `protobuf:"bytes,9,opt,name=prodInfo,proto3" json:"prodInfo,omitempty"`
	Memo                  string `protobuf:"bytes,10,opt,name=memo,proto3" json:"memo,omitempty"`
	Reserved              string `protobuf:"bytes,11,opt,name=reserved,proto3" json:"reserved,omitempty"`
	Creator               string `protobuf:"bytes,12,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Registry) Reset()         { *m = Registry{} }
func (m *Registry) String() string { return proto.CompactTextString(m) }
func (*Registry) ProtoMessage()    {}
func (*Registry) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3f64dc2c9630278, []int{0}
}
func (m *Registry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Registry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Registry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Registry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Registry.Merge(m, src)
}
func (m *Registry) XXX_Size() int {
	return m.Size()
}
func (m *Registry) XXX_DiscardUnknown() {
	xxx_messageInfo_Registry.DiscardUnknown(m)
}

var xxx_messageInfo_Registry proto.InternalMessageInfo

func (m *Registry) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Registry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Registry) GetStakedAmount() string {
	if m != nil {
		return m.StakedAmount
	}
	return ""
}

func (m *Registry) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Registry) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *Registry) GetQuorum() string {
	if m != nil {
		return m.Quorum
	}
	return ""
}

func (m *Registry) GetConsensusExpiringTime() string {
	if m != nil {
		return m.ConsensusExpiringTime
	}
	return ""
}

func (m *Registry) GetActiveMembers() string {
	if m != nil {
		return m.ActiveMembers
	}
	return ""
}

func (m *Registry) GetProdInfo() string {
	if m != nil {
		return m.ProdInfo
	}
	return ""
}

func (m *Registry) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *Registry) GetReserved() string {
	if m != nil {
		return m.Reserved
	}
	return ""
}

func (m *Registry) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Registry)(nil), "pricechain.cprc.registry.Registry")
}

func init() { proto.RegisterFile("registry/registry.proto", fileDescriptor_f3f64dc2c9630278) }

var fileDescriptor_f3f64dc2c9630278 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xbd, 0x4a, 0x33, 0x41,
	0x14, 0x86, 0xb3, 0xfb, 0xe5, 0x77, 0xbe, 0x68, 0x31, 0xf8, 0x73, 0xb0, 0x58, 0x42, 0xb0, 0x08,
	0x08, 0xbb, 0x85, 0xde, 0x80, 0x4a, 0x0a, 0x0b, 0x41, 0x82, 0x95, 0xdd, 0x66, 0xf6, 0x98, 0x0c,
	0x32, 0x3b, 0xeb, 0x99, 0xd9, 0x90, 0xdc, 0x85, 0x97, 0x65, 0x99, 0x52, 0xac, 0x24, 0xb9, 0x11,
	0x99, 0xd9, 0x6c, 0x20, 0x60, 0x77, 0x9e, 0xf3, 0x3e, 0xc5, 0x9c, 0x79, 0xd9, 0x39, 0xe1, 0x4c,
	0x1a, 0x4b, 0xab, 0xa4, 0x1e, 0xe2, 0x82, 0xb4, 0xd5, 0x1c, 0x0a, 0x92, 0x02, 0xc5, 0x3c, 0x95,
	0x79, 0x2c, 0x0a, 0x12, 0x71, 0x9d, 0x0f, 0xbf, 0x43, 0xd6, 0x9d, 0xec, 0x80, 0x1f, 0xb3, 0x50,
	0x66, 0x10, 0x0c, 0x82, 0x51, 0x73, 0x12, 0xca, 0x8c, 0x73, 0xd6, 0xcc, 0x53, 0x85, 0x10, 0x0e,
	0x82, 0x51, 0x6f, 0xe2, 0x67, 0x3e, 0x64, 0x7d, 0x63, 0xd3, 0x37, 0xcc, 0x6e, 0x95, 0x2e, 0x73,
	0x0b, 0xff, 0x7c, 0x76, 0xb0, 0xe3, 0x67, 0xac, 0x6d, 0x6c, 0x6a, 0x4b, 0x03, 0x4d, 0x9f, 0xee,
	0x88, 0x9f, 0xb0, 0x96, 0x7f, 0x08, 0xb4, 0xfc, 0xba, 0x02, 0x67, 0xbf, 0x97, 0x9a, 0x4a, 0x05,
	0xed, 0xca, 0xae, 0x88, 0xdf, 0xb0, 0x53, 0xa1, 0x73, 0x83, 0xb9, 0x29, 0xcd, 0x78, 0x59, 0x48,
	0x92, 0xf9, 0xec, 0x59, 0x2a, 0x84, 0x8e, 0xd7, 0xfe, 0x0e, 0xf9, 0x25, 0x3b, 0x4a, 0x85, 0x95,
	0x0b, 0x7c, 0x44, 0x35, 0x45, 0x32, 0xd0, 0xf5, 0xf6, 0xe1, 0x92, 0x5f, 0xb0, 0x6e, 0x41, 0x3a,
	0x7b, 0xc8, 0x5f, 0x35, 0xf4, 0xbc, 0xb0, 0x67, 0x77, 0xb5, 0x42, 0xa5, 0x81, 0x55, 0x57, 0xbb,
	0xd9, 0xf9, 0x84, 0x06, 0x69, 0x81, 0x19, 0xfc, 0xaf, 0xfc, 0x9a, 0x39, 0xb0, 0x8e, 0x20, 0x4c,
	0xad, 0x26, 0xe8, 0xfb, 0xa8, 0xc6, 0xbb, 0xf1, 0xe7, 0x26, 0x0a, 0xd6, 0x9b, 0x28, 0xf8, 0xd9,
	0x44, 0xc1, 0xc7, 0x36, 0x6a, 0xac, 0xb7, 0x51, 0xe3, 0x6b, 0x1b, 0x35, 0x5e, 0xae, 0x66, 0xd2,
	0xce, 0xcb, 0x69, 0x2c, 0xb4, 0x4a, 0x9e, 0xdc, 0x2f, 0xdc, 0xbb, 0x6e, 0x12, 0xd7, 0x4d, 0xb2,
	0xdc, 0xb7, 0x97, 0xd8, 0x55, 0x81, 0x66, 0xda, 0xf6, 0x25, 0x5e, 0xff, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x7f, 0xa4, 0x8a, 0xe4, 0xdf, 0x01, 0x00, 0x00,
}

func (m *Registry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Registry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Registry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintRegistry(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.Reserved) > 0 {
		i -= len(m.Reserved)
		copy(dAtA[i:], m.Reserved)
		i = encodeVarintRegistry(dAtA, i, uint64(len(m.Reserved)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.Memo) > 0 {
		i -= len(m.Memo)
		copy(dAtA[i:], m.Memo)
		i = encodeVarintRegistry(dAtA, i, uint64(len(m.Memo)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.ProdInfo) > 0 {
		i -= len(m.ProdInfo)
		copy(dAtA[i:], m.ProdInfo)
		i = encodeVarintRegistry(dAtA, i, uint64(len(m.ProdInfo)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.ActiveMembers) > 0 {
		i -= len(m.ActiveMembers)
		copy(dAtA[i:], m.ActiveMembers)
		i = encodeVarintRegistry(dAtA, i, uint64(len(m.ActiveMembers)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.ConsensusExpiringTime) > 0 {
		i -= len(m.ConsensusExpiringTime)
		copy(dAtA[i:], m.ConsensusExpiringTime)
		i = encodeVarintRegistry(dAtA, i, uint64(len(m.ConsensusExpiringTime)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Quorum) > 0 {
		i -= len(m.Quorum)
		copy(dAtA[i:], m.Quorum)
		i = encodeVarintRegistry(dAtA, i, uint64(len(m.Quorum)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Price) > 0 {
		i -= len(m.Price)
		copy(dAtA[i:], m.Price)
		i = encodeVarintRegistry(dAtA, i, uint64(len(m.Price)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintRegistry(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.StakedAmount) > 0 {
		i -= len(m.StakedAmount)
		copy(dAtA[i:], m.StakedAmount)
		i = encodeVarintRegistry(dAtA, i, uint64(len(m.StakedAmount)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintRegistry(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintRegistry(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRegistry(dAtA []byte, offset int, v uint64) int {
	offset -= sovRegistry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Registry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovRegistry(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	l = len(m.StakedAmount)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	l = len(m.Price)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	l = len(m.Quorum)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	l = len(m.ConsensusExpiringTime)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	l = len(m.ActiveMembers)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	l = len(m.ProdInfo)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	l = len(m.Reserved)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	return n
}

func sovRegistry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRegistry(x uint64) (n int) {
	return sovRegistry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Registry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRegistry
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
			return fmt.Errorf("proto: Registry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Registry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
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
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
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
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakedAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
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
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakedAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
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
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
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
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quorum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
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
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Quorum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusExpiringTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
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
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsensusExpiringTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveMembers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
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
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ActiveMembers = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProdInfo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
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
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProdInfo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
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
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Memo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reserved", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
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
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reserved = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegistry
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
				return ErrInvalidLengthRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRegistry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRegistry
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
func skipRegistry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRegistry
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
					return 0, ErrIntOverflowRegistry
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
					return 0, ErrIntOverflowRegistry
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
				return 0, ErrInvalidLengthRegistry
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRegistry
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRegistry
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRegistry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRegistry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRegistry = fmt.Errorf("proto: unexpected end of group")
)
