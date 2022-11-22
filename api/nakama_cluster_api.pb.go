// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: nakama_cluster_api.proto

package api

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

type Envelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Node string `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	Cid  string `protobuf:"bytes,3,opt,name=cid,proto3" json:"cid,omitempty"`
	// Types that are assignable to Payload:
	//	*Envelope_Bytes
	//	*Envelope_Error
	//	*Envelope_Track
	//	*Envelope_Untrack
	//	*Envelope_UntrackAll
	//	*Envelope_Message
	Payload isEnvelope_Payload `protobuf_oneof:"payload"`
}

func (x *Envelope) Reset() {
	*x = Envelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nakama_cluster_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envelope) ProtoMessage() {}

func (x *Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_nakama_cluster_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Envelope.ProtoReflect.Descriptor instead.
func (*Envelope) Descriptor() ([]byte, []int) {
	return file_nakama_cluster_api_proto_rawDescGZIP(), []int{0}
}

func (x *Envelope) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Envelope) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *Envelope) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (m *Envelope) GetPayload() isEnvelope_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Envelope) GetBytes() []byte {
	if x, ok := x.GetPayload().(*Envelope_Bytes); ok {
		return x.Bytes
	}
	return nil
}

func (x *Envelope) GetError() *Error {
	if x, ok := x.GetPayload().(*Envelope_Error); ok {
		return x.Error
	}
	return nil
}

func (x *Envelope) GetTrack() *Track {
	if x, ok := x.GetPayload().(*Envelope_Track); ok {
		return x.Track
	}
	return nil
}

func (x *Envelope) GetUntrack() *Untrack {
	if x, ok := x.GetPayload().(*Envelope_Untrack); ok {
		return x.Untrack
	}
	return nil
}

func (x *Envelope) GetUntrackAll() *UntrackAll {
	if x, ok := x.GetPayload().(*Envelope_UntrackAll); ok {
		return x.UntrackAll
	}
	return nil
}

func (x *Envelope) GetMessage() *Message {
	if x, ok := x.GetPayload().(*Envelope_Message); ok {
		return x.Message
	}
	return nil
}

type isEnvelope_Payload interface {
	isEnvelope_Payload()
}

type Envelope_Bytes struct {
	Bytes []byte `protobuf:"bytes,4,opt,name=bytes,proto3,oneof"`
}

type Envelope_Error struct {
	Error *Error `protobuf:"bytes,5,opt,name=error,proto3,oneof"`
}

type Envelope_Track struct {
	Track *Track `protobuf:"bytes,6,opt,name=track,proto3,oneof"`
}

type Envelope_Untrack struct {
	Untrack *Untrack `protobuf:"bytes,7,opt,name=untrack,proto3,oneof"`
}

type Envelope_UntrackAll struct {
	UntrackAll *UntrackAll `protobuf:"bytes,8,opt,name=untrackAll,proto3,oneof"`
}

type Envelope_Message struct {
	Message *Message `protobuf:"bytes,9,opt,name=message,proto3,oneof"`
}

func (*Envelope_Bytes) isEnvelope_Payload() {}

func (*Envelope_Error) isEnvelope_Payload() {}

func (*Envelope_Track) isEnvelope_Payload() {}

func (*Envelope_Untrack) isEnvelope_Payload() {}

func (*Envelope_UntrackAll) isEnvelope_Payload() {}

func (*Envelope_Message) isEnvelope_Payload() {}

// error
type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The error code which should be one of "Error.Code" enums.
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// A message in English to help developers debug the response.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// Additional error details which may be different for each response.
	Context map[string]string `protobuf:"bytes,3,rep,name=context,proto3" json:"context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nakama_cluster_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_nakama_cluster_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_nakama_cluster_api_proto_rawDescGZIP(), []int{1}
}

