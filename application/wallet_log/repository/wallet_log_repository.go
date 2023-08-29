package repository

import (
	"game-online/domain"

	"gorm.io/gorm"
)

type WalletLogRepository struct{}

func NewWalletLogRepository() *WalletLogRepository {
	return &WalletLogRepository{}
}

func (obj *WalletLogRepository) SaveDetail(db *gorm.DB, data domain.WalletLog) (err error) {
	return db.Debug().Table("wallet_log").Create(&data).Error
}
