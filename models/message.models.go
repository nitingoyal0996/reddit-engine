package models

import (
	pb "github.com/nitingoyal0996/reddit-clone/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"time"

	"gorm.io/gorm"
)

type Message struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement"`
	Text      string         `gorm:"type:text;not null"`
	FromId    uint64         `gorm:"not null"`
	ToId      uint64         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *Message) ToProto() *pb.Message {
	return &pb.Message{
		Id:        m.ID,
		Text:      m.Text,
		FromId:    m.FromId,
		ToId:      m.ToId,
		CreatedAt: timestamppb.New(m.CreatedAt),
	}
}
