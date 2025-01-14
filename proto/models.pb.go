// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: models.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username          string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email             string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Karma             int64                  `protobuf:"varint,4,opt,name=karma,proto3" json:"karma,omitempty"`
	CreatedAt         *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	LastLogin         *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=last_login,json=lastLogin,proto3" json:"last_login,omitempty"`
	CreatedSubreddits []*Subreddit           `protobuf:"bytes,7,rep,name=created_subreddits,json=createdSubreddits,proto3" json:"created_subreddits,omitempty"`
	Subscriptions     []*Subreddit           `protobuf:"bytes,8,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_models_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetKarma() int64 {
	if x != nil {
		return x.Karma
	}
	return 0
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetLastLogin() *timestamppb.Timestamp {
	if x != nil {
		return x.LastLogin
	}
	return nil
}

func (x *User) GetCreatedSubreddits() []*Subreddit {
	if x != nil {
		return x.CreatedSubreddits
	}
	return nil
}

func (x *User) GetSubscriptions() []*Subreddit {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

// subreddit and relationship
type Subreddit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description     string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatorId       uint64                 `protobuf:"varint,4,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	SubscriberCount int64                  `protobuf:"varint,5,opt,name=subscriber_count,json=subscriberCount,proto3" json:"subscriber_count,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Creator         *User                  `protobuf:"bytes,9,opt,name=creator,proto3" json:"creator,omitempty"`
	PostCount       int64                  `protobuf:"varint,10,opt,name=post_count,json=postCount,proto3" json:"post_count,omitempty"`
}

func (x *Subreddit) Reset() {
	*x = Subreddit{}
	mi := &file_models_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Subreddit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subreddit) ProtoMessage() {}

func (x *Subreddit) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subreddit.ProtoReflect.Descriptor instead.
func (*Subreddit) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{1}
}

func (x *Subreddit) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Subreddit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Subreddit) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Subreddit) GetCreatorId() uint64 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *Subreddit) GetSubscriberCount() int64 {
	if x != nil {
		return x.SubscriberCount
	}
	return 0
}

func (x *Subreddit) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Subreddit) GetCreator() *User {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *Subreddit) GetPostCount() int64 {
	if x != nil {
		return x.PostCount
	}
	return 0
}

type UserSubredditSubscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       uint64                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SubredditId  uint64                 `protobuf:"varint,2,opt,name=subreddit_id,json=subredditId,proto3" json:"subreddit_id,omitempty"`
	SubscribedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=subscribed_at,json=subscribedAt,proto3" json:"subscribed_at,omitempty"`
	User         *User                  `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	Subreddit    *Subreddit             `protobuf:"bytes,5,opt,name=subreddit,proto3" json:"subreddit,omitempty"`
}

func (x *UserSubredditSubscription) Reset() {
	*x = UserSubredditSubscription{}
	mi := &file_models_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserSubredditSubscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSubredditSubscription) ProtoMessage() {}

func (x *UserSubredditSubscription) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSubredditSubscription.ProtoReflect.Descriptor instead.
func (*UserSubredditSubscription) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{2}
}

func (x *UserSubredditSubscription) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserSubredditSubscription) GetSubredditId() uint64 {
	if x != nil {
		return x.SubredditId
	}
	return 0
}

func (x *UserSubredditSubscription) GetSubscribedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.SubscribedAt
	}
	return nil
}

func (x *UserSubredditSubscription) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserSubredditSubscription) GetSubreddit() *Subreddit {
	if x != nil {
		return x.Subreddit
	}
	return nil
}

// direct messaging
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text      string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	FromId    uint64                 `protobuf:"varint,3,opt,name=from_id,json=fromId,proto3" json:"from_id,omitempty"`
	ToId      uint64                 `protobuf:"varint,4,opt,name=to_id,json=toId,proto3" json:"to_id,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_models_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[3]
	if x != nil {
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
	return file_models_proto_rawDescGZIP(), []int{3}
}

func (x *Message) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Message) GetFromId() uint64 {
	if x != nil {
		return x.FromId
	}
	return 0
}

func (x *Message) GetToId() uint64 {
	if x != nil {
		return x.ToId
	}
	return 0
}

func (x *Message) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title        string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content      string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	AuthorId     uint64                 `protobuf:"varint,4,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	SubredditId  uint64                 `protobuf:"varint,5,opt,name=subreddit_id,json=subredditId,proto3" json:"subreddit_id,omitempty"`
	Score        int64                  `protobuf:"varint,6,opt,name=score,proto3" json:"score,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Author       *User                  `protobuf:"bytes,9,opt,name=author,proto3" json:"author,omitempty"`
	Subreddit    *Subreddit             `protobuf:"bytes,10,opt,name=subreddit,proto3" json:"subreddit,omitempty"`
	Comments     []*Comment             `protobuf:"bytes,11,rep,name=comments,proto3" json:"comments,omitempty"`
	CommentCount int64                  `protobuf:"varint,12,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	mi := &file_models_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{4}
}

func (x *Post) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Post) GetAuthorId() uint64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *Post) GetSubredditId() uint64 {
	if x != nil {
		return x.SubredditId
	}
	return 0
}

func (x *Post) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Post) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Post) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Post) GetAuthor() *User {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Post) GetSubreddit() *Subreddit {
	if x != nil {
		return x.Subreddit
	}
	return nil
}

func (x *Post) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *Post) GetCommentCount() int64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Content   string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	UserId    uint64                 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PostId    uint64                 `protobuf:"varint,4,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	ParentId  *uint64                `protobuf:"varint,5,opt,name=parent_id,json=parentId,proto3,oneof" json:"parent_id,omitempty"`
	Children  []*Comment             `protobuf:"bytes,6,rep,name=children,proto3" json:"children,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Author    *User                  `protobuf:"bytes,9,opt,name=author,proto3" json:"author,omitempty"`
	Post      *Post                  `protobuf:"bytes,10,opt,name=post,proto3" json:"post,omitempty"`
	Votes     int32                  `protobuf:"varint,11,opt,name=votes,proto3" json:"votes,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	mi := &file_models_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{5}
}

