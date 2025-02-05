// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: api/proto/wa.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

// ✅ Request for sending a text message
type TextMessageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Conversation  string                 `protobuf:"bytes,1,opt,name=conversation,proto3" json:"conversation,omitempty"`
	Metadata      *MessageMetadata       `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TextMessageRequest) Reset() {
	*x = TextMessageRequest{}
	mi := &file_api_proto_wa_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TextMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextMessageRequest) ProtoMessage() {}

func (x *TextMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_wa_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextMessageRequest.ProtoReflect.Descriptor instead.
func (*TextMessageRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_wa_proto_rawDescGZIP(), []int{0}
}

func (x *TextMessageRequest) GetConversation() string {
	if x != nil {
		return x.Conversation
	}
	return ""
}

func (x *TextMessageRequest) GetMetadata() *MessageMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// ✅ Request for sending an extended text message
type ExtendedTextMessageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Text          string                 `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	ContextInfo   *ContextInfo           `protobuf:"bytes,2,opt,name=contextInfo,proto3" json:"contextInfo,omitempty"`
	Metadata      *MessageMetadata       `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExtendedTextMessageRequest) Reset() {
	*x = ExtendedTextMessageRequest{}
	mi := &file_api_proto_wa_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExtendedTextMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendedTextMessageRequest) ProtoMessage() {}

func (x *ExtendedTextMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_wa_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendedTextMessageRequest.ProtoReflect.Descriptor instead.
func (*ExtendedTextMessageRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_wa_proto_rawDescGZIP(), []int{1}
}

func (x *ExtendedTextMessageRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ExtendedTextMessageRequest) GetContextInfo() *ContextInfo {
	if x != nil {
		return x.ContextInfo
	}
	return nil
}

func (x *ExtendedTextMessageRequest) GetMetadata() *MessageMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// ✅ Common response for sending a message
type CommonMessageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CommonMessageResponse) Reset() {
	*x = CommonMessageResponse{}
	mi := &file_api_proto_wa_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommonMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonMessageResponse) ProtoMessage() {}

func (x *CommonMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_wa_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonMessageResponse.ProtoReflect.Descriptor instead.
func (*CommonMessageResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_wa_proto_rawDescGZIP(), []int{2}
}

func (x *CommonMessageResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// ✅ Request for media upload (sent in chunks)
type MediaUploadRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Chunk         []byte                 `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"` // The media file sent in chunks
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MediaUploadRequest) Reset() {
	*x = MediaUploadRequest{}
	mi := &file_api_proto_wa_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MediaUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaUploadRequest) ProtoMessage() {}

func (x *MediaUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_wa_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaUploadRequest.ProtoReflect.Descriptor instead.
func (*MediaUploadRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_wa_proto_rawDescGZIP(), []int{3}
}

func (x *MediaUploadRequest) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

// ✅ Response after successful media upload
type MediaUploadResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileUrl       string                 `protobuf:"bytes,1,opt,name=fileUrl,proto3" json:"fileUrl,omitempty"`   // URL of the uploaded file
	FilePath      string                 `protobuf:"bytes,2,opt,name=filePath,proto3" json:"filePath,omitempty"` // File name
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MediaUploadResponse) Reset() {
	*x = MediaUploadResponse{}
	mi := &file_api_proto_wa_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MediaUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaUploadResponse) ProtoMessage() {}

