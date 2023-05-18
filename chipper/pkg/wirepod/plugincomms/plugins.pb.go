// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: plugins.proto

package plugincomms

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Intent       string            `protobuf:"bytes,1,opt,name=intent,proto3" json:"intent,omitempty"`
	IntentParams map[string]string `protobuf:"bytes,2,rep,name=intent_params,json=intentParams,proto3" json:"intent_params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Esn          string            `protobuf:"bytes,3,opt,name=esn,proto3" json:"esn,omitempty"`
	Ip           string            `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`
	Guid         string            `protobuf:"bytes,5,opt,name=guid,proto3" json:"guid,omitempty"`
}

func (x *PluginRequest) Reset() {
	*x = PluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginRequest) ProtoMessage() {}

func (x *PluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginRequest.ProtoReflect.Descriptor instead.
func (*PluginRequest) Descriptor() ([]byte, []int) {
	return file_plugins_proto_rawDescGZIP(), []int{0}
}

func (x *PluginRequest) GetIntent() string {
	if x != nil {
		return x.Intent
	}
	return ""
}

func (x *PluginRequest) GetIntentParams() map[string]string {
	if x != nil {
		return x.IntentParams
	}
	return nil
}

func (x *PluginRequest) GetEsn() string {
	if x != nil {
		return x.Esn
	}
	return ""
}

func (x *PluginRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *PluginRequest) GetGuid() string {
	if x != nil {
		return x.Guid
	}
	return ""
}

type PluginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsRunning        bool              `protobuf:"varint,1,opt,name=is_running,json=isRunning,proto3" json:"is_running,omitempty"`
	Intent           string            `protobuf:"bytes,2,opt,name=intent,proto3" json:"intent,omitempty"`
	IntentParams     map[string]string `protobuf:"bytes,3,rep,name=intent_params,json=intentParams,proto3" json:"intent_params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IsKnowledgeGraph bool              `protobuf:"varint,4,opt,name=is_knowledge_graph,json=isKnowledgeGraph,proto3" json:"is_knowledge_graph,omitempty"`
	SpeechOutput     string            `protobuf:"bytes,5,opt,name=speech_output,json=speechOutput,proto3" json:"speech_output,omitempty"`
}

func (x *PluginResponse) Reset() {
	*x = PluginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginResponse) ProtoMessage() {}

func (x *PluginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginResponse.ProtoReflect.Descriptor instead.
func (*PluginResponse) Descriptor() ([]byte, []int) {
	return file_plugins_proto_rawDescGZIP(), []int{1}
}

func (x *PluginResponse) GetIsRunning() bool {
	if x != nil {
		return x.IsRunning
	}
	return false
}

func (x *PluginResponse) GetIntent() string {
	if x != nil {
		return x.Intent
	}
	return ""
}

func (x *PluginResponse) GetIntentParams() map[string]string {
	if x != nil {
		return x.IntentParams
	}
	return nil
}

func (x *PluginResponse) GetIsKnowledgeGraph() bool {
	if x != nil {
		return x.IsKnowledgeGraph
	}
	return false
}

func (x *PluginResponse) GetSpeechOutput() string {
	if x != nil {
		return x.SpeechOutput
	}
	return ""
}

type ConnTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Test bool `protobuf:"varint,1,opt,name=test,proto3" json:"test,omitempty"`
}

func (x *ConnTestRequest) Reset() {
	*x = ConnTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnTestRequest) ProtoMessage() {}

func (x *ConnTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnTestRequest.ProtoReflect.Descriptor instead.
func (*ConnTestRequest) Descriptor() ([]byte, []int) {
	return file_plugins_proto_rawDescGZIP(), []int{2}
}

func (x *ConnTestRequest) GetTest() bool {
	if x != nil {
		return x.Test
	}
	return false
}

type ConnTestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Test bool `protobuf:"varint,1,opt,name=test,proto3" json:"test,omitempty"`
}

