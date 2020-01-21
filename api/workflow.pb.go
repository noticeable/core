// Code generated by protoc-gen-go. DO NOT EDIT.
// source: workflow.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateWorkflowRequest struct {
	Namespace            string    `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Workflow             *Workflow `protobuf:"bytes,2,opt,name=workflow,proto3" json:"workflow,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateWorkflowRequest) Reset()         { *m = CreateWorkflowRequest{} }
func (m *CreateWorkflowRequest) String() string { return proto.CompactTextString(m) }
func (*CreateWorkflowRequest) ProtoMessage()    {}
func (*CreateWorkflowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{0}
}

func (m *CreateWorkflowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateWorkflowRequest.Unmarshal(m, b)
}
func (m *CreateWorkflowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateWorkflowRequest.Marshal(b, m, deterministic)
}
func (m *CreateWorkflowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWorkflowRequest.Merge(m, src)
}
func (m *CreateWorkflowRequest) XXX_Size() int {
	return xxx_messageInfo_CreateWorkflowRequest.Size(m)
}
func (m *CreateWorkflowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWorkflowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWorkflowRequest proto.InternalMessageInfo

func (m *CreateWorkflowRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *CreateWorkflowRequest) GetWorkflow() *Workflow {
	if m != nil {
		return m.Workflow
	}
	return nil
}

type GetWorkflowRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWorkflowRequest) Reset()         { *m = GetWorkflowRequest{} }
func (m *GetWorkflowRequest) String() string { return proto.CompactTextString(m) }
func (*GetWorkflowRequest) ProtoMessage()    {}
func (*GetWorkflowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{1}
}

