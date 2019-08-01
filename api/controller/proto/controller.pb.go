// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/controller/proto/controller.proto

package controller

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "global/global_proto/proto"
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

type ControllerObj struct {
	Name                 string         `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Status               int32          `protobuf:"varint,2,opt,name=Status,proto3" json:"Status,omitempty"`
	Ovn                  *OvnController `protobuf:"bytes,3,opt,name=ovn,proto3" json:"ovn,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ControllerObj) Reset()         { *m = ControllerObj{} }
func (m *ControllerObj) String() string { return proto.CompactTextString(m) }
func (*ControllerObj) ProtoMessage()    {}
func (*ControllerObj) Descriptor() ([]byte, []int) {
	return fileDescriptor_15818d6fb51d82bd, []int{0}
}

func (m *ControllerObj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControllerObj.Unmarshal(m, b)
}
func (m *ControllerObj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControllerObj.Marshal(b, m, deterministic)
}
func (m *ControllerObj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControllerObj.Merge(m, src)
}
func (m *ControllerObj) XXX_Size() int {
	return xxx_messageInfo_ControllerObj.Size(m)
}
func (m *ControllerObj) XXX_DiscardUnknown() {
	xxx_messageInfo_ControllerObj.DiscardUnknown(m)
}

var xxx_messageInfo_ControllerObj proto.InternalMessageInfo

func (m *ControllerObj) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ControllerObj) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ControllerObj) GetOvn() *OvnController {
	if m != nil {
		return m.Ovn
	}
	return nil
}