func (x *ConnTestResponse) Reset() {
	*x = ConnTestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnTestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnTestResponse) ProtoMessage() {}

func (x *ConnTestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnTestResponse.ProtoReflect.Descriptor instead.
func (*ConnTestResponse) Descriptor() ([]byte, []int) {
	return file_plugins_proto_rawDescGZIP(), []int{3}
}

func (x *ConnTestResponse) GetTest() bool {
	if x != nil {
		return x.Test
	}
	return false
}

type AddPluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddPluginRequest) Reset() {
	*x = AddPluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPluginRequest) ProtoMessage() {}

func (x *AddPluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPluginRequest.ProtoReflect.Descriptor instead.
func (*AddPluginRequest) Descriptor() ([]byte, []int) {
	return file_plugins_proto_rawDescGZIP(), []int{4}
}

func (x *AddPluginRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AddPluginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AddPluginResponse) Reset() {
	*x = AddPluginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPluginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPluginResponse) ProtoMessage() {}

func (x *AddPluginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPluginResponse.ProtoReflect.Descriptor instead.
func (*AddPluginResponse) Descriptor() ([]byte, []int) {
	return file_plugins_proto_rawDescGZIP(), []int{5}
}

func (x *AddPluginResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AddPluginResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type RemovePluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemovePluginRequest) Reset() {
	*x = RemovePluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePluginRequest) ProtoMessage() {}

func (x *RemovePluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePluginRequest.ProtoReflect.Descriptor instead.
func (*RemovePluginRequest) Descriptor() ([]byte, []int) {
	return file_plugins_proto_rawDescGZIP(), []int{6}
}

func (x *RemovePluginRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RemovePluginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *RemovePluginResponse) Reset() {
	*x = RemovePluginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePluginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePluginResponse) ProtoMessage() {}

func (x *RemovePluginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePluginResponse.ProtoReflect.Descriptor instead.
func (*RemovePluginResponse) Descriptor() ([]byte, []int) {
	return file_plugins_proto_rawDescGZIP(), []int{7}
}

func (x *RemovePluginResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RemovePluginResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ListPluginsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []uint32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *ListPluginsResponse) Reset() {
	*x = ListPluginsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPluginsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPluginsResponse) ProtoMessage() {}

func (x *ListPluginsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPluginsResponse.ProtoReflect.Descriptor instead.
func (*ListPluginsResponse) Descriptor() ([]byte, []int) {
	return file_plugins_proto_rawDescGZIP(), []int{8}
}

func (x *ListPluginsResponse) GetIds() []uint32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

var File_plugins_proto protoreflect.FileDescriptor

var file_plugins_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x01, 0x0a, 0x0d, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x4d, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x49,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x73, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x73, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x67, 0x75, 0x69, 0x64, 0x1a, 0x3f, 0x0a, 0x11, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xab, 0x02, 0x0a, 0x0e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x72,
	0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73,
	0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x4e, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x2c, 0x0a, 0x12, 0x69, 0x73, 0x5f, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x5f,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x4b,
	0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x23, 0x0a,
	0x0d, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x1a, 0x3f, 0x0a, 0x11, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x25, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x74, 0x65, 0x73, 0x74, 0x22, 0x26, 0x0a, 0x10, 0x43, 0x6f,
	0x6e, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x74, 0x65,
	0x73, 0x74, 0x22, 0x22, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x25, 0x0a, 0x13, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x46, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x2b, 0x0a, 0x13, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x42, 0x02,
	0x10, 0x01, 0x52, 0x03, 0x69, 0x64, 0x73, 0x32, 0xf1, 0x02, 0x0a, 0x0d, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a,
	0x08, 0x43, 0x6f, 0x6e, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x43, 0x6f,
	0x6e, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x43, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x19, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x1c, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x65, 0x72, 0x63, 0x72, 0x65,
	0x31, 0x32, 0x33, 0x2f, 0x63, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x77, 0x69, 0x72, 0x65, 0x70, 0x6f, 0x64, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x63, 0x6f,
	0x6d, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugins_proto_rawDescOnce sync.Once
	file_plugins_proto_rawDescData = file_plugins_proto_rawDesc
)

func file_plugins_proto_rawDescGZIP() []byte {
	file_plugins_proto_rawDescOnce.Do(func() {
		file_plugins_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugins_proto_rawDescData)
	})
	return file_plugins_proto_rawDescData
}

var file_plugins_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_plugins_proto_goTypes = []interface{}{
	(*PluginRequest)(nil),        // 0: plugins.PluginRequest
	(*PluginResponse)(nil),       // 1: plugins.PluginResponse
	(*ConnTestRequest)(nil),      // 2: plugins.ConnTestRequest
	(*ConnTestResponse)(nil),     // 3: plugins.ConnTestResponse
	(*AddPluginRequest)(nil),     // 4: plugins.AddPluginRequest
	(*AddPluginResponse)(nil),    // 5: plugins.AddPluginResponse
	(*RemovePluginRequest)(nil),  // 6: plugins.RemovePluginRequest
	(*RemovePluginResponse)(nil), // 7: plugins.RemovePluginResponse
	(*ListPluginsResponse)(nil),  // 8: plugins.ListPluginsResponse
	nil,                          // 9: plugins.PluginRequest.IntentParamsEntry
	nil,                          // 10: plugins.PluginResponse.IntentParamsEntry
	(*emptypb.Empty)(nil),        // 11: google.protobuf.Empty
}
var file_plugins_proto_depIdxs = []int32{
	9,  // 0: plugins.PluginRequest.intent_params:type_name -> plugins.PluginRequest.IntentParamsEntry
	10, // 1: plugins.PluginResponse.intent_params:type_name -> plugins.PluginResponse.IntentParamsEntry
	0,  // 2: plugins.PluginService.ProcessPlugin:input_type -> plugins.PluginRequest
	2,  // 3: plugins.PluginService.ConnTest:input_type -> plugins.ConnTestRequest
	4,  // 4: plugins.PluginService.AddPlugin:input_type -> plugins.AddPluginRequest
	6,  // 5: plugins.PluginService.RemovePlugin:input_type -> plugins.RemovePluginRequest
	11, // 6: plugins.PluginService.ListPlugins:input_type -> google.protobuf.Empty
	1,  // 7: plugins.PluginService.ProcessPlugin:output_type -> plugins.PluginResponse
	3,  // 8: plugins.PluginService.ConnTest:output_type -> plugins.ConnTestResponse
	3,  // 9: plugins.PluginService.AddPlugin:output_type -> plugins.ConnTestResponse
	7,  // 10: plugins.PluginService.RemovePlugin:output_type -> plugins.RemovePluginResponse
	8,  // 11: plugins.PluginService.ListPlugins:output_type -> plugins.ListPluginsResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_plugins_proto_init() }
func file_plugins_proto_init() {
	if File_plugins_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugins_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginRequest); i {
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
		file_plugins_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginResponse); i {
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
		file_plugins_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnTestRequest); i {
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
		file_plugins_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnTestResponse); i {
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
		file_plugins_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPluginRequest); i {
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
		file_plugins_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPluginResponse); i {
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
		file_plugins_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePluginRequest); i {
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
		file_plugins_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePluginResponse); i {
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
		file_plugins_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPluginsResponse); i {
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
			RawDescriptor: file_plugins_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_plugins_proto_goTypes,
		DependencyIndexes: file_plugins_proto_depIdxs,
		MessageInfos:      file_plugins_proto_msgTypes,
	}.Build()
	File_plugins_proto = out.File
	file_plugins_proto_rawDesc = nil
	file_plugins_proto_goTypes = nil
	file_plugins_proto_depIdxs = nil
}
