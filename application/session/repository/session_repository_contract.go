package repository

import (
	"game-online/domain"

	"gorm.io/gorm"
)

type SessionRepositoryContract interface {
	Create(db *gorm.DB, data domain.Session) (err error)
	DeleteByUserID(db *gorm.DB, userId int) (err error)
}
