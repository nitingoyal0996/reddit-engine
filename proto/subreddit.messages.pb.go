// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: subreddit.messages.proto

package proto

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

type CreateSubredditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatorId   uint64 `protobuf:"varint,4,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
}

func (x *CreateSubredditRequest) Reset() {
	*x = CreateSubredditRequest{}
	mi := &file_subreddit_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSubredditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSubredditRequest) ProtoMessage() {}

func (x *CreateSubredditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subreddit_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSubredditRequest.ProtoReflect.Descriptor instead.
func (*CreateSubredditRequest) Descriptor() ([]byte, []int) {
	return file_subreddit_messages_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSubredditRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreateSubredditRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateSubredditRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateSubredditRequest) GetCreatorId() uint64 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

type CreateSubredditResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateSubredditResponse) Reset() {
	*x = CreateSubredditResponse{}
	mi := &file_subreddit_messages_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSubredditResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSubredditResponse) ProtoMessage() {}

func (x *CreateSubredditResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subreddit_messages_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSubredditResponse.ProtoReflect.Descriptor instead.
func (*CreateSubredditResponse) Descriptor() ([]byte, []int) {
	return file_subreddit_messages_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSubredditResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetSubredditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetSubredditRequest) Reset() {
	*x = GetSubredditRequest{}
	mi := &file_subreddit_messages_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSubredditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubredditRequest) ProtoMessage() {}

func (x *GetSubredditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subreddit_messages_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubredditRequest.ProtoReflect.Descriptor instead.
func (*GetSubredditRequest) Descriptor() ([]byte, []int) {
	return file_subreddit_messages_proto_rawDescGZIP(), []int{2}
}

func (x *GetSubredditRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetSubredditRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId      uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SubredditId uint64 `protobuf:"varint,3,opt,name=subreddit_id,json=subredditId,proto3" json:"subreddit_id,omitempty"`
}

func (x *SubscriptionRequest) Reset() {
	*x = SubscriptionRequest{}
	mi := &file_subreddit_messages_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionRequest) ProtoMessage() {}

func (x *SubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subreddit_messages_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionRequest.ProtoReflect.Descriptor instead.
func (*SubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_subreddit_messages_proto_rawDescGZIP(), []int{3}
}

func (x *SubscriptionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SubscriptionRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SubscriptionRequest) GetSubredditId() uint64 {
	if x != nil {
		return x.SubredditId
	}
	return 0
}

type SubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SubscriptionResponse) Reset() {
	*x = SubscriptionResponse{}
	mi := &file_subreddit_messages_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionResponse) ProtoMessage() {}

func (x *SubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subreddit_messages_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionResponse.ProtoReflect.Descriptor instead.
func (*SubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_subreddit_messages_proto_rawDescGZIP(), []int{4}
}

func (x *SubscriptionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SubscriptionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UnsubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId      uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SubredditId uint64 `protobuf:"varint,3,opt,name=subreddit_id,json=subredditId,proto3" json:"subreddit_id,omitempty"`
}

func (x *UnsubscribeRequest) Reset() {
	*x = UnsubscribeRequest{}
	mi := &file_subreddit_messages_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnsubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeRequest) ProtoMessage() {}

func (x *UnsubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subreddit_messages_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeRequest.ProtoReflect.Descriptor instead.
func (*UnsubscribeRequest) Descriptor() ([]byte, []int) {
	return file_subreddit_messages_proto_rawDescGZIP(), []int{5}
}

func (x *UnsubscribeRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UnsubscribeRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UnsubscribeRequest) GetSubredditId() uint64 {
	if x != nil {
		return x.SubredditId
	}
	return 0
}

type UnsubscribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UnsubscribeResponse) Reset() {
	*x = UnsubscribeResponse{}
	mi := &file_subreddit_messages_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnsubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeResponse) ProtoMessage() {}

func (x *UnsubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subreddit_messages_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeResponse.ProtoReflect.Descriptor instead.
func (*UnsubscribeResponse) Descriptor() ([]byte, []int) {
	return file_subreddit_messages_proto_rawDescGZIP(), []int{6}
}

func (x *UnsubscribeResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UnsubscribeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Query string `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	Limit int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	mi := &file_subreddit_messages_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subreddit_messages_proto_msgTypes[7]
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
	return file_subreddit_messages_proto_rawDescGZIP(), []int{7}
}

