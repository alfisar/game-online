package repository

import (
	"game-online/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (obj *UserRepository) Create(db *gorm.DB, data domain.User) (result domain.User, err error) {
	err = db.Debug().Table("user").Save(&data).Error
	result = data
	return
}

func (obj *UserRepository) GetByUsername(db *gorm.DB, username string) (result domain.User, err error) {
	err = db.Debug().Table("user").Where("username = ?", username).First(&result).Error
	return
}
