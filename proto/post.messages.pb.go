// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: post.messages.proto

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

type CreatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content     string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	AuthorId    uint64 `protobuf:"varint,4,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	SubredditId uint64 `protobuf:"varint,5,opt,name=subreddit_id,json=subredditId,proto3" json:"subreddit_id,omitempty"`
}

func (x *CreatePostRequest) Reset() {
	*x = CreatePostRequest{}
	mi := &file_post_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostRequest) ProtoMessage() {}

func (x *CreatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostRequest.ProtoReflect.Descriptor instead.
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return file_post_messages_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePostRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreatePostRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreatePostRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreatePostRequest) GetAuthorId() uint64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *CreatePostRequest) GetSubredditId() uint64 {
	if x != nil {
		return x.SubredditId
	}
	return 0
}

type CreatePostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreatePostResponse) Reset() {
	*x = CreatePostResponse{}
	mi := &file_post_messages_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostResponse) ProtoMessage() {}

func (x *CreatePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_messages_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostResponse.ProtoReflect.Descriptor instead.
func (*CreatePostResponse) Descriptor() ([]byte, []int) {
	return file_post_messages_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePostResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetPostsBySubredditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	SubredditId uint64 `protobuf:"varint,2,opt,name=subreddit_id,json=subredditId,proto3" json:"subreddit_id,omitempty"`
	Limit       int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset      int32  `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetPostsBySubredditRequest) Reset() {
	*x = GetPostsBySubredditRequest{}
	mi := &file_post_messages_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPostsBySubredditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostsBySubredditRequest) ProtoMessage() {}

