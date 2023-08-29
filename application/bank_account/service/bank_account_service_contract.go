package service

import (
	"game-online/domain"
	"game-online/internal/errorhandler"

	"gorm.io/gorm"
)

type BankAccountServiceContract interface {
	Create(db *gorm.DB, data domain.BankAccount) (err errorhandler.ErrorData)
}