func (x *MediaUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_wa_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaUploadResponse.ProtoReflect.Descriptor instead.
func (*MediaUploadResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_wa_proto_rawDescGZIP(), []int{4}
}

func (x *MediaUploadResponse) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

func (x *MediaUploadResponse) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

type MessageMetadata struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RemoteJid     string                 `protobuf:"bytes,1,opt,name=remoteJid,proto3" json:"remoteJid,omitempty"`
	FromMe        bool                   `protobuf:"varint,2,opt,name=fromMe,proto3" json:"fromMe,omitempty"`
	Id            string                 `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageMetadata) Reset() {
	*x = MessageMetadata{}
	mi := &file_api_proto_wa_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageMetadata) ProtoMessage() {}

func (x *MessageMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_wa_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageMetadata.ProtoReflect.Descriptor instead.
func (*MessageMetadata) Descriptor() ([]byte, []int) {
	return file_api_proto_wa_proto_rawDescGZIP(), []int{5}
}

func (x *MessageMetadata) GetRemoteJid() string {
	if x != nil {
		return x.RemoteJid
	}
	return ""
}

func (x *MessageMetadata) GetFromMe() bool {
	if x != nil {
		return x.FromMe
	}
	return false
}

func (x *MessageMetadata) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ImageMessage struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Caption       *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=caption,proto3" json:"caption,omitempty"`
	MimeType      string                  `protobuf:"bytes,2,opt,name=mimeType,proto3" json:"mimeType,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImageMessage) Reset() {
	*x = ImageMessage{}
	mi := &file_api_proto_wa_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageMessage) ProtoMessage() {}

func (x *ImageMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_wa_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageMessage.ProtoReflect.Descriptor instead.
func (*ImageMessage) Descriptor() ([]byte, []int) {
	return file_api_proto_wa_proto_rawDescGZIP(), []int{6}
}

func (x *ImageMessage) GetCaption() *wrapperspb.StringValue {
	if x != nil {
		return x.Caption
	}
	return nil
}

func (x *ImageMessage) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

type QuotedMessage struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Conversation  *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=conversation,proto3" json:"conversation,omitempty"`
	ImageMessage  *ImageMessage           `protobuf:"bytes,2,opt,name=imageMessage,proto3,oneof" json:"imageMessage,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *QuotedMessage) Reset() {
	*x = QuotedMessage{}
	mi := &file_api_proto_wa_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QuotedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuotedMessage) ProtoMessage() {}

func (x *QuotedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_wa_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuotedMessage.ProtoReflect.Descriptor instead.
func (*QuotedMessage) Descriptor() ([]byte, []int) {
	return file_api_proto_wa_proto_rawDescGZIP(), []int{7}
}

func (x *QuotedMessage) GetConversation() *wrapperspb.StringValue {
	if x != nil {
		return x.Conversation
	}
	return nil
}

func (x *QuotedMessage) GetImageMessage() *ImageMessage {
	if x != nil {
		return x.ImageMessage
	}
	return nil
}

type ContextInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StanzaId      string                 `protobuf:"bytes,1,opt,name=stanzaId,proto3" json:"stanzaId,omitempty"`
	Participant   string                 `protobuf:"bytes,2,opt,name=participant,proto3" json:"participant,omitempty"`
	QuotedMessage *QuotedMessage         `protobuf:"bytes,3,opt,name=quotedMessage,proto3,oneof" json:"quotedMessage,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ContextInfo) Reset() {
	*x = ContextInfo{}
	mi := &file_api_proto_wa_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContextInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextInfo) ProtoMessage() {}

func (x *ContextInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_wa_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextInfo.ProtoReflect.Descriptor instead.
func (*ContextInfo) Descriptor() ([]byte, []int) {
	return file_api_proto_wa_proto_rawDescGZIP(), []int{8}
}

func (x *ContextInfo) GetStanzaId() string {
	if x != nil {
		return x.StanzaId
	}
	return ""
}

func (x *ContextInfo) GetParticipant() string {
	if x != nil {
		return x.Participant
	}
	return ""
}

func (x *ContextInfo) GetQuotedMessage() *QuotedMessage {
	if x != nil {
		return x.QuotedMessage
	}
	return nil
}

type FileMetadataRequest struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	FileName      string                  `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`             // File name
	Caption       *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=caption,proto3" json:"caption,omitempty"`               // Caption of the file (optional)
	ContextInfo   *ContextInfo            `protobuf:"bytes,3,opt,name=contextInfo,proto3,oneof" json:"contextInfo,omitempty"` // Context information
	Metadata      *MessageMetadata        `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`             // Metadata for the file
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileMetadataRequest) Reset() {
	*x = FileMetadataRequest{}
	mi := &file_api_proto_wa_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMetadataRequest) ProtoMessage() {}

func (x *FileMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_wa_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMetadataRequest.ProtoReflect.Descriptor instead.
func (*FileMetadataRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_wa_proto_rawDescGZIP(), []int{9}
}

func (x *FileMetadataRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileMetadataRequest) GetCaption() *wrapperspb.StringValue {
	if x != nil {
		return x.Caption
	}
	return nil
}

func (x *FileMetadataRequest) GetContextInfo() *ContextInfo {
	if x != nil {
		return x.ContextInfo
	}
	return nil
}

func (x *FileMetadataRequest) GetMetadata() *MessageMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type FileMetadataResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // Status of the file upload
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileMetadataResponse) Reset() {
	*x = FileMetadataResponse{}
	mi := &file_api_proto_wa_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMetadataResponse) ProtoMessage() {}