func (m *GetWorkflowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWorkflowRequest.Unmarshal(m, b)
}
func (m *GetWorkflowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWorkflowRequest.Marshal(b, m, deterministic)
}
func (m *GetWorkflowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWorkflowRequest.Merge(m, src)
}
func (m *GetWorkflowRequest) XXX_Size() int {
	return xxx_messageInfo_GetWorkflowRequest.Size(m)
}
func (m *GetWorkflowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWorkflowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetWorkflowRequest proto.InternalMessageInfo

func (m *GetWorkflowRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetWorkflowRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type WatchWorkflowRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatchWorkflowRequest) Reset()         { *m = WatchWorkflowRequest{} }
func (m *WatchWorkflowRequest) String() string { return proto.CompactTextString(m) }
func (*WatchWorkflowRequest) ProtoMessage()    {}
func (*WatchWorkflowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{2}
}

func (m *WatchWorkflowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchWorkflowRequest.Unmarshal(m, b)
}
func (m *WatchWorkflowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchWorkflowRequest.Marshal(b, m, deterministic)
}
func (m *WatchWorkflowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchWorkflowRequest.Merge(m, src)
}
func (m *WatchWorkflowRequest) XXX_Size() int {
	return xxx_messageInfo_WatchWorkflowRequest.Size(m)
}
func (m *WatchWorkflowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchWorkflowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WatchWorkflowRequest proto.InternalMessageInfo

func (m *WatchWorkflowRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *WatchWorkflowRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetWorkflowLogsRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PodName              string   `protobuf:"bytes,3,opt,name=podName,proto3" json:"podName,omitempty"`
	ContainerName        string   `protobuf:"bytes,4,opt,name=containerName,proto3" json:"containerName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWorkflowLogsRequest) Reset()         { *m = GetWorkflowLogsRequest{} }
func (m *GetWorkflowLogsRequest) String() string { return proto.CompactTextString(m) }
func (*GetWorkflowLogsRequest) ProtoMessage()    {}
func (*GetWorkflowLogsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{3}
}

func (m *GetWorkflowLogsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWorkflowLogsRequest.Unmarshal(m, b)
}
func (m *GetWorkflowLogsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWorkflowLogsRequest.Marshal(b, m, deterministic)
}
func (m *GetWorkflowLogsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWorkflowLogsRequest.Merge(m, src)
}
func (m *GetWorkflowLogsRequest) XXX_Size() int {
	return xxx_messageInfo_GetWorkflowLogsRequest.Size(m)
}
func (m *GetWorkflowLogsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWorkflowLogsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetWorkflowLogsRequest proto.InternalMessageInfo

func (m *GetWorkflowLogsRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetWorkflowLogsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetWorkflowLogsRequest) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

func (m *GetWorkflowLogsRequest) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

type ListWorkflowsRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	WorkflowTemplateUid  string   `protobuf:"bytes,2,opt,name=workflowTemplateUid,proto3" json:"workflowTemplateUid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListWorkflowsRequest) Reset()         { *m = ListWorkflowsRequest{} }
func (m *ListWorkflowsRequest) String() string { return proto.CompactTextString(m) }
func (*ListWorkflowsRequest) ProtoMessage()    {}
func (*ListWorkflowsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{4}
}

func (m *ListWorkflowsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWorkflowsRequest.Unmarshal(m, b)
}
func (m *ListWorkflowsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWorkflowsRequest.Marshal(b, m, deterministic)
}
func (m *ListWorkflowsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWorkflowsRequest.Merge(m, src)
}
func (m *ListWorkflowsRequest) XXX_Size() int {
	return xxx_messageInfo_ListWorkflowsRequest.Size(m)
}
func (m *ListWorkflowsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWorkflowsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListWorkflowsRequest proto.InternalMessageInfo

func (m *ListWorkflowsRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ListWorkflowsRequest) GetWorkflowTemplateUid() string {
	if m != nil {
		return m.WorkflowTemplateUid
	}
	return ""
}

type ListWorkflowsResponse struct {
	Count                int32       `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Workflows            []*Workflow `protobuf:"bytes,2,rep,name=workflows,proto3" json:"workflows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListWorkflowsResponse) Reset()         { *m = ListWorkflowsResponse{} }
func (m *ListWorkflowsResponse) String() string { return proto.CompactTextString(m) }
func (*ListWorkflowsResponse) ProtoMessage()    {}
func (*ListWorkflowsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{5}
}

func (m *ListWorkflowsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWorkflowsResponse.Unmarshal(m, b)
}
func (m *ListWorkflowsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWorkflowsResponse.Marshal(b, m, deterministic)
}
func (m *ListWorkflowsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWorkflowsResponse.Merge(m, src)
}
func (m *ListWorkflowsResponse) XXX_Size() int {
	return xxx_messageInfo_ListWorkflowsResponse.Size(m)
}
func (m *ListWorkflowsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWorkflowsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListWorkflowsResponse proto.InternalMessageInfo

func (m *ListWorkflowsResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListWorkflowsResponse) GetWorkflows() []*Workflow {
	if m != nil {
		return m.Workflows
	}
	return nil
}

type LogEntry struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogEntry) Reset()         { *m = LogEntry{} }
func (m *LogEntry) String() string { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()    {}
func (*LogEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{6}
}

func (m *LogEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEntry.Unmarshal(m, b)
}
func (m *LogEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEntry.Marshal(b, m, deterministic)
}
func (m *LogEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEntry.Merge(m, src)
}
func (m *LogEntry) XXX_Size() int {
	return xxx_messageInfo_LogEntry.Size(m)
}
func (m *LogEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEntry.DiscardUnknown(m)
}

var xxx_messageInfo_LogEntry proto.InternalMessageInfo

func (m *LogEntry) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type Workflow struct {
	Uid                  string               `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status               string               `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Parameters           []*WorkflowParameter `protobuf:"bytes,4,rep,name=parameters,proto3" json:"parameters,omitempty"`
	WorkflowTemplate     *WorkflowTemplate    `protobuf:"bytes,5,opt,name=workflowTemplate,proto3" json:"workflowTemplate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Workflow) Reset()         { *m = Workflow{} }
func (m *Workflow) String() string { return proto.CompactTextString(m) }
func (*Workflow) ProtoMessage()    {}
func (*Workflow) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{7}
}

func (m *Workflow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Workflow.Unmarshal(m, b)
}
func (m *Workflow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Workflow.Marshal(b, m, deterministic)
}
func (m *Workflow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Workflow.Merge(m, src)
}
func (m *Workflow) XXX_Size() int {
	return xxx_messageInfo_Workflow.Size(m)
}
func (m *Workflow) XXX_DiscardUnknown() {
	xxx_messageInfo_Workflow.DiscardUnknown(m)
}

var xxx_messageInfo_Workflow proto.InternalMessageInfo

func (m *Workflow) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *Workflow) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Workflow) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Workflow) GetParameters() []*WorkflowParameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *Workflow) GetWorkflowTemplate() *WorkflowTemplate {
	if m != nil {
		return m.WorkflowTemplate
	}
	return nil
}

type WorkflowParameter struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowParameter) Reset()         { *m = WorkflowParameter{} }
func (m *WorkflowParameter) String() string { return proto.CompactTextString(m) }
func (*WorkflowParameter) ProtoMessage()    {}
func (*WorkflowParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_892c7f566756b0be, []int{8}
}

func (m *WorkflowParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowParameter.Unmarshal(m, b)
}
func (m *WorkflowParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowParameter.Marshal(b, m, deterministic)
}
func (m *WorkflowParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowParameter.Merge(m, src)
}
func (m *WorkflowParameter) XXX_Size() int {
	return xxx_messageInfo_WorkflowParameter.Size(m)
}
func (m *WorkflowParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowParameter.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowParameter proto.InternalMessageInfo

func (m *WorkflowParameter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WorkflowParameter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateWorkflowRequest)(nil), "api.CreateWorkflowRequest")
	proto.RegisterType((*GetWorkflowRequest)(nil), "api.GetWorkflowRequest")
	proto.RegisterType((*WatchWorkflowRequest)(nil), "api.WatchWorkflowRequest")
	proto.RegisterType((*GetWorkflowLogsRequest)(nil), "api.GetWorkflowLogsRequest")
	proto.RegisterType((*ListWorkflowsRequest)(nil), "api.ListWorkflowsRequest")
	proto.RegisterType((*ListWorkflowsResponse)(nil), "api.ListWorkflowsResponse")
	proto.RegisterType((*LogEntry)(nil), "api.LogEntry")
	proto.RegisterType((*Workflow)(nil), "api.Workflow")
	proto.RegisterType((*WorkflowParameter)(nil), "api.WorkflowParameter")
}

func init() { proto.RegisterFile("workflow.proto", fileDescriptor_892c7f566756b0be) }

var fileDescriptor_892c7f566756b0be = []byte{
	// 773 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x96, 0xdd, 0x4e, 0x13, 0x41,
	0x14, 0xc7, 0xb3, 0x94, 0x22, 0x3d, 0x58, 0xc0, 0x43, 0x29, 0xb5, 0x62, 0xc4, 0x11, 0x13, 0x10,
	0xd2, 0x2d, 0xf5, 0x23, 0x06, 0x43, 0xa2, 0x21, 0x80, 0x17, 0x8d, 0x31, 0xf5, 0x83, 0x88, 0x17,
	0x32, 0xb4, 0x43, 0xdd, 0x58, 0x76, 0xd6, 0x9d, 0x69, 0x89, 0x69, 0x7a, 0xa3, 0x77, 0xc6, 0x3b,
	0xbd, 0xf3, 0xce, 0xa7, 0xf0, 0x5a, 0x1f, 0x81, 0x57, 0xf0, 0x41, 0xcc, 0x4e, 0x77, 0x96, 0xee,
	0x76, 0xd5, 0x56, 0xee, 0x76, 0xe6, 0x9c, 0x39, 0xe7, 0x77, 0x66, 0xce, 0xff, 0xb4, 0x30, 0x79,
	0xcc, 0xdd, 0x37, 0x87, 0x0d, 0x7e, 0x5c, 0x70, 0x5c, 0x2e, 0x39, 0x26, 0xa8, 0x63, 0xe5, 0xe7,
	0xeb, 0x9c, 0xd7, 0x1b, 0xcc, 0xa4, 0x8e, 0x65, 0x52, 0xdb, 0xe6, 0x92, 0x4a, 0x8b, 0xdb, 0xa2,
	0xeb, 0x92, 0x9f, 0xd3, 0x47, 0x5e, 0x49, 0x76, 0xe4, 0x34, 0xa8, 0x64, 0x5d, 0x03, 0xd9, 0x87,
	0xd9, 0x4d, 0x97, 0x51, 0xc9, 0x76, 0x7d, 0x87, 0x0a, 0x7b, 0xdb, 0x64, 0x42, 0xe2, 0x3c, 0xa4,
	0x6c, 0x7a, 0xc4, 0x84, 0x43, 0xab, 0x2c, 0x67, 0x2c, 0x18, 0x4b, 0xa9, 0xca, 0xe9, 0x06, 0x2e,
	0xc3, 0xb8, 0x8e, 0x98, 0x1b, 0x59, 0x30, 0x96, 0x26, 0x4a, 0xe9, 0x02, 0x75, 0xac, 0x42, 0x10,
	0x25, 0x30, 0x93, 0x6d, 0xc0, 0x1d, 0x26, 0x87, 0x0b, 0x8f, 0x30, 0xea, 0x2d, 0x54, 0xe8, 0x54,
	0x45, 0x7d, 0x93, 0x87, 0x90, 0xd9, 0xa5, 0xb2, 0xfa, 0xfa, 0xec, 0x91, 0x3e, 0x1a, 0x90, 0xed,
	0x41, 0x2a, 0xf3, 0xba, 0xf8, 0xef, 0x60, 0x98, 0x83, 0x73, 0x0e, 0xaf, 0x3d, 0xf2, 0xb6, 0x13,
	0x6a, 0x5b, 0x2f, 0x71, 0x11, 0xd2, 0x55, 0x6e, 0x4b, 0x6a, 0xd9, 0xcc, 0x55, 0xf6, 0x51, 0x65,
	0x0f, 0x6f, 0x92, 0x43, 0xc8, 0x94, 0x2d, 0x11, 0xc0, 0x0c, 0x48, 0x52, 0x84, 0x19, 0x7d, 0xc1,
	0x4f, 0xfd, 0x07, 0x7d, 0x66, 0xd5, 0x7c, 0xb0, 0x38, 0x13, 0xd9, 0x83, 0xd9, 0x48, 0x1e, 0xe1,
	0x70, 0x5b, 0x30, 0xcc, 0x40, 0xb2, 0xca, 0x9b, 0xb6, 0x54, 0x49, 0x92, 0x95, 0xee, 0x02, 0x57,
	0x20, 0xa5, 0xa3, 0x88, 0xdc, 0xc8, 0x42, 0xa2, 0xff, 0x85, 0x4f, 0xed, 0x64, 0x11, 0xc6, 0xcb,
	0xbc, 0xbe, 0x65, 0x4b, 0xf7, 0x9d, 0x77, 0x1f, 0x5e, 0x81, 0xcc, 0x0f, 0x98, 0xaa, 0xe8, 0x25,
	0xf9, 0x69, 0xc0, 0xb8, 0x3e, 0x8d, 0xd3, 0x90, 0x68, 0x5a, 0x35, 0xdf, 0xc5, 0xfb, 0x8c, 0xbd,
	0xdc, 0x2c, 0x8c, 0x09, 0x49, 0x65, 0x53, 0xf8, 0x77, 0xeb, 0xaf, 0xf0, 0x0e, 0x80, 0x43, 0x5d,
	0x7a, 0xc4, 0x24, 0x73, 0x45, 0x6e, 0x54, 0xe1, 0x65, 0x43, 0x78, 0x8f, 0xb5, 0xb9, 0xd2, 0xe3,
	0x89, 0x0f, 0x60, 0x3a, 0x7a, 0x37, 0xb9, 0xa4, 0x6a, 0xdf, 0xd9, 0xd0, 0x69, 0x6d, 0xac, 0xf4,
	0xb9, 0x93, 0x0d, 0xb8, 0xd0, 0x97, 0x23, 0x60, 0x37, 0x7a, 0xd8, 0x33, 0x90, 0x6c, 0xd1, 0x46,
	0x53, 0x17, 0xd4, 0x5d, 0x94, 0x3e, 0x9d, 0x87, 0x29, 0x7d, 0xfe, 0x09, 0x73, 0x5b, 0x56, 0x95,
	0xa1, 0x0b, 0x93, 0x61, 0x0d, 0x62, 0x5e, 0xd1, 0xc4, 0x0a, 0x33, 0x1f, 0x7e, 0x06, 0x72, 0xfb,
	0xfd, 0xc9, 0xaf, 0xcf, 0x23, 0x26, 0xb9, 0xe6, 0x29, 0x5f, 0x98, 0xad, 0xb5, 0x03, 0x26, 0xe9,
	0x9a, 0xd9, 0x0e, 0x7a, 0xa5, 0x63, 0x06, 0x0f, 0xb5, 0x1e, 0xa8, 0x12, 0x2d, 0x98, 0xe8, 0x91,
	0x00, 0xce, 0xa9, 0xa0, 0xfd, 0x3a, 0x8d, 0x66, 0x2b, 0xa9, 0x6c, 0xab, 0x78, 0x63, 0x80, 0x6c,
	0xdd, 0xdd, 0x0e, 0x0a, 0x48, 0x87, 0x3a, 0x0f, 0x2f, 0xaa, 0x98, 0x71, 0x5d, 0x9f, 0xcf, 0xc7,
	0x99, 0xba, 0x8d, 0x4a, 0x56, 0x54, 0xee, 0xeb, 0x38, 0x48, 0xa5, 0xd8, 0x84, 0x74, 0x68, 0x5a,
	0xf8, 0x49, 0xe3, 0x26, 0x48, 0xb4, 0xc6, 0xbb, 0x2a, 0x4f, 0x09, 0x8b, 0x83, 0xd7, 0x68, 0x1e,
	0x7b, 0x71, 0x8b, 0x06, 0x7e, 0x33, 0x60, 0x2a, 0x32, 0x5a, 0xf0, 0x52, 0xf4, 0x6e, 0x7b, 0x06,
	0x8e, 0x9f, 0x5b, 0xab, 0x87, 0x50, 0x95, 0xfb, 0x25, 0xbe, 0x18, 0x22, 0xb7, 0xc3, 0x6b, 0xc2,
	0x6c, 0xfb, 0x53, 0xa7, 0x63, 0x06, 0xe3, 0x45, 0x98, 0xed, 0xd0, 0xa8, 0xe9, 0x98, 0x0d, 0x5e,
	0x17, 0x45, 0x03, 0xbf, 0x1a, 0x90, 0x0d, 0xf7, 0x96, 0xee, 0x6e, 0x24, 0x31, 0x8d, 0x17, 0x88,
	0xc1, 0x47, 0x8e, 0x97, 0x0a, 0xd9, 0x56, 0xe8, 0xf7, 0xc9, 0xea, 0xbf, 0xd1, 0x83, 0x1f, 0x21,
	0xb1, 0xde, 0x27, 0x30, 0xfc, 0x61, 0xc0, 0xe5, 0x78, 0x80, 0xe7, 0xcc, 0x15, 0x16, 0xb7, 0xcf,
	0x02, 0x79, 0xa8, 0x20, 0xf7, 0x49, 0x79, 0x18, 0x48, 0xb3, 0x1d, 0x85, 0x2c, 0x34, 0xad, 0x5a,
	0xc7, 0x6c, 0x75, 0x79, 0xe2, 0x8a, 0x38, 0x31, 0x60, 0xa6, 0xe7, 0xc5, 0x83, 0xfd, 0x2b, 0xd1,
	0x5e, 0x18, 0x90, 0xfb, 0x83, 0xa1, 0xc0, 0x3b, 0x58, 0x1a, 0x0e, 0xdc, 0xe3, 0xdc, 0xdb, 0xc1,
	0xad, 0xe1, 0x4f, 0x05, 0xd5, 0x99, 0x6d, 0xff, 0xab, 0x83, 0xdf, 0x0d, 0x98, 0xef, 0xd5, 0x66,
	0xe4, 0x61, 0x04, 0x2e, 0xf5, 0xc9, 0x37, 0xea, 0xa2, 0xeb, 0x5c, 0x1e, 0xc0, 0xd3, 0xd7, 0xfd,
	0xa6, 0x2a, 0x7d, 0x03, 0xef, 0x9d, 0xa1, 0x08, 0xfc, 0x62, 0x84, 0x7f, 0xff, 0x74, 0x36, 0x81,
	0x57, 0xff, 0x48, 0x12, 0xc0, 0x92, 0xbf, 0xb9, 0xf8, 0x94, 0xb7, 0x14, 0x65, 0x01, 0x87, 0x6a,
	0xff, 0x83, 0x31, 0xf5, 0x2f, 0xec, 0xe6, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x18, 0x2a, 0x86,
	0x9a, 0xd3, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WorkflowServiceClient is the client API for WorkflowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkflowServiceClient interface {
	// Creates a Workflow
	CreateWorkflow(ctx context.Context, in *CreateWorkflowRequest, opts ...grpc.CallOption) (*Workflow, error)
	GetWorkflow(ctx context.Context, in *GetWorkflowRequest, opts ...grpc.CallOption) (*Workflow, error)
	ListWorkflows(ctx context.Context, in *ListWorkflowsRequest, opts ...grpc.CallOption) (*ListWorkflowsResponse, error)
	WatchWorkflow(ctx context.Context, in *WatchWorkflowRequest, opts ...grpc.CallOption) (WorkflowService_WatchWorkflowClient, error)
	GetWorkflowLogs(ctx context.Context, in *GetWorkflowLogsRequest, opts ...grpc.CallOption) (WorkflowService_GetWorkflowLogsClient, error)
	CreateWorkflowTemplate(ctx context.Context, in *CreateWorkflowTemplateRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error)
	CreateWorkflowTemplateVersion(ctx context.Context, in *CreateWorkflowTemplateRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error)
	GetWorkflowTemplate(ctx context.Context, in *GetWorkflowTemplateRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error)
	ListWorkflowTemplateVersions(ctx context.Context, in *ListWorkflowTemplateVersionsRequest, opts ...grpc.CallOption) (*ListWorkflowTemplateVersionsResponse, error)
	ListWorkflowTemplates(ctx context.Context, in *ListWorkflowTemplatesRequest, opts ...grpc.CallOption) (*ListWorkflowTemplatesResponse, error)
}

type workflowServiceClient struct {
	cc *grpc.ClientConn
}

func NewWorkflowServiceClient(cc *grpc.ClientConn) WorkflowServiceClient {
	return &workflowServiceClient{cc}
}

func (c *workflowServiceClient) CreateWorkflow(ctx context.Context, in *CreateWorkflowRequest, opts ...grpc.CallOption) (*Workflow, error) {
	out := new(Workflow)
	err := c.cc.Invoke(ctx, "/api.WorkflowService/CreateWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowServiceClient) GetWorkflow(ctx context.Context, in *GetWorkflowRequest, opts ...grpc.CallOption) (*Workflow, error) {
	out := new(Workflow)
	err := c.cc.Invoke(ctx, "/api.WorkflowService/GetWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowServiceClient) ListWorkflows(ctx context.Context, in *ListWorkflowsRequest, opts ...grpc.CallOption) (*ListWorkflowsResponse, error) {
	out := new(ListWorkflowsResponse)
	err := c.cc.Invoke(ctx, "/api.WorkflowService/ListWorkflows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowServiceClient) WatchWorkflow(ctx context.Context, in *WatchWorkflowRequest, opts ...grpc.CallOption) (WorkflowService_WatchWorkflowClient, error) {
	stream, err := c.cc.NewStream(ctx, &_WorkflowService_serviceDesc.Streams[0], "/api.WorkflowService/WatchWorkflow", opts...)
	if err != nil {
		return nil, err
	}
	x := &workflowServiceWatchWorkflowClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WorkflowService_WatchWorkflowClient interface {
	Recv() (*Workflow, error)
	grpc.ClientStream
}

type workflowServiceWatchWorkflowClient struct {
	grpc.ClientStream
}

func (x *workflowServiceWatchWorkflowClient) Recv() (*Workflow, error) {
	m := new(Workflow)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *workflowServiceClient) GetWorkflowLogs(ctx context.Context, in *GetWorkflowLogsRequest, opts ...grpc.CallOption) (WorkflowService_GetWorkflowLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_WorkflowService_serviceDesc.Streams[1], "/api.WorkflowService/GetWorkflowLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &workflowServiceGetWorkflowLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WorkflowService_GetWorkflowLogsClient interface {
	Recv() (*LogEntry, error)
	grpc.ClientStream
}

type workflowServiceGetWorkflowLogsClient struct {
	grpc.ClientStream
}

func (x *workflowServiceGetWorkflowLogsClient) Recv() (*LogEntry, error) {
	m := new(LogEntry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *workflowServiceClient) CreateWorkflowTemplate(ctx context.Context, in *CreateWorkflowTemplateRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error) {
	out := new(WorkflowTemplate)
	err := c.cc.Invoke(ctx, "/api.WorkflowService/CreateWorkflowTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowServiceClient) CreateWorkflowTemplateVersion(ctx context.Context, in *CreateWorkflowTemplateRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error) {
	out := new(WorkflowTemplate)
	err := c.cc.Invoke(ctx, "/api.WorkflowService/CreateWorkflowTemplateVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowServiceClient) GetWorkflowTemplate(ctx context.Context, in *GetWorkflowTemplateRequest, opts ...grpc.CallOption) (*WorkflowTemplate, error) {
	out := new(WorkflowTemplate)
	err := c.cc.Invoke(ctx, "/api.WorkflowService/GetWorkflowTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowServiceClient) ListWorkflowTemplateVersions(ctx context.Context, in *ListWorkflowTemplateVersionsRequest, opts ...grpc.CallOption) (*ListWorkflowTemplateVersionsResponse, error) {
	out := new(ListWorkflowTemplateVersionsResponse)
	err := c.cc.Invoke(ctx, "/api.WorkflowService/ListWorkflowTemplateVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowServiceClient) ListWorkflowTemplates(ctx context.Context, in *ListWorkflowTemplatesRequest, opts ...grpc.CallOption) (*ListWorkflowTemplatesResponse, error) {
	out := new(ListWorkflowTemplatesResponse)
	err := c.cc.Invoke(ctx, "/api.WorkflowService/ListWorkflowTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkflowServiceServer is the server API for WorkflowService service.
type WorkflowServiceServer interface {
	// Creates a Workflow
	CreateWorkflow(context.Context, *CreateWorkflowRequest) (*Workflow, error)
	GetWorkflow(context.Context, *GetWorkflowRequest) (*Workflow, error)
	ListWorkflows(context.Context, *ListWorkflowsRequest) (*ListWorkflowsResponse, error)
	WatchWorkflow(*WatchWorkflowRequest, WorkflowService_WatchWorkflowServer) error
	GetWorkflowLogs(*GetWorkflowLogsRequest, WorkflowService_GetWorkflowLogsServer) error
	CreateWorkflowTemplate(context.Context, *CreateWorkflowTemplateRequest) (*WorkflowTemplate, error)
	CreateWorkflowTemplateVersion(context.Context, *CreateWorkflowTemplateRequest) (*WorkflowTemplate, error)
	GetWorkflowTemplate(context.Context, *GetWorkflowTemplateRequest) (*WorkflowTemplate, error)
	ListWorkflowTemplateVersions(context.Context, *ListWorkflowTemplateVersionsRequest) (*ListWorkflowTemplateVersionsResponse, error)
	ListWorkflowTemplates(context.Context, *ListWorkflowTemplatesRequest) (*ListWorkflowTemplatesResponse, error)
}

// UnimplementedWorkflowServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWorkflowServiceServer struct {
}

func (*UnimplementedWorkflowServiceServer) CreateWorkflow(ctx context.Context, req *CreateWorkflowRequest) (*Workflow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkflow not implemented")
}
func (*UnimplementedWorkflowServiceServer) GetWorkflow(ctx context.Context, req *GetWorkflowRequest) (*Workflow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkflow not implemented")
}
func (*UnimplementedWorkflowServiceServer) ListWorkflows(ctx context.Context, req *ListWorkflowsRequest) (*ListWorkflowsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkflows not implemented")
}
func (*UnimplementedWorkflowServiceServer) WatchWorkflow(req *WatchWorkflowRequest, srv WorkflowService_WatchWorkflowServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchWorkflow not implemented")
}
func (*UnimplementedWorkflowServiceServer) GetWorkflowLogs(req *GetWorkflowLogsRequest, srv WorkflowService_GetWorkflowLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetWorkflowLogs not implemented")
}
func (*UnimplementedWorkflowServiceServer) CreateWorkflowTemplate(ctx context.Context, req *CreateWorkflowTemplateRequest) (*WorkflowTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkflowTemplate not implemented")
}
func (*UnimplementedWorkflowServiceServer) CreateWorkflowTemplateVersion(ctx context.Context, req *CreateWorkflowTemplateRequest) (*WorkflowTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkflowTemplateVersion not implemented")
}
func (*UnimplementedWorkflowServiceServer) GetWorkflowTemplate(ctx context.Context, req *GetWorkflowTemplateRequest) (*WorkflowTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkflowTemplate not implemented")
}
func (*UnimplementedWorkflowServiceServer) ListWorkflowTemplateVersions(ctx context.Context, req *ListWorkflowTemplateVersionsRequest) (*ListWorkflowTemplateVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkflowTemplateVersions not implemented")
}
func (*UnimplementedWorkflowServiceServer) ListWorkflowTemplates(ctx context.Context, req *ListWorkflowTemplatesRequest) (*ListWorkflowTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkflowTemplates not implemented")
}

func RegisterWorkflowServiceServer(s *grpc.Server, srv WorkflowServiceServer) {
	s.RegisterService(&_WorkflowService_serviceDesc, srv)
}

func _WorkflowService_CreateWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServiceServer).CreateWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkflowService/CreateWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServiceServer).CreateWorkflow(ctx, req.(*CreateWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowService_GetWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServiceServer).GetWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkflowService/GetWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServiceServer).GetWorkflow(ctx, req.(*GetWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowService_ListWorkflows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkflowsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServiceServer).ListWorkflows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkflowService/ListWorkflows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServiceServer).ListWorkflows(ctx, req.(*ListWorkflowsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowService_WatchWorkflow_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchWorkflowRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WorkflowServiceServer).WatchWorkflow(m, &workflowServiceWatchWorkflowServer{stream})
}

type WorkflowService_WatchWorkflowServer interface {
	Send(*Workflow) error
	grpc.ServerStream
}

type workflowServiceWatchWorkflowServer struct {
	grpc.ServerStream
}

func (x *workflowServiceWatchWorkflowServer) Send(m *Workflow) error {
	return x.ServerStream.SendMsg(m)
}

func _WorkflowService_GetWorkflowLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetWorkflowLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WorkflowServiceServer).GetWorkflowLogs(m, &workflowServiceGetWorkflowLogsServer{stream})
}

type WorkflowService_GetWorkflowLogsServer interface {
	Send(*LogEntry) error
	grpc.ServerStream
}

type workflowServiceGetWorkflowLogsServer struct {
	grpc.ServerStream
}

func (x *workflowServiceGetWorkflowLogsServer) Send(m *LogEntry) error {
	return x.ServerStream.SendMsg(m)
}

func _WorkflowService_CreateWorkflowTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkflowTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServiceServer).CreateWorkflowTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkflowService/CreateWorkflowTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServiceServer).CreateWorkflowTemplate(ctx, req.(*CreateWorkflowTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowService_CreateWorkflowTemplateVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkflowTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServiceServer).CreateWorkflowTemplateVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkflowService/CreateWorkflowTemplateVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServiceServer).CreateWorkflowTemplateVersion(ctx, req.(*CreateWorkflowTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowService_GetWorkflowTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkflowTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServiceServer).GetWorkflowTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkflowService/GetWorkflowTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServiceServer).GetWorkflowTemplate(ctx, req.(*GetWorkflowTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowService_ListWorkflowTemplateVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkflowTemplateVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServiceServer).ListWorkflowTemplateVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkflowService/ListWorkflowTemplateVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServiceServer).ListWorkflowTemplateVersions(ctx, req.(*ListWorkflowTemplateVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowService_ListWorkflowTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkflowTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServiceServer).ListWorkflowTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkflowService/ListWorkflowTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServiceServer).ListWorkflowTemplates(ctx, req.(*ListWorkflowTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkflowService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.WorkflowService",
	HandlerType: (*WorkflowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWorkflow",
			Handler:    _WorkflowService_CreateWorkflow_Handler,
		},
		{
			MethodName: "GetWorkflow",
			Handler:    _WorkflowService_GetWorkflow_Handler,
		},
		{
			MethodName: "ListWorkflows",
			Handler:    _WorkflowService_ListWorkflows_Handler,
		},
		{
			MethodName: "CreateWorkflowTemplate",
			Handler:    _WorkflowService_CreateWorkflowTemplate_Handler,
		},
		{
			MethodName: "CreateWorkflowTemplateVersion",
			Handler:    _WorkflowService_CreateWorkflowTemplateVersion_Handler,
		},
		{
			MethodName: "GetWorkflowTemplate",
			Handler:    _WorkflowService_GetWorkflowTemplate_Handler,
		},
		{
			MethodName: "ListWorkflowTemplateVersions",
			Handler:    _WorkflowService_ListWorkflowTemplateVersions_Handler,
		},
		{
			MethodName: "ListWorkflowTemplates",
			Handler:    _WorkflowService_ListWorkflowTemplates_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchWorkflow",
			Handler:       _WorkflowService_WatchWorkflow_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetWorkflowLogs",
			Handler:       _WorkflowService_GetWorkflowLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "workflow.proto",
}