func (x *Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetContext() map[string]string {
	if x != nil {
		return x.Context
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID []string `protobuf:"bytes,1,rep,name=sessionID,proto3" json:"sessionID,omitempty"`
	Content   []byte   `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nakama_cluster_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_nakama_cluster_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_nakama_cluster_api_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetSessionID() []string {
	if x != nil {
		return x.SessionID
	}
	return nil
}

func (x *Message) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type PresenceID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node      string `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	SessionID string `protobuf:"bytes,2,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
}

func (x *PresenceID) Reset() {
	*x = PresenceID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nakama_cluster_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PresenceID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PresenceID) ProtoMessage() {}

func (x *PresenceID) ProtoReflect() protoreflect.Message {
	mi := &file_nakama_cluster_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PresenceID.ProtoReflect.Descriptor instead.
func (*PresenceID) Descriptor() ([]byte, []int) {
	return file_nakama_cluster_api_proto_rawDescGZIP(), []int{3}
}

func (x *PresenceID) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *PresenceID) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

type PresenceStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode       int32  `protobuf:"varint,1,opt,name=mode,proto3" json:"mode,omitempty"`
	Subject    string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Subcontext string `protobuf:"bytes,3,opt,name=subcontext,proto3" json:"subcontext,omitempty"`
	Label      string `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *PresenceStream) Reset() {
	*x = PresenceStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nakama_cluster_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PresenceStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PresenceStream) ProtoMessage() {}

func (x *PresenceStream) ProtoReflect() protoreflect.Message {
	mi := &file_nakama_cluster_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PresenceStream.ProtoReflect.Descriptor instead.
func (*PresenceStream) Descriptor() ([]byte, []int) {
	return file_nakama_cluster_api_proto_rawDescGZIP(), []int{4}
}

func (x *PresenceStream) GetMode() int32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

func (x *PresenceStream) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *PresenceStream) GetSubcontext() string {
	if x != nil {
		return x.Subcontext
	}
	return ""
}

func (x *PresenceStream) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type PresenceMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionFormat int32  `protobuf:"varint,1,opt,name=sessionFormat,proto3" json:"sessionFormat,omitempty"`
	Hidden        bool   `protobuf:"varint,2,opt,name=hidden,proto3" json:"hidden,omitempty"`
	Persistence   bool   `protobuf:"varint,3,opt,name=persistence,proto3" json:"persistence,omitempty"`
	Username      string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Status        string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Reason        int32  `protobuf:"varint,6,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *PresenceMeta) Reset() {
	*x = PresenceMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nakama_cluster_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PresenceMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PresenceMeta) ProtoMessage() {}

func (x *PresenceMeta) ProtoReflect() protoreflect.Message {
	mi := &file_nakama_cluster_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PresenceMeta.ProtoReflect.Descriptor instead.
func (*PresenceMeta) Descriptor() ([]byte, []int) {
	return file_nakama_cluster_api_proto_rawDescGZIP(), []int{5}
}

func (x *PresenceMeta) GetSessionFormat() int32 {
	if x != nil {
		return x.SessionFormat
	}
	return 0
}

func (x *PresenceMeta) GetHidden() bool {
	if x != nil {
		return x.Hidden
	}
	return false
}

func (x *PresenceMeta) GetPersistence() bool {
	if x != nil {
		return x.Persistence
	}
	return false
}

func (x *PresenceMeta) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *PresenceMeta) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PresenceMeta) GetReason() int32 {
	if x != nil {
		return x.Reason
	}
	return 0
}

type Presence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     *PresenceID     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Stream *PresenceStream `protobuf:"bytes,2,opt,name=stream,proto3" json:"stream,omitempty"`
	UserID string          `protobuf:"bytes,3,opt,name=userID,proto3" json:"userID,omitempty"`
	Meta   *PresenceMeta   `protobuf:"bytes,4,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *Presence) Reset() {
	*x = Presence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nakama_cluster_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Presence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Presence) ProtoMessage() {}

func (x *Presence) ProtoReflect() protoreflect.Message {
	mi := &file_nakama_cluster_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Presence.ProtoReflect.Descriptor instead.
func (*Presence) Descriptor() ([]byte, []int) {
	return file_nakama_cluster_api_proto_rawDescGZIP(), []int{6}
}

func (x *Presence) GetId() *PresenceID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Presence) GetStream() *PresenceStream {
	if x != nil {
		return x.Stream
	}
	return nil
}

func (x *Presence) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Presence) GetMeta() *PresenceMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type Presences struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Presences []*Presence `protobuf:"bytes,1,rep,name=Presences,proto3" json:"Presences,omitempty"`
}

func (x *Presences) Reset() {
	*x = Presences{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nakama_cluster_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Presences) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Presences) ProtoMessage() {}

func (x *Presences) ProtoReflect() protoreflect.Message {
	mi := &file_nakama_cluster_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Presences.ProtoReflect.Descriptor instead.
func (*Presences) Descriptor() ([]byte, []int) {
	return file_nakama_cluster_api_proto_rawDescGZIP(), []int{7}
}

func (x *Presences) GetPresences() []*Presence {
	if x != nil {
		return x.Presences
	}
	return nil
}

type Track struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Presences []*Presence `protobuf:"bytes,1,rep,name=Presences,proto3" json:"Presences,omitempty"`
}

func (x *Track) Reset() {
	*x = Track{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nakama_cluster_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Track) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Track) ProtoMessage() {}

