// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: workflow/processing.proto

package workflow

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

type CancelProcessingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CancelProcessingRequest) Reset() {
	*x = CancelProcessingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_processing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelProcessingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelProcessingRequest) ProtoMessage() {}

func (x *CancelProcessingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_processing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelProcessingRequest.ProtoReflect.Descriptor instead.
func (*CancelProcessingRequest) Descriptor() ([]byte, []int) {
	return file_workflow_processing_proto_rawDescGZIP(), []int{0}
}

func (x *CancelProcessingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CancelProcessingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelProcessingResponse) Reset() {
	*x = CancelProcessingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_processing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelProcessingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelProcessingResponse) ProtoMessage() {}

func (x *CancelProcessingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_processing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelProcessingResponse.ProtoReflect.Descriptor instead.
func (*CancelProcessingResponse) Descriptor() ([]byte, []int) {
	return file_workflow_processing_proto_rawDescGZIP(), []int{1}
}

type StartProcessingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaId  uint32            `protobuf:"varint,1,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"`
	DomainId  int64             `protobuf:"varint,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	Variables map[string]string `protobuf:"bytes,3,rep,name=variables,proto3" json:"variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *StartProcessingRequest) Reset() {
	*x = StartProcessingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_processing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartProcessingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartProcessingRequest) ProtoMessage() {}

func (x *StartProcessingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_processing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartProcessingRequest.ProtoReflect.Descriptor instead.
func (*StartProcessingRequest) Descriptor() ([]byte, []int) {
	return file_workflow_processing_proto_rawDescGZIP(), []int{2}
}

func (x *StartProcessingRequest) GetSchemaId() uint32 {
	if x != nil {
		return x.SchemaId
	}
	return 0
}

func (x *StartProcessingRequest) GetDomainId() int64 {
	if x != nil {
		return x.DomainId
	}
	return 0
}

func (x *StartProcessingRequest) GetVariables() map[string]string {
	if x != nil {
		return x.Variables
	}
	return nil
}

type FormActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DomainId  int64             `protobuf:"varint,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	Action    string            `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	Variables map[string]string `protobuf:"bytes,4,rep,name=variables,proto3" json:"variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FormActionRequest) Reset() {
	*x = FormActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_processing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FormActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FormActionRequest) ProtoMessage() {}

func (x *FormActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_processing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FormActionRequest.ProtoReflect.Descriptor instead.
func (*FormActionRequest) Descriptor() ([]byte, []int) {
	return file_workflow_processing_proto_rawDescGZIP(), []int{3}
}

func (x *FormActionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FormActionRequest) GetDomainId() int64 {
	if x != nil {
		return x.DomainId
	}
	return 0
}

func (x *FormActionRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *FormActionRequest) GetVariables() map[string]string {
	if x != nil {
		return x.Variables
	}
	return nil
}

type ErrorForm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ErrorForm) Reset() {
	*x = ErrorForm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_processing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorForm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorForm) ProtoMessage() {}

func (x *ErrorForm) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_processing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorForm.ProtoReflect.Descriptor instead.
func (*ErrorForm) Descriptor() ([]byte, []int) {
	return file_workflow_processing_proto_rawDescGZIP(), []int{4}
}

func (x *ErrorForm) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ErrorForm) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Form struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AppId   string     `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Form    []byte     `protobuf:"bytes,3,opt,name=form,proto3" json:"form,omitempty"`
	Timeout uint64     `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Stop    bool       `protobuf:"varint,5,opt,name=stop,proto3" json:"stop,omitempty"`
	Error   *ErrorForm `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Form) Reset() {
	*x = Form{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_processing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Form) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Form) ProtoMessage() {}

func (x *Form) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_processing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Form.ProtoReflect.Descriptor instead.
func (*Form) Descriptor() ([]byte, []int) {
	return file_workflow_processing_proto_rawDescGZIP(), []int{5}
}

func (x *Form) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Form) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *Form) GetForm() []byte {
	if x != nil {
		return x.Form
	}
	return nil
}

func (x *Form) GetTimeout() uint64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *Form) GetStop() bool {
	if x != nil {
		return x.Stop
	}
	return false
}