func (x *SearchRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SearchRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subreddits []*Subreddit `protobuf:"bytes,1,rep,name=subreddits,proto3" json:"subreddits,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	mi := &file_subreddit_messages_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subreddit_messages_proto_msgTypes[8]
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
	return file_subreddit_messages_proto_rawDescGZIP(), []int{8}
}

func (x *SearchResponse) GetSubreddits() []*Subreddit {
	if x != nil {
		return x.Subreddits
	}
	return nil
}

type GetUserSubscriptionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserSubscriptionsRequest) Reset() {
	*x = GetUserSubscriptionsRequest{}
	mi := &file_subreddit_messages_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserSubscriptionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserSubscriptionsRequest) ProtoMessage() {}

func (x *GetUserSubscriptionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subreddit_messages_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserSubscriptionsRequest.ProtoReflect.Descriptor instead.
func (*GetUserSubscriptionsRequest) Descriptor() ([]byte, []int) {
	return file_subreddit_messages_proto_rawDescGZIP(), []int{9}
}

func (x *GetUserSubscriptionsRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetUserSubscriptionsRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetUserSubscriptionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subreddits []*Subreddit `protobuf:"bytes,1,rep,name=subreddits,proto3" json:"subreddits,omitempty"`
}

func (x *GetUserSubscriptionsResponse) Reset() {
	*x = GetUserSubscriptionsResponse{}
	mi := &file_subreddit_messages_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserSubscriptionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserSubscriptionsResponse) ProtoMessage() {}

func (x *GetUserSubscriptionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subreddit_messages_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserSubscriptionsResponse.ProtoReflect.Descriptor instead.
func (*GetUserSubscriptionsResponse) Descriptor() ([]byte, []int) {
	return file_subreddit_messages_proto_rawDescGZIP(), []int{10}
}

func (x *GetUserSubscriptionsResponse) GetSubreddits() []*Subreddit {
	if x != nil {
		return x.Subreddits
	}
	return nil
}

var File_subreddit_messages_proto protoreflect.FileDescriptor

var file_subreddit_messages_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x83, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x72, 0x65, 0x64,
	0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x3f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62,
	0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x67, 0x0a, 0x13, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x49, 0x64,
	0x22, 0x4a, 0x0a, 0x14, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x66, 0x0a, 0x12,
	0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64,
	0x69, 0x74, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x13, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x51, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x42, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x72,
	0x65, 0x64, 0x64, 0x69, 0x74, 0x73, 0x22, 0x4c, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x50, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x72,
	0x65, 0x64, 0x64, 0x69, 0x74, 0x73, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x6f, 0x79, 0x61, 0x6c, 0x30,
	0x39, 0x39, 0x36, 0x2f, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x2d, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_subreddit_messages_proto_rawDescOnce sync.Once
	file_subreddit_messages_proto_rawDescData = file_subreddit_messages_proto_rawDesc
)

func file_subreddit_messages_proto_rawDescGZIP() []byte {
	file_subreddit_messages_proto_rawDescOnce.Do(func() {
		file_subreddit_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_subreddit_messages_proto_rawDescData)
	})
	return file_subreddit_messages_proto_rawDescData
}

var file_subreddit_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_subreddit_messages_proto_goTypes = []any{
	(*CreateSubredditRequest)(nil),       // 0: proto.CreateSubredditRequest
	(*CreateSubredditResponse)(nil),      // 1: proto.CreateSubredditResponse
	(*GetSubredditRequest)(nil),          // 2: proto.GetSubredditRequest
	(*SubscriptionRequest)(nil),          // 3: proto.SubscriptionRequest
	(*SubscriptionResponse)(nil),         // 4: proto.SubscriptionResponse
	(*UnsubscribeRequest)(nil),           // 5: proto.UnsubscribeRequest
	(*UnsubscribeResponse)(nil),          // 6: proto.UnsubscribeResponse
	(*SearchRequest)(nil),                // 7: proto.SearchRequest
	(*SearchResponse)(nil),               // 8: proto.SearchResponse
	(*GetUserSubscriptionsRequest)(nil),  // 9: proto.GetUserSubscriptionsRequest
	(*GetUserSubscriptionsResponse)(nil), // 10: proto.GetUserSubscriptionsResponse
	(*Subreddit)(nil),                    // 11: proto.Subreddit
}
var file_subreddit_messages_proto_depIdxs = []int32{
	11, // 0: proto.SearchResponse.subreddits:type_name -> proto.Subreddit
	11, // 1: proto.GetUserSubscriptionsResponse.subreddits:type_name -> proto.Subreddit
	2,  // [2:2] is the sub-list for method output_type
	2,  // [2:2] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_subreddit_messages_proto_init() }
func file_subreddit_messages_proto_init() {
	if File_subreddit_messages_proto != nil {
		return
	}
	file_models_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_subreddit_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_subreddit_messages_proto_goTypes,
		DependencyIndexes: file_subreddit_messages_proto_depIdxs,
		MessageInfos:      file_subreddit_messages_proto_msgTypes,
	}.Build()
	File_subreddit_messages_proto = out.File
	file_subreddit_messages_proto_rawDesc = nil
	file_subreddit_messages_proto_goTypes = nil
	file_subreddit_messages_proto_depIdxs = nil
}
