package repository

import (
	"game-online/domain"

	"gorm.io/gorm"
)

type UserRepositoryContract interface {
	Create(db *gorm.DB, data domain.User) (result domain.User, err error)
	GetByUsername(db *gorm.DB, username string) (result domain.User, err error)
}
