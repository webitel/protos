// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.0
// source: storage/file_transcript.proto

package storage

import (
	engine "github.com/webitel/protos/engine"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type DeleteFileTranscriptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   []int64  `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
	Uuid []string `protobuf:"bytes,2,rep,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *DeleteFileTranscriptRequest) Reset() {
	*x = DeleteFileTranscriptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_file_transcript_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFileTranscriptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileTranscriptRequest) ProtoMessage() {}

func (x *DeleteFileTranscriptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_file_transcript_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileTranscriptRequest.ProtoReflect.Descriptor instead.
func (*DeleteFileTranscriptRequest) Descriptor() ([]byte, []int) {
	return file_storage_file_transcript_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteFileTranscriptRequest) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *DeleteFileTranscriptRequest) GetUuid() []string {
	if x != nil {
		return x.Uuid
	}
	return nil
}

type DeleteFileTranscriptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []int64 `protobuf:"varint,1,rep,packed,name=items,proto3" json:"items,omitempty"`
}

func (x *DeleteFileTranscriptResponse) Reset() {
	*x = DeleteFileTranscriptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_file_transcript_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFileTranscriptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileTranscriptResponse) ProtoMessage() {}

func (x *DeleteFileTranscriptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_file_transcript_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileTranscriptResponse.ProtoReflect.Descriptor instead.
func (*DeleteFileTranscriptResponse) Descriptor() ([]byte, []int) {
	return file_storage_file_transcript_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteFileTranscriptResponse) GetItems() []int64 {
	if x != nil {
		return x.Items
	}
	return nil
}

type TranscriptPhrase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartSec float32 `protobuf:"fixed32,1,opt,name=start_sec,json=startSec,proto3" json:"start_sec,omitempty"`
	EndSec   float32 `protobuf:"fixed32,2,opt,name=end_sec,json=endSec,proto3" json:"end_sec,omitempty"`
	Channel  uint32  `protobuf:"varint,3,opt,name=channel,proto3" json:"channel,omitempty"`
	Phrase   string  `protobuf:"bytes,4,opt,name=phrase,proto3" json:"phrase,omitempty"`
}

func (x *TranscriptPhrase) Reset() {
	*x = TranscriptPhrase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_file_transcript_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranscriptPhrase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranscriptPhrase) ProtoMessage() {}

func (x *TranscriptPhrase) ProtoReflect() protoreflect.Message {
	mi := &file_storage_file_transcript_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranscriptPhrase.ProtoReflect.Descriptor instead.
func (*TranscriptPhrase) Descriptor() ([]byte, []int) {
	return file_storage_file_transcript_proto_rawDescGZIP(), []int{2}
}

func (x *TranscriptPhrase) GetStartSec() float32 {
	if x != nil {
		return x.StartSec
	}
	return 0
}

func (x *TranscriptPhrase) GetEndSec() float32 {
	if x != nil {
		return x.EndSec
	}
	return 0
}

func (x *TranscriptPhrase) GetChannel() uint32 {
	if x != nil {
		return x.Channel
	}
	return 0
}

func (x *TranscriptPhrase) GetPhrase() string {
	if x != nil {
		return x.Phrase
	}
	return ""
}

type GetFileTranscriptPhrasesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size int32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Id   int64 `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetFileTranscriptPhrasesRequest) Reset() {
	*x = GetFileTranscriptPhrasesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_file_transcript_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileTranscriptPhrasesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileTranscriptPhrasesRequest) ProtoMessage() {}

func (x *GetFileTranscriptPhrasesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_file_transcript_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileTranscriptPhrasesRequest.ProtoReflect.Descriptor instead.
func (*GetFileTranscriptPhrasesRequest) Descriptor() ([]byte, []int) {
	return file_storage_file_transcript_proto_rawDescGZIP(), []int{3}
}

func (x *GetFileTranscriptPhrasesRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetFileTranscriptPhrasesRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *GetFileTranscriptPhrasesRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListPhrases struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Next  bool                `protobuf:"varint,1,opt,name=next,proto3" json:"next,omitempty"`
	Items []*TranscriptPhrase `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListPhrases) Reset() {
	*x = ListPhrases{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_file_transcript_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPhrases) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPhrases) ProtoMessage() {}

