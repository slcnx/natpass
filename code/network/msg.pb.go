// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: msg.proto

package network

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MsgType int32

const (
	Msg_unknown     MsgType = 0
	Msg_handshake   MsgType = 1
	Msg_keepalive   MsgType = 2
	Msg_connect_req MsgType = 3
	Msg_connect_rep MsgType = 4
	Msg_disconnect  MsgType = 5
	Msg_forward     MsgType = 6
	// shell
	Msg_shell_resize MsgType = 10
	Msg_shell_data   MsgType = 11
	// vnc
	Msg_vnc_ctrl     MsgType = 20
	Msg_vnc_image    MsgType = 21
	Msg_vnc_mouse    MsgType = 22
	Msg_vnc_keyboard MsgType = 23
	Msg_vnc_cad      MsgType = 24 // ctrl+alt+del
)

// Enum value maps for MsgType.
var (
	MsgType_name = map[int32]string{
		0:  "unknown",
		1:  "handshake",
		2:  "keepalive",
		3:  "connect_req",
		4:  "connect_rep",
		5:  "disconnect",
		6:  "forward",
		10: "shell_resize",
		11: "shell_data",
		20: "vnc_ctrl",
		21: "vnc_image",
		22: "vnc_mouse",
		23: "vnc_keyboard",
		24: "vnc_cad",
	}
	MsgType_value = map[string]int32{
		"unknown":      0,
		"handshake":    1,
		"keepalive":    2,
		"connect_req":  3,
		"connect_rep":  4,
		"disconnect":   5,
		"forward":      6,
		"shell_resize": 10,
		"shell_data":   11,
		"vnc_ctrl":     20,
		"vnc_image":    21,
		"vnc_mouse":    22,
		"vnc_keyboard": 23,
		"vnc_cad":      24,
	}
)

func (x MsgType) Enum() *MsgType {
	p := new(MsgType)
	*p = x
	return p
}

func (x MsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_msg_proto_enumTypes[0].Descriptor()
}

func (MsgType) Type() protoreflect.EnumType {
	return &file_msg_proto_enumTypes[0]
}

func (x MsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgType.Descriptor instead.
func (MsgType) EnumDescriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{1, 0}
}

type HandshakePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enc []byte `protobuf:"bytes,1,opt,name=enc,proto3" json:"enc,omitempty"`
}

func (x *HandshakePayload) Reset() {
	*x = HandshakePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandshakePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandshakePayload) ProtoMessage() {}

func (x *HandshakePayload) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandshakePayload.ProtoReflect.Descriptor instead.
func (*HandshakePayload) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

