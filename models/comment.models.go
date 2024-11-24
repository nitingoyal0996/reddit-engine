package models

import "time"

type Comment struct {
    ID        uint64         `gorm:"primarykey"`
    Content   string         `gorm:"type:text;not null"`
    AuthorID  uint64         `gorm:"not null"`
    PostID    uint64         `gorm:"not null"`
    ParentID  *uint64        `gorm:"default:null"` // nullable for top-level comments
    Score     int64          `gorm:"default:0"`
    CreatedAt time.Time      `gorm:"not null"`
    UpdatedAt time.Time      `gorm:"not null"`
    
    // Relationships
    Author    User           `gorm:"foreignKey:AuthorID"`
    Post      Post           `gorm:"foreignKey:PostID"`
    Parent    *Comment       `gorm:"foreignKey:ParentID"`
    Replies   []Comment      `gorm:"foreignKey:ParentID"`
}