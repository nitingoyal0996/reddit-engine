package models

import (
	"errors"
	"regexp"
	"time"

	"github.com/nitingoyal0996/reddit-clone/proto"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"unique;not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Karma     int64     `gorm:"default:0"`
	LastLogin time.Time `gorm:"autoUpdateTime"`

	// Relationships
	CreatedSubreddits []Subreddit `gorm:"foreignKey:CreatorID"`
	Subscriptions     []Subreddit `gorm:"many2many:user_subreddit_subscriptions"`
}

func (u *User) ToProto() *proto.User {
	userProto := &proto.User{
		Id:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Karma:     u.Karma,
		CreatedAt: timestamppb.New(u.CreatedAt),
		LastLogin: timestamppb.New(u.LastLogin),
	}

	// Convert created subreddits
	if len(u.CreatedSubreddits) > 0 {
		userProto.CreatedSubreddits = make([]*proto.Subreddit, len(u.CreatedSubreddits))
		for i, subreddit := range u.CreatedSubreddits {
			userProto.CreatedSubreddits[i] = subreddit.ToProto()
		}
	}

	// Convert subscriptions
	if len(u.Subscriptions) > 0 {
		userProto.Subscriptions = make([]*proto.Subreddit, len(u.Subscriptions))
		for i, subscription := range u.Subscriptions {
			userProto.Subscriptions[i] = subscription.ToProto()
		}
	}

	return userProto
}

func (u *User) Validate() error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !emailRegex.MatchString(u.Email) {
		return errors.New("invalid email format")
	}
	if len(u.Username) < 3 || len(u.Username) > 50 {
		return errors.New("username must be between 3 and 50 characters")
	}

	return nil
}

func (u *User) HashPassword() error {
	if len(u.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)

	return nil
}

func (u *User) CheckPassword(providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
	return err == nil
}

func (u *User) SafeUser() map[string]interface{} {
	return map[string]interface{}{
		"username": u.Username,
		"email":    u.Email,
		"karma":    u.Karma,
	}
}
