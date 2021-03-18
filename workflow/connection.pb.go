// Code generated by protoc-gen-go. DO NOT EDIT.
// source: workflow/connection.proto

package workflow

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type StartFlowRequest struct {
	SchemaId             uint32            `protobuf:"varint,1,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"`
	DomainId             int64             `protobuf:"varint,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	Variables            map[string]string `protobuf:"bytes,3,rep,name=variables,proto3" json:"variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StartFlowRequest) Reset()         { *m = StartFlowRequest{} }
func (m *StartFlowRequest) String() string { return proto.CompactTextString(m) }
func (*StartFlowRequest) ProtoMessage()    {}
func (*StartFlowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab8abf1dcbab6634, []int{0}
}

func (m *StartFlowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartFlowRequest.Unmarshal(m, b)
}
func (m *StartFlowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartFlowRequest.Marshal(b, m, deterministic)
}
func (m *StartFlowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartFlowRequest.Merge(m, src)
}
func (m *StartFlowRequest) XXX_Size() int {
	return xxx_messageInfo_StartFlowRequest.Size(m)
}
func (m *StartFlowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartFlowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartFlowRequest proto.InternalMessageInfo

func (m *StartFlowRequest) GetSchemaId() uint32 {
	if m != nil {
		return m.SchemaId
	}
	return 0
}

func (m *StartFlowRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

func (m *StartFlowRequest) GetVariables() map[string]string {
	if m != nil {
		return m.Variables
	}
	return nil
}

type StartFlowResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartFlowResponse) Reset()         { *m = StartFlowResponse{} }
func (m *StartFlowResponse) String() string { return proto.CompactTextString(m) }
func (*StartFlowResponse) ProtoMessage()    {}
func (*StartFlowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab8abf1dcbab6634, []int{1}
}

func (m *StartFlowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartFlowResponse.Unmarshal(m, b)
}
func (m *StartFlowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartFlowResponse.Marshal(b, m, deterministic)
}
func (m *StartFlowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartFlowResponse.Merge(m, src)
}
func (m *StartFlowResponse) XXX_Size() int {
	return xxx_messageInfo_StartFlowResponse.Size(m)
}
func (m *StartFlowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartFlowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartFlowResponse proto.InternalMessageInfo

func (m *StartFlowResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DistributeAttemptRequest struct {
	SchemaId             int32             `protobuf:"varint,1,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"`
	DomainId             int64             `protobuf:"varint,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	Destination          string            `protobuf:"bytes,3,opt,name=destination,proto3" json:"destination,omitempty"`
	Display              string            `protobuf:"bytes,4,opt,name=display,proto3" json:"display,omitempty"`
	Variables            map[string]string `protobuf:"bytes,5,rep,name=variables,proto3" json:"variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DistributeAttemptRequest) Reset()         { *m = DistributeAttemptRequest{} }
func (m *DistributeAttemptRequest) String() string { return proto.CompactTextString(m) }
func (*DistributeAttemptRequest) ProtoMessage()    {}
func (*DistributeAttemptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab8abf1dcbab6634, []int{2}
}

func (m *DistributeAttemptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributeAttemptRequest.Unmarshal(m, b)
}
func (m *DistributeAttemptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributeAttemptRequest.Marshal(b, m, deterministic)
}
func (m *DistributeAttemptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributeAttemptRequest.Merge(m, src)
}
func (m *DistributeAttemptRequest) XXX_Size() int {
	return xxx_messageInfo_DistributeAttemptRequest.Size(m)
}
func (m *DistributeAttemptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributeAttemptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DistributeAttemptRequest proto.InternalMessageInfo

func (m *DistributeAttemptRequest) GetSchemaId() int32 {
	if m != nil {
		return m.SchemaId
	}
	return 0
}

func (m *DistributeAttemptRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

func (m *DistributeAttemptRequest) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *DistributeAttemptRequest) GetDisplay() string {
	if m != nil {
		return m.Display
	}
	return ""
}

func (m *DistributeAttemptRequest) GetVariables() map[string]string {
	if m != nil {
		return m.Variables
	}
	return nil
}

type DistributeAttemptResponse struct {
	// Types that are valid to be assigned to Result:
	//	*DistributeAttemptResponse_Cancel_
	//	*DistributeAttemptResponse_Confirm_
	Result               isDistributeAttemptResponse_Result `protobuf_oneof:"result"`
	Variables            map[string]string                  `protobuf:"bytes,3,rep,name=variables,proto3" json:"variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *DistributeAttemptResponse) Reset()         { *m = DistributeAttemptResponse{} }
func (m *DistributeAttemptResponse) String() string { return proto.CompactTextString(m) }
func (*DistributeAttemptResponse) ProtoMessage()    {}
func (*DistributeAttemptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab8abf1dcbab6634, []int{3}
}

func (m *DistributeAttemptResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributeAttemptResponse.Unmarshal(m, b)
}
func (m *DistributeAttemptResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributeAttemptResponse.Marshal(b, m, deterministic)
}
func (m *DistributeAttemptResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributeAttemptResponse.Merge(m, src)
}
func (m *DistributeAttemptResponse) XXX_Size() int {
	return xxx_messageInfo_DistributeAttemptResponse.Size(m)
}
func (m *DistributeAttemptResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributeAttemptResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DistributeAttemptResponse proto.InternalMessageInfo

type isDistributeAttemptResponse_Result interface {
	isDistributeAttemptResponse_Result()
}

type DistributeAttemptResponse_Cancel_ struct {
	Cancel *DistributeAttemptResponse_Cancel `protobuf:"bytes,1,opt,name=cancel,proto3,oneof"`
}

type DistributeAttemptResponse_Confirm_ struct {
	Confirm *DistributeAttemptResponse_Confirm `protobuf:"bytes,2,opt,name=confirm,proto3,oneof"`
}

func (*DistributeAttemptResponse_Cancel_) isDistributeAttemptResponse_Result() {}

func (*DistributeAttemptResponse_Confirm_) isDistributeAttemptResponse_Result() {}

func (m *DistributeAttemptResponse) GetResult() isDistributeAttemptResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *DistributeAttemptResponse) GetCancel() *DistributeAttemptResponse_Cancel {
	if x, ok := m.GetResult().(*DistributeAttemptResponse_Cancel_); ok {
		return x.Cancel
	}
	return nil
}

func (m *DistributeAttemptResponse) GetConfirm() *DistributeAttemptResponse_Confirm {
	if x, ok := m.GetResult().(*DistributeAttemptResponse_Confirm_); ok {
		return x.Confirm
	}
	return nil
}

func (m *DistributeAttemptResponse) GetVariables() map[string]string {
	if m != nil {
		return m.Variables
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DistributeAttemptResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DistributeAttemptResponse_Cancel_)(nil),
		(*DistributeAttemptResponse_Confirm_)(nil),
	}
}

type DistributeAttemptResponse_Cancel struct {
	Description          string   `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	NextDistributeSec    uint32   `protobuf:"varint,2,opt,name=next_distribute_sec,json=nextDistributeSec,proto3" json:"next_distribute_sec,omitempty"`
	Stop                 bool     `protobuf:"varint,3,opt,name=stop,proto3" json:"stop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DistributeAttemptResponse_Cancel) Reset()         { *m = DistributeAttemptResponse_Cancel{} }
func (m *DistributeAttemptResponse_Cancel) String() string { return proto.CompactTextString(m) }
func (*DistributeAttemptResponse_Cancel) ProtoMessage()    {}
func (*DistributeAttemptResponse_Cancel) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab8abf1dcbab6634, []int{3, 0}
}

func (m *DistributeAttemptResponse_Cancel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributeAttemptResponse_Cancel.Unmarshal(m, b)
}
func (m *DistributeAttemptResponse_Cancel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributeAttemptResponse_Cancel.Marshal(b, m, deterministic)
}
func (m *DistributeAttemptResponse_Cancel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributeAttemptResponse_Cancel.Merge(m, src)
}
func (m *DistributeAttemptResponse_Cancel) XXX_Size() int {
	return xxx_messageInfo_DistributeAttemptResponse_Cancel.Size(m)
}
func (m *DistributeAttemptResponse_Cancel) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributeAttemptResponse_Cancel.DiscardUnknown(m)
}

var xxx_messageInfo_DistributeAttemptResponse_Cancel proto.InternalMessageInfo

func (m *DistributeAttemptResponse_Cancel) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DistributeAttemptResponse_Cancel) GetNextDistributeSec() uint32 {
	if m != nil {
		return m.NextDistributeSec
	}
	return 0
}

func (m *DistributeAttemptResponse_Cancel) GetStop() bool {
	if m != nil {
		return m.Stop
	}
	return false
}

type DistributeAttemptResponse_Confirm struct {
	Destination          string   `protobuf:"bytes,1,opt,name=destination,proto3" json:"destination,omitempty"`
	Display              string   `protobuf:"bytes,2,opt,name=display,proto3" json:"display,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DistributeAttemptResponse_Confirm) Reset()         { *m = DistributeAttemptResponse_Confirm{} }
func (m *DistributeAttemptResponse_Confirm) String() string { return proto.CompactTextString(m) }
func (*DistributeAttemptResponse_Confirm) ProtoMessage()    {}
func (*DistributeAttemptResponse_Confirm) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab8abf1dcbab6634, []int{3, 1}
}

func (m *DistributeAttemptResponse_Confirm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributeAttemptResponse_Confirm.Unmarshal(m, b)
}
func (m *DistributeAttemptResponse_Confirm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributeAttemptResponse_Confirm.Marshal(b, m, deterministic)
}
func (m *DistributeAttemptResponse_Confirm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributeAttemptResponse_Confirm.Merge(m, src)
}
func (m *DistributeAttemptResponse_Confirm) XXX_Size() int {
	return xxx_messageInfo_DistributeAttemptResponse_Confirm.Size(m)
}
func (m *DistributeAttemptResponse_Confirm) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributeAttemptResponse_Confirm.DiscardUnknown(m)
}

var xxx_messageInfo_DistributeAttemptResponse_Confirm proto.InternalMessageInfo

func (m *DistributeAttemptResponse_Confirm) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *DistributeAttemptResponse_Confirm) GetDisplay() string {
	if m != nil {
		return m.Display
	}
	return ""
}

func init() {
	proto.RegisterType((*StartFlowRequest)(nil), "workflow.StartFlowRequest")
	proto.RegisterMapType((map[string]string)(nil), "workflow.StartFlowRequest.VariablesEntry")
	proto.RegisterType((*StartFlowResponse)(nil), "workflow.StartFlowResponse")
	proto.RegisterType((*DistributeAttemptRequest)(nil), "workflow.DistributeAttemptRequest")
	proto.RegisterMapType((map[string]string)(nil), "workflow.DistributeAttemptRequest.VariablesEntry")
	proto.RegisterType((*DistributeAttemptResponse)(nil), "workflow.DistributeAttemptResponse")
	proto.RegisterMapType((map[string]string)(nil), "workflow.DistributeAttemptResponse.VariablesEntry")
	proto.RegisterType((*DistributeAttemptResponse_Cancel)(nil), "workflow.DistributeAttemptResponse.Cancel")
	proto.RegisterType((*DistributeAttemptResponse_Confirm)(nil), "workflow.DistributeAttemptResponse.Confirm")
}

func init() { proto.RegisterFile("workflow/connection.proto", fileDescriptor_ab8abf1dcbab6634) }

var fileDescriptor_ab8abf1dcbab6634 = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0x8e, 0xed, 0xe6, 0xc3, 0x13, 0xb5, 0x6a, 0xf6, 0x7d, 0x0f, 0x6e, 0x7a, 0x89, 0x5c, 0x0e,
	0x01, 0x24, 0x47, 0x84, 0x0b, 0x42, 0x5c, 0x28, 0xfd, 0x3c, 0x81, 0x36, 0x12, 0x07, 0x0e, 0x44,
	0xce, 0x7a, 0x4a, 0x57, 0x75, 0x76, 0xcd, 0xee, 0x26, 0x21, 0xbf, 0x86, 0x5f, 0xc1, 0x2f, 0xe1,
	0xe7, 0x70, 0x41, 0x5e, 0xc7, 0x49, 0xfa, 0x91, 0x12, 0x09, 0x71, 0xf3, 0xce, 0xe3, 0x67, 0x9e,
	0x99, 0x67, 0x66, 0x17, 0x0e, 0x66, 0x52, 0xdd, 0x5c, 0xa5, 0x72, 0xd6, 0x63, 0x52, 0x08, 0x64,
	0x86, 0x4b, 0x11, 0x65, 0x4a, 0x1a, 0x49, 0x1a, 0x25, 0x14, 0xfe, 0x74, 0x60, 0x7f, 0x60, 0x62,
	0x65, 0xce, 0x52, 0x39, 0xa3, 0xf8, 0x75, 0x82, 0xda, 0x90, 0x43, 0xf0, 0x35, 0xbb, 0xc6, 0x71,
	0x3c, 0xe4, 0x49, 0xe0, 0x74, 0x9c, 0xee, 0x2e, 0x6d, 0x14, 0x81, 0xcb, 0x24, 0x07, 0x13, 0x39,
	0x8e, 0xb9, 0xc8, 0x41, 0xb7, 0xe3, 0x74, 0x3d, 0xda, 0x28, 0x02, 0x97, 0x09, 0x39, 0x07, 0x7f,
	0x1a, 0x2b, 0x1e, 0x8f, 0x52, 0xd4, 0x81, 0xd7, 0xf1, 0xba, 0xcd, 0xfe, 0xd3, 0xa8, 0x14, 0x8b,
	0xee, 0x0a, 0x45, 0x1f, 0xcb, 0x7f, 0x4f, 0x85, 0x51, 0x73, 0xba, 0xe2, 0xb6, 0xdf, 0xc0, 0xde,
	0x6d, 0x90, 0xec, 0x83, 0x77, 0x83, 0x73, 0x5b, 0x8e, 0x4f, 0xf3, 0x4f, 0xf2, 0x3f, 0x54, 0xa7,
	0x71, 0x3a, 0x41, 0x5b, 0x85, 0x4f, 0x8b, 0xc3, 0x6b, 0xf7, 0x95, 0x13, 0x1e, 0x41, 0x6b, 0x4d,
	0x4b, 0x67, 0x52, 0x68, 0x24, 0x7b, 0xe0, 0x2e, 0xda, 0xf1, 0xa9, 0xcb, 0x93, 0xf0, 0xbb, 0x0b,
	0xc1, 0x09, 0xd7, 0x46, 0xf1, 0xd1, 0xc4, 0xe0, 0x5b, 0x63, 0x70, 0x9c, 0x99, 0x8d, 0x16, 0x54,
	0xb7, 0xb5, 0xa0, 0x03, 0xcd, 0x04, 0xb5, 0xe1, 0x22, 0xce, 0x0d, 0x0f, 0x3c, 0xab, 0xb7, 0x1e,
	0x22, 0x01, 0xd4, 0x13, 0xae, 0xb3, 0x34, 0x9e, 0x07, 0x3b, 0x16, 0x2d, 0x8f, 0xe4, 0xfd, 0xba,
	0x7d, 0x55, 0x6b, 0xdf, 0x8b, 0x95, 0x7d, 0x9b, 0x8a, 0xfd, 0x67, 0x36, 0xfe, 0xf2, 0xe0, 0xe0,
	0x01, 0xd1, 0x85, 0x9f, 0x27, 0x50, 0x63, 0xb1, 0x60, 0x98, 0xda, 0x64, 0xcd, 0xfe, 0xb3, 0x47,
	0x2b, 0x2d, 0x48, 0xd1, 0x3b, 0xcb, 0xb8, 0xa8, 0xd0, 0x05, 0x97, 0x9c, 0x43, 0x9d, 0x49, 0x71,
	0xc5, 0xd5, 0xd8, 0xea, 0x37, 0xfb, 0xcf, 0xb7, 0x4a, 0x53, 0x50, 0x2e, 0x2a, 0xb4, 0x64, 0x93,
	0x0f, 0xf7, 0x57, 0xaf, 0xbf, 0x4d, 0xaa, 0xcd, 0xe6, 0x09, 0xa8, 0x15, 0xe5, 0x2e, 0x66, 0xca,
	0x14, 0xcf, 0xec, 0x4c, 0x9d, 0xe5, 0x4c, 0xcb, 0x10, 0x89, 0xe0, 0x3f, 0x81, 0xdf, 0xcc, 0x30,
	0x59, 0xea, 0x0c, 0x35, 0x32, 0xdb, 0xd2, 0x2e, 0x6d, 0xe5, 0xd0, 0xaa, 0x82, 0x01, 0x32, 0x42,
	0x60, 0x47, 0x1b, 0x99, 0xd9, 0xf5, 0x68, 0x50, 0xfb, 0xdd, 0x3e, 0x85, 0xfa, 0xa2, 0xaf, 0xbb,
	0x4b, 0xe4, 0x3c, 0xba, 0x44, 0xee, 0xad, 0x25, 0xfa, 0xbb, 0x99, 0x1f, 0x37, 0xa0, 0xa6, 0x50,
	0x4f, 0x52, 0xd3, 0xff, 0xe1, 0x40, 0x33, 0xbf, 0x40, 0x03, 0x54, 0x53, 0xce, 0x90, 0x7c, 0x86,
	0xd6, 0x3d, 0x17, 0x49, 0xf8, 0xe7, 0xf5, 0x6c, 0x1f, 0x6d, 0x31, 0x86, 0xb0, 0x42, 0xce, 0xc0,
	0x5f, 0x5e, 0x5a, 0xd2, 0xde, 0xfc, 0x6a, 0xb4, 0x0f, 0x1f, 0xc4, 0xca, 0x3c, 0xc7, 0x4f, 0x3e,
	0x85, 0x5f, 0xb8, 0xb9, 0x9e, 0x8c, 0x22, 0x26, 0xc7, 0xbd, 0x19, 0x8e, 0xb8, 0xc1, 0xb4, 0x67,
	0x1f, 0x3e, 0xdd, 0x2b, 0x99, 0xa3, 0x9a, 0x0d, 0xbc, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xa3,
	0xf1, 0x82, 0xe2, 0x26, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FlowServiceClient is the client API for FlowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FlowServiceClient interface {
	DistributeAttempt(ctx context.Context, in *DistributeAttemptRequest, opts ...grpc.CallOption) (*DistributeAttemptResponse, error)
	StartFlow(ctx context.Context, in *StartFlowRequest, opts ...grpc.CallOption) (*StartFlowResponse, error)
}

type flowServiceClient struct {
	cc *grpc.ClientConn
}

func NewFlowServiceClient(cc *grpc.ClientConn) FlowServiceClient {
	return &flowServiceClient{cc}
}

func (c *flowServiceClient) DistributeAttempt(ctx context.Context, in *DistributeAttemptRequest, opts ...grpc.CallOption) (*DistributeAttemptResponse, error) {
	out := new(DistributeAttemptResponse)
	err := c.cc.Invoke(ctx, "/workflow.FlowService/DistributeAttempt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowServiceClient) StartFlow(ctx context.Context, in *StartFlowRequest, opts ...grpc.CallOption) (*StartFlowResponse, error) {
	out := new(StartFlowResponse)
	err := c.cc.Invoke(ctx, "/workflow.FlowService/StartFlow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlowServiceServer is the server API for FlowService service.
type FlowServiceServer interface {
	DistributeAttempt(context.Context, *DistributeAttemptRequest) (*DistributeAttemptResponse, error)
	StartFlow(context.Context, *StartFlowRequest) (*StartFlowResponse, error)
}

// UnimplementedFlowServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFlowServiceServer struct {
}

func (*UnimplementedFlowServiceServer) DistributeAttempt(ctx context.Context, req *DistributeAttemptRequest) (*DistributeAttemptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DistributeAttempt not implemented")
}
func (*UnimplementedFlowServiceServer) StartFlow(ctx context.Context, req *StartFlowRequest) (*StartFlowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartFlow not implemented")
}

func RegisterFlowServiceServer(s *grpc.Server, srv FlowServiceServer) {
	s.RegisterService(&_FlowService_serviceDesc, srv)
}

func _FlowService_DistributeAttempt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DistributeAttemptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowServiceServer).DistributeAttempt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.FlowService/DistributeAttempt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowServiceServer).DistributeAttempt(ctx, req.(*DistributeAttemptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowService_StartFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowServiceServer).StartFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.FlowService/StartFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowServiceServer).StartFlow(ctx, req.(*StartFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FlowService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "workflow.FlowService",
	HandlerType: (*FlowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DistributeAttempt",
			Handler:    _FlowService_DistributeAttempt_Handler,
		},
		{
			MethodName: "StartFlow",
			Handler:    _FlowService_StartFlow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workflow/connection.proto",
}
