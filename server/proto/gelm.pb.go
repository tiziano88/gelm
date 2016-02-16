// Code generated by protoc-gen-go.
// source: proto/gelm.proto
// DO NOT EDIT!

/*
Package gelm is a generated protocol buffer package.

It is generated from these files:
	proto/gelm.proto

It has these top-level messages:
	SubMessage
	Message
*/
package gelm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Enum int32

const (
	Enum_ENUM_VALUE_DEFAULT Enum = 0
	Enum_ENUM_VALUE_1       Enum = 1
	Enum_ENUM_VALUE_2       Enum = 2
)

var Enum_name = map[int32]string{
	0: "ENUM_VALUE_DEFAULT",
	1: "ENUM_VALUE_1",
	2: "ENUM_VALUE_2",
}
var Enum_value = map[string]int32{
	"ENUM_VALUE_DEFAULT": 0,
	"ENUM_VALUE_1":       1,
	"ENUM_VALUE_2":       2,
}

func (x Enum) String() string {
	return proto.EnumName(Enum_name, int32(x))
}
func (Enum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SubMessage struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *SubMessage) Reset()                    { *m = SubMessage{} }
func (m *SubMessage) String() string            { return proto.CompactTextString(m) }
func (*SubMessage) ProtoMessage()               {}
func (*SubMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Message struct {
	Id                int64       `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	FieldWithLongName string      `protobuf:"bytes,2,opt,name=field_with_long_name,json=fieldWithLongName" json:"field_with_long_name,omitempty"`
	Enum              Enum        `protobuf:"varint,3,opt,name=enum,enum=Enum" json:"enum,omitempty"`
	SubMessage        *SubMessage `protobuf:"bytes,4,opt,name=sub_message,json=subMessage" json:"sub_message,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Message) GetSubMessage() *SubMessage {
	if m != nil {
		return m.SubMessage
	}
	return nil
}

func init() {
	proto.RegisterType((*SubMessage)(nil), "SubMessage")
	proto.RegisterType((*Message)(nil), "Message")
	proto.RegisterEnum("Enum", Enum_name, Enum_value)
}

var fileDescriptor0 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x4f, 0xcd, 0xc9, 0xd5, 0x03, 0x33, 0x95, 0x64, 0xb8, 0xb8, 0x82, 0x4b, 0x93,
	0x7c, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x98, 0x83, 0x80, 0x2c, 0xa5, 0xc9, 0x8c, 0x5c, 0xec, 0x38, 0xe4, 0x84, 0xf4, 0xb9,
	0x44, 0xd2, 0x32, 0x53, 0x73, 0x52, 0xe2, 0xcb, 0x33, 0x4b, 0x32, 0xe2, 0x73, 0xf2, 0xf3, 0xd2,
	0xe3, 0xf3, 0x12, 0x73, 0x53, 0x25, 0x98, 0x80, 0x2a, 0x38, 0x83, 0x04, 0xc1, 0x72, 0xe1, 0x40,
	0x29, 0x1f, 0xa0, 0x8c, 0x1f, 0x50, 0x42, 0x48, 0x92, 0x8b, 0x25, 0x35, 0xaf, 0x34, 0x57, 0x82,
	0x19, 0xa8, 0x80, 0xcf, 0x88, 0x55, 0xcf, 0x15, 0xc8, 0x09, 0x02, 0x0b, 0x09, 0xe9, 0x70, 0x71,
	0x17, 0x97, 0x26, 0xc5, 0xe7, 0x42, 0xac, 0x92, 0x60, 0x01, 0xaa, 0xe0, 0x36, 0xe2, 0xd6, 0x43,
	0xb8, 0x2c, 0x88, 0xab, 0x18, 0xce, 0xd6, 0x72, 0xe2, 0x62, 0x01, 0xe9, 0x15, 0x12, 0xe3, 0x12,
	0x72, 0xf5, 0x0b, 0xf5, 0x8d, 0x0f, 0x73, 0xf4, 0x09, 0x75, 0x8d, 0x77, 0x71, 0x75, 0x73, 0x0c,
	0xf5, 0x09, 0x11, 0x60, 0x10, 0x12, 0xe0, 0xe2, 0x41, 0x12, 0x37, 0x14, 0x60, 0x44, 0x13, 0x31,
	0x12, 0x60, 0x4a, 0x62, 0x03, 0x7b, 0xdf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xd3, 0x61,
	0x7e, 0x12, 0x01, 0x00, 0x00,
}
