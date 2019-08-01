// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv/appcron/proto/appcron.proto

package appcron

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

type CronTask struct {
	TaskName             string    `protobuf:"bytes,1,opt,name=TaskName,proto3" json:"TaskName,omitempty"`
	TaskPeriod           string    `protobuf:"bytes,2,opt,name=TaskPeriod,proto3" json:"TaskPeriod,omitempty"`
	Status               string    `protobuf:"bytes,3,opt,name=Status,proto3" json:"Status,omitempty"`
	TaskType             string    `protobuf:"bytes,4,opt,name=TaskType,proto3" json:"TaskType,omitempty"`
	Tm                   *TaskMeta `protobuf:"bytes,5,opt,name=tm,proto3" json:"tm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CronTask) Reset()         { *m = CronTask{} }
func (m *CronTask) String() string { return proto.CompactTextString(m) }
func (*CronTask) ProtoMessage()    {}
func (*CronTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_73d61ef8d3da2acd, []int{0}
}

func (m *CronTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CronTask.Unmarshal(m, b)
}
func (m *CronTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CronTask.Marshal(b, m, deterministic)
}
func (m *CronTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CronTask.Merge(m, src)
}
func (m *CronTask) XXX_Size() int {
	return xxx_messageInfo_CronTask.Size(m)
}
func (m *CronTask) XXX_DiscardUnknown() {
	xxx_messageInfo_CronTask.DiscardUnknown(m)
}

var xxx_messageInfo_CronTask proto.InternalMessageInfo

func (m *CronTask) GetTaskName() string {
	if m != nil {
		return m.TaskName
	}
	return ""
}

func (m *CronTask) GetTaskPeriod() string {
	if m != nil {
		return m.TaskPeriod
	}
	return ""
}

func (m *CronTask) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *CronTask) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

func (m *CronTask) GetTm() *TaskMeta {
	if m != nil {
		return m.Tm
	}
	return nil
}

type TaskMeta struct {
	TaskCron             string   `protobuf:"bytes,1,opt,name=TaskCron,proto3" json:"TaskCron,omitempty"`
	TaskParameter        string   `protobuf:"bytes,2,opt,name=TaskParameter,proto3" json:"TaskParameter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskMeta) Reset()         { *m = TaskMeta{} }
func (m *TaskMeta) String() string { return proto.CompactTextString(m) }
func (*TaskMeta) ProtoMessage()    {}
func (*TaskMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_73d61ef8d3da2acd, []int{1}
}

func (m *TaskMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskMeta.Unmarshal(m, b)
}
func (m *TaskMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskMeta.Marshal(b, m, deterministic)
}
func (m *TaskMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskMeta.Merge(m, src)
}
func (m *TaskMeta) XXX_Size() int {
	return xxx_messageInfo_TaskMeta.Size(m)
}
func (m *TaskMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskMeta.DiscardUnknown(m)
}

var xxx_messageInfo_TaskMeta proto.InternalMessageInfo

func (m *TaskMeta) GetTaskCron() string {
	if m != nil {
		return m.TaskCron
	}
	return ""
}

func (m *TaskMeta) GetTaskParameter() string {
	if m != nil {
		return m.TaskParameter
	}
	return ""
}

type ListCronTask struct {
	Ct                   []*CronTask `protobuf:"bytes,1,rep,name=ct,proto3" json:"ct,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListCronTask) Reset()         { *m = ListCronTask{} }
func (m *ListCronTask) String() string { return proto.CompactTextString(m) }
func (*ListCronTask) ProtoMessage()    {}
func (*ListCronTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_73d61ef8d3da2acd, []int{2}
}

func (m *ListCronTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCronTask.Unmarshal(m, b)
}
func (m *ListCronTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCronTask.Marshal(b, m, deterministic)
}
func (m *ListCronTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCronTask.Merge(m, src)
}
func (m *ListCronTask) XXX_Size() int {
	return xxx_messageInfo_ListCronTask.Size(m)
}
func (m *ListCronTask) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCronTask.DiscardUnknown(m)
}

var xxx_messageInfo_ListCronTask proto.InternalMessageInfo

func (m *ListCronTask) GetCt() []*CronTask {
	if m != nil {
		return m.Ct
	}
	return nil
}

func init() {
	proto.RegisterType((*CronTask)(nil), "CronTask")
	proto.RegisterType((*TaskMeta)(nil), "TaskMeta")
	proto.RegisterType((*ListCronTask)(nil), "ListCronTask")
}

func init() { proto.RegisterFile("srv/appcron/proto/appcron.proto", fileDescriptor_73d61ef8d3da2acd) }

var fileDescriptor_73d61ef8d3da2acd = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xc1, 0x4e, 0x02, 0x31,
	0x10, 0xa5, 0x8b, 0x22, 0x0c, 0x62, 0x48, 0x0f, 0xa6, 0xec, 0x41, 0xc9, 0xc6, 0x18, 0x3c, 0xb8,
	0x1b, 0xf1, 0x0b, 0x8c, 0x31, 0x5e, 0xd0, 0x98, 0x15, 0xcf, 0xa6, 0xc0, 0x84, 0x10, 0x59, 0xda,
	0xb4, 0xa3, 0x09, 0x1f, 0xe2, 0xcd, 0x8f, 0x35, 0x6d, 0x59, 0x76, 0xbd, 0x79, 0xd9, 0xd9, 0xf7,
	0x66, 0xde, 0xeb, 0xeb, 0x14, 0xce, 0xad, 0xf9, 0xca, 0xa4, 0xd6, 0x73, 0xa3, 0x36, 0x99, 0x36,
	0x8a, 0x54, 0x89, 0x52, 0x8f, 0xe2, 0xcb, 0xe5, 0x5a, 0xcd, 0xe4, 0x3a, 0x0b, 0xe5, 0x3d, 0x8c,
	0x84, 0x6f, 0xa0, 0xc2, 0x5c, 0xf2, 0xcd, 0xa0, 0x7d, 0x6f, 0xd4, 0x66, 0x2a, 0xed, 0x07, 0x8f,
	0xa1, 0xed, 0xea, 0xb3, 0x2c, 0x50, 0xb0, 0x21, 0x1b, 0x75, 0xf2, 0x3d, 0xe6, 0x67, 0x00, 0xee,
	0xff, 0x05, 0xcd, 0x4a, 0x2d, 0x44, 0xe4, 0xbb, 0x35, 0x86, 0x9f, 0x42, 0xeb, 0x95, 0x24, 0x7d,
	0x5a, 0xd1, 0xf4, 0xbd, 0x1d, 0x2a, 0x3d, 0xa7, 0x5b, 0x8d, 0xe2, 0xa0, 0xf2, 0x74, 0x98, 0x0f,
	0x20, 0xa2, 0x42, 0x1c, 0x0e, 0xd9, 0xa8, 0x3b, 0xee, 0xa4, 0x8e, 0x7e, 0x42, 0x92, 0x79, 0x44,
	0x45, 0x32, 0x09, 0x32, 0x87, 0x4b, 0x0b, 0x17, 0xb3, 0x1e, 0xcb, 0x61, 0x7e, 0x01, 0x3d, 0x1f,
	0x42, 0x1a, 0x59, 0x20, 0xa1, 0xd9, 0x25, 0xfb, 0x4b, 0x26, 0x57, 0x70, 0x3c, 0x59, 0x59, 0xda,
	0x5f, 0x74, 0x00, 0xd1, 0x9c, 0x04, 0x1b, 0x36, 0xfd, 0xc1, 0x25, 0x9d, 0x47, 0x73, 0x1a, 0xff,
	0x30, 0x38, 0xba, 0x0b, 0xab, 0xe4, 0xd7, 0xd0, 0x7d, 0xc4, 0x4a, 0xd5, 0x4b, 0x77, 0xab, 0x7b,
	0x28, 0x34, 0x6d, 0xe3, 0x5e, 0x5a, 0xf7, 0x4c, 0x1a, 0xfc, 0x06, 0xfa, 0x39, 0x2e, 0x57, 0x96,
	0xd0, 0xec, 0x35, 0x95, 0x7b, 0xdc, 0x2f, 0xe5, 0x39, 0x5a, 0xad, 0x36, 0x16, 0x93, 0x06, 0xcf,
	0xe0, 0xe4, 0x4d, 0x2f, 0x24, 0xe1, 0x3f, 0x05, 0xb3, 0x96, 0x7f, 0xb6, 0xdb, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x95, 0xd0, 0xaf, 0x3d, 0x01, 0x02, 0x00, 0x00,
}
