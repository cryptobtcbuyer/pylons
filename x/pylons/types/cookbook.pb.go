// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pylons/cookbook.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

type Cookbook struct {
	Creator      string      `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ID           string      `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	NodeVersion  string      `protobuf:"bytes,3,opt,name=nodeVersion,proto3" json:"nodeVersion,omitempty"`
	Name         string      `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description  string      `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Developer    string      `protobuf:"bytes,6,opt,name=developer,proto3" json:"developer,omitempty"`
	Version      string      `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
	SupportEmail string      `protobuf:"bytes,8,opt,name=supportEmail,proto3" json:"supportEmail,omitempty"`
	CostPerBlock *types.Coin `protobuf:"bytes,9,opt,name=costPerBlock,proto3" json:"costPerBlock,omitempty"`
	Enabled      bool        `protobuf:"varint,10,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (m *Cookbook) Reset()         { *m = Cookbook{} }
func (m *Cookbook) String() string { return proto.CompactTextString(m) }
func (*Cookbook) ProtoMessage()    {}
func (*Cookbook) Descriptor() ([]byte, []int) {
	return fileDescriptor_54566b7965d848c1, []int{0}
}
func (m *Cookbook) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Cookbook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Cookbook.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Cookbook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cookbook.Merge(m, src)
}
func (m *Cookbook) XXX_Size() int {
	return m.Size()
}
func (m *Cookbook) XXX_DiscardUnknown() {
	xxx_messageInfo_Cookbook.DiscardUnknown(m)
}

var xxx_messageInfo_Cookbook proto.InternalMessageInfo

func (m *Cookbook) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Cookbook) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Cookbook) GetNodeVersion() string {
	if m != nil {
		return m.NodeVersion
	}
	return ""
}

func (m *Cookbook) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Cookbook) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Cookbook) GetDeveloper() string {
	if m != nil {
		return m.Developer
	}
	return ""
}

func (m *Cookbook) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Cookbook) GetSupportEmail() string {
	if m != nil {
		return m.SupportEmail
	}
	return ""
}

func (m *Cookbook) GetCostPerBlock() *types.Coin {
	if m != nil {
		return m.CostPerBlock
	}
	return nil
}

func (m *Cookbook) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func init() {
	proto.RegisterType((*Cookbook)(nil), "Pylonstech.pylons.pylons.Cookbook")
}

func init() { proto.RegisterFile("pylons/cookbook.proto", fileDescriptor_54566b7965d848c1) }

var fileDescriptor_54566b7965d848c1 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x51, 0x3f, 0x6f, 0xfa, 0x30,
	0x10, 0x25, 0xf9, 0xf1, 0xe3, 0x8f, 0x41, 0x1d, 0x2c, 0x55, 0x72, 0xab, 0x2a, 0x8a, 0x98, 0x18,
	0x5a, 0x47, 0xb4, 0x73, 0x17, 0xa0, 0x95, 0xd8, 0x10, 0x43, 0x87, 0x6e, 0x89, 0x73, 0x2a, 0x11,
	0x89, 0xcf, 0xb2, 0x0d, 0x2a, 0xdf, 0xa2, 0x5f, 0xa9, 0x5b, 0x47, 0xc6, 0x8e, 0x15, 0x7c, 0x91,
	0x2a, 0x4e, 0x50, 0x61, 0xf2, 0xdd, 0x7b, 0xf7, 0x7c, 0xa7, 0xf7, 0xc8, 0xa5, 0xda, 0xe6, 0x28,
	0x4d, 0x24, 0x10, 0x57, 0x09, 0xe2, 0x8a, 0x2b, 0x8d, 0x16, 0x29, 0x9b, 0x3b, 0xd8, 0x82, 0x58,
	0xf2, 0x6a, 0xa2, 0x7e, 0xae, 0x03, 0x81, 0xa6, 0x40, 0x13, 0x25, 0xb1, 0x81, 0x68, 0x33, 0x4a,
	0xc0, 0xc6, 0xa3, 0x48, 0x60, 0x26, 0x2b, 0xe5, 0xe0, 0xd3, 0x27, 0x9d, 0x49, 0xfd, 0x19, 0x65,
	0xa4, 0x2d, 0x34, 0xc4, 0x16, 0x35, 0xf3, 0x42, 0x6f, 0xd8, 0x5d, 0x1c, 0x5b, 0x7a, 0x41, 0xfc,
	0xd9, 0x94, 0xf9, 0x0e, 0xf4, 0x67, 0x53, 0x1a, 0x92, 0x9e, 0xc4, 0x14, 0x5e, 0x40, 0x9b, 0x0c,
	0x25, 0xfb, 0xe7, 0x88, 0x53, 0x88, 0x52, 0xd2, 0x94, 0x71, 0x01, 0xac, 0xe9, 0x28, 0x57, 0x97,
	0xaa, 0x14, 0x8c, 0xd0, 0x99, 0xb2, 0xa5, 0xea, 0x7f, 0xa5, 0x3a, 0x81, 0xe8, 0x0d, 0xe9, 0xa6,
	0xb0, 0x81, 0x1c, 0x15, 0x68, 0xd6, 0x72, 0xfc, 0x1f, 0x50, 0xde, 0xb7, 0xa9, 0x37, 0xb6, 0xab,
	0xfb, 0xea, 0x96, 0x0e, 0x48, 0xdf, 0xac, 0x95, 0x42, 0x6d, 0x9f, 0x8a, 0x38, 0xcb, 0x59, 0xc7,
	0xd1, 0x67, 0x18, 0x7d, 0x24, 0x7d, 0x81, 0xc6, 0xce, 0x41, 0x8f, 0x73, 0x14, 0x2b, 0xd6, 0x0d,
	0xbd, 0x61, 0xef, 0xfe, 0x8a, 0x57, 0x0e, 0xf1, 0xd2, 0x21, 0x5e, 0x3b, 0xc4, 0x27, 0x98, 0xc9,
	0xc5, 0xd9, 0x78, 0xb9, 0x1c, 0x64, 0x9c, 0xe4, 0x90, 0x32, 0x12, 0x7a, 0xc3, 0xce, 0xe2, 0xd8,
	0x8e, 0x9f, 0xbf, 0xf6, 0x81, 0xb7, 0xdb, 0x07, 0xde, 0xcf, 0x3e, 0xf0, 0x3e, 0x0e, 0x41, 0x63,
	0x77, 0x08, 0x1a, 0xdf, 0x87, 0xa0, 0xf1, 0x7a, 0xfb, 0x96, 0xd9, 0xe5, 0x3a, 0xe1, 0x02, 0x8b,
	0xa8, 0x8a, 0xe8, 0xae, 0xcc, 0x28, 0xaa, 0x53, 0x7c, 0x3f, 0x16, 0x76, 0xab, 0xc0, 0x24, 0x2d,
	0x17, 0xc9, 0xc3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x57, 0x72, 0xa8, 0xe5, 0x01, 0x00,
	0x00,
}

