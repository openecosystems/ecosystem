// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: platform/audit/v2alpha/audit.proto

package auditv2alphapb

import (
	_ "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/options/v2"
	v2 "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/spec/v2"
	v21 "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/type/v2"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// *
// Commands used with the Audit System
type AuditCommands int32

const (
	AuditCommands_AUDIT_COMMANDS_UNSPECIFIED AuditCommands = 0 // No command specified.
)

// Enum value maps for AuditCommands.
var (
	AuditCommands_name = map[int32]string{
		0: "AUDIT_COMMANDS_UNSPECIFIED",
	}
	AuditCommands_value = map[string]int32{
		"AUDIT_COMMANDS_UNSPECIFIED": 0,
	}
)

func (x AuditCommands) Enum() *AuditCommands {
	p := new(AuditCommands)
	*p = x
	return p
}

func (x AuditCommands) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuditCommands) Descriptor() protoreflect.EnumDescriptor {
	return file_platform_audit_v2alpha_audit_proto_enumTypes[0].Descriptor()
}

func (AuditCommands) Type() protoreflect.EnumType {
	return &file_platform_audit_v2alpha_audit_proto_enumTypes[0]
}

func (x AuditCommands) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuditCommands.Descriptor instead.
func (AuditCommands) EnumDescriptor() ([]byte, []int) {
	return file_platform_audit_v2alpha_audit_proto_rawDescGZIP(), []int{0}
}

// *
// Event associated with Audit System.
type AuditEvents int32

const (
	AuditEvents_AUDIT_EVENTS_UNSPECIFIED AuditEvents = 0 // No event specified
	AuditEvents_AUDIT_EVENTS_CREATED     AuditEvents = 1
)

// Enum value maps for AuditEvents.
var (
	AuditEvents_name = map[int32]string{
		0: "AUDIT_EVENTS_UNSPECIFIED",
		1: "AUDIT_EVENTS_CREATED",
	}
	AuditEvents_value = map[string]int32{
		"AUDIT_EVENTS_UNSPECIFIED": 0,
		"AUDIT_EVENTS_CREATED":     1,
	}
)

func (x AuditEvents) Enum() *AuditEvents {
	p := new(AuditEvents)
	*p = x
	return p
}

func (x AuditEvents) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuditEvents) Descriptor() protoreflect.EnumDescriptor {
	return file_platform_audit_v2alpha_audit_proto_enumTypes[1].Descriptor()
}

func (AuditEvents) Type() protoreflect.EnumType {
	return &file_platform_audit_v2alpha_audit_proto_enumTypes[1]
}

func (x AuditEvents) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuditEvents.Descriptor instead.
func (AuditEvents) EnumDescriptor() ([]byte, []int) {
	return file_platform_audit_v2alpha_audit_proto_rawDescGZIP(), []int{1}
}

type AuditConfiguration struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuditConfiguration) Reset() {
	*x = AuditConfiguration{}
	mi := &file_platform_audit_v2alpha_audit_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuditConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditConfiguration) ProtoMessage() {}

func (x *AuditConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_platform_audit_v2alpha_audit_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditConfiguration.ProtoReflect.Descriptor instead.
func (*AuditConfiguration) Descriptor() ([]byte, []int) {
	return file_platform_audit_v2alpha_audit_proto_rawDescGZIP(), []int{0}
}

func (x *AuditConfiguration) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// *
// Message request for a search.
type SearchRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Timestamp to begin searching.
	StartAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	// Timestamp to end searching
	EndAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	// Indicates the page size
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token provided in the response
	PageToken     string `protobuf:"bytes,5,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	mi := &file_platform_audit_v2alpha_audit_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_platform_audit_v2alpha_audit_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_platform_audit_v2alpha_audit_proto_rawDescGZIP(), []int{1}
}

func (x *SearchRequest) GetStartAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartAt
	}
	return nil
}

func (x *SearchRequest) GetEndAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndAt
	}
	return nil
}

func (x *SearchRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SearchRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Message response from a search.
type SearchResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The response context
	SpecContext *v2.SpecResponseContext `protobuf:"bytes,1,opt,name=spec_context,json=specContext,proto3" json:"spec_context,omitempty"`
	// Total count of audits
	TotalSize int32 `protobuf:"varint,2,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	// Token to retrieve the next page
	NextPageToken string `protobuf:"bytes,3,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// List of audits
	Audits        []*Audit `protobuf:"bytes,4,rep,name=audits,proto3" json:"audits,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	mi := &file_platform_audit_v2alpha_audit_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_platform_audit_v2alpha_audit_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_platform_audit_v2alpha_audit_proto_rawDescGZIP(), []int{2}
}

func (x *SearchResponse) GetSpecContext() *v2.SpecResponseContext {
	if x != nil {
		return x.SpecContext
	}
	return nil
}

func (x *SearchResponse) GetTotalSize() int32 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

func (x *SearchResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *SearchResponse) GetAudits() []*Audit {
	if x != nil {
		return x.Audits
	}
	return nil
}

