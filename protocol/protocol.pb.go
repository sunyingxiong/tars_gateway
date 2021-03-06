// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol.proto

package protocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The request message containing the user's name.
type Request struct {
	Version              uint32   `protobuf:"fixed32,1,opt,name=version,proto3" json:"version,omitempty"`
	Servant              uint32   `protobuf:"fixed32,2,opt,name=servant,proto3" json:"servant,omitempty"`
	Seq                  uint32   `protobuf:"fixed32,3,opt,name=seq,proto3" json:"seq,omitempty"`
	Uid                  uint64   `protobuf:"fixed64,4,opt,name=uid,proto3" json:"uid,omitempty"`
	Body                 []byte   `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Request) GetServant() uint32 {
	if m != nil {
		return m.Servant
	}
	return 0
}

func (m *Request) GetSeq() uint32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Request) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *Request) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

// The response message containing the greetings
type Respond struct {
	Body                 []byte   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	Extend               []byte   `protobuf:"bytes,2,opt,name=extend,proto3" json:"extend,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Respond) Reset()         { *m = Respond{} }
func (m *Respond) String() string { return proto.CompactTextString(m) }
func (*Respond) ProtoMessage()    {}
func (*Respond) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{1}
}

func (m *Respond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Respond.Unmarshal(m, b)
}
func (m *Respond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Respond.Marshal(b, m, deterministic)
}
func (m *Respond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Respond.Merge(m, src)
}
func (m *Respond) XXX_Size() int {
	return xxx_messageInfo_Respond.Size(m)
}
func (m *Respond) XXX_DiscardUnknown() {
	xxx_messageInfo_Respond.DiscardUnknown(m)
}

var xxx_messageInfo_Respond proto.InternalMessageInfo

func (m *Respond) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Respond) GetExtend() []byte {
	if m != nil {
		return m.Extend
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "protocol.Request")
	proto.RegisterType((*Respond)(nil), "protocol.Respond")
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptor_2bc2336598a3f7e0) }

var fileDescriptor_2bc2336598a3f7e0 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0x3d, 0x0e, 0x82, 0x40,
	0x10, 0x46, 0xb3, 0x82, 0x60, 0x26, 0xc4, 0x98, 0x29, 0xcc, 0x94, 0x84, 0x8a, 0xca, 0xc6, 0x78,
	0x91, 0xbd, 0x81, 0xb8, 0x53, 0x90, 0x98, 0x1d, 0xd8, 0x1f, 0xa2, 0xb7, 0x37, 0xbb, 0x42, 0xe8,
	0xde, 0x9b, 0x57, 0x7c, 0x03, 0xe7, 0xc9, 0x49, 0x90, 0x97, 0xbc, 0x6f, 0x19, 0xf0, 0xb4, 0x79,
	0x17, 0xa1, 0xd6, 0x3c, 0x47, 0xf6, 0x01, 0x09, 0xea, 0x85, 0x9d, 0x1f, 0xc5, 0x92, 0x6a, 0x55,
	0x5f, 0xeb, 0x4d, 0x53, 0xf1, 0xec, 0x96, 0xa7, 0x0d, 0x74, 0xf8, 0x97, 0x55, 0xf1, 0x02, 0x85,
	0xe7, 0x99, 0x8a, 0x7c, 0x4d, 0x98, 0x2e, 0x71, 0x34, 0x54, 0xb6, 0xaa, 0xaf, 0x74, 0x42, 0x44,
	0x28, 0x07, 0x31, 0x5f, 0x3a, 0xb6, 0xaa, 0x6f, 0x74, 0xe6, 0xee, 0x91, 0x66, 0xfd, 0x24, 0x76,
	0xcf, 0x6a, 0xcf, 0x78, 0x85, 0x8a, 0x3f, 0x81, 0xad, 0xc9, 0x7b, 0x8d, 0x5e, 0x6d, 0xa8, 0xf2,
	0xdf, 0xf7, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x28, 0x61, 0x40, 0xd0, 0x00, 0x00, 0x00,
}