func (m *Cookbook) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Cookbook) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Cookbook) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x50
	}
	if m.CostPerBlock != nil {
		{
			size, err := m.CostPerBlock.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCookbook(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if len(m.SupportEmail) > 0 {
		i -= len(m.SupportEmail)
		copy(dAtA[i:], m.SupportEmail)
		i = encodeVarintCookbook(dAtA, i, uint64(len(m.SupportEmail)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintCookbook(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Developer) > 0 {
		i -= len(m.Developer)
		copy(dAtA[i:], m.Developer)
		i = encodeVarintCookbook(dAtA, i, uint64(len(m.Developer)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintCookbook(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintCookbook(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.NodeVersion) > 0 {
		i -= len(m.NodeVersion)
		copy(dAtA[i:], m.NodeVersion)
		i = encodeVarintCookbook(dAtA, i, uint64(len(m.NodeVersion)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintCookbook(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintCookbook(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCookbook(dAtA []byte, offset int, v uint64) int {
	offset -= sovCookbook(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Cookbook) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovCookbook(uint64(l))
	}
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovCookbook(uint64(l))
	}
	l = len(m.NodeVersion)
	if l > 0 {
		n += 1 + l + sovCookbook(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCookbook(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovCookbook(uint64(l))
	}
	l = len(m.Developer)
	if l > 0 {
		n += 1 + l + sovCookbook(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovCookbook(uint64(l))
	}
	l = len(m.SupportEmail)
	if l > 0 {
		n += 1 + l + sovCookbook(uint64(l))
	}
	if m.CostPerBlock != nil {
		l = m.CostPerBlock.Size()
		n += 1 + l + sovCookbook(uint64(l))
	}
	if m.Enabled {
		n += 2
	}
	return n
}

func sovCookbook(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCookbook(x uint64) (n int) {
	return sovCookbook(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Cookbook) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCookbook
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
			return fmt.Errorf("proto: Cookbook: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Cookbook: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCookbook
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
				return ErrInvalidLengthCookbook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCookbook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCookbook
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
				return ErrInvalidLengthCookbook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCookbook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCookbook
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
				return ErrInvalidLengthCookbook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCookbook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCookbook
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
				return ErrInvalidLengthCookbook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCookbook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCookbook
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
				return ErrInvalidLengthCookbook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCookbook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Developer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCookbook
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
				return ErrInvalidLengthCookbook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCookbook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Developer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCookbook
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
				return ErrInvalidLengthCookbook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCookbook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SupportEmail", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCookbook
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
				return ErrInvalidLengthCookbook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCookbook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SupportEmail = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CostPerBlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCookbook
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
				return ErrInvalidLengthCookbook
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCookbook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CostPerBlock == nil {
				m.CostPerBlock = &types.Coin{}
			}
			if err := m.CostPerBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCookbook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCookbook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCookbook
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
func skipCookbook(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCookbook
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
					return 0, ErrIntOverflowCookbook
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
					return 0, ErrIntOverflowCookbook
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
				return 0, ErrInvalidLengthCookbook
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCookbook
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCookbook
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCookbook        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCookbook          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCookbook = fmt.Errorf("proto: unexpected end of group")
)