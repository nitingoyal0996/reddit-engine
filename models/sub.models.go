package models

import "time"

type Subreddit struct {
	ID 			uint   		`gorm:"primaryKey"`
	Name 		string 		`gorm:"unique;not null"`
	Description string 		`gorm:"not null"`
	Owner 		string 		`gorm:"not null"`
	CreatedAt 	time.Time 	`gorm:"autoCreateTime"`
	UpdatedAt 	time.Time 	`gorm:"autoUpdateTime"`
}
