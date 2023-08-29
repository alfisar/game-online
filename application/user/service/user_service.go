package service

import (
	_repositoryRedis "game-online/application/redis/repository"
	"game-online/application/user/repository"
	_repositoryWallet "game-online/application/wallet/repository"
	_database "game-online/internal/database"
	"strconv"

	"game-online/config"
	"game-online/domain"
	"game-online/internal/errorhandler"
	"game-online/internal/hashgenerator"
	"game-online/internal/jwthandler"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	repo       repository.UserRepositoryContract
	repoWallet _repositoryWallet.WalletRepositoryContract
	repoRedis  _repositoryRedis.RedisRepositoryContract
}

func NewUserService(repo repository.UserRepositoryContract, repoWallet _repositoryWallet.WalletRepositoryContract, repoRedis _repositoryRedis.RedisRepositoryContract) *UserService {
	return &UserService{
		repo:       repo,
		repoWallet: repoWallet,
		repoRedis:  repoRedis,
	}
}

func (obj *UserService) Create(db *gorm.DB, data domain.User) (err errorhandler.ErrorData) {
	pass, errData := hashgenerator.Generate(data.Password)
	if errData != nil {
		return errorhandler.ErrorMapping(1101, "terjadi kesalahan - GO-1101", errData)
	}
	data.Password = pass

	trx := db.Begin()
	userData, errData := obj.repo.Create(trx, data)
	if errData != nil {
		trx.Rollback()
		return errorhandler.ErrorMapping(1001, "terjadi kesalahan - GO-1001", errData)
	}

	walletData := domain.Wallet{
		UserID:  userData.ID,
		Balance: 0,
	}

	errData = obj.repoWallet.Create(trx, walletData)
	if errData != nil {
		trx.Rollback()
		return errorhandler.ErrorMapping(1001, "terjadi kesalahan - GO-1001", errData)
	}

	trx.Commit()
	return
}

func (obj *UserService) Login(db *gorm.DB, username string, pass string) (token string, err errorhandler.ErrorData) {

	userData, errData := obj.repo.GetByUsername(db, username)
	if errData != nil {
		err = errorhandler.ErrorMapping(1001, "terjadi kesalahan - GO-1001", errData)
		return
	}

	result, errData := hashgenerator.Verify(userData.Password, pass)
	if errData == bcrypt.ErrMismatchedHashAndPassword {
		err = errorhandler.ErrorMapping(1102, "terjadi kesalahan - GO-1102", errData)
		return
	} else if !result {
		err = errorhandler.ErrorMapping(1103, "terjadi kesalahan - GO-1103", errData)
		return
	}

	key := "token_user_" + strconv.Itoa(userData.ID)
	token, errData = obj.repoRedis.Get(key)
	if errData != nil && errData.Error() != "redis: nil" {
		err = errorhandler.ErrorMapping(1002, "terjadi kesalahan - GO-1002", errData)
		return
	} else if token != "" {
		return
	}

	cfg := config.GetConfig()

	objJwt := jwthandler.JwtHandler{
		Secret: cfg.JWTSecret,
	}

	sessionDuration := time.Duration(cfg.SessionEXP) * time.Minute
	token, errData = objJwt.GetToken(sessionDuration, userData.ID)
	if errData != nil {
		err = errorhandler.ErrorMapping(1002, "terjadi kesalahan - GO-1002", errData)
		return
	}

	go obj.repoRedis.Insert(key, token, sessionDuration)

	return
}

func (obj *UserService) Logout(db *gorm.DB, userID int) (err errorhandler.ErrorData) {
	key := "token_user_" + strconv.Itoa(userID)
	errData := obj.repoRedis.Delete(key)
	if errData != nil {
		err = errorhandler.ErrorMapping(1001, "terjadi kesalahan - GO-1001", errData)
		return
	}
	return
}

func (obj *UserService) GetList(db *gorm.DB, page int, limit int, search string) (result []domain.UserDetail, pagination _database.Pagination, err errorhandler.ErrorData) {
	result, pagination, errData := obj.repo.GetList(db, page, limit, search)
	if errData != nil {
		err = errorhandler.ErrorMapping(1001, "Terjadi kesalahan - GO-1001", errData)
		return
	}

	return
}

func (obj *UserService) GetByID(db *gorm.DB, userID int) (result domain.UserDetail, err errorhandler.ErrorData) {
	result, errData := obj.repo.GetByID(db, userID)
	if errData != nil {
		err = errorhandler.ErrorMapping(1001, "Terjadi kesalahan - GO-1001", errData)
		return
	}

	return
}
