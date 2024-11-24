package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID        uint64    `gorm:"primaryKey"`
	Content   string    `gorm:"type:text;not null"`
	UserID    uint64    `gorm:"not null"`
	PostID    uint64    `gorm:"not null"`
	ParentID  *uint     // Nullable parent comment ID
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Votes     int32     `gorm:"default:0"`
	// relationships
	Children []Comment `gorm:"foreignkey:ParentID"`
	User     User      `gorm:"foreignkey:UserID"`
	Post     Post      `gorm:"foreignkey:PostID"`
}
