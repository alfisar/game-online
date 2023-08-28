package repository

import (
	"game-online/domain"

	"gorm.io/gorm"
)

type WalletRepository struct{}

func NewUserRepository() *WalletRepository {
	return &WalletRepository{}
}

func (obj WalletRepository) Create(db *gorm.DB, data domain.Wallet) (err error) {
	err = db.Debug().Table("wallet").Create(&data).Error
	return
}

func (obj WalletRepository) UpdateBalace(db *gorm.DB, price float64) (err error) {
	err = db.Debug().Table("wallet").UpdateColumn("balance", gorm.Expr("balance + ?", price)).Error
	return
}
