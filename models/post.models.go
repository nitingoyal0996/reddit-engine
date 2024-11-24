package models

import "time"

type Post struct {
    ID          uint64         `gorm:"primarykey"`
    Title       string         `gorm:"size:300;not null"`
    Content     string         `gorm:"type:text;not null"`
    AuthorID    uint64         `gorm:"not null"`
    SubredditID uint64         `gorm:"not null"`
    Score       int64          `gorm:"default:0"`
    CreatedAt   time.Time      `gorm:"not null"`
    UpdatedAt   time.Time      `gorm:"not null"`
    
    // Relationships
    Author      User           `gorm:"foreignKey:AuthorID"`
    Subreddit   Subreddit      `gorm:"foreignKey:SubredditID"`
    Comments    []Comment      `gorm:"foreignKey:PostID"`
}
