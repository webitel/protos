// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.6
// source: gateway/contacts/imclients.proto

package contacts

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/api/visibility"
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

// A contact's [I]nstant[M]essaging client.
type IMClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique ID of the association. Never changes.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Version of the latest update. Numeric sequence.
	Ver int32 `protobuf:"varint,2,opt,name=ver,proto3" json:"ver,omitempty"`
	// Unique ID of the latest version of the update.
	// This ID changes after any update to the underlying value(s).
	Etag string `protobuf:"bytes,3,opt,name=etag,proto3" json:"etag,omitempty"`
	// The user who created this Field.
	CreatedAt int64 `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Timestamp(milli) of the Field creation.
	CreatedBy *Lookup `protobuf:"bytes,6,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	// Timestamp(milli) of the last Field update.
	// Take part in Etag generation.
	UpdatedAt int64 `protobuf:"varint,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// The user who performed last Update.
	UpdatedBy *Lookup `protobuf:"bytes,8,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	// External user which contacted to us.
	// Id will be from external service.
	// Name will be from external service.
	User *Lookup `protobuf:"bytes,11,opt,name=user,proto3" json:"user,omitempty"`
	// App (Text-Gateway) used to connect the IM client.
	// Id will be internal id of gateway.
	// Name will be name of the gateway.
	App *Lookup `protobuf:"bytes,12,opt,name=app,proto3" json:"app,omitempty"`
	// Protocol used to connect the IM client.
	Protocol string `protobuf:"bytes,13,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// [Via] App(-specific) peer(-id) to connect[from] the IM client.
	Via string `protobuf:"bytes,14,opt,name=via,proto3" json:"via,omitempty"`
}

func (x *IMClient) Reset() {
	*x = IMClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_contacts_imclients_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IMClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IMClient) ProtoMessage() {}

func (x *IMClient) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_contacts_imclients_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IMClient.ProtoReflect.Descriptor instead.
func (*IMClient) Descriptor() ([]byte, []int) {
	return file_gateway_contacts_imclients_proto_rawDescGZIP(), []int{0}
}

func (x *IMClient) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *IMClient) GetVer() int32 {
	if x != nil {
		return x.Ver
	}
	return 0
}

func (x *IMClient) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *IMClient) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *IMClient) GetCreatedBy() *Lookup {
	if x != nil {
		return x.CreatedBy
	}
	return nil
}

func (x *IMClient) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *IMClient) GetUpdatedBy() *Lookup {
	if x != nil {
		return x.UpdatedBy
	}
	return nil
}

func (x *IMClient) GetUser() *Lookup {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *IMClient) GetApp() *Lookup {
	if x != nil {
		return x.App
	}
	return nil
}

func (x *IMClient) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *IMClient) GetVia() string {
	if x != nil {
		return x.Via
	}
	return ""
}

// Input of the contact IM client.
type InputIMClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol string `protobuf:"bytes,9,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// Id of Agent created this IM client.
	CreatedBy string `protobuf:"bytes,10,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	// External user id
	ExternalUser string `protobuf:"bytes,11,opt,name=external_user,json=externalUser,proto3" json:"external_user,omitempty"`
	// App (Text-Gateway) used to connect the IM client.
	GatewayId string `protobuf:"bytes,12,opt,name=gateway_id,json=gatewayId,proto3" json:"gateway_id,omitempty"`
	// [Via] App(-specific) peer(-id) to connect[from] the IM client.
	Via string `protobuf:"bytes,13,opt,name=via,proto3" json:"via,omitempty"`
}

func (x *InputIMClient) Reset() {
	*x = InputIMClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_contacts_imclients_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputIMClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputIMClient) ProtoMessage() {}

func (x *InputIMClient) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_contacts_imclients_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputIMClient.ProtoReflect.Descriptor instead.
func (*InputIMClient) Descriptor() ([]byte, []int) {
	return file_gateway_contacts_imclients_proto_rawDescGZIP(), []int{1}
}

func (x *InputIMClient) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *InputIMClient) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *InputIMClient) GetExternalUser() string {
	if x != nil {
		return x.ExternalUser
	}
	return ""
}

func (x *InputIMClient) GetGatewayId() string {
	if x != nil {
		return x.GatewayId
	}
	return ""
}

func (x *InputIMClient) GetVia() string {
	if x != nil {
		return x.Via
	}
	return ""
}

type IMClientList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// IMClient dataset page.
	Data []*IMClient `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	// The page number of the partial result.
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	// Indicates that this is a partial result.
	// More data available upon query: ?size=${data.len}&page=${page++}
	Next bool `protobuf:"varint,3,opt,name=next,proto3" json:"next,omitempty"`
}

func (x *IMClientList) Reset() {
	*x = IMClientList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_contacts_imclients_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IMClientList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IMClientList) ProtoMessage() {}

