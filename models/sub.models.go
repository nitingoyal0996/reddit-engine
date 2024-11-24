package models

import (
	"time"

	pb "github.com/nitingoyal0996/reddit-clone/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type Subreddit struct {
	gorm.Model             // This will include ID, CreatedAt, UpdatedAt, DeletedAt
	ID              uint64 `gorm:"primaryKey;autoIncrement"`
	Name            string `gorm:"uniqueIndex;not null"`
	Description     string `gorm:"type:text"`
	CreatorID       uint64 `gorm:"not null"` // Changed to uint to match gorm.Model ID type
	Creator         User   `gorm:"foreignKey:CreatorID"`
	Subscribers     []User `gorm:"many2many:user_subreddit_subscriptions"`
	Posts           []Post `gorm:"foreignKey:SubredditID"`
	SubscriberCount int64  `gorm:"default:0"`
}

// Join table for user subscriptions with additional metadata
type UserSubredditSubscription struct {
	UserID       uint64    `gorm:"primaryKey"`
	SubredditID  uint64    `gorm:"primaryKey"`
	SubscribedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`

	User      User      `gorm:"foreignKey:UserID"`
	Subreddit Subreddit `gorm:"foreignKey:SubredditID"`
}

func (s *Subreddit) ToProto() *pb.Subreddit {
	return &pb.Subreddit{
		Id:              s.ID,
		Name:            s.Name,
		Description:     s.Description,
		CreatorId:       s.CreatorID,
		SubscriberCount: s.SubscriberCount,
		CreatedAt:       timestamppb.New(s.CreatedAt),
		Creator:         s.Creator.ToProto(),
	}
}
