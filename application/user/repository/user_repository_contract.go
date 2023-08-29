package repository

import (
	"game-online/domain"
	_database "game-online/internal/database"

	"gorm.io/gorm"
)

type UserRepositoryContract interface {
	Create(db *gorm.DB, data domain.User) (result domain.User, err error)
	GetByUsername(db *gorm.DB, username string) (result domain.User, err error)
	GetList(db *gorm.DB, page int, limit int, search string) (result []domain.UserDetail, pagination _database.Pagination, err error)
	GetByID(db *gorm.DB, userID int) (result domain.UserDetail, err error)
}
