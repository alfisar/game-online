package repository

import (
	"game-online/domain"

	"gorm.io/gorm"
)

type BankAccountRepository struct{}

func NewBankAccountRepository() *BankAccountRepository {
	return &BankAccountRepository{}
}

func (obj *BankAccountRepository) Create(db *gorm.DB, data domain.BankAccount) (err error) {
	err = db.Debug().Table("bank_account").Create(&data).Error
	return
}