type OvnController struct {
	Nbip                 string   `protobuf:"bytes,1,opt,name=Nbip,proto3" json:"Nbip,omitempty"`
	Sbip                 string   `protobuf:"bytes,2,opt,name=Sbip,proto3" json:"Sbip,omitempty"`
	Nbport               string   `protobuf:"bytes,3,opt,name=Nbport,proto3" json:"Nbport,omitempty"`
	Sbport               string   `protobuf:"bytes,4,opt,name=Sbport,proto3" json:"Sbport,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OvnController) Reset()         { *m = OvnController{} }
func (m *OvnController) String() string { return proto.CompactTextString(m) }
func (*OvnController) ProtoMessage()    {}
func (*OvnController) Descriptor() ([]byte, []int) {
	return fileDescriptor_15818d6fb51d82bd, []int{1}
}

func (m *OvnController) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OvnController.Unmarshal(m, b)
}
func (m *OvnController) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OvnController.Marshal(b, m, deterministic)
}
func (m *OvnController) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OvnController.Merge(m, src)
}
func (m *OvnController) XXX_Size() int {
	return xxx_messageInfo_OvnController.Size(m)
}
func (m *OvnController) XXX_DiscardUnknown() {
	xxx_messageInfo_OvnController.DiscardUnknown(m)
}

var xxx_messageInfo_OvnController proto.InternalMessageInfo

func (m *OvnController) GetNbip() string {
	if m != nil {
		return m.Nbip
	}
	return ""
}

func (m *OvnController) GetSbip() string {
	if m != nil {
		return m.Sbip
	}
	return ""
}

func (m *OvnController) GetNbport() string {
	if m != nil {
		return m.Nbport
	}
	return ""
}

func (m *OvnController) GetSbport() string {
	if m != nil {
		return m.Sbport
	}
	return ""
}

type ControllerList struct {
	Co                   []*ControllerObj `protobuf:"bytes,1,rep,name=co,proto3" json:"co,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ControllerList) Reset()         { *m = ControllerList{} }
func (m *ControllerList) String() string { return proto.CompactTextString(m) }
func (*ControllerList) ProtoMessage()    {}
func (*ControllerList) Descriptor() ([]byte, []int) {
	return fileDescriptor_15818d6fb51d82bd, []int{2}
}

func (m *ControllerList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControllerList.Unmarshal(m, b)
}
func (m *ControllerList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControllerList.Marshal(b, m, deterministic)
}
func (m *ControllerList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControllerList.Merge(m, src)
}
func (m *ControllerList) XXX_Size() int {
	return xxx_messageInfo_ControllerList.Size(m)
}
func (m *ControllerList) XXX_DiscardUnknown() {
	xxx_messageInfo_ControllerList.DiscardUnknown(m)
}

var xxx_messageInfo_ControllerList proto.InternalMessageInfo

func (m *ControllerList) GetCo() []*ControllerObj {
	if m != nil {
		return m.Co
	}
	return nil
}

func init() {
	proto.RegisterType((*ControllerObj)(nil), "controller.ControllerObj")
	proto.RegisterType((*OvnController)(nil), "controller.OvnController")
	proto.RegisterType((*ControllerList)(nil), "controller.ControllerList")
}

func init() {
	proto.RegisterFile("api/controller/proto/controller.proto", fileDescriptor_15818d6fb51d82bd)
}

var fileDescriptor_15818d6fb51d82bd = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0xbb, 0x69, 0xff, 0x85, 0xce, 0xdf, 0x48, 0xd9, 0x83, 0xc4, 0x9c, 0x42, 0x40, 0x89,
	0x08, 0x09, 0xd4, 0x63, 0x4f, 0xa5, 0x88, 0x17, 0x69, 0x60, 0x8b, 0x67, 0xc9, 0xc6, 0x25, 0xa6,
	0xa4, 0xd9, 0x25, 0x19, 0x0b, 0x7e, 0x4c, 0xbf, 0x91, 0xec, 0x6e, 0x6c, 0x52, 0xd1, 0x8b, 0x97,
	0x64, 0xe6, 0xc7, 0xcc, 0x7b, 0x6f, 0x12, 0xb8, 0xca, 0x54, 0x99, 0xe4, 0xb2, 0xc6, 0x46, 0x56,
	0x95, 0x68, 0x12, 0xd5, 0x48, 0x94, 0x03, 0x10, 0x1b, 0x40, 0xa1, 0x27, 0xfe, 0x75, 0x51, 0x49,
	0x9e, 0x55, 0x89, 0x7d, 0x3d, 0xdb, 0x0d, 0xfb, 0xb4, 0xc8, 0xee, 0x84, 0xaf, 0xe0, 0xae, 0x8f,
	0x5b, 0x29, 0xdf, 0x51, 0x0a, 0x93, 0x4d, 0xb6, 0x17, 0x1e, 0x09, 0x48, 0x34, 0x63, 0xa6, 0xa6,
	0x17, 0x30, 0xdd, 0x62, 0x86, 0x6f, 0xad, 0xe7, 0x04, 0x24, 0xfa, 0xc7, 0xba, 0x8e, 0xde, 0xc2,
	0x58, 0x1e, 0x6a, 0x6f, 0x1c, 0x90, 0xe8, 0xff, 0xe2, 0x32, 0x1e, 0x04, 0x4a, 0x0f, 0x75, 0x2f,
	0xcb, 0xf4, 0x54, 0x58, 0x80, 0x7b, 0x42, 0x8d, 0x13, 0x2f, 0xd5, 0xd1, 0x89, 0x97, 0x4a, 0xb3,
	0xad, 0x66, 0x8e, 0x65, 0xba, 0xd6, 0xee, 0x1b, 0xae, 0x64, 0x83, 0xc6, 0x68, 0xc6, 0xba, 0xce,
	0xa4, 0xb2, 0x7c, 0x62, 0xb9, 0xed, 0xc2, 0x25, 0x9c, 0xf7, 0x2e, 0x8f, 0x65, 0x8b, 0xf4, 0x06,
	0x9c, 0x5c, 0x7a, 0x24, 0x18, 0x7f, 0x8f, 0x79, 0x72, 0x3a, 0x73, 0x72, 0xb9, 0xf8, 0x20, 0x00,
	0x83, 0x8c, 0x6b, 0xa0, 0x4c, 0x14, 0x65, 0x8b, 0xa2, 0x19, 0xd0, 0xdf, 0x35, 0xfc, 0x79, 0xdc,
	0x7d, 0x5e, 0x26, 0x5a, 0x25, 0xeb, 0x56, 0x84, 0x23, 0xba, 0x82, 0xf9, 0x93, 0x7a, 0xc9, 0x50,
	0xfc, 0x5d, 0x62, 0x09, 0x67, 0x0f, 0x02, 0x57, 0x4a, 0xa5, 0x7c, 0x27, 0x72, 0xa4, 0xee, 0xd7,
	0xcc, 0xfd, 0x5e, 0xe1, 0xbb, 0xef, 0xff, 0xac, 0xa6, 0x8f, 0x0f, 0x47, 0x7c, 0x6a, 0x7e, 0xf5,
	0xdd, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x34, 0xdd, 0x82, 0x82, 0x47, 0x02, 0x00, 0x00,
}