func (x *HandshakePayload) GetEnc() []byte {
	if x != nil {
		return x.Enc
	}
	return nil
}

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XType   MsgType `protobuf:"varint,1,opt,name=_type,json=Type,proto3,enum=network.MsgType" json:"_type,omitempty"`
	From    string  `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	FromIdx uint32  `protobuf:"varint,3,opt,name=from_idx,json=fromIdx,proto3" json:"from_idx,omitempty"`
	To      string  `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	ToIdx   uint32  `protobuf:"varint,5,opt,name=to_idx,json=toIdx,proto3" json:"to_idx,omitempty"`
	LinkId  string  `protobuf:"bytes,6,opt,name=link_id,json=linkId,proto3" json:"link_id,omitempty"`
	// Types that are assignable to Payload:
	//	*Msg_Hsp
	//	*Msg_Creq
	//	*Msg_Crep
	//	*Msg_XData
	//	*Msg_Sresize
	//	*Msg_Sdata
	//	*Msg_Vctrl
	//	*Msg_Vimg
	//	*Msg_Vmouse
	//	*Msg_Vkbd
	Payload isMsg_Payload `protobuf_oneof:"payload"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{1}
}

func (x *Msg) GetXType() MsgType {
	if x != nil {
		return x.XType
	}
	return Msg_unknown
}

func (x *Msg) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Msg) GetFromIdx() uint32 {
	if x != nil {
		return x.FromIdx
	}
	return 0
}

func (x *Msg) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Msg) GetToIdx() uint32 {
	if x != nil {
		return x.ToIdx
	}
	return 0
}

func (x *Msg) GetLinkId() string {
	if x != nil {
		return x.LinkId
	}
	return ""
}

func (m *Msg) GetPayload() isMsg_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Msg) GetHsp() *HandshakePayload {
	if x, ok := x.GetPayload().(*Msg_Hsp); ok {
		return x.Hsp
	}
	return nil
}

func (x *Msg) GetCreq() *ConnectRequest {
	if x, ok := x.GetPayload().(*Msg_Creq); ok {
		return x.Creq
	}
	return nil
}

func (x *Msg) GetCrep() *ConnectResponse {
	if x, ok := x.GetPayload().(*Msg_Crep); ok {
		return x.Crep
	}
	return nil
}

func (x *Msg) GetXData() *Data {
	if x, ok := x.GetPayload().(*Msg_XData); ok {
		return x.XData
	}
	return nil
}

func (x *Msg) GetSresize() *ShellResize {
	if x, ok := x.GetPayload().(*Msg_Sresize); ok {
		return x.Sresize
	}
	return nil
}

func (x *Msg) GetSdata() *ShellData {
	if x, ok := x.GetPayload().(*Msg_Sdata); ok {
		return x.Sdata
	}
	return nil
}

func (x *Msg) GetVctrl() *VncControl {
	if x, ok := x.GetPayload().(*Msg_Vctrl); ok {
		return x.Vctrl
	}
	return nil
}

func (x *Msg) GetVimg() *VncImage {
	if x, ok := x.GetPayload().(*Msg_Vimg); ok {
		return x.Vimg
	}
	return nil
}

func (x *Msg) GetVmouse() *VncMouse {
	if x, ok := x.GetPayload().(*Msg_Vmouse); ok {
		return x.Vmouse
	}
	return nil
}

func (x *Msg) GetVkbd() *VncKeyboard {
	if x, ok := x.GetPayload().(*Msg_Vkbd); ok {
		return x.Vkbd
	}
	return nil
}

type isMsg_Payload interface {
	isMsg_Payload()
}

type Msg_Hsp struct {
	Hsp *HandshakePayload `protobuf:"bytes,10,opt,name=hsp,proto3,oneof"`
}

type Msg_Creq struct {
	Creq *ConnectRequest `protobuf:"bytes,11,opt,name=creq,proto3,oneof"`
}

type Msg_Crep struct {
	Crep *ConnectResponse `protobuf:"bytes,12,opt,name=crep,proto3,oneof"`
}

type Msg_XData struct {
	XData *Data `protobuf:"bytes,13,opt,name=_data,json=Data,proto3,oneof"`
}

type Msg_Sresize struct {
	// shell
	Sresize *ShellResize `protobuf:"bytes,20,opt,name=sresize,proto3,oneof"`
}

type Msg_Sdata struct {
	Sdata *ShellData `protobuf:"bytes,21,opt,name=sdata,proto3,oneof"`
}

type Msg_Vctrl struct {
	// vnc
	Vctrl *VncControl `protobuf:"bytes,30,opt,name=vctrl,proto3,oneof"`
}

type Msg_Vimg struct {
	Vimg *VncImage `protobuf:"bytes,31,opt,name=vimg,proto3,oneof"`
}

type Msg_Vmouse struct {
	Vmouse *VncMouse `protobuf:"bytes,32,opt,name=vmouse,proto3,oneof"`
}

type Msg_Vkbd struct {
	Vkbd *VncKeyboard `protobuf:"bytes,33,opt,name=vkbd,proto3,oneof"`
}

func (*Msg_Hsp) isMsg_Payload() {}

func (*Msg_Creq) isMsg_Payload() {}

func (*Msg_Crep) isMsg_Payload() {}

func (*Msg_XData) isMsg_Payload() {}

func (*Msg_Sresize) isMsg_Payload() {}

func (*Msg_Sdata) isMsg_Payload() {}

func (*Msg_Vctrl) isMsg_Payload() {}

func (*Msg_Vimg) isMsg_Payload() {}

func (*Msg_Vmouse) isMsg_Payload() {}

func (*Msg_Vkbd) isMsg_Payload() {}

var File_msg_proto protoreflect.FileDescriptor

var file_msg_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x1a, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0b, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x09, 0x76, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x11, 0x68, 0x61,
	0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x6e, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x65, 0x6e,
	0x63, 0x22, 0xd1, 0x06, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x26, 0x0a, 0x05, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x69, 0x64,
	0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x64, 0x78,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f,
	0x12, 0x15, 0x0a, 0x06, 0x74, 0x6f, 0x5f, 0x69, 0x64, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x74, 0x6f, 0x49, 0x64, 0x78, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x69, 0x6e, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x69, 0x6e, 0x6b, 0x49, 0x64,
	0x12, 0x2e, 0x0a, 0x03, 0x68, 0x73, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b,
	0x65, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x00, 0x52, 0x03, 0x68, 0x73, 0x70,
	0x12, 0x2e, 0x0a, 0x04, 0x63, 0x72, 0x65, 0x71, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x04, 0x63, 0x72, 0x65, 0x71,
	0x12, 0x2f, 0x0a, 0x04, 0x63, 0x72, 0x65, 0x70, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x04, 0x63, 0x72, 0x65,
	0x70, 0x12, 0x24, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x48,
	0x00, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x31, 0x0a, 0x07, 0x73, 0x72, 0x65, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x48,
	0x00, 0x52, 0x07, 0x73, 0x72, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x73, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00,
	0x52, 0x05, 0x73, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x63, 0x74, 0x72, 0x6c,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x76, 0x6e, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x48, 0x00, 0x52, 0x05,
	0x76, 0x63, 0x74, 0x72, 0x6c, 0x12, 0x28, 0x0a, 0x04, 0x76, 0x69, 0x6d, 0x67, 0x18, 0x1f, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x6e,
	0x63, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x04, 0x76, 0x69, 0x6d, 0x67, 0x12,
	0x2c, 0x0a, 0x06, 0x76, 0x6d, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x20, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x6e, 0x63, 0x5f, 0x6d, 0x6f,
	0x75, 0x73, 0x65, 0x48, 0x00, 0x52, 0x06, 0x76, 0x6d, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x2b, 0x0a,
	0x04, 0x76, 0x6b, 0x62, 0x64, 0x18, 0x21, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x6e, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x48, 0x00, 0x52, 0x04, 0x76, 0x6b, 0x62, 0x64, 0x22, 0xdd, 0x01, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00,
	0x12, 0x0d, 0x0a, 0x09, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x10, 0x01, 0x12,
	0x0d, 0x0a, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x10, 0x02, 0x12, 0x0f,
	0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x10, 0x03, 0x12,
	0x0f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x10, 0x04,
	0x12, 0x0e, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x10, 0x05,
	0x12, 0x0b, 0x0a, 0x07, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x10, 0x06, 0x12, 0x10, 0x0a,
	0x0c, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x10, 0x0a, 0x12,
	0x0e, 0x0a, 0x0a, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x10, 0x0b, 0x12,
	0x0c, 0x0a, 0x08, 0x76, 0x6e, 0x63, 0x5f, 0x63, 0x74, 0x72, 0x6c, 0x10, 0x14, 0x12, 0x0d, 0x0a,
	0x09, 0x76, 0x6e, 0x63, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x10, 0x15, 0x12, 0x0d, 0x0a, 0x09,
	0x76, 0x6e, 0x63, 0x5f, 0x6d, 0x6f, 0x75, 0x73, 0x65, 0x10, 0x16, 0x12, 0x10, 0x0a, 0x0c, 0x76,
	0x6e, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x10, 0x17, 0x12, 0x0b, 0x0a,
	0x07, 0x76, 0x6e, 0x63, 0x5f, 0x63, 0x61, 0x64, 0x10, 0x18, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_proto_rawDescOnce sync.Once
	file_msg_proto_rawDescData = file_msg_proto_rawDesc
)

func file_msg_proto_rawDescGZIP() []byte {
	file_msg_proto_rawDescOnce.Do(func() {
		file_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_proto_rawDescData)
	})
	return file_msg_proto_rawDescData
}

var file_msg_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_msg_proto_goTypes = []interface{}{
	(MsgType)(0),             // 0: network.msg.type
	(*HandshakePayload)(nil), // 1: network.handshake_payload
	(*Msg)(nil),              // 2: network.msg
	(*ConnectRequest)(nil),   // 3: network.connect_request
	(*ConnectResponse)(nil),  // 4: network.connect_response
	(*Data)(nil),             // 5: network.data
	(*ShellResize)(nil),      // 6: network.shell_resize
	(*ShellData)(nil),        // 7: network.shell_data
	(*VncControl)(nil),       // 8: network.vnc_control
	(*VncImage)(nil),         // 9: network.vnc_image
	(*VncMouse)(nil),         // 10: network.vnc_mouse
	(*VncKeyboard)(nil),      // 11: network.vnc_keyboard
}
var file_msg_proto_depIdxs = []int32{
	0,  // 0: network.msg._type:type_name -> network.msg.type
	1,  // 1: network.msg.hsp:type_name -> network.handshake_payload
	3,  // 2: network.msg.creq:type_name -> network.connect_request
	4,  // 3: network.msg.crep:type_name -> network.connect_response
	5,  // 4: network.msg._data:type_name -> network.data
	6,  // 5: network.msg.sresize:type_name -> network.shell_resize
	7,  // 6: network.msg.sdata:type_name -> network.shell_data
	8,  // 7: network.msg.vctrl:type_name -> network.vnc_control
	9,  // 8: network.msg.vimg:type_name -> network.vnc_image
	10, // 9: network.msg.vmouse:type_name -> network.vnc_mouse
	11, // 10: network.msg.vkbd:type_name -> network.vnc_keyboard
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_msg_proto_init() }
func file_msg_proto_init() {
	if File_msg_proto != nil {
		return
	}
	file_connect_proto_init()
	file_forward_proto_init()
	file_shell_proto_init()
	file_vnc_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandshakePayload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_msg_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Msg_Hsp)(nil),
		(*Msg_Creq)(nil),
		(*Msg_Crep)(nil),
		(*Msg_XData)(nil),
		(*Msg_Sresize)(nil),
		(*Msg_Sdata)(nil),
		(*Msg_Vctrl)(nil),
		(*Msg_Vimg)(nil),
		(*Msg_Vmouse)(nil),
		(*Msg_Vkbd)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msg_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_proto_goTypes,
		DependencyIndexes: file_msg_proto_depIdxs,
		EnumInfos:         file_msg_proto_enumTypes,
		MessageInfos:      file_msg_proto_msgTypes,
	}.Build()
	File_msg_proto = out.File
	file_msg_proto_rawDesc = nil
	file_msg_proto_goTypes = nil
	file_msg_proto_depIdxs = nil
}
