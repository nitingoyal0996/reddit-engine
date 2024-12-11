package repositories

import (
	"errors"
	"time"

	"github.com/nitingoyal0996/reddit-clone/models"
	"gorm.io/gorm"
)

type SqliteUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *SqliteUserRepository {
	return &SqliteUserRepository{db: db}
}

func (r *SqliteUserRepository) Create(user *models.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.HashPassword(); err != nil {
		return err
	}

	var existingUser models.User
	if r.db.Where("username = ? or email = ?", user.Username, user.Email).First(&existingUser); existingUser.ID != 0 {
		return errors.New("username or email already exists")
	}

	return r.db.Create(user).Error
}

func (r *SqliteUserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *SqliteUserRepository) GetUserById(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *SqliteUserRepository) Update(user *models.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	return r.db.Save(user).
			Error
}

func (r *SqliteUserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).
			Error
}

func (r *SqliteUserRepository) CheckPassword(username string, password string) (*models.User, error) {
	user, err := r.GetUserByUsername(username); if err != nil {
		return nil, errors.New("user not found")
	}
	isValid, err := user.CheckPassword(password)
	if err != nil || !isValid {
		return nil, errors.New("invalid password")
	}
	// save timestamp of last login
	user.LastLogin = time.Now()
	r.Update(user)

	return user, nil
}

func (r *SqliteUserRepository) UpdateKarma(userId uint, amount int) error {
	return r.db.Model(models.User{}).
			Where("ID = ?", userId).
			Update("karma", gorm.Expr("karma + ?", amount)).
			Error
}