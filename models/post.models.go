package models

import "time"

type Post struct {
	ID        uint      `gorm:"primaryKey"`
	Subreddit string    `gorm:"not null"`
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"not null"`
	Author    string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	PostKarma int       `gorm:"default:0"`
	IsDeleted bool      `gorm:"default:false"`
}