// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: go-internal/log_entry_setting.proto

package group

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/gogo/protobuf/proto"
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

type SettingsEntryPayload_PayloadType int32

const (
	SettingsEntryPayload_PayloadTypeUnknown       SettingsEntryPayload_PayloadType = 0
	SettingsEntryPayload_PayloadTypeGroupSetting  SettingsEntryPayload_PayloadType = 1
	SettingsEntryPayload_PayloadTypeMemberSetting SettingsEntryPayload_PayloadType = 2
)

var SettingsEntryPayload_PayloadType_name = map[int32]string{
	0: "PayloadTypeUnknown",
	1: "PayloadTypeGroupSetting",
	2: "PayloadTypeMemberSetting",
}

var SettingsEntryPayload_PayloadType_value = map[string]int32{
	"PayloadTypeUnknown":       0,
	"PayloadTypeGroupSetting":  1,
	"PayloadTypeMemberSetting": 2,
}

func (x SettingsEntryPayload_PayloadType) String() string {
	return proto.EnumName(SettingsEntryPayload_PayloadType_name, int32(x))
}

func (SettingsEntryPayload_PayloadType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f48654281ae787cf, []int{1, 0}
}

type SettingsEntryEnvelope struct {
	EncryptedSettingsPayload []byte `protobuf:"bytes,1,opt,name=encrypted_settings_payload,json=encryptedSettingsPayload,proto3" json:"encrypted_settings_payload,omitempty"`
	SettingsPayloadSignature []byte `protobuf:"bytes,2,opt,name=settings_payload_signature,json=settingsPayloadSignature,proto3" json:"settings_payload_signature,omitempty"`
}

func (m *SettingsEntryEnvelope) Reset()         { *m = SettingsEntryEnvelope{} }
func (m *SettingsEntryEnvelope) String() string { return proto.CompactTextString(m) }
func (*SettingsEntryEnvelope) ProtoMessage()    {}
func (*SettingsEntryEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_f48654281ae787cf, []int{0}
}
func (m *SettingsEntryEnvelope) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SettingsEntryEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SettingsEntryEnvelope.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SettingsEntryEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettingsEntryEnvelope.Merge(m, src)
}
func (m *SettingsEntryEnvelope) XXX_Size() int {
	return m.Size()
}
func (m *SettingsEntryEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_SettingsEntryEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_SettingsEntryEnvelope proto.InternalMessageInfo

func (m *SettingsEntryEnvelope) GetEncryptedSettingsPayload() []byte {
	if m != nil {
		return m.EncryptedSettingsPayload
	}
	return nil
}

func (m *SettingsEntryEnvelope) GetSettingsPayloadSignature() []byte {
	if m != nil {
		return m.SettingsPayloadSignature
	}
	return nil
}

type SettingsEntryPayload struct {
	Type         SettingsEntryPayload_PayloadType `protobuf:"varint,1,opt,name=type,proto3,enum=SettingsEntryPayload_PayloadType" json:"type,omitempty"`
	MemberPubKey []byte                           `protobuf:"bytes,2,opt,name=member_pub_key,json=memberPubKey,proto3" json:"member_pub_key,omitempty"`
	Key          []byte                           `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value        []byte                           `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *SettingsEntryPayload) Reset()         { *m = SettingsEntryPayload{} }
func (m *SettingsEntryPayload) String() string { return proto.CompactTextString(m) }
func (*SettingsEntryPayload) ProtoMessage()    {}
func (*SettingsEntryPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_f48654281ae787cf, []int{1}
}
func (m *SettingsEntryPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SettingsEntryPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SettingsEntryPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SettingsEntryPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettingsEntryPayload.Merge(m, src)
}
func (m *SettingsEntryPayload) XXX_Size() int {
	return m.Size()
}
func (m *SettingsEntryPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_SettingsEntryPayload.DiscardUnknown(m)
}

var xxx_messageInfo_SettingsEntryPayload proto.InternalMessageInfo

func (m *SettingsEntryPayload) GetType() SettingsEntryPayload_PayloadType {
	if m != nil {
		return m.Type
	}
	return SettingsEntryPayload_PayloadTypeUnknown
}

func (m *SettingsEntryPayload) GetMemberPubKey() []byte {
	if m != nil {
		return m.MemberPubKey
	}
	return nil
}

func (m *SettingsEntryPayload) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *SettingsEntryPayload) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterEnum("SettingsEntryPayload_PayloadType", SettingsEntryPayload_PayloadType_name, SettingsEntryPayload_PayloadType_value)
	proto.RegisterType((*SettingsEntryEnvelope)(nil), "SettingsEntryEnvelope")
	proto.RegisterType((*SettingsEntryPayload)(nil), "SettingsEntryPayload")
}

