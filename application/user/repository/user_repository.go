package repository

import (
	"game-online/domain"
	_database "game-online/internal/database"

	"gorm.io/gorm"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (obj *UserRepository) Create(db *gorm.DB, data domain.User) (result domain.User, err error) {
	err = db.Debug().Table("user").Save(&data).Error
	result = data
	return
}

func (obj *UserRepository) GetByUsername(db *gorm.DB, username string) (result domain.User, err error) {
	err = db.Debug().Table("user").Where("username = ?", username).First(&result).Error
	return
}

func (obj *UserRepository) GetList(db *gorm.DB, page int, limit int, search string) (result []domain.UserDetail, pagination _database.Pagination, err error) {
	pagination = _database.Pagination{}
	paginate := _database.Paginations("user", page, limit, &pagination, db)
	err = paginate(db).Select(`"user".*, wallet.balance, bank_account.accountNumber ,bank_account.accountName,bank_account.bankName`).Joins(`LEFT JOIN wallet ON wallet.userId = "user".id`).Joins(`LEFT JOIN bank_account ON bank_account.userId = "user".id`).Where(`"user".username like ?`, "%"+search+"%").Find(&result).Error
	pagination.Data = result
	return
}

func (obj *UserRepository) GetByID(db *gorm.DB, userID int) (result domain.UserDetail, err error) {
	err = db.Debug().Table("user").Select(`"user".*, wallet.id as walletId,wallet.balance, bank_account.accountNumber ,bank_account.accountName,bank_account.bankName`).Joins(`LEFT JOIN wallet ON wallet.userId = "user".id`).Joins(`LEFT JOIN bank_account ON bank_account.userId = "user".id`).Where(`"user".id = ?`, userID).Find(&result).Error
	return
}
