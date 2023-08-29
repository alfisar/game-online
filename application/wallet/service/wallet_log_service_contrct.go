package service

import (
	"game-online/internal/errorhandler"

	"gorm.io/gorm"
)

type WalletLogServiceContract interface {
	Transaction(db *gorm.DB, price float64, userID int, walletID int) (err errorhandler.ErrorData)
}
