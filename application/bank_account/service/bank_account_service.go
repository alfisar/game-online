package service

import (
	"game-online/application/bank_account/repository"
	"game-online/domain"
	"game-online/internal/errorhandler"

	"gorm.io/gorm"
)

type BankAccountService struct {
	repo repository.BankAccountRepositoryContrct
}

func NewBankAccountService(repo repository.BankAccountRepositoryContrct) *BankAccountService {
	return &BankAccountService{
		repo: repo,
	}
}

func (obj *BankAccountService) Create(db *gorm.DB, data domain.BankAccount) (err errorhandler.ErrorData) {
	errData := obj.repo.Create(db, data)
	if errData != nil {
		return errorhandler.ErrorMapping(1001, "Terjadi kesalahan - GO-1001", errData)
	}

	return
}
