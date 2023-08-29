package service

import (
	"game-online/domain"
	_database "game-online/internal/database"
	"game-online/internal/errorhandler"

	"gorm.io/gorm"
)

type UserServiceContract interface {
	Create(db *gorm.DB, data domain.User) (err errorhandler.ErrorData)
	Login(db *gorm.DB, username string, pass string) (token string, err errorhandler.ErrorData)
	Logout(db *gorm.DB, userID int) (err errorhandler.ErrorData)
	GetList(db *gorm.DB, page int, limit int, search string) (result []domain.UserDetail, pagination _database.Pagination, err errorhandler.ErrorData)
	GetByID(db *gorm.DB, userID int) (result domain.UserDetail, err errorhandler.ErrorData)
}
