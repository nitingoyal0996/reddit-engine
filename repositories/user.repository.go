package repositories

import "github.com/nitingoyal0996/reddit-clone/models"

type UserRepository interface {
	Create(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)
	GetUserById(id uint) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
	CheckPassword(username string, password string) (*models.User, error)
	UpdateKarma(userId uint, amount int) error
}