func (x *Comment) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Comment) GetPostId() uint64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *Comment) GetParentId() uint64 {
	if x != nil && x.ParentId != nil {
		return *x.ParentId
	}
	return 0
}

func (x *Comment) GetChildren() []*Comment {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *Comment) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Comment) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Comment) GetAuthor() *User {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Comment) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

func (x *Comment) GetVotes() int32 {
	if x != nil {
		return x.Votes
	}
	return 0
}

var File_models_proto protoreflect.FileDescriptor

var file_models_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6b, 0x61, 0x72, 0x6d, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x6b, 0x61, 0x72, 0x6d, 0x61, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x3f, 0x0a,
	0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64,
	0x69, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x52, 0x11, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x53, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x73, 0x12, 0x36,
	0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x75,
	0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x9c, 0x02, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x72, 0x65,
	0x64, 0x64, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x25, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xe9, 0x01, 0x0a, 0x19, 0x55, 0x73, 0x65, 0x72, 0x53, 0x75,
	0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x49, 0x64, 0x12,
	0x3f, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x2e, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x75, 0x62,
	0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x52, 0x09, 0x73, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69,
	0x74, 0x22, 0x96, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x6f,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x6f, 0x49, 0x64, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xb8, 0x03, 0x0a, 0x04, 0x50,
	0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69,
	0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x23, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x52, 0x09, 0x73, 0x75, 0x62, 0x72, 0x65,
	0x64, 0x64, 0x69, 0x74, 0x12, 0x2a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x93, 0x03, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x48, 0x00, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x2a, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x23, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x73,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x36, 0x5a, 0x34, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x74, 0x69, 0x6e, 0x67,
	0x6f, 0x79, 0x61, 0x6c, 0x30, 0x39, 0x39, 0x36, 0x2f, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x2d,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_proto_rawDescOnce sync.Once
	file_models_proto_rawDescData = file_models_proto_rawDesc
)

func file_models_proto_rawDescGZIP() []byte {
	file_models_proto_rawDescOnce.Do(func() {
		file_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_proto_rawDescData)
	})
	return file_models_proto_rawDescData
}

var file_models_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_models_proto_goTypes = []any{
	(*User)(nil),                      // 0: proto.User
	(*Subreddit)(nil),                 // 1: proto.Subreddit
	(*UserSubredditSubscription)(nil), // 2: proto.UserSubredditSubscription
	(*Message)(nil),                   // 3: proto.Message
	(*Post)(nil),                      // 4: proto.Post
	(*Comment)(nil),                   // 5: proto.Comment
	(*timestamppb.Timestamp)(nil),     // 6: google.protobuf.Timestamp
}
var file_models_proto_depIdxs = []int32{
	6,  // 0: proto.User.created_at:type_name -> google.protobuf.Timestamp
	6,  // 1: proto.User.last_login:type_name -> google.protobuf.Timestamp
	1,  // 2: proto.User.created_subreddits:type_name -> proto.Subreddit
	1,  // 3: proto.User.subscriptions:type_name -> proto.Subreddit
	6,  // 4: proto.Subreddit.created_at:type_name -> google.protobuf.Timestamp
	0,  // 5: proto.Subreddit.creator:type_name -> proto.User
	6,  // 6: proto.UserSubredditSubscription.subscribed_at:type_name -> google.protobuf.Timestamp
	0,  // 7: proto.UserSubredditSubscription.user:type_name -> proto.User
	1,  // 8: proto.UserSubredditSubscription.subreddit:type_name -> proto.Subreddit
	6,  // 9: proto.Message.created_at:type_name -> google.protobuf.Timestamp
	6,  // 10: proto.Post.created_at:type_name -> google.protobuf.Timestamp
	6,  // 11: proto.Post.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 12: proto.Post.author:type_name -> proto.User
	1,  // 13: proto.Post.subreddit:type_name -> proto.Subreddit
	5,  // 14: proto.Post.comments:type_name -> proto.Comment
	5,  // 15: proto.Comment.children:type_name -> proto.Comment
	6,  // 16: proto.Comment.created_at:type_name -> google.protobuf.Timestamp
	6,  // 17: proto.Comment.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 18: proto.Comment.author:type_name -> proto.User
	4,  // 19: proto.Comment.post:type_name -> proto.Post
	20, // [20:20] is the sub-list for method output_type
	20, // [20:20] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
}

func init() { file_models_proto_init() }
func file_models_proto_init() {
	if File_models_proto != nil {
		return
	}
	file_models_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_proto_goTypes,
		DependencyIndexes: file_models_proto_depIdxs,
		MessageInfos:      file_models_proto_msgTypes,
	}.Build()
	File_models_proto = out.File
	file_models_proto_rawDesc = nil
	file_models_proto_goTypes = nil
	file_models_proto_depIdxs = nil
}
