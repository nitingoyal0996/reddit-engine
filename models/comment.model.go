package models

import (
	"time"
)	

type Comment struct {
	ID        uint      `gorm:"primaryKey"`
	PostId    uint      `gorm:"not null"`
	Author    string    `gorm:"not null"`
	Content   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	CommentKarma uint   `gorm:"default:0"`
	IsDeleted bool      `gorm:"default:false"`
}