func (x *Track) ProtoReflect() protoreflect.Message {
	mi := &file_nakama_cluster_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Track.ProtoReflect.Descriptor instead.
func (*Track) Descriptor() ([]byte, []int) {
	return file_nakama_cluster_api_proto_rawDescGZIP(), []int{8}
}

func (x *Track) GetPresences() []*Presence {
	if x != nil {
		return x.Presences
	}
	return nil
}

type Untrack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Presences []*Presence `protobuf:"bytes,1,rep,name=Presences,proto3" json:"Presences,omitempty"`
}

func (x *Untrack) Reset() {
	*x = Untrack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nakama_cluster_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Untrack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Untrack) ProtoMessage() {}

func (x *Untrack) ProtoReflect() protoreflect.Message {
	mi := &file_nakama_cluster_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Untrack.ProtoReflect.Descriptor instead.
func (*Untrack) Descriptor() ([]byte, []int) {
	return file_nakama_cluster_api_proto_rawDescGZIP(), []int{9}
}

func (x *Untrack) GetPresences() []*Presence {
	if x != nil {
		return x.Presences
	}
	return nil
}

type UntrackAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID string `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	Reason    int32  `protobuf:"varint,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *UntrackAll) Reset() {
	*x = UntrackAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nakama_cluster_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UntrackAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UntrackAll) ProtoMessage() {}

func (x *UntrackAll) ProtoReflect() protoreflect.Message {
	mi := &file_nakama_cluster_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UntrackAll.ProtoReflect.Descriptor instead.
func (*UntrackAll) Descriptor() ([]byte, []int) {
	return file_nakama_cluster_api_proto_rawDescGZIP(), []int{10}
}

func (x *UntrackAll) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

func (x *UntrackAll) GetReason() int32 {
	if x != nil {
		return x.Reason
	}
	return 0
}

var File_nakama_cluster_api_proto protoreflect.FileDescriptor

var file_nakama_cluster_api_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6e, 0x61, 0x6b, 0x61, 0x6d, 0x61, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6e, 0x61, 0x6b, 0x61,
	0x6d, 0x61, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0xe9, 0x02, 0x0a, 0x08, 0x45,
	0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6e, 0x61, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x2d, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6e, 0x61, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x05, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x12, 0x33, 0x0a, 0x07, 0x75, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x61, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x48, 0x00, 0x52,
	0x07, 0x75, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x3c, 0x0a, 0x0a, 0x75, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x41, 0x6c, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e,
	0x61, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x41, 0x6c, 0x6c, 0x48, 0x00, 0x52, 0x0a, 0x75, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x41, 0x6c, 0x6c, 0x12, 0x33, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x61, 0x6b, 0x61, 0x6d, 0x61,
	0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xaf, 0x01, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3c,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x6e, 0x61, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x1a, 0x3a, 0x0a, 0x0c,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x41, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x3e, 0x0a, 0x0a, 0x50,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0x74, 0x0a, 0x0e, 0x50,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x12, 0x0a,
	0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6d, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x75, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x75, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x22, 0xba, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x69, 0x64, 0x64,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0xb8,
	0x01, 0x0a, 0x08, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x61, 0x6b, 0x61, 0x6d, 0x61,
	0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63,
	0x65, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x61, 0x6b, 0x61, 0x6d, 0x61,
	0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63,
	0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x61, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x43, 0x0a, 0x09, 0x50, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x09, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x61, 0x6b, 0x61,
	0x6d, 0x61, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x63, 0x65, 0x52, 0x09, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x3f,
	0x0a, 0x05, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x36, 0x0a, 0x09, 0x50, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x61, 0x6b,
	0x61, 0x6d, 0x61, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x09, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x22,
	0x41, 0x0a, 0x07, 0x55, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x36, 0x0a, 0x09, 0x50, 0x72,
	0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x6e, 0x61, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x09, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x22, 0x42, 0x0a, 0x0a, 0x55, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x41, 0x6c, 0x6c,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x32, 0x8d, 0x01, 0x0a, 0x09, 0x41, 0x70, 0x69, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x18, 0x2e, 0x6e,
	0x61, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x45, 0x6e,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x1a, 0x18, 0x2e, 0x6e, 0x61, 0x6b, 0x61, 0x6d, 0x61, 0x2e,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65,
	0x22, 0x00, 0x12, 0x42, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x18, 0x2e, 0x6e,
	0x61, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x45, 0x6e,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x1a, 0x18, 0x2e, 0x6e, 0x61, 0x6b, 0x61, 0x6d, 0x61, 0x2e,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x6d, 0x6f, 0x2f, 0x6e, 0x61,
	0x6b, 0x61, 0x6d, 0x61, 0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nakama_cluster_api_proto_rawDescOnce sync.Once
	file_nakama_cluster_api_proto_rawDescData = file_nakama_cluster_api_proto_rawDesc
)

