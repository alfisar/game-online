package repository

import (
	"game-online/domain"

	"gorm.io/gorm"
)

type SessionRepository struct{}

func NewSessionRepository() *SessionRepository {
	return &SessionRepository{}
}

func (obj *SessionRepository) Create(db *gorm.DB, data domain.Session) (err error) {
	err = db.Debug().Table("session").Create(&data).Error
	return
}

func (obj *SessionRepository) DeleteByUserID(db *gorm.DB, userId int) (err error) {
	err = db.Debug().Table("session").Delete(&domain.Session{}).Where("username = ?", userId).Error
	return
}