func (x *FileMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_wa_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMetadataResponse.ProtoReflect.Descriptor instead.
func (*FileMetadataResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_wa_proto_rawDescGZIP(), []int{10}
}

func (x *FileMetadataResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_api_proto_wa_proto protoreflect.FileDescriptor

var file_api_proto_wa_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x12, 0x54,
	0x65, 0x78, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9a, 0x01, 0x0a, 0x1a, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x34, 0x0a, 0x0b,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x32, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2f, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2a, 0x0a, 0x12, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x22, 0x4b, 0x0a, 0x13, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69,
	0x6c, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x22, 0x57, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4a, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4a, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x72, 0x6f, 0x6d, 0x4d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x66, 0x72, 0x6f, 0x6d, 0x4d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x62, 0x0a, 0x0c, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x63, 0x61, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0xa0, 0x01,
	0x0a, 0x0d, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x40, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3c, 0x0a, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0c,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42,
	0x0f, 0x0a, 0x0d, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x9e, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x6e, 0x7a, 0x61, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x61, 0x6e, 0x7a, 0x61, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x3f,
	0x0a, 0x0d, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75,
	0x6f, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x71,
	0x75, 0x6f, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42,
	0x10, 0x0a, 0x0e, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0xe8, 0x01, 0x0a, 0x13, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2e, 0x0a, 0x14,
	0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xcf, 0x02, 0x0a,
	0x0f, 0x57, 0x68, 0x61, 0x74, 0x73, 0x41, 0x70, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4a, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x78, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x17,
	0x53, 0x65, 0x6e, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01,
	0x12, 0x4c, 0x0a, 0x11, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d,
	0x5a, 0x0b, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_proto_wa_proto_rawDescOnce sync.Once
	file_api_proto_wa_proto_rawDescData []byte
)

func file_api_proto_wa_proto_rawDescGZIP() []byte {
	file_api_proto_wa_proto_rawDescOnce.Do(func() {
		file_api_proto_wa_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_proto_wa_proto_rawDesc), len(file_api_proto_wa_proto_rawDesc)))
	})
	return file_api_proto_wa_proto_rawDescData
}

var file_api_proto_wa_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_proto_wa_proto_goTypes = []any{
	(*TextMessageRequest)(nil),         // 0: proto.TextMessageRequest
	(*ExtendedTextMessageRequest)(nil), // 1: proto.ExtendedTextMessageRequest
	(*CommonMessageResponse)(nil),      // 2: proto.CommonMessageResponse
	(*MediaUploadRequest)(nil),         // 3: proto.MediaUploadRequest
	(*MediaUploadResponse)(nil),        // 4: proto.MediaUploadResponse
	(*MessageMetadata)(nil),            // 5: proto.MessageMetadata
	(*ImageMessage)(nil),               // 6: proto.ImageMessage
	(*QuotedMessage)(nil),              // 7: proto.QuotedMessage
	(*ContextInfo)(nil),                // 8: proto.ContextInfo
	(*FileMetadataRequest)(nil),        // 9: proto.FileMetadataRequest
	(*FileMetadataResponse)(nil),       // 10: proto.FileMetadataResponse
	(*wrapperspb.StringValue)(nil),     // 11: google.protobuf.StringValue
}
var file_api_proto_wa_proto_depIdxs = []int32{
	5,  // 0: proto.TextMessageRequest.metadata:type_name -> proto.MessageMetadata
	8,  // 1: proto.ExtendedTextMessageRequest.contextInfo:type_name -> proto.ContextInfo
	5,  // 2: proto.ExtendedTextMessageRequest.metadata:type_name -> proto.MessageMetadata
	11, // 3: proto.ImageMessage.caption:type_name -> google.protobuf.StringValue
	11, // 4: proto.QuotedMessage.conversation:type_name -> google.protobuf.StringValue
	6,  // 5: proto.QuotedMessage.imageMessage:type_name -> proto.ImageMessage
	7,  // 6: proto.ContextInfo.quotedMessage:type_name -> proto.QuotedMessage
	11, // 7: proto.FileMetadataRequest.caption:type_name -> google.protobuf.StringValue
	8,  // 8: proto.FileMetadataRequest.contextInfo:type_name -> proto.ContextInfo
	5,  // 9: proto.FileMetadataRequest.metadata:type_name -> proto.MessageMetadata
	0,  // 10: proto.WhatsAppService.SendTextMessage:input_type -> proto.TextMessageRequest
	1,  // 11: proto.WhatsAppService.SendExtendedTextMessage:input_type -> proto.ExtendedTextMessageRequest
	3,  // 12: proto.WhatsAppService.UploadMedia:input_type -> proto.MediaUploadRequest
	9,  // 13: proto.WhatsAppService.StoreFileMetadata:input_type -> proto.FileMetadataRequest
	2,  // 14: proto.WhatsAppService.SendTextMessage:output_type -> proto.CommonMessageResponse
	2,  // 15: proto.WhatsAppService.SendExtendedTextMessage:output_type -> proto.CommonMessageResponse
	4,  // 16: proto.WhatsAppService.UploadMedia:output_type -> proto.MediaUploadResponse
	10, // 17: proto.WhatsAppService.StoreFileMetadata:output_type -> proto.FileMetadataResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_api_proto_wa_proto_init() }
func file_api_proto_wa_proto_init() {
	if File_api_proto_wa_proto != nil {
		return
	}
	file_api_proto_wa_proto_msgTypes[7].OneofWrappers = []any{}
	file_api_proto_wa_proto_msgTypes[8].OneofWrappers = []any{}
	file_api_proto_wa_proto_msgTypes[9].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_proto_wa_proto_rawDesc), len(file_api_proto_wa_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_wa_proto_goTypes,
		DependencyIndexes: file_api_proto_wa_proto_depIdxs,
		MessageInfos:      file_api_proto_wa_proto_msgTypes,
	}.Build()
	File_api_proto_wa_proto = out.File
	file_api_proto_wa_proto_goTypes = nil
	file_api_proto_wa_proto_depIdxs = nil
}