func file_nakama_cluster_api_proto_rawDescGZIP() []byte {
	file_nakama_cluster_api_proto_rawDescOnce.Do(func() {
		file_nakama_cluster_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_nakama_cluster_api_proto_rawDescData)
	})
	return file_nakama_cluster_api_proto_rawDescData
}

var file_nakama_cluster_api_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_nakama_cluster_api_proto_goTypes = []interface{}{
	(*Envelope)(nil),       // 0: nakama.cluster.Envelope
	(*Error)(nil),          // 1: nakama.cluster.Error
	(*Message)(nil),        // 2: nakama.cluster.Message
	(*PresenceID)(nil),     // 3: nakama.cluster.PresenceID
	(*PresenceStream)(nil), // 4: nakama.cluster.PresenceStream
	(*PresenceMeta)(nil),   // 5: nakama.cluster.PresenceMeta
	(*Presence)(nil),       // 6: nakama.cluster.Presence
	(*Presences)(nil),      // 7: nakama.cluster.Presences
	(*Track)(nil),          // 8: nakama.cluster.Track
	(*Untrack)(nil),        // 9: nakama.cluster.Untrack
	(*UntrackAll)(nil),     // 10: nakama.cluster.UntrackAll
	nil,                    // 11: nakama.cluster.Error.ContextEntry
}
var file_nakama_cluster_api_proto_depIdxs = []int32{
	1,  // 0: nakama.cluster.Envelope.error:type_name -> nakama.cluster.Error
	8,  // 1: nakama.cluster.Envelope.track:type_name -> nakama.cluster.Track
	9,  // 2: nakama.cluster.Envelope.untrack:type_name -> nakama.cluster.Untrack
	10, // 3: nakama.cluster.Envelope.untrackAll:type_name -> nakama.cluster.UntrackAll
	2,  // 4: nakama.cluster.Envelope.message:type_name -> nakama.cluster.Message
	11, // 5: nakama.cluster.Error.context:type_name -> nakama.cluster.Error.ContextEntry
	3,  // 6: nakama.cluster.Presence.id:type_name -> nakama.cluster.PresenceID
	4,  // 7: nakama.cluster.Presence.stream:type_name -> nakama.cluster.PresenceStream
	5,  // 8: nakama.cluster.Presence.meta:type_name -> nakama.cluster.PresenceMeta
	6,  // 9: nakama.cluster.Presences.Presences:type_name -> nakama.cluster.Presence
	6,  // 10: nakama.cluster.Track.Presences:type_name -> nakama.cluster.Presence
	6,  // 11: nakama.cluster.Untrack.Presences:type_name -> nakama.cluster.Presence
	0,  // 12: nakama.cluster.ApiServer.Call:input_type -> nakama.cluster.Envelope
	0,  // 13: nakama.cluster.ApiServer.Stream:input_type -> nakama.cluster.Envelope
	0,  // 14: nakama.cluster.ApiServer.Call:output_type -> nakama.cluster.Envelope
	0,  // 15: nakama.cluster.ApiServer.Stream:output_type -> nakama.cluster.Envelope
	14, // [14:16] is the sub-list for method output_type
	12, // [12:14] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_nakama_cluster_api_proto_init() }
func file_nakama_cluster_api_proto_init() {
	if File_nakama_cluster_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nakama_cluster_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Envelope); i {
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
		file_nakama_cluster_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_nakama_cluster_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_nakama_cluster_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PresenceID); i {
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
		file_nakama_cluster_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PresenceStream); i {
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
		file_nakama_cluster_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PresenceMeta); i {
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
		file_nakama_cluster_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Presence); i {
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
		file_nakama_cluster_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Presences); i {
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
		file_nakama_cluster_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Track); i {
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
		file_nakama_cluster_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Untrack); i {
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
		file_nakama_cluster_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UntrackAll); i {
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
	file_nakama_cluster_api_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Envelope_Bytes)(nil),
		(*Envelope_Error)(nil),
		(*Envelope_Track)(nil),
		(*Envelope_Untrack)(nil),
		(*Envelope_UntrackAll)(nil),
		(*Envelope_Message)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nakama_cluster_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nakama_cluster_api_proto_goTypes,
		DependencyIndexes: file_nakama_cluster_api_proto_depIdxs,
		MessageInfos:      file_nakama_cluster_api_proto_msgTypes,
	}.Build()
	File_nakama_cluster_api_proto = out.File
	file_nakama_cluster_api_proto_rawDesc = nil
	file_nakama_cluster_api_proto_goTypes = nil
	file_nakama_cluster_api_proto_depIdxs = nil
}