func (x *Form) GetError() *ErrorForm {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_workflow_processing_proto protoreflect.FileDescriptor

var file_workflow_processing_proto_rawDesc = []byte{
	0x0a, 0x19, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0x29, 0x0a, 0x17, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x1a, 0x0a, 0x18, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xdf, 0x01, 0x0a,
	0x16, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49,
	0x64, 0x12, 0x4d, 0x0a, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x1a, 0x3c, 0x0a, 0x0e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe0,
	0x01, 0x0a, 0x11, 0x46, 0x6f, 0x72, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x48, 0x0a, 0x09, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x1a, 0x3c, 0x0a, 0x0e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x35, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x9a, 0x01, 0x0a, 0x04, 0x46, 0x6f, 0x72,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6f, 0x72, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x12, 0x29, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xf8, 0x01, 0x0a, 0x15, 0x46, 0x6c, 0x6f, 0x77, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x45, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x12, 0x20, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x46, 0x6f, 0x72, 0x6d, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0a, 0x46, 0x6f, 0x72, 0x6d, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x46, 0x6f, 0x72, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0e, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x6f, 0x72,
	0x6d, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x10, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77,
	0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_workflow_processing_proto_rawDescOnce sync.Once
	file_workflow_processing_proto_rawDescData = file_workflow_processing_proto_rawDesc
)

func file_workflow_processing_proto_rawDescGZIP() []byte {
	file_workflow_processing_proto_rawDescOnce.Do(func() {
		file_workflow_processing_proto_rawDescData = protoimpl.X.CompressGZIP(file_workflow_processing_proto_rawDescData)
	})
	return file_workflow_processing_proto_rawDescData
}

var file_workflow_processing_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_workflow_processing_proto_goTypes = []interface{}{
	(*CancelProcessingRequest)(nil),  // 0: workflow.CancelProcessingRequest
	(*CancelProcessingResponse)(nil), // 1: workflow.CancelProcessingResponse
	(*StartProcessingRequest)(nil),   // 2: workflow.StartProcessingRequest
	(*FormActionRequest)(nil),        // 3: workflow.FormActionRequest
	(*ErrorForm)(nil),                // 4: workflow.ErrorForm
	(*Form)(nil),                     // 5: workflow.Form
	nil,                              // 6: workflow.StartProcessingRequest.VariablesEntry
	nil,                              // 7: workflow.FormActionRequest.VariablesEntry
}
var file_workflow_processing_proto_depIdxs = []int32{
	6, // 0: workflow.StartProcessingRequest.variables:type_name -> workflow.StartProcessingRequest.VariablesEntry
	7, // 1: workflow.FormActionRequest.variables:type_name -> workflow.FormActionRequest.VariablesEntry
	4, // 2: workflow.Form.error:type_name -> workflow.ErrorForm
	2, // 3: workflow.FlowProcessingService.StartProcessing:input_type -> workflow.StartProcessingRequest
	3, // 4: workflow.FlowProcessingService.FormAction:input_type -> workflow.FormActionRequest
	0, // 5: workflow.FlowProcessingService.CancelProcessing:input_type -> workflow.CancelProcessingRequest
	5, // 6: workflow.FlowProcessingService.StartProcessing:output_type -> workflow.Form
	5, // 7: workflow.FlowProcessingService.FormAction:output_type -> workflow.Form
	1, // 8: workflow.FlowProcessingService.CancelProcessing:output_type -> workflow.CancelProcessingResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_workflow_processing_proto_init() }
func file_workflow_processing_proto_init() {
	if File_workflow_processing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_workflow_processing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelProcessingRequest); i {
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
		file_workflow_processing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelProcessingResponse); i {
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
		file_workflow_processing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartProcessingRequest); i {
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
		file_workflow_processing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FormActionRequest); i {
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
		file_workflow_processing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorForm); i {
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
		file_workflow_processing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Form); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_workflow_processing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_workflow_processing_proto_goTypes,
		DependencyIndexes: file_workflow_processing_proto_depIdxs,
		MessageInfos:      file_workflow_processing_proto_msgTypes,
	}.Build()
	File_workflow_processing_proto = out.File
	file_workflow_processing_proto_rawDesc = nil
	file_workflow_processing_proto_goTypes = nil
	file_workflow_processing_proto_depIdxs = nil
}
