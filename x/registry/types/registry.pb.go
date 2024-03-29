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
	Id           uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	StakedAmount string   `protobuf:"bytes,3,opt,name=stakedAmount,proto3" json:"stakedAmount,omitempty"`
	Status       string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Description  string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	ImageUrl     string   `protobuf:"bytes,6,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	PriceCount   string   `protobuf:"bytes,7,opt,name=priceCount,proto3" json:"priceCount,omitempty"`
	ReviewCount  string   `protobuf:"bytes,8,opt,name=reviewCount,proto3" json:"reviewCount,omitempty"`
	Timestamp    string   `protobuf:"bytes,9,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Reserved     string   `protobuf:"bytes,10,opt,name=reserved,proto3" json:"reserved,omitempty"`
	Owners       []string `protobuf:"bytes,11,rep,name=owners,proto3" json:"owners,omitempty"`
	Validators   []string `protobuf:"bytes,12,rep,name=validators,proto3" json:"validators,omitempty"`
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

func (m *Registry) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Registry) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *Registry) GetPriceCount() string {
	if m != nil {
		return m.PriceCount
	}
	return ""
}

func (m *Registry) GetReviewCount() string {
	if m != nil {
		return m.ReviewCount
	}
	return ""
}

func (m *Registry) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *Registry) GetReserved() string {
	if m != nil {
		return m.Reserved
	}
	return ""
}

func (m *Registry) GetOwners() []string {
	if m != nil {
		return m.Owners
	}
	return nil
}

func (m *Registry) GetValidators() []string {
	if m != nil {
		return m.Validators
	}
	return nil
}

func init() {
	proto.RegisterType((*Registry)(nil), "pricechain.cprc.registry.Registry")
}

func init() { proto.RegisterFile("registry/registry.proto", fileDescriptor_f3f64dc2c9630278) }

