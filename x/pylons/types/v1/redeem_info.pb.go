// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pylons/pylons/v1/redeem_info.proto

package v1

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type RedeemInfo struct {
	ID            string                                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ProcessorName string                                 `protobuf:"bytes,2,opt,name=processorName,proto3" json:"processorName,omitempty"`
	Address       string                                 `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Amount        github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
	Signature     string                                 `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *RedeemInfo) Reset()         { *m = RedeemInfo{} }
func (m *RedeemInfo) String() string { return proto.CompactTextString(m) }
func (*RedeemInfo) ProtoMessage()    {}
func (*RedeemInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fd1bc6684bea649, []int{0}
}
func (m *RedeemInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RedeemInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RedeemInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RedeemInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedeemInfo.Merge(m, src)
}
func (m *RedeemInfo) XXX_Size() int {
	return m.Size()
}
func (m *RedeemInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RedeemInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RedeemInfo proto.InternalMessageInfo

func (m *RedeemInfo) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *RedeemInfo) GetProcessorName() string {
	if m != nil {
		return m.ProcessorName
	}
	return ""
}

func (m *RedeemInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RedeemInfo) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type CreatePaymentAccount struct {
	Address   string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Token     string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Signature string `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *CreatePaymentAccount) Reset()         { *m = CreatePaymentAccount{} }
func (m *CreatePaymentAccount) String() string { return proto.CompactTextString(m) }
func (*CreatePaymentAccount) ProtoMessage()    {}
func (*CreatePaymentAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fd1bc6684bea649, []int{1}
}
func (m *CreatePaymentAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreatePaymentAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreatePaymentAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreatePaymentAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePaymentAccount.Merge(m, src)
}
func (m *CreatePaymentAccount) XXX_Size() int {
	return m.Size()
}
func (m *CreatePaymentAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePaymentAccount.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePaymentAccount proto.InternalMessageInfo

func (m *CreatePaymentAccount) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *CreatePaymentAccount) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CreatePaymentAccount) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func init() {
	proto.RegisterType((*RedeemInfo)(nil), "pylons.pylons.v1.RedeemInfo")
	proto.RegisterType((*CreatePaymentAccount)(nil), "pylons.pylons.v1.CreatePaymentAccount")
}

func init() {
	proto.RegisterFile("pylons/pylons/v1/redeem_info.proto", fileDescriptor_7fd1bc6684bea649)
}

var fileDescriptor_7fd1bc6684bea649 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x51, 0xc1, 0x6a, 0x02, 0x31,
	0x10, 0xdd, 0x68, 0xb5, 0x18, 0x68, 0x29, 0x8b, 0x87, 0xa5, 0x94, 0x58, 0xa4, 0x94, 0x5e, 0xdc,
	0x20, 0xfd, 0x82, 0x5a, 0x29, 0xec, 0xa5, 0x88, 0xc7, 0x5e, 0x4a, 0xcc, 0x8e, 0xab, 0xd8, 0x64,
	0x96, 0x24, 0x4a, 0xfd, 0x8b, 0x7e, 0x53, 0x4f, 0x1e, 0x3d, 0x96, 0x1e, 0xa4, 0xe8, 0x8f, 0x14,
	0xe3, 0x4a, 0xb5, 0xa7, 0x37, 0x33, 0x2f, 0x8f, 0x37, 0x6f, 0x42, 0x9b, 0xf9, 0xfc, 0x0d, 0xb5,
	0xe5, 0x05, 0xcc, 0xda, 0xdc, 0x40, 0x0a, 0xa0, 0x5e, 0xc7, 0x7a, 0x88, 0x71, 0x6e, 0xd0, 0x61,
	0x78, 0xb1, 0x23, 0xe3, 0x02, 0x66, 0xed, 0xcb, 0x7a, 0x86, 0x19, 0x7a, 0x92, 0x6f, 0xab, 0xdd,
	0xbb, 0xe6, 0x27, 0xa1, 0xb4, 0xef, 0xd5, 0x89, 0x1e, 0x62, 0x78, 0x4e, 0x4b, 0x49, 0x37, 0x22,
	0xd7, 0xe4, 0xae, 0xd6, 0x2f, 0x25, 0xdd, 0xf0, 0x86, 0x9e, 0xe5, 0x06, 0x25, 0x58, 0x8b, 0xe6,
	0x59, 0x28, 0x88, 0x4a, 0x9e, 0x3a, 0x1e, 0x86, 0x11, 0x3d, 0x15, 0x69, 0x6a, 0xc0, 0xda, 0xa8,
	0xec, 0xf9, 0x7d, 0x1b, 0x3e, 0xd1, 0xaa, 0x50, 0x38, 0xd5, 0x2e, 0x3a, 0xd9, 0x12, 0x9d, 0x78,
	0xb1, 0x6a, 0x04, 0xdf, 0xab, 0xc6, 0x6d, 0x36, 0x76, 0xa3, 0xe9, 0x20, 0x96, 0xa8, 0xb8, 0x44,
	0xab, 0xd0, 0x16, 0xd0, 0xb2, 0xe9, 0x84, 0xbb, 0x79, 0x0e, 0x36, 0x4e, 0xb4, 0xeb, 0x17, 0xea,
	0xf0, 0x8a, 0xd6, 0xec, 0x38, 0xd3, 0xc2, 0x4d, 0x0d, 0x44, 0x15, 0xef, 0xf1, 0x37, 0x68, 0xa6,
	0xb4, 0xfe, 0x68, 0x40, 0x38, 0xe8, 0x89, 0xb9, 0x02, 0xed, 0x1e, 0xa4, 0xf4, 0xaa, 0x83, 0xbd,
	0xc8, 0xf1, 0x5e, 0x75, 0x5a, 0x71, 0x38, 0x01, 0x5d, 0xe4, 0xd9, 0x35, 0xc7, 0x2e, 0xe5, 0x7f,
	0x2e, 0x9d, 0x64, 0xb1, 0x66, 0x64, 0xb9, 0x66, 0xe4, 0x67, 0xcd, 0xc8, 0xc7, 0x86, 0x05, 0xcb,
	0x0d, 0x0b, 0xbe, 0x36, 0x2c, 0x78, 0xe1, 0x07, 0x69, 0x7a, 0xfe, 0xe0, 0x2d, 0x07, 0x72, 0xb4,
	0xff, 0xa0, 0xf7, 0x7d, 0xe1, 0x73, 0xf1, 0x59, 0x7b, 0x50, 0xf5, 0xc7, 0xbf, 0xff, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0x36, 0x0b, 0xdf, 0xf9, 0xca, 0x01, 0x00, 0x00,
}

func (m *RedeemInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RedeemInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RedeemInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintRedeemInfo(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x2a
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintRedeemInfo(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintRedeemInfo(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ProcessorName) > 0 {
		i -= len(m.ProcessorName)
		copy(dAtA[i:], m.ProcessorName)
		i = encodeVarintRedeemInfo(dAtA, i, uint64(len(m.ProcessorName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintRedeemInfo(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CreatePaymentAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreatePaymentAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreatePaymentAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintRedeemInfo(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintRedeemInfo(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintRedeemInfo(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRedeemInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovRedeemInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RedeemInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovRedeemInfo(uint64(l))
	}
	l = len(m.ProcessorName)
	if l > 0 {
		n += 1 + l + sovRedeemInfo(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovRedeemInfo(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovRedeemInfo(uint64(l))
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovRedeemInfo(uint64(l))
	}
	return n
}

func (m *CreatePaymentAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovRedeemInfo(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovRedeemInfo(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovRedeemInfo(uint64(l))
	}
	return n
}

func sovRedeemInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRedeemInfo(x uint64) (n int) {
	return sovRedeemInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RedeemInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRedeemInfo
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
			return fmt.Errorf("proto: RedeemInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RedeemInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRedeemInfo
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
				return ErrInvalidLengthRedeemInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRedeemInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProcessorName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRedeemInfo
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
				return ErrInvalidLengthRedeemInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRedeemInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProcessorName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRedeemInfo
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
				return ErrInvalidLengthRedeemInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRedeemInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRedeemInfo
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
				return ErrInvalidLengthRedeemInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRedeemInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRedeemInfo
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
				return ErrInvalidLengthRedeemInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRedeemInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRedeemInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRedeemInfo
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
func (m *CreatePaymentAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRedeemInfo
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
			return fmt.Errorf("proto: CreatePaymentAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreatePaymentAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRedeemInfo
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
				return ErrInvalidLengthRedeemInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRedeemInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRedeemInfo
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
				return ErrInvalidLengthRedeemInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRedeemInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRedeemInfo
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
				return ErrInvalidLengthRedeemInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRedeemInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRedeemInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRedeemInfo
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
func skipRedeemInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRedeemInfo
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
					return 0, ErrIntOverflowRedeemInfo
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
					return 0, ErrIntOverflowRedeemInfo
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
				return 0, ErrInvalidLengthRedeemInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRedeemInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRedeemInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRedeemInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRedeemInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRedeemInfo = fmt.Errorf("proto: unexpected end of group")
)