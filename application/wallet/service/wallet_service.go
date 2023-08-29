package service

import (
	"game-online/application/wallet/repository"
	_repoWalletLog "game-online/application/wallet_log/repository"
	"game-online/domain"
	"game-online/internal/errorhandler"

	"gorm.io/gorm"
)

type WalletService struct {
	repo          repository.WalletRepositoryContract
	repoWalletLog _repoWalletLog.WalletLogRepositoryContract
}

func NewWalletSevice(repo repository.WalletRepositoryContract, repoWalletLog _repoWalletLog.WalletLogRepositoryContract) *WalletService {
	return &WalletService{
		repo:          repo,
		repoWalletLog: repoWalletLog,
	}
}

func (obj *WalletService) Transaction(db *gorm.DB, price float64, userID int, walletID int) (err errorhandler.ErrorData) {
	trx := db.Begin()

	errData := obj.repo.UpdateBalace(trx, price, userID, walletID)
	if errData != nil {
		trx.Rollback()
		return errorhandler.ErrorMapping(1001, "Terjadi kesalahan - GO-1001", errData)
	}

	dataLog := domain.WalletLog{
		UserID:   userID,
		WalletID: walletID,
	}

	if price < 0 {
		dataLog.WalletOut = price
	} else if price > 0 {
		dataLog.WalletIn = price
	}

	errData = obj.repoWalletLog.SaveDetail(trx, dataLog)
	if errData != nil {
		trx.Rollback()
		return errorhandler.ErrorMapping(1001, "Terjadi kesalahan - GO-1001", errData)
	}
	trx.Commit()
	return
}
