package repository

import (
	"game-online/domain"

	"gorm.io/gorm"
)

type WalletLogRepositoryContract interface {
	SaveDetail(db *gorm.DB, data domain.WalletLog) (err error)
}