func (x *ListPhrases) ProtoReflect() protoreflect.Message {
	mi := &file_storage_file_transcript_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPhrases.ProtoReflect.Descriptor instead.
func (*ListPhrases) Descriptor() ([]byte, []int) {
	return file_storage_file_transcript_proto_rawDescGZIP(), []int{4}
}

func (x *ListPhrases) GetNext() bool {
	if x != nil {
		return x.Next
	}
	return false
}

func (x *ListPhrases) GetItems() []*TranscriptPhrase {
	if x != nil {
		return x.Items
	}
	return nil
}

type StartFileTranscriptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId  []int64        `protobuf:"varint,1,rep,packed,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	Locale  string         `protobuf:"bytes,2,opt,name=locale,proto3" json:"locale,omitempty"`
	Profile *engine.Lookup `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile,omitempty"`
	Uuid    []string       `protobuf:"bytes,4,rep,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *StartFileTranscriptRequest) Reset() {
	*x = StartFileTranscriptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_file_transcript_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartFileTranscriptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartFileTranscriptRequest) ProtoMessage() {}

func (x *StartFileTranscriptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_file_transcript_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartFileTranscriptRequest.ProtoReflect.Descriptor instead.
func (*StartFileTranscriptRequest) Descriptor() ([]byte, []int) {
	return file_storage_file_transcript_proto_rawDescGZIP(), []int{5}
}

func (x *StartFileTranscriptRequest) GetFileId() []int64 {
	if x != nil {
		return x.FileId
	}
	return nil
}

func (x *StartFileTranscriptRequest) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *StartFileTranscriptRequest) GetProfile() *engine.Lookup {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *StartFileTranscriptRequest) GetUuid() []string {
	if x != nil {
		return x.Uuid
	}
	return nil
}

type StartFileTranscriptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*StartFileTranscriptResponse_TranscriptJob `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *StartFileTranscriptResponse) Reset() {
	*x = StartFileTranscriptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_file_transcript_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartFileTranscriptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartFileTranscriptResponse) ProtoMessage() {}