func init() {
	proto.RegisterFile("go-internal/log_entry_setting.proto", fileDescriptor_f48654281ae787cf)
}

var fileDescriptor_f48654281ae787cf = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xb3, 0x6d, 0xf5, 0x30, 0x96, 0x52, 0x96, 0xaa, 0x41, 0x4b, 0xd0, 0xea, 0xc1, 0x8b,
	0x29, 0x28, 0x7a, 0xf2, 0x24, 0x14, 0x0f, 0x22, 0x94, 0x56, 0x2f, 0x5e, 0x62, 0xd2, 0x0e, 0xb1,
	0x34, 0xdd, 0x5d, 0x36, 0x9b, 0xca, 0xbe, 0x85, 0xe0, 0x4b, 0x79, 0xec, 0xd1, 0xa3, 0xb4, 0xcf,
	0x21, 0x48, 0x36, 0x49, 0x49, 0xc5, 0xd3, 0xee, 0xfc, 0xdf, 0xfc, 0x33, 0x3f, 0x0c, 0x9c, 0x84,
	0xfc, 0x7c, 0xc2, 0x14, 0x4a, 0xe6, 0x47, 0xdd, 0x88, 0x87, 0x1e, 0x32, 0x25, 0xb5, 0x17, 0xa3,
	0x52, 0x13, 0x16, 0xba, 0x42, 0x72, 0xc5, 0x3b, 0x1f, 0x04, 0x76, 0x87, 0x99, 0x12, 0xf7, 0x52,
	0xde, 0x63, 0x73, 0x8c, 0xb8, 0x40, 0x7a, 0x03, 0x07, 0xc8, 0x46, 0x52, 0x0b, 0x85, 0xe3, 0xc2,
	0x14, 0x7b, 0xc2, 0xd7, 0x11, 0xf7, 0xc7, 0x36, 0x39, 0x22, 0x67, 0xf5, 0x81, 0xbd, 0xee, 0x28,
	0x66, 0xf4, 0x33, 0x9e, 0xba, 0xff, 0x7a, 0xbc, 0x78, 0x12, 0x32, 0x5f, 0x25, 0x12, 0xed, 0x4a,
	0xe6, 0x8e, 0x37, 0x4d, 0xc3, 0x82, 0x77, 0x7e, 0x08, 0xb4, 0x36, 0x52, 0x15, 0x63, 0xaf, 0xa0,
	0xa6, 0xb4, 0x40, 0xb3, 0xbe, 0x71, 0x71, 0xec, 0xfe, 0xd7, 0xe4, 0xe6, 0xef, 0xa3, 0x16, 0x38,
	0x30, 0xed, 0xf4, 0x14, 0x1a, 0x33, 0x9c, 0x05, 0x28, 0x3d, 0x91, 0x04, 0xde, 0x14, 0x75, 0x9e,
	0xa0, 0x9e, 0xa9, 0xfd, 0x24, 0xb8, 0x47, 0x4d, 0x9b, 0x50, 0x4d, 0x51, 0xd5, 0xa0, 0xf4, 0x4b,
	0x5b, 0xb0, 0x35, 0xf7, 0xa3, 0x04, 0xed, 0x9a, 0xd1, 0xb2, 0xa2, 0xf3, 0x02, 0x3b, 0xa5, 0x15,
	0x74, 0x0f, 0x68, 0xa9, 0x7c, 0x62, 0x53, 0xc6, 0xdf, 0x58, 0xd3, 0xa2, 0x87, 0xb0, 0x5f, 0xd2,
	0xef, 0x24, 0x4f, 0x44, 0x1e, 0xb7, 0x49, 0x68, 0x1b, 0xec, 0x12, 0x7c, 0x30, 0x31, 0x0a, 0x5a,
	0xb9, 0xbd, 0xfe, 0x5c, 0x3a, 0x64, 0xb1, 0x74, 0xc8, 0xf7, 0xd2, 0x21, 0xef, 0x2b, 0xc7, 0x5a,
	0xac, 0x1c, 0xeb, 0x6b, 0xe5, 0x58, 0xcf, 0xed, 0x00, 0xa5, 0xd2, 0xae, 0xc2, 0xd1, 0x6b, 0x37,
	0xe4, 0xdd, 0xf5, 0x7d, 0xc3, 0x74, 0x7c, 0xb0, 0x6d, 0x8e, 0x7a, 0xf9, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0xb1, 0x0a, 0x45, 0x7a, 0xfb, 0x01, 0x00, 0x00,
}

