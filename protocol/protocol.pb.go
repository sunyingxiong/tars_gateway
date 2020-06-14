// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol.proto

package protocol

import (
	context "context"
	fmt "fmt"
	tars "github.com/TarsCloud/TarsGo/tars"
	model "github.com/TarsCloud/TarsGo/tars/model"
	requestf "github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
	tools "github.com/TarsCloud/TarsGo/tars/util/tools"
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

type ECmd int32

const (
	ECmd_UNKNOWN            ECmd = 0
	ECmd_E_LOGIN_NOTIFY_REQ ECmd = 601
)

var ECmd_name = map[int32]string{
	0:   "UNKNOWN",
	601: "E_LOGIN_NOTIFY_REQ",
}

var ECmd_value = map[string]int32{
	"UNKNOWN":            0,
	"E_LOGIN_NOTIFY_REQ": 601,
}

func (x ECmd) String() string {
	return proto.EnumName(ECmd_name, int32(x))
}

func (ECmd) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{0}
}

// The request message containing the user's name.
type Request struct {
	Version              uint32   `protobuf:"fixed32,1,opt,name=version,proto3" json:"version,omitempty"`
	Servant              uint32   `protobuf:"fixed32,2,opt,name=servant,proto3" json:"servant,omitempty"`
	Seq                  uint32   `protobuf:"fixed32,3,opt,name=seq,proto3" json:"seq,omitempty"`
	Uid                  uint32   `protobuf:"fixed32,4,opt,name=uid,proto3" json:"uid,omitempty"`
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

func (m *Request) GetUid() uint32 {
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

// The request message containing the user's name.
type MsgHead struct {
	BodyLen              uint32   `protobuf:"fixed32,1,opt,name=body_len,json=bodyLen,proto3" json:"body_len,omitempty"`
	Version              uint32   `protobuf:"fixed32,2,opt,name=version,proto3" json:"version,omitempty"`
	App                  uint32   `protobuf:"fixed32,3,opt,name=app,proto3" json:"app,omitempty"`
	Server               uint32   `protobuf:"fixed32,4,opt,name=server,proto3" json:"server,omitempty"`
	Servant              uint32   `protobuf:"fixed32,5,opt,name=servant,proto3" json:"servant,omitempty"`
	Seq                  uint32   `protobuf:"fixed32,6,opt,name=seq,proto3" json:"seq,omitempty"`
	RouteId              uint64   `protobuf:"fixed64,7,opt,name=route_id,json=routeId,proto3" json:"route_id,omitempty"`
	Encrypt              uint32   `protobuf:"fixed32,8,opt,name=encrypt,proto3" json:"encrypt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgHead) Reset()         { *m = MsgHead{} }
func (m *MsgHead) String() string { return proto.CompactTextString(m) }
func (*MsgHead) ProtoMessage()    {}
func (*MsgHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{2}
}

func (m *MsgHead) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgHead.Unmarshal(m, b)
}
func (m *MsgHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgHead.Marshal(b, m, deterministic)
}
func (m *MsgHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgHead.Merge(m, src)
}
func (m *MsgHead) XXX_Size() int {
	return xxx_messageInfo_MsgHead.Size(m)
}
func (m *MsgHead) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgHead.DiscardUnknown(m)
}

var xxx_messageInfo_MsgHead proto.InternalMessageInfo

func (m *MsgHead) GetBodyLen() uint32 {
	if m != nil {
		return m.BodyLen
	}
	return 0
}

func (m *MsgHead) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *MsgHead) GetApp() uint32 {
	if m != nil {
		return m.App
	}
	return 0
}

func (m *MsgHead) GetServer() uint32 {
	if m != nil {
		return m.Server
	}
	return 0
}

func (m *MsgHead) GetServant() uint32 {
	if m != nil {
		return m.Servant
	}
	return 0
}

func (m *MsgHead) GetSeq() uint32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *MsgHead) GetRouteId() uint64 {
	if m != nil {
		return m.RouteId
	}
	return 0
}

func (m *MsgHead) GetEncrypt() uint32 {
	if m != nil {
		return m.Encrypt
	}
	return 0
}

type MsgBody struct {
	Body                 []byte   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	Extend               []byte   `protobuf:"bytes,2,opt,name=extend,proto3" json:"extend,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgBody) Reset()         { *m = MsgBody{} }
func (m *MsgBody) String() string { return proto.CompactTextString(m) }
func (*MsgBody) ProtoMessage()    {}
func (*MsgBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{3}
}

func (m *MsgBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgBody.Unmarshal(m, b)
}
func (m *MsgBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgBody.Marshal(b, m, deterministic)
}
func (m *MsgBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBody.Merge(m, src)
}
func (m *MsgBody) XXX_Size() int {
	return xxx_messageInfo_MsgBody.Size(m)
}
func (m *MsgBody) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBody.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBody proto.InternalMessageInfo

func (m *MsgBody) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *MsgBody) GetExtend() []byte {
	if m != nil {
		return m.Extend
	}
	return nil
}

type Errorinfo struct {
	ErrorCode            uint32   `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorInfo            []byte   `protobuf:"bytes,2,opt,name=error_info,json=errorInfo,proto3" json:"error_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Errorinfo) Reset()         { *m = Errorinfo{} }
func (m *Errorinfo) String() string { return proto.CompactTextString(m) }
func (*Errorinfo) ProtoMessage()    {}
func (*Errorinfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{4}
}

func (m *Errorinfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Errorinfo.Unmarshal(m, b)
}
func (m *Errorinfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Errorinfo.Marshal(b, m, deterministic)
}
func (m *Errorinfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Errorinfo.Merge(m, src)
}
func (m *Errorinfo) XXX_Size() int {
	return xxx_messageInfo_Errorinfo.Size(m)
}
func (m *Errorinfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Errorinfo.DiscardUnknown(m)
}

var xxx_messageInfo_Errorinfo proto.InternalMessageInfo

func (m *Errorinfo) GetErrorCode() uint32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *Errorinfo) GetErrorInfo() []byte {
	if m != nil {
		return m.ErrorInfo
	}
	return nil
}

type ErrorRsp struct {
	Errinfo              *Errorinfo `protobuf:"bytes,1,opt,name=errinfo,proto3" json:"errinfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ErrorRsp) Reset()         { *m = ErrorRsp{} }
func (m *ErrorRsp) String() string { return proto.CompactTextString(m) }
func (*ErrorRsp) ProtoMessage()    {}
func (*ErrorRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{5}
}

func (m *ErrorRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorRsp.Unmarshal(m, b)
}
func (m *ErrorRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorRsp.Marshal(b, m, deterministic)
}
func (m *ErrorRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorRsp.Merge(m, src)
}
func (m *ErrorRsp) XXX_Size() int {
	return xxx_messageInfo_ErrorRsp.Size(m)
}
func (m *ErrorRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorRsp.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorRsp proto.InternalMessageInfo

func (m *ErrorRsp) GetErrinfo() *Errorinfo {
	if m != nil {
		return m.Errinfo
	}
	return nil
}

type LoginNotifyReq struct {
	Uid                  uint64   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginNotifyReq) Reset()         { *m = LoginNotifyReq{} }
func (m *LoginNotifyReq) String() string { return proto.CompactTextString(m) }
func (*LoginNotifyReq) ProtoMessage()    {}
func (*LoginNotifyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{6}
}

func (m *LoginNotifyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginNotifyReq.Unmarshal(m, b)
}
func (m *LoginNotifyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginNotifyReq.Marshal(b, m, deterministic)
}
func (m *LoginNotifyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginNotifyReq.Merge(m, src)
}
func (m *LoginNotifyReq) XXX_Size() int {
	return xxx_messageInfo_LoginNotifyReq.Size(m)
}
func (m *LoginNotifyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginNotifyReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginNotifyReq proto.InternalMessageInfo

func (m *LoginNotifyReq) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func init() {
	proto.RegisterEnum("protocol.ECmd", ECmd_name, ECmd_value)
	proto.RegisterType((*Request)(nil), "protocol.Request")
	proto.RegisterType((*Respond)(nil), "protocol.Respond")
	proto.RegisterType((*MsgHead)(nil), "protocol.MsgHead")
	proto.RegisterType((*MsgBody)(nil), "protocol.MsgBody")
	proto.RegisterType((*Errorinfo)(nil), "protocol.errorinfo")
	proto.RegisterType((*ErrorRsp)(nil), "protocol.error_rsp")
	proto.RegisterType((*LoginNotifyReq)(nil), "protocol.login_notify_req")
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptor_2bc2336598a3f7e0) }

var fileDescriptor_2bc2336598a3f7e0 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x51, 0x6b, 0xd4, 0x40,
	0x10, 0x6e, 0x7a, 0x69, 0x72, 0x4e, 0xab, 0xc4, 0x15, 0x74, 0x2b, 0x08, 0x47, 0xf0, 0xe1, 0x10,
	0x2d, 0x52, 0xf1, 0xa5, 0x8f, 0x96, 0xd3, 0x06, 0xdb, 0x1c, 0x2e, 0x8a, 0xf8, 0xb4, 0xa4, 0xb7,
	0x73, 0x47, 0x20, 0xee, 0x6e, 0x77, 0x93, 0xe2, 0xfd, 0x44, 0x7f, 0x82, 0xff, 0x46, 0x76, 0xb3,
	0xb9, 0xb4, 0xe0, 0x8b, 0x6f, 0xf3, 0xcd, 0x37, 0xfb, 0xcd, 0x7c, 0x3b, 0x03, 0x8f, 0xb4, 0x51,
	0xad, 0x5a, 0xa9, 0xe6, 0xc4, 0x07, 0x64, 0x3a, 0xe0, 0xbc, 0x83, 0x94, 0xe1, 0x4d, 0x87, 0xb6,
	0x25, 0x14, 0xd2, 0x5b, 0x34, 0xb6, 0x56, 0x92, 0x46, 0xb3, 0x68, 0x9e, 0xb2, 0x01, 0x3a, 0xc6,
	0xa2, 0xb9, 0xad, 0x64, 0x4b, 0xf7, 0x7b, 0x26, 0x40, 0x92, 0xc1, 0xc4, 0xe2, 0x0d, 0x9d, 0xf8,
	0xac, 0x0b, 0x5d, 0xa6, 0xab, 0x05, 0x8d, 0xfb, 0x4c, 0x57, 0x0b, 0x42, 0x20, 0xbe, 0x56, 0x62,
	0x4b, 0x0f, 0x66, 0xd1, 0xfc, 0x88, 0xf9, 0x38, 0x7f, 0xef, 0xda, 0x5a, 0xad, 0xe4, 0x48, 0x47,
	0x23, 0x4d, 0x9e, 0x42, 0x82, 0xbf, 0x5a, 0x94, 0xc2, 0xf7, 0x3b, 0x62, 0x01, 0xe5, 0xbf, 0x23,
	0x48, 0xaf, 0xec, 0xe6, 0x02, 0x2b, 0x41, 0x8e, 0x61, 0xea, 0x6a, 0x79, 0x83, 0xbb, 0x79, 0x1d,
	0xbe, 0x44, 0x79, 0xd7, 0xc9, 0xfe, 0x7d, 0x27, 0x19, 0x4c, 0x2a, 0xad, 0x87, 0x79, 0x2b, 0xad,
	0x5d, 0x2b, 0x67, 0x06, 0x4d, 0x18, 0x39, 0xa0, 0xbb, 0x9e, 0x0f, 0xfe, 0xe9, 0x39, 0x19, 0x3d,
	0x1f, 0xc3, 0xd4, 0xa8, 0xae, 0x45, 0x5e, 0x0b, 0x9a, 0xce, 0xa2, 0x79, 0xc2, 0x52, 0x8f, 0x0b,
	0xe1, 0x64, 0x50, 0xae, 0xcc, 0x56, 0xb7, 0x74, 0xda, 0xcb, 0x04, 0xe8, 0xbe, 0xe0, 0xca, 0x6e,
	0x3e, 0x38, 0xbb, 0xff, 0xf3, 0x05, 0x05, 0x3c, 0x40, 0x63, 0x94, 0xa9, 0xe5, 0x5a, 0x91, 0x17,
	0x00, 0x1e, 0xf0, 0x95, 0x12, 0xe8, 0x9f, 0x3f, 0x64, 0x3d, 0x7d, 0xae, 0x04, 0x8e, 0xb4, 0x2b,
	0x0e, 0x3a, 0x3d, 0x5d, 0xc8, 0xb5, 0xca, 0xcf, 0x82, 0x14, 0x37, 0x56, 0x93, 0x37, 0x90, 0xa2,
	0xf1, 0xaa, 0x5e, 0xe7, 0xf0, 0xf4, 0xc9, 0xc9, 0xee, 0x68, 0x76, 0x0d, 0xd9, 0x50, 0x93, 0xbf,
	0x84, 0xac, 0x51, 0x9b, 0x5a, 0x72, 0xa9, 0xda, 0x7a, 0xbd, 0xe5, 0x66, 0x5c, 0xbd, 0x7b, 0x1e,
	0xfb, 0xd5, 0xbf, 0x7a, 0x0d, 0xf1, 0xe2, 0xfc, 0xa7, 0x20, 0x87, 0x90, 0x7e, 0x2b, 0x3f, 0x97,
	0xcb, 0xef, 0x65, 0xb6, 0x47, 0x9e, 0x01, 0x59, 0xf0, 0xcb, 0xe5, 0xa7, 0xa2, 0xe4, 0xe5, 0xf2,
	0x6b, 0xf1, 0xf1, 0x07, 0x67, 0x8b, 0x2f, 0xd9, 0x9f, 0xf8, 0xf4, 0x6c, 0x58, 0x05, 0x79, 0x0b,
	0xc9, 0x45, 0x25, 0x45, 0x83, 0xe4, 0xf1, 0x38, 0x45, 0xb8, 0xd3, 0xe7, 0xf7, 0x52, 0xfe, 0x86,
	0xf2, 0xbd, 0xeb, 0xc4, 0xe7, 0xde, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x29, 0x47, 0xe9, 0xd0,
	0xea, 0x02, 0x00, 0x00,
}

// This following code was generated by tarsrpc
// Gernerated from protocol.proto
type Server struct {
	s model.Servant
}

//SetServant is required by the servant interface.
func (obj *Server) SetServant(s model.Servant) {
	obj.s = s
}

//AddServant is required by the servant interface
func (obj *Server) AddServant(imp impServer, objStr string) {
	tars.AddServant(obj, imp, objStr)
}

//TarsSetTimeout is required by the servant interface. t is the timeout in ms.
func (obj *Server) TarsSetTimeout(t int) {
	obj.s.TarsSetTimeout(t)
}

type impServer interface {
	Handle(input Request) (output Respond, err error)
}

//Dispatch is used to call the user implement of the defined method.
func (obj *Server) Dispatch(ctx context.Context, val interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket, withContext bool) (err error) {
	input := tools.Int8ToByte(req.SBuffer)
	var output []byte
	imp := val.(impServer)
	funcName := req.SFuncName
	switch funcName {

	case "Handle":
		inputDefine := Request{}
		if err = proto.Unmarshal(input, &inputDefine); err != nil {
			return err
		}
		res, err := imp.Handle(inputDefine)
		if err != nil {
			return err
		}
		output, err = proto.Marshal(&res)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("func mismatch")
	}
	var status map[string]string
	*resp = requestf.ResponsePacket{
		IVersion:     1,
		CPacketType:  0,
		IRequestId:   req.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(output),
		Status:       status,
		SResultDesc:  "",
		Context:      req.Context,
	}
	return nil
}

// Handle is client rpc method as defined
func (obj *Server) Handle(input Request) (output Respond, err error) {
	var _status map[string]string
	var _context map[string]string
	var inputMarshal []byte
	inputMarshal, err = proto.Marshal(&input)
	if err != nil {
		return output, err
	}
	resp := new(requestf.ResponsePacket)
	ctx := context.Background()
	err = obj.s.Tars_invoke(ctx, 0, "Handle", inputMarshal, _status, _context, resp)
	if err != nil {
		return output, err
	}
	if err = proto.Unmarshal(tools.Int8ToByte(resp.SBuffer), &output); err != nil {
		return output, err
	}
	return output, nil
}
