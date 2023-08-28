package repository

import (
	"game-online/domain"

	"gorm.io/gorm"
)

type BankAccountRepositoryContrct interface {
	Create(db *gorm.DB, data domain.BankAccount) (err error)
}
