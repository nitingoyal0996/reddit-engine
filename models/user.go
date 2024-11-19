package models

import (
	"errors"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string    `gorm:"unique;not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Karma     int       `gorm:"default:0"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	LastLogin time.Time `gorm:"autoUpdateTime"`
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