func (m *SettingsEntryEnvelope) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SettingsEntryEnvelope) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SettingsEntryEnvelope) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SettingsPayloadSignature) > 0 {
		i -= len(m.SettingsPayloadSignature)
		copy(dAtA[i:], m.SettingsPayloadSignature)
		i = encodeVarintLogEntrySetting(dAtA, i, uint64(len(m.SettingsPayloadSignature)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.EncryptedSettingsPayload) > 0 {
		i -= len(m.EncryptedSettingsPayload)
		copy(dAtA[i:], m.EncryptedSettingsPayload)
		i = encodeVarintLogEntrySetting(dAtA, i, uint64(len(m.EncryptedSettingsPayload)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SettingsEntryPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SettingsEntryPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SettingsEntryPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintLogEntrySetting(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintLogEntrySetting(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.MemberPubKey) > 0 {
		i -= len(m.MemberPubKey)
		copy(dAtA[i:], m.MemberPubKey)
		i = encodeVarintLogEntrySetting(dAtA, i, uint64(len(m.MemberPubKey)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintLogEntrySetting(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLogEntrySetting(dAtA []byte, offset int, v uint64) int {
	offset -= sovLogEntrySetting(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SettingsEntryEnvelope) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EncryptedSettingsPayload)
	if l > 0 {
		n += 1 + l + sovLogEntrySetting(uint64(l))
	}
	l = len(m.SettingsPayloadSignature)
	if l > 0 {
		n += 1 + l + sovLogEntrySetting(uint64(l))
	}
	return n
}

func (m *SettingsEntryPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovLogEntrySetting(uint64(m.Type))
	}
	l = len(m.MemberPubKey)
	if l > 0 {
		n += 1 + l + sovLogEntrySetting(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovLogEntrySetting(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovLogEntrySetting(uint64(l))
	}
	return n
}

func sovLogEntrySetting(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLogEntrySetting(x uint64) (n int) {
	return sovLogEntrySetting(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SettingsEntryEnvelope) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogEntrySetting
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
			return fmt.Errorf("proto: SettingsEntryEnvelope: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SettingsEntryEnvelope: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptedSettingsPayload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntrySetting
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
				return ErrInvalidLengthLogEntrySetting
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLogEntrySetting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncryptedSettingsPayload = append(m.EncryptedSettingsPayload[:0], dAtA[iNdEx:postIndex]...)
			if m.EncryptedSettingsPayload == nil {
				m.EncryptedSettingsPayload = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SettingsPayloadSignature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntrySetting
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
				return ErrInvalidLengthLogEntrySetting
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLogEntrySetting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SettingsPayloadSignature = append(m.SettingsPayloadSignature[:0], dAtA[iNdEx:postIndex]...)
			if m.SettingsPayloadSignature == nil {
				m.SettingsPayloadSignature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogEntrySetting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogEntrySetting
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogEntrySetting
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
func (m *SettingsEntryPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogEntrySetting
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
			return fmt.Errorf("proto: SettingsEntryPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SettingsEntryPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntrySetting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= SettingsEntryPayload_PayloadType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemberPubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntrySetting
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
				return ErrInvalidLengthLogEntrySetting
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLogEntrySetting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MemberPubKey = append(m.MemberPubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.MemberPubKey == nil {
				m.MemberPubKey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntrySetting
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
				return ErrInvalidLengthLogEntrySetting
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLogEntrySetting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogEntrySetting
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
				return ErrInvalidLengthLogEntrySetting
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLogEntrySetting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogEntrySetting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLogEntrySetting
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLogEntrySetting
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
func skipLogEntrySetting(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLogEntrySetting
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
					return 0, ErrIntOverflowLogEntrySetting
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
					return 0, ErrIntOverflowLogEntrySetting
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
				return 0, ErrInvalidLengthLogEntrySetting
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLogEntrySetting
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLogEntrySetting
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLogEntrySetting        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLogEntrySetting          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLogEntrySetting = fmt.Errorf("proto: unexpected end of group")
)