var fileDescriptor_f3f64dc2c9630278 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xc1, 0x4e, 0x02, 0x31,
	0x10, 0x86, 0xd9, 0x05, 0x11, 0x0a, 0xf1, 0xd0, 0x83, 0x36, 0xc6, 0x34, 0x1b, 0x4e, 0x24, 0x26,
	0xbb, 0x07, 0x9f, 0x40, 0x8d, 0x77, 0x43, 0xe2, 0xc5, 0x5b, 0xd9, 0x9d, 0x40, 0x23, 0xbb, 0x6d,
	0xda, 0x01, 0xe4, 0x2d, 0x7c, 0x2c, 0x8f, 0x1c, 0x3d, 0x78, 0x30, 0xf0, 0x22, 0xa6, 0xb3, 0x80,
	0xeb, 0x6d, 0xfe, 0xff, 0x6b, 0x66, 0xfa, 0xe7, 0x67, 0x57, 0x0e, 0x66, 0xda, 0xa3, 0xdb, 0x64,
	0xc7, 0x21, 0xb5, 0xce, 0xa0, 0xe1, 0xc2, 0x3a, 0x9d, 0x43, 0x3e, 0x57, 0xba, 0x4a, 0x73, 0xeb,
	0xf2, 0xf4, 0xc8, 0x47, 0xdf, 0x31, 0xeb, 0x4d, 0x0e, 0x82, 0x5f, 0xb0, 0x58, 0x17, 0x22, 0x4a,
	0xa2, 0x71, 0x67, 0x12, 0xeb, 0x82, 0x73, 0xd6, 0xa9, 0x54, 0x09, 0x22, 0x4e, 0xa2, 0x71, 0x7f,
	0x42, 0x33, 0x1f, 0xb1, 0xa1, 0x47, 0xf5, 0x06, 0xc5, 0x7d, 0x69, 0x96, 0x15, 0x8a, 0x36, 0xb1,
	0x7f, 0x1e, 0xbf, 0x64, 0x5d, 0x8f, 0x0a, 0x97, 0x5e, 0x74, 0x88, 0x1e, 0x14, 0x4f, 0xd8, 0xa0,
	0x00, 0x9f, 0x3b, 0x6d, 0x51, 0x9b, 0x4a, 0x9c, 0x11, 0x6c, 0x5a, 0xfc, 0x9a, 0xf5, 0x74, 0xa9,
	0x66, 0xf0, 0xe2, 0x16, 0xa2, 0x4b, 0xf8, 0xa4, 0xb9, 0x64, 0x8c, 0x62, 0x3c, 0xd2, 0xdd, 0x73,
	0xa2, 0x0d, 0x27, 0x6c, 0x77, 0xb0, 0xd2, 0xb0, 0xae, 0x1f, 0xf4, 0xea, 0xed, 0x0d, 0x8b, 0xdf,
	0xb0, 0x3e, 0xea, 0x12, 0x3c, 0xaa, 0xd2, 0x8a, 0x3e, 0xf1, 0x3f, 0x23, 0xdc, 0x76, 0xe0, 0xc1,
	0xad, 0xa0, 0x10, 0xac, 0xbe, 0x7d, 0xd4, 0x21, 0x91, 0x59, 0x57, 0xe0, 0xbc, 0x18, 0x24, 0xed,
	0x90, 0xa8, 0x56, 0xe1, 0x4f, 0x2b, 0xb5, 0xd0, 0x85, 0x42, 0xe3, 0xbc, 0x18, 0x12, 0x6b, 0x38,
	0x0f, 0x4f, 0x9f, 0x3b, 0x19, 0x6d, 0x77, 0x32, 0xfa, 0xd9, 0xc9, 0xe8, 0x63, 0x2f, 0x5b, 0xdb,
	0xbd, 0x6c, 0x7d, 0xed, 0x65, 0xeb, 0xf5, 0x76, 0xa6, 0x71, 0xbe, 0x9c, 0xa6, 0xb9, 0x29, 0xb3,
	0x67, 0x0a, 0x11, 0xda, 0xc9, 0x42, 0x3b, 0xd9, 0xfb, 0xa9, 0xbf, 0x0c, 0x37, 0x16, 0xfc, 0xb4,
	0x4b, 0x35, 0xde, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0x05, 0x91, 0x22, 0x47, 0xe1, 0x01, 0x00,
	0x00,
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
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Validators[iNdEx])
			copy(dAtA[i:], m.Validators[iNdEx])
			i = encodeVarintRegistry(dAtA, i, uint64(len(m.Validators[iNdEx])))
			i--
			dAtA[i] = 0x62
		}
	}
	if len(m.Owners) > 0 {
		for iNdEx := len(m.Owners) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Owners[iNdEx])
			copy(dAtA[i:], m.Owners[iNdEx])
			i = encodeVarintRegistry(dAtA, i, uint64(len(m.Owners[iNdEx])))
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.Reserved) > 0 {
		i -= len(m.Reserved)
		copy(dAtA[i:], m.Reserved)
		i = encodeVarintRegistry(dAtA, i, uint64(len(m.Reserved)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Timestamp) > 0 {
		i -= len(m.Timestamp)
		copy(dAtA[i:], m.Timestamp)
		i = encodeVarintRegistry(dAtA, i, uint64(len(m.Timestamp)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.ReviewCount) > 0 {
		i -= len(m.ReviewCount)
		copy(dAtA[i:], m.ReviewCount)
		i = encodeVarintRegistry(dAtA, i, uint64(len(m.ReviewCount)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.PriceCount) > 0 {
		i -= len(m.PriceCount)
		copy(dAtA[i:], m.PriceCount)
		i = encodeVarintRegistry(dAtA, i, uint64(len(m.PriceCount)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ImageUrl) > 0 {
		i -= len(m.ImageUrl)
		copy(dAtA[i:], m.ImageUrl)
		i = encodeVarintRegistry(dAtA, i, uint64(len(m.ImageUrl)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintRegistry(dAtA, i, uint64(len(m.Description)))
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
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	l = len(m.ImageUrl)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	l = len(m.PriceCount)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	l = len(m.ReviewCount)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	l = len(m.Timestamp)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	l = len(m.Reserved)
	if l > 0 {
		n += 1 + l + sovRegistry(uint64(l))
	}
	if len(m.Owners) > 0 {
		for _, s := range m.Owners {
			l = len(s)
			n += 1 + l + sovRegistry(uint64(l))
		}
	}
	if len(m.Validators) > 0 {
		for _, s := range m.Validators {
			l = len(s)
			n += 1 + l + sovRegistry(uint64(l))
		}
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
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
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
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImageUrl", wireType)
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
			m.ImageUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceCount", wireType)
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
			m.PriceCount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReviewCount", wireType)
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
			m.ReviewCount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
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
			m.Timestamp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
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
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owners", wireType)
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
			m.Owners = append(m.Owners, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
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
			m.Validators = append(m.Validators, string(dAtA[iNdEx:postIndex]))
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