func (x *StartFileTranscriptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_file_transcript_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartFileTranscriptResponse.ProtoReflect.Descriptor instead.
func (*StartFileTranscriptResponse) Descriptor() ([]byte, []int) {
	return file_storage_file_transcript_proto_rawDescGZIP(), []int{6}
}

func (x *StartFileTranscriptResponse) GetItems() []*StartFileTranscriptResponse_TranscriptJob {
	if x != nil {
		return x.Items
	}
	return nil
}

type StartFileTranscriptResponse_TranscriptJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FileId    int64  `protobuf:"varint,2,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	CreatedAt int64  `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Action    string `protobuf:"bytes,4,opt,name=action,proto3" json:"action,omitempty"`
	State     string `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *StartFileTranscriptResponse_TranscriptJob) Reset() {
	*x = StartFileTranscriptResponse_TranscriptJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_file_transcript_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartFileTranscriptResponse_TranscriptJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartFileTranscriptResponse_TranscriptJob) ProtoMessage() {}

func (x *StartFileTranscriptResponse_TranscriptJob) ProtoReflect() protoreflect.Message {
	mi := &file_storage_file_transcript_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartFileTranscriptResponse_TranscriptJob.ProtoReflect.Descriptor instead.
func (*StartFileTranscriptResponse_TranscriptJob) Descriptor() ([]byte, []int) {
	return file_storage_file_transcript_proto_rawDescGZIP(), []int{6, 0}
}

func (x *StartFileTranscriptResponse_TranscriptJob) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StartFileTranscriptResponse_TranscriptJob) GetFileId() int64 {
	if x != nil {
		return x.FileId
	}
	return 0
}

func (x *StartFileTranscriptResponse_TranscriptJob) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *StartFileTranscriptResponse_TranscriptJob) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *StartFileTranscriptResponse_TranscriptJob) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

var File_storage_file_transcript_proto protoreflect.FileDescriptor

var file_storage_file_transcript_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x12, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x1b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x34, 0x0a,
	0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x7a, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x50, 0x68, 0x72, 0x61, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x73, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x53, 0x65, 0x63, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x63, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x68, 0x72, 0x61, 0x73,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x22,
	0x59, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x50, 0x68, 0x72, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x52, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x68, 0x72, 0x61, 0x73, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x12, 0x2f, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x50, 0x68, 0x72, 0x61, 0x73, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x8b,
	0x01, 0x0a, 0x1a, 0x53, 0x74, 0x61, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x28,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0xef, 0x01, 0x0a,
	0x1b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x4a, 0x6f, 0x62, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x85, 0x01, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x4a, 0x6f, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x32, 0xb7,
	0x03, 0x0a, 0x15, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x12, 0x23, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x3a, 0x01,
	0x2a, 0x12, 0x89, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x50, 0x68, 0x72, 0x61, 0x73, 0x65, 0x73, 0x12, 0x28,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x50, 0x68, 0x72, 0x61, 0x73, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x68, 0x72, 0x61, 0x73, 0x65, 0x73, 0x22, 0x2d,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12, 0x25, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x73, 0x12, 0x88, 0x01,
	0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x24, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x2a, 0x18, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_file_transcript_proto_rawDescOnce sync.Once
	file_storage_file_transcript_proto_rawDescData = file_storage_file_transcript_proto_rawDesc
)

func file_storage_file_transcript_proto_rawDescGZIP() []byte {
	file_storage_file_transcript_proto_rawDescOnce.Do(func() {
		file_storage_file_transcript_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_file_transcript_proto_rawDescData)
	})
	return file_storage_file_transcript_proto_rawDescData
}

var file_storage_file_transcript_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_storage_file_transcript_proto_goTypes = []interface{}{
	(*DeleteFileTranscriptRequest)(nil),               // 0: storage.DeleteFileTranscriptRequest
	(*DeleteFileTranscriptResponse)(nil),              // 1: storage.DeleteFileTranscriptResponse
	(*TranscriptPhrase)(nil),                          // 2: storage.TranscriptPhrase
	(*GetFileTranscriptPhrasesRequest)(nil),           // 3: storage.GetFileTranscriptPhrasesRequest
	(*ListPhrases)(nil),                               // 4: storage.ListPhrases
	(*StartFileTranscriptRequest)(nil),                // 5: storage.StartFileTranscriptRequest
	(*StartFileTranscriptResponse)(nil),               // 6: storage.StartFileTranscriptResponse
	(*StartFileTranscriptResponse_TranscriptJob)(nil), // 7: storage.StartFileTranscriptResponse.TranscriptJob
	(*engine.Lookup)(nil),                             // 8: engine.Lookup
}
var file_storage_file_transcript_proto_depIdxs = []int32{
	2, // 0: storage.ListPhrases.items:type_name -> storage.TranscriptPhrase
	8, // 1: storage.StartFileTranscriptRequest.profile:type_name -> engine.Lookup
	7, // 2: storage.StartFileTranscriptResponse.items:type_name -> storage.StartFileTranscriptResponse.TranscriptJob
	5, // 3: storage.FileTranscriptService.CreateFileTranscript:input_type -> storage.StartFileTranscriptRequest
	3, // 4: storage.FileTranscriptService.GetFileTranscriptPhrases:input_type -> storage.GetFileTranscriptPhrasesRequest
	0, // 5: storage.FileTranscriptService.DeleteFileTranscript:input_type -> storage.DeleteFileTranscriptRequest
	6, // 6: storage.FileTranscriptService.CreateFileTranscript:output_type -> storage.StartFileTranscriptResponse
	4, // 7: storage.FileTranscriptService.GetFileTranscriptPhrases:output_type -> storage.ListPhrases
	1, // 8: storage.FileTranscriptService.DeleteFileTranscript:output_type -> storage.DeleteFileTranscriptResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_storage_file_transcript_proto_init() }
func file_storage_file_transcript_proto_init() {
	if File_storage_file_transcript_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storage_file_transcript_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFileTranscriptRequest); i {
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
		file_storage_file_transcript_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFileTranscriptResponse); i {
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
		file_storage_file_transcript_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranscriptPhrase); i {
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
		file_storage_file_transcript_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileTranscriptPhrasesRequest); i {
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
		file_storage_file_transcript_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPhrases); i {
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
		file_storage_file_transcript_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartFileTranscriptRequest); i {
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
		file_storage_file_transcript_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartFileTranscriptResponse); i {
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
		file_storage_file_transcript_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartFileTranscriptResponse_TranscriptJob); i {
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
			RawDescriptor: file_storage_file_transcript_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_storage_file_transcript_proto_goTypes,
		DependencyIndexes: file_storage_file_transcript_proto_depIdxs,
		MessageInfos:      file_storage_file_transcript_proto_msgTypes,
	}.Build()
	File_storage_file_transcript_proto = out.File
	file_storage_file_transcript_proto_rawDesc = nil
	file_storage_file_transcript_proto_goTypes = nil
	file_storage_file_transcript_proto_depIdxs = nil
}
