// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: registry/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// GenesisState defines the registry module's genesis state.
type GenesisState struct {
	Params                            Params                          `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	RegistryList                      []Registry                      `protobuf:"bytes,2,rep,name=registryList,proto3" json:"registryList"`
	RegistryCount                     uint64                          `protobuf:"varint,3,opt,name=registryCount,proto3" json:"registryCount,omitempty"`
	RegistryOwnerList                 []RegistryOwner                 `protobuf:"bytes,4,rep,name=registryOwnerList,proto3" json:"registryOwnerList"`
	RegistryOwnerCount                uint64                          `protobuf:"varint,5,opt,name=registryOwnerCount,proto3" json:"registryOwnerCount,omitempty"`
	RegistryMemberList                []RegistryMember                `protobuf:"bytes,6,rep,name=registryMemberList,proto3" json:"registryMemberList"`
	RegistryMemberCount               uint64                          `protobuf:"varint,7,opt,name=registryMemberCount,proto3" json:"registryMemberCount,omitempty"`
	RegistryStakedAmountList          []RegistryStakedAmount          `protobuf:"bytes,8,rep,name=registryStakedAmountList,proto3" json:"registryStakedAmountList"`
	RegistryStakedAmountPerWalletList []RegistryStakedAmountPerWallet `protobuf:"bytes,9,rep,name=registryStakedAmountPerWalletList,proto3" json:"registryStakedAmountPerWalletList"`
	PriceDataList                     []PriceData                     `protobuf:"bytes,10,rep,name=priceDataList,proto3" json:"priceDataList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed2e112cbf381ddc, []int{0}
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

func (m *GenesisState) GetRegistryList() []Registry {
	if m != nil {
		return m.RegistryList
	}
	return nil
}

func (m *GenesisState) GetRegistryCount() uint64 {
	if m != nil {
		return m.RegistryCount
	}
	return 0
}

func (m *GenesisState) GetRegistryOwnerList() []RegistryOwner {
	if m != nil {
		return m.RegistryOwnerList
	}
	return nil
}

func (m *GenesisState) GetRegistryOwnerCount() uint64 {
	if m != nil {
		return m.RegistryOwnerCount
	}
	return 0
}

func (m *GenesisState) GetRegistryMemberList() []RegistryMember {
	if m != nil {
		return m.RegistryMemberList
	}
	return nil
}

func (m *GenesisState) GetRegistryMemberCount() uint64 {
	if m != nil {
		return m.RegistryMemberCount
	}
	return 0
}

func (m *GenesisState) GetRegistryStakedAmountList() []RegistryStakedAmount {
	if m != nil {
		return m.RegistryStakedAmountList
	}
	return nil
}

func (m *GenesisState) GetRegistryStakedAmountPerWalletList() []RegistryStakedAmountPerWallet {
	if m != nil {
		return m.RegistryStakedAmountPerWalletList
	}
	return nil
}

func (m *GenesisState) GetPriceDataList() []PriceData {
	if m != nil {
		return m.PriceDataList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "pricechain.cprc.registry.GenesisState")
}

func init() { proto.RegisterFile("registry/genesis.proto", fileDescriptor_ed2e112cbf381ddc) }

var fileDescriptor_ed2e112cbf381ddc = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0xab, 0xd3, 0x40,
	0x14, 0x4d, 0x7c, 0x35, 0xea, 0xbc, 0xf7, 0x16, 0x8e, 0x5f, 0xb1, 0x60, 0x8c, 0x4f, 0xc5, 0x80,
	0x30, 0x79, 0x3c, 0x17, 0xee, 0x04, 0x5f, 0x15, 0x37, 0x95, 0x96, 0x76, 0x21, 0x28, 0x18, 0xa6,
	0xe9, 0x90, 0x06, 0x9b, 0x0f, 0x26, 0x53, 0x6a, 0x7f, 0x83, 0x1b, 0x7f, 0x8b, 0xbf, 0xa2, 0xcb,
	0x2e, 0x5d, 0x89, 0xb4, 0x7f, 0x44, 0x72, 0x33, 0x93, 0x1a, 0x92, 0x1a, 0xdc, 0x4d, 0xee, 0x39,
	0xf7, 0x9c, 0x33, 0x37, 0x77, 0xd0, 0x5d, 0xce, 0x82, 0x30, 0x13, 0x7c, 0xe5, 0x06, 0x2c, 0x66,
	0x59, 0x98, 0x91, 0x94, 0x27, 0x22, 0xc1, 0x66, 0xca, 0x43, 0x9f, 0xf9, 0x33, 0x1a, 0xc6, 0xc4,
	0x4f, 0xb9, 0x4f, 0x14, 0xaf, 0x7b, 0x3b, 0x48, 0x82, 0x04, 0x48, 0x6e, 0x7e, 0x2a, 0xf8, 0xdd,
	0x3b, 0xa5, 0x4e, 0x4a, 0x39, 0x8d, 0xa4, 0x4c, 0xf7, 0x5e, 0x59, 0x56, 0x07, 0x09, 0x3c, 0xa8,
	0x01, 0x5e, 0xb2, 0x8c, 0x19, 0x97, 0xb0, 0x55, 0x87, 0x23, 0x16, 0x4d, 0x4a, 0xfc, 0x69, 0x1d,
	0xcf, 0x04, 0xfd, 0xc2, 0xa6, 0x1e, 0x8d, 0x92, 0x45, 0x2c, 0x24, 0xed, 0xbc, 0x85, 0xe6, 0xa5,
	0x8c, 0x7b, 0x4b, 0x3a, 0x9f, 0x33, 0xd5, 0x71, 0x7f, 0x7f, 0x8f, 0x7c, 0x00, 0xde, 0x94, 0x0a,
	0x5a, 0x40, 0x67, 0x3f, 0x0c, 0x74, 0xf2, 0xae, 0x18, 0xd2, 0x58, 0x50, 0xc1, 0xf0, 0x2b, 0x64,
	0x14, 0x97, 0x35, 0x75, 0x5b, 0x77, 0x8e, 0x2f, 0x6c, 0x72, 0x68, 0x68, 0x64, 0x08, 0xbc, 0xcb,
	0xce, 0xfa, 0xd7, 0x43, 0x6d, 0x24, 0xbb, 0x70, 0x1f, 0x9d, 0x28, 0x42, 0x3f, 0xcc, 0x84, 0x79,
	0xc5, 0x3e, 0x72, 0x8e, 0x2f, 0xce, 0x0e, 0xab, 0x8c, 0xe4, 0x41, 0xea, 0x54, 0xba, 0xf1, 0x13,
	0x74, 0xaa, 0xbe, 0x7b, 0xf9, 0xdd, 0xcc, 0x23, 0x5b, 0x77, 0x3a, 0xa3, 0x6a, 0x11, 0x7f, 0x42,
	0x37, 0x55, 0x61, 0x90, 0xcf, 0x1b, 0x8c, 0x3b, 0x60, 0xfc, 0xac, 0xdd, 0x18, 0x5a, 0xa4, 0x7b,
	0x5d, 0x07, 0x13, 0x84, 0x2b, 0xc5, 0x22, 0xc7, 0x55, 0xc8, 0xd1, 0x80, 0xe0, 0xcf, 0x7b, 0xfe,
	0x7b, 0xf8, 0xbb, 0x90, 0xc6, 0x80, 0x34, 0x4e, 0x7b, 0x9a, 0xa2, 0x47, 0xc6, 0x69, 0x50, 0xc2,
	0xe7, 0xe8, 0x56, 0xb5, 0x5a, 0x04, 0xba, 0x06, 0x81, 0x9a, 0x20, 0x9c, 0x22, 0x53, 0x95, 0xc7,
	0xb0, 0x28, 0xaf, 0x61, 0x4f, 0x20, 0xd7, 0x75, 0xc8, 0x45, 0xda, 0x73, 0xfd, 0xdd, 0x29, 0xd3,
	0x1d, 0x54, 0xc5, 0xdf, 0x74, 0xf4, 0xa8, 0x09, 0x1c, 0x32, 0xfe, 0x01, 0x16, 0x13, 0xbc, 0x6f,
	0x80, 0xf7, 0xcb, 0xff, 0xf3, 0x2e, 0x25, 0x64, 0x88, 0x76, 0x1f, 0x3c, 0x40, 0xa7, 0x60, 0xf1,
	0x86, 0x0a, 0x0a, 0xc6, 0x08, 0x8c, 0x1f, 0xff, 0x63, 0xb3, 0x15, 0x5d, 0x9a, 0x54, 0xfb, 0x2f,
	0xdf, 0xae, 0xb7, 0x96, 0xbe, 0xd9, 0x5a, 0xfa, 0xef, 0xad, 0xa5, 0x7f, 0xdf, 0x59, 0xda, 0x66,
	0x67, 0x69, 0x3f, 0x77, 0x96, 0xf6, 0xf1, 0x79, 0x10, 0x8a, 0xd9, 0x62, 0x42, 0xfc, 0x24, 0x72,
	0x41, 0xa4, 0x97, 0xab, 0xbb, 0xb9, 0xba, 0xfb, 0xb5, 0x7c, 0xaf, 0xae, 0x58, 0xa5, 0x2c, 0x9b,
	0x18, 0xf0, 0x04, 0x5f, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x04, 0x7c, 0xa1, 0xaf, 0x04,
	0x00, 0x00,
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
	if len(m.PriceDataList) > 0 {
		for iNdEx := len(m.PriceDataList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PriceDataList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.RegistryStakedAmountPerWalletList) > 0 {
		for iNdEx := len(m.RegistryStakedAmountPerWalletList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RegistryStakedAmountPerWalletList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.RegistryStakedAmountList) > 0 {
		for iNdEx := len(m.RegistryStakedAmountList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RegistryStakedAmountList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if m.RegistryMemberCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.RegistryMemberCount))
		i--
		dAtA[i] = 0x38
	}
	if len(m.RegistryMemberList) > 0 {
		for iNdEx := len(m.RegistryMemberList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RegistryMemberList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.RegistryOwnerCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.RegistryOwnerCount))
		i--
		dAtA[i] = 0x28
	}
	if len(m.RegistryOwnerList) > 0 {
		for iNdEx := len(m.RegistryOwnerList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RegistryOwnerList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.RegistryCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.RegistryCount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.RegistryList) > 0 {
		for iNdEx := len(m.RegistryList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RegistryList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.RegistryList) > 0 {
		for _, e := range m.RegistryList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.RegistryCount != 0 {
		n += 1 + sovGenesis(uint64(m.RegistryCount))
	}
	if len(m.RegistryOwnerList) > 0 {
		for _, e := range m.RegistryOwnerList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.RegistryOwnerCount != 0 {
		n += 1 + sovGenesis(uint64(m.RegistryOwnerCount))
	}
	if len(m.RegistryMemberList) > 0 {
		for _, e := range m.RegistryMemberList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.RegistryMemberCount != 0 {
		n += 1 + sovGenesis(uint64(m.RegistryMemberCount))
	}
	if len(m.RegistryStakedAmountList) > 0 {
		for _, e := range m.RegistryStakedAmountList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RegistryStakedAmountPerWalletList) > 0 {
		for _, e := range m.RegistryStakedAmountPerWalletList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PriceDataList) > 0 {
		for _, e := range m.PriceDataList {
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
				return fmt.Errorf("proto: wrong wireType = %d for field RegistryList", wireType)
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
			m.RegistryList = append(m.RegistryList, Registry{})
			if err := m.RegistryList[len(m.RegistryList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegistryCount", wireType)
			}
			m.RegistryCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RegistryCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegistryOwnerList", wireType)
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
			m.RegistryOwnerList = append(m.RegistryOwnerList, RegistryOwner{})
			if err := m.RegistryOwnerList[len(m.RegistryOwnerList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegistryOwnerCount", wireType)
			}
			m.RegistryOwnerCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RegistryOwnerCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegistryMemberList", wireType)
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
			m.RegistryMemberList = append(m.RegistryMemberList, RegistryMember{})
			if err := m.RegistryMemberList[len(m.RegistryMemberList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegistryMemberCount", wireType)
			}
			m.RegistryMemberCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RegistryMemberCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegistryStakedAmountList", wireType)
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
			m.RegistryStakedAmountList = append(m.RegistryStakedAmountList, RegistryStakedAmount{})
			if err := m.RegistryStakedAmountList[len(m.RegistryStakedAmountList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegistryStakedAmountPerWalletList", wireType)
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
			m.RegistryStakedAmountPerWalletList = append(m.RegistryStakedAmountPerWalletList, RegistryStakedAmountPerWallet{})
			if err := m.RegistryStakedAmountPerWalletList[len(m.RegistryStakedAmountPerWalletList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceDataList", wireType)
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
			m.PriceDataList = append(m.PriceDataList, PriceData{})
			if err := m.PriceDataList[len(m.PriceDataList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