// Audit Entry
type AuditEntry struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Data block recorded for Audit System
	Data          *anypb.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuditEntry) Reset() {
	*x = AuditEntry{}
	mi := &file_platform_audit_v2alpha_audit_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuditEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditEntry) ProtoMessage() {}

func (x *AuditEntry) ProtoReflect() protoreflect.Message {
	mi := &file_platform_audit_v2alpha_audit_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditEntry.ProtoReflect.Descriptor instead.
func (*AuditEntry) Descriptor() ([]byte, []int) {
	return file_platform_audit_v2alpha_audit_proto_rawDescGZIP(), []int{3}
}

func (x *AuditEntry) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

// Audit Message
type Audit struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID used to identify this message
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Timestamp indicating when this message was created.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Timestamp indicating when this message was last updated.
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Audit Entry to record.
	Entry         *AuditEntry      `protobuf:"bytes,4,opt,name=entry,proto3" json:"entry,omitempty"`
	Jurisdiction  v21.Jurisdiction `protobuf:"varint,5,opt,name=jurisdiction,proto3,enum=platform.type.v2.Jurisdiction" json:"jurisdiction,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Audit) Reset() {
	*x = Audit{}
	mi := &file_platform_audit_v2alpha_audit_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Audit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Audit) ProtoMessage() {}

func (x *Audit) ProtoReflect() protoreflect.Message {
	mi := &file_platform_audit_v2alpha_audit_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Audit.ProtoReflect.Descriptor instead.
func (*Audit) Descriptor() ([]byte, []int) {
	return file_platform_audit_v2alpha_audit_proto_rawDescGZIP(), []int{4}
}

func (x *Audit) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Audit) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Audit) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Audit) GetEntry() *AuditEntry {
	if x != nil {
		return x.Entry
	}
	return nil
}

func (x *Audit) GetJurisdiction() v21.Jurisdiction {
	if x != nil {
		return x.Jurisdiction
	}
	return v21.Jurisdiction(0)
}

var File_platform_audit_v2alpha_audit_proto protoreflect.FileDescriptor

var file_platform_audit_v2alpha_audit_proto_rawDesc = string([]byte{
	0x0a, 0x22, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x76, 0x32, 0x2f,
	0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x6a, 0x75, 0x72,
	0x69, 0x73, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x42, 0x0a, 0x12, 0x41, 0x75, 0x64, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x1c, 0xd2, 0xb7, 0x18, 0x18, 0x0a, 0x16, 0x1a, 0x12, 0x54, 0x68, 0x69, 0x73, 0x20,
	0x69, 0x73, 0x20, 0x61, 0x20, 0x61, 0x75, 0x64, 0x69, 0x74, 0x20, 0x69, 0x64, 0x28, 0x01, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xbd, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x31, 0x0a, 0x06,
	0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x41, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x3a, 0x06, 0xfa, 0xb6, 0x18,
	0x02, 0x08, 0x01, 0x22, 0xe0, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0c, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x76, 0x32, 0x2e,
	0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x52, 0x0b, 0x73, 0x70, 0x65, 0x63, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x35, 0x0a, 0x06, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x06, 0x61, 0x75, 0x64, 0x69, 0x74, 0x73, 0x3a, 0x06,
	0xfa, 0xb6, 0x18, 0x02, 0x08, 0x02, 0x22, 0x3e, 0x0a, 0x0a, 0x41, 0x75, 0x64, 0x69, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x06,
	0xfa, 0xb6, 0x18, 0x02, 0x08, 0x02, 0x22, 0x9b, 0x02, 0x0a, 0x05, 0x41, 0x75, 0x64, 0x69, 0x74,
	0x12, 0x16, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xca, 0xb7,
	0x18, 0x02, 0x08, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38,
	0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x42, 0x0a, 0x0c, 0x6a, 0x75, 0x72, 0x69,
	0x73, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e,
	0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76,
	0x32, 0x2e, 0x4a, 0x75, 0x72, 0x69, 0x73, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c,
	0x6a, 0x75, 0x72, 0x69, 0x73, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x06, 0xfa, 0xb6,
	0x18, 0x02, 0x08, 0x02, 0x2a, 0x37, 0x0a, 0x0d, 0x41, 0x75, 0x64, 0x69, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x1e, 0x0a, 0x1a, 0x41, 0x55, 0x44, 0x49, 0x54, 0x5f, 0x43,
	0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x1a, 0x06, 0x92, 0xb8, 0x18, 0x02, 0x08, 0x03, 0x2a, 0x55, 0x0a,
	0x0b, 0x41, 0x75, 0x64, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x18,
	0x41, 0x55, 0x44, 0x49, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x14, 0x41, 0x55,
	0x44, 0x49, 0x54, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x53, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x44, 0x10, 0x01, 0x1a, 0x06, 0xe2, 0xb8, 0x18, 0x02, 0x08, 0x01, 0x1a, 0x06, 0x92, 0xb8,
	0x18, 0x02, 0x08, 0x04, 0x32, 0x91, 0x01, 0x0a, 0x0c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x80, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x12, 0x25, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x27, 0xa2, 0xb6, 0x18, 0x08, 0x2a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0xaa, 0xb6, 0x18,
	0x02, 0x08, 0x08, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x76,
	0x32, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0xa0, 0x01, 0x82, 0xc4, 0x13, 0x02, 0x08,
	0x01, 0x82, 0xb5, 0x18, 0x06, 0x08, 0x01, 0x10, 0x01, 0x18, 0x02, 0x8a, 0xb5, 0x18, 0x16, 0x0a,
	0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x12, 0x06, 0x61, 0x75, 0x64, 0x69, 0x74, 0x73, 0x22, 0x03,
	0x6a, 0x61, 0x6e, 0x28, 0x02, 0x92, 0xb5, 0x18, 0x03, 0x0a, 0x01, 0x03, 0x9a, 0xb5, 0x18, 0x02,
	0x08, 0x01, 0xa2, 0xb5, 0x18, 0x02, 0x08, 0x01, 0x5a, 0x61, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x73, 0x2f, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x6c, 0x69,
	0x62, 0x73, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x64,
	0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_platform_audit_v2alpha_audit_proto_rawDescOnce sync.Once
	file_platform_audit_v2alpha_audit_proto_rawDescData []byte
)

func file_platform_audit_v2alpha_audit_proto_rawDescGZIP() []byte {
	file_platform_audit_v2alpha_audit_proto_rawDescOnce.Do(func() {
		file_platform_audit_v2alpha_audit_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_platform_audit_v2alpha_audit_proto_rawDesc), len(file_platform_audit_v2alpha_audit_proto_rawDesc)))
	})
	return file_platform_audit_v2alpha_audit_proto_rawDescData
}

var file_platform_audit_v2alpha_audit_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_platform_audit_v2alpha_audit_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_platform_audit_v2alpha_audit_proto_goTypes = []any{
	(AuditCommands)(0),             // 0: platform.audit.v2alpha.AuditCommands
	(AuditEvents)(0),               // 1: platform.audit.v2alpha.AuditEvents
	(*AuditConfiguration)(nil),     // 2: platform.audit.v2alpha.AuditConfiguration
	(*SearchRequest)(nil),          // 3: platform.audit.v2alpha.SearchRequest
	(*SearchResponse)(nil),         // 4: platform.audit.v2alpha.SearchResponse
	(*AuditEntry)(nil),             // 5: platform.audit.v2alpha.AuditEntry
	(*Audit)(nil),                  // 6: platform.audit.v2alpha.Audit
	(*timestamppb.Timestamp)(nil),  // 7: google.protobuf.Timestamp
	(*v2.SpecResponseContext)(nil), // 8: platform.spec.v2.SpecResponseContext
	(*anypb.Any)(nil),              // 9: google.protobuf.Any
	(v21.Jurisdiction)(0),          // 10: platform.type.v2.Jurisdiction
}
var file_platform_audit_v2alpha_audit_proto_depIdxs = []int32{
	7,  // 0: platform.audit.v2alpha.SearchRequest.start_at:type_name -> google.protobuf.Timestamp
	7,  // 1: platform.audit.v2alpha.SearchRequest.end_at:type_name -> google.protobuf.Timestamp
	8,  // 2: platform.audit.v2alpha.SearchResponse.spec_context:type_name -> platform.spec.v2.SpecResponseContext
	6,  // 3: platform.audit.v2alpha.SearchResponse.audits:type_name -> platform.audit.v2alpha.Audit
	9,  // 4: platform.audit.v2alpha.AuditEntry.data:type_name -> google.protobuf.Any
	7,  // 5: platform.audit.v2alpha.Audit.created_at:type_name -> google.protobuf.Timestamp
	7,  // 6: platform.audit.v2alpha.Audit.updated_at:type_name -> google.protobuf.Timestamp
	5,  // 7: platform.audit.v2alpha.Audit.entry:type_name -> platform.audit.v2alpha.AuditEntry
	10, // 8: platform.audit.v2alpha.Audit.jurisdiction:type_name -> platform.type.v2.Jurisdiction
	3,  // 9: platform.audit.v2alpha.AuditService.Search:input_type -> platform.audit.v2alpha.SearchRequest
	4,  // 10: platform.audit.v2alpha.AuditService.Search:output_type -> platform.audit.v2alpha.SearchResponse
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_platform_audit_v2alpha_audit_proto_init() }
func file_platform_audit_v2alpha_audit_proto_init() {
	if File_platform_audit_v2alpha_audit_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_platform_audit_v2alpha_audit_proto_rawDesc), len(file_platform_audit_v2alpha_audit_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_platform_audit_v2alpha_audit_proto_goTypes,
		DependencyIndexes: file_platform_audit_v2alpha_audit_proto_depIdxs,
		EnumInfos:         file_platform_audit_v2alpha_audit_proto_enumTypes,
		MessageInfos:      file_platform_audit_v2alpha_audit_proto_msgTypes,
	}.Build()
	File_platform_audit_v2alpha_audit_proto = out.File
	file_platform_audit_v2alpha_audit_proto_goTypes = nil
	file_platform_audit_v2alpha_audit_proto_depIdxs = nil
}