func (x *GetPostsBySubredditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_messages_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostsBySubredditRequest.ProtoReflect.Descriptor instead.
func (*GetPostsBySubredditRequest) Descriptor() ([]byte, []int) {
	return file_post_messages_proto_rawDescGZIP(), []int{2}
}

func (x *GetPostsBySubredditRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetPostsBySubredditRequest) GetSubredditId() uint64 {
	if x != nil {
		return x.SubredditId
	}
	return 0
}

func (x *GetPostsBySubredditRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetPostsBySubredditRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetPostsBySubredditResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *GetPostsBySubredditResponse) Reset() {
	*x = GetPostsBySubredditResponse{}
	mi := &file_post_messages_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPostsBySubredditResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostsBySubredditResponse) ProtoMessage() {}

func (x *GetPostsBySubredditResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_messages_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostsBySubredditResponse.ProtoReflect.Descriptor instead.
func (*GetPostsBySubredditResponse) Descriptor() ([]byte, []int) {
	return file_post_messages_proto_rawDescGZIP(), []int{3}
}

func (x *GetPostsBySubredditResponse) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type GetPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	PostId uint64 `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
}

func (x *GetPostRequest) Reset() {
	*x = GetPostRequest{}
	mi := &file_post_messages_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostRequest) ProtoMessage() {}

func (x *GetPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_messages_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostRequest.ProtoReflect.Descriptor instead.
func (*GetPostRequest) Descriptor() ([]byte, []int) {
	return file_post_messages_proto_rawDescGZIP(), []int{4}
}

func (x *GetPostRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetPostRequest) GetPostId() uint64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

type GetPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *GetPostResponse) Reset() {
	*x = GetPostResponse{}
	mi := &file_post_messages_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostResponse) ProtoMessage() {}

func (x *GetPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_messages_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostResponse.ProtoReflect.Descriptor instead.
func (*GetPostResponse) Descriptor() ([]byte, []int) {
	return file_post_messages_proto_rawDescGZIP(), []int{5}
}

func (x *GetPostResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type GetPostByUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Limit  int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32  `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetPostByUserRequest) Reset() {
	*x = GetPostByUserRequest{}
	mi := &file_post_messages_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPostByUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostByUserRequest) ProtoMessage() {}

func (x *GetPostByUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_messages_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostByUserRequest.ProtoReflect.Descriptor instead.
func (*GetPostByUserRequest) Descriptor() ([]byte, []int) {
	return file_post_messages_proto_rawDescGZIP(), []int{6}
}

func (x *GetPostByUserRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetPostByUserRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetPostByUserRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetPostByUserRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetPostByUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *GetPostByUserResponse) Reset() {
	*x = GetPostByUserResponse{}
	mi := &file_post_messages_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPostByUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostByUserResponse) ProtoMessage() {}

func (x *GetPostByUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_messages_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostByUserResponse.ProtoReflect.Descriptor instead.
func (*GetPostByUserResponse) Descriptor() ([]byte, []int) {
	return file_post_messages_proto_rawDescGZIP(), []int{7}
}

func (x *GetPostByUserResponse) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type UpdatePostVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	PostId uint64 `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Upvote bool   `protobuf:"varint,3,opt,name=upvote,proto3" json:"upvote,omitempty"`
}

func (x *UpdatePostVoteRequest) Reset() {
	*x = UpdatePostVoteRequest{}
	mi := &file_post_messages_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePostVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostVoteRequest) ProtoMessage() {}

func (x *UpdatePostVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_messages_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostVoteRequest.ProtoReflect.Descriptor instead.
func (*UpdatePostVoteRequest) Descriptor() ([]byte, []int) {
	return file_post_messages_proto_rawDescGZIP(), []int{8}
}

func (x *UpdatePostVoteRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UpdatePostVoteRequest) GetPostId() uint64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *UpdatePostVoteRequest) GetUpvote() bool {
	if x != nil {
		return x.Upvote
	}
	return false
}

type UpdatePostVoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *UpdatePostVoteResponse) Reset() {
	*x = UpdatePostVoteResponse{}
	mi := &file_post_messages_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePostVoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostVoteResponse) ProtoMessage() {}

func (x *UpdatePostVoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_messages_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostVoteResponse.ProtoReflect.Descriptor instead.
func (*UpdatePostVoteResponse) Descriptor() ([]byte, []int) {
	return file_post_messages_proto_rawDescGZIP(), []int{9}
}

func (x *UpdatePostVoteResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_post_messages_proto protoreflect.FileDescriptor

var file_post_messages_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x72, 0x65,
	0x64, 0x64, 0x69, 0x74, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x83, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x42,
	0x79, 0x53, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x72, 0x65,
	0x64, 0x64, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73,
	0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x40, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x73, 0x42, 0x79, 0x53, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x3f, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22,
	0x73, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x22, 0x3a, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a,
	0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73,
	0x22, 0x5e, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x76, 0x6f,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65,
	0x22, 0x2e, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x56, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e,
	0x69, 0x74, 0x69, 0x6e, 0x67, 0x6f, 0x79, 0x61, 0x6c, 0x30, 0x39, 0x39, 0x36, 0x2f, 0x72, 0x65,
	0x64, 0x64, 0x69, 0x74, 0x2d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_post_messages_proto_rawDescOnce sync.Once
	file_post_messages_proto_rawDescData = file_post_messages_proto_rawDesc
)

func file_post_messages_proto_rawDescGZIP() []byte {
	file_post_messages_proto_rawDescOnce.Do(func() {
		file_post_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_post_messages_proto_rawDescData)
	})
	return file_post_messages_proto_rawDescData
}

var file_post_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_post_messages_proto_goTypes = []any{
	(*CreatePostRequest)(nil),           // 0: proto.CreatePostRequest
	(*CreatePostResponse)(nil),          // 1: proto.CreatePostResponse
	(*GetPostsBySubredditRequest)(nil),  // 2: proto.GetPostsBySubredditRequest
	(*GetPostsBySubredditResponse)(nil), // 3: proto.GetPostsBySubredditResponse
	(*GetPostRequest)(nil),              // 4: proto.GetPostRequest
	(*GetPostResponse)(nil),             // 5: proto.GetPostResponse
	(*GetPostByUserRequest)(nil),        // 6: proto.GetPostByUserRequest
	(*GetPostByUserResponse)(nil),       // 7: proto.GetPostByUserResponse
	(*UpdatePostVoteRequest)(nil),       // 8: proto.UpdatePostVoteRequest
	(*UpdatePostVoteResponse)(nil),      // 9: proto.UpdatePostVoteResponse
	(*Post)(nil),                        // 10: proto.Post
}
var file_post_messages_proto_depIdxs = []int32{
	10, // 0: proto.GetPostsBySubredditResponse.posts:type_name -> proto.Post
	10, // 1: proto.GetPostResponse.post:type_name -> proto.Post
	10, // 2: proto.GetPostByUserResponse.posts:type_name -> proto.Post
	3,  // [3:3] is the sub-list for method output_type
	3,  // [3:3] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_post_messages_proto_init() }
func file_post_messages_proto_init() {
	if File_post_messages_proto != nil {
		return
	}
	file_models_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_post_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_post_messages_proto_goTypes,
		DependencyIndexes: file_post_messages_proto_depIdxs,
		MessageInfos:      file_post_messages_proto_msgTypes,
	}.Build()
	File_post_messages_proto = out.File
	file_post_messages_proto_rawDesc = nil
	file_post_messages_proto_goTypes = nil
	file_post_messages_proto_depIdxs = nil
}
