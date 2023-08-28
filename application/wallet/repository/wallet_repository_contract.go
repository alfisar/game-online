package repository

import (
	"game-online/domain"

	"gorm.io/gorm"
)

type WalletRepositoryContract interface {
	Create(db *gorm.DB, data domain.Wallet) (err error)
	UpdateBalace(db *gorm.DB, price float64) (err error)
}