func (x *IMClientList) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_contacts_imclients_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IMClientList.ProtoReflect.Descriptor instead.
func (*IMClientList) Descriptor() ([]byte, []int) {
	return file_gateway_contacts_imclients_proto_rawDescGZIP(), []int{2}
}

func (x *IMClientList) GetData() []*IMClient {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *IMClientList) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *IMClientList) GetNext() bool {
	if x != nil {
		return x.Next
	}
	return false
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_contacts_imclients_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_contacts_imclients_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_gateway_contacts_imclients_proto_rawDescGZIP(), []int{3}
}

type ListIMClientsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Page number of result dataset records. offset = (page*size)
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	// Size count of records on result page. limit = (size++)
	Size int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// Search term: IMClient.
	// `?` - matches any one character
	// `*` - matches 0 or more characters
	// Search fields: {user,app}
	Q string `protobuf:"bytes,6,opt,name=q,proto3" json:"q,omitempty"`
	// Sort the result according to fields.
	Sort []string `protobuf:"bytes,3,rep,name=sort,proto3" json:"sort,omitempty"`
	// Fields to be retrieved into result.
	Fields []string `protobuf:"bytes,4,rep,name=fields,proto3" json:"fields,omitempty"`
	// Link contact ID.
	ContactId string `protobuf:"bytes,5,opt,name=contact_id,json=contactId,proto3" json:"contact_id,omitempty"`
	// Link(s) with unique ID only.
	Id []string `protobuf:"bytes,7,rep,name=id,proto3" json:"id,omitempty"`
}

func (x *ListIMClientsRequest) Reset() {
	*x = ListIMClientsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_contacts_imclients_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIMClientsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIMClientsRequest) ProtoMessage() {}

func (x *ListIMClientsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_contacts_imclients_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIMClientsRequest.ProtoReflect.Descriptor instead.
func (*ListIMClientsRequest) Descriptor() ([]byte, []int) {
	return file_gateway_contacts_imclients_proto_rawDescGZIP(), []int{4}
}

func (x *ListIMClientsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListIMClientsRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListIMClientsRequest) GetQ() string {
	if x != nil {
		return x.Q
	}
	return ""
}

func (x *ListIMClientsRequest) GetSort() []string {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *ListIMClientsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *ListIMClientsRequest) GetContactId() string {
	if x != nil {
		return x.ContactId
	}
	return ""
}

func (x *ListIMClientsRequest) GetId() []string {
	if x != nil {
		return x.Id
	}
	return nil
}

type CreateIMClientsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Link contact ID.
	ContactId string `protobuf:"bytes,1,opt,name=contact_id,json=contactId,proto3" json:"contact_id,omitempty"`
	DomainId  int64  `protobuf:"varint,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	// Fixed set of IM client(s) to be linked with the contact.
	// IM client(s) that conflicts(user.id) with already linked will be updated.
	Input []*InputIMClient `protobuf:"bytes,3,rep,name=input,proto3" json:"input,omitempty"`
}

func (x *CreateIMClientsRequest) Reset() {
	*x = CreateIMClientsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_contacts_imclients_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIMClientsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIMClientsRequest) ProtoMessage() {}

func (x *CreateIMClientsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_contacts_imclients_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIMClientsRequest.ProtoReflect.Descriptor instead.
func (*CreateIMClientsRequest) Descriptor() ([]byte, []int) {
	return file_gateway_contacts_imclients_proto_rawDescGZIP(), []int{5}
}

func (x *CreateIMClientsRequest) GetContactId() string {
	if x != nil {
		return x.ContactId
	}
	return ""
}

func (x *CreateIMClientsRequest) GetDomainId() int64 {
	if x != nil {
		return x.DomainId
	}
	return 0
}

func (x *CreateIMClientsRequest) GetInput() []*InputIMClient {
	if x != nil {
		return x.Input
	}
	return nil
}

var File_gateway_contacts_imclients_proto protoreflect.FileDescriptor

var file_gateway_contacts_imclients_proto_rawDesc = []byte{
	0x0a, 0x20, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x73, 0x2f, 0x69, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x73, 0x1a, 0x1d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5,
	0x03, 0x0a, 0x08, 0x49, 0x4d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x76,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x76, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x74, 0x61,
	0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x37, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x37, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x77,
	0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x2e,
	0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x12, 0x69, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x73, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x42, 0x3b, 0x92, 0x41, 0x38, 0x4a, 0x36,
	0x7b, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x22, 0x35, 0x39, 0x36, 0x34, 0x31, 0x37, 0x33, 0x34, 0x33,
	0x22, 0x2c, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72,
	0x61, 0x6d, 0x22, 0x2c, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x22, 0x4a, 0x6f, 0x68, 0x6e,
	0x20, 0x44, 0x6f, 0x65, 0x22, 0x7d, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x03,
	0x61, 0x70, 0x70, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x77, 0x65, 0x62, 0x69,
	0x74, 0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x4c, 0x6f, 0x6f,
	0x6b, 0x75, 0x70, 0x52, 0x03, 0x61, 0x70, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x69, 0x61, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x76, 0x69, 0x61, 0x22, 0xa0, 0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x49, 0x4d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x62, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x69, 0x61, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x69, 0x61, 0x22, 0x66, 0x0a, 0x0c, 0x49, 0x4d, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65,
	0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x49, 0x4d, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6e, 0x65, 0x78,
	0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x4d, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x71, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8b, 0x01, 0x0a,
	0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x4d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x49, 0x4d, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x32, 0xef, 0x01, 0x0a, 0x09, 0x49,
	0x4d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x81, 0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x4d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x2e, 0x77, 0x65, 0x62,
	0x69, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x4d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x49, 0x4d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x7d, 0x2f, 0x69, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x5e, 0x0a, 0x0f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x4d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x28, 0x2e, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x4d, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x77, 0x65, 0x62, 0x69,
	0x74, 0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x22, 0x5a, 0x20,
	0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2e, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_contacts_imclients_proto_rawDescOnce sync.Once
	file_gateway_contacts_imclients_proto_rawDescData = file_gateway_contacts_imclients_proto_rawDesc
)

func file_gateway_contacts_imclients_proto_rawDescGZIP() []byte {
	file_gateway_contacts_imclients_proto_rawDescOnce.Do(func() {
		file_gateway_contacts_imclients_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_contacts_imclients_proto_rawDescData)
	})
	return file_gateway_contacts_imclients_proto_rawDescData
}

var file_gateway_contacts_imclients_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_gateway_contacts_imclients_proto_goTypes = []interface{}{
	(*IMClient)(nil),               // 0: webitel.contacts.IMClient
	(*InputIMClient)(nil),          // 1: webitel.contacts.InputIMClient
	(*IMClientList)(nil),           // 2: webitel.contacts.IMClientList
	(*EmptyResponse)(nil),          // 3: webitel.contacts.EmptyResponse
	(*ListIMClientsRequest)(nil),   // 4: webitel.contacts.ListIMClientsRequest
	(*CreateIMClientsRequest)(nil), // 5: webitel.contacts.CreateIMClientsRequest
	(*Lookup)(nil),                 // 6: webitel.contacts.Lookup
}
var file_gateway_contacts_imclients_proto_depIdxs = []int32{
	6, // 0: webitel.contacts.IMClient.created_by:type_name -> webitel.contacts.Lookup
	6, // 1: webitel.contacts.IMClient.updated_by:type_name -> webitel.contacts.Lookup
	6, // 2: webitel.contacts.IMClient.user:type_name -> webitel.contacts.Lookup
	6, // 3: webitel.contacts.IMClient.app:type_name -> webitel.contacts.Lookup
	0, // 4: webitel.contacts.IMClientList.data:type_name -> webitel.contacts.IMClient
	1, // 5: webitel.contacts.CreateIMClientsRequest.input:type_name -> webitel.contacts.InputIMClient
	4, // 6: webitel.contacts.IMClients.ListIMClients:input_type -> webitel.contacts.ListIMClientsRequest
	5, // 7: webitel.contacts.IMClients.CreateIMClients:input_type -> webitel.contacts.CreateIMClientsRequest
	2, // 8: webitel.contacts.IMClients.ListIMClients:output_type -> webitel.contacts.IMClientList
	3, // 9: webitel.contacts.IMClients.CreateIMClients:output_type -> webitel.contacts.EmptyResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_gateway_contacts_imclients_proto_init() }
func file_gateway_contacts_imclients_proto_init() {
	if File_gateway_contacts_imclients_proto != nil {
		return
	}
	file_gateway_contacts_fields_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_gateway_contacts_imclients_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IMClient); i {
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
		file_gateway_contacts_imclients_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputIMClient); i {
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
		file_gateway_contacts_imclients_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IMClientList); i {
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
		file_gateway_contacts_imclients_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
		file_gateway_contacts_imclients_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListIMClientsRequest); i {
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
		file_gateway_contacts_imclients_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIMClientsRequest); i {
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
			RawDescriptor: file_gateway_contacts_imclients_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gateway_contacts_imclients_proto_goTypes,
		DependencyIndexes: file_gateway_contacts_imclients_proto_depIdxs,
		MessageInfos:      file_gateway_contacts_imclients_proto_msgTypes,
	}.Build()
	File_gateway_contacts_imclients_proto = out.File
	file_gateway_contacts_imclients_proto_rawDesc = nil
	file_gateway_contacts_imclients_proto_goTypes = nil
	file_gateway_contacts_imclients_proto_depIdxs = nil
}
