package service

import (
	_repositorySession "game-online/application/session/repository"
	"game-online/application/user/repository"
	_repositoryWallet "game-online/application/wallet/repository"
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
	repo        repository.UserRepositoryContract
	repoWallet  _repositoryWallet.WalletRepositoryContract
	repoSession _repositorySession.SessionRepositoryContract
}

func NewUserService(repo repository.UserRepositoryContract, repoWallet _repositoryWallet.WalletRepositoryContract, repoSession _repositorySession.SessionRepositoryContract) *UserService {
	return &UserService{
		repo:        repo,
		repoWallet:  repoWallet,
		repoSession: repoSession,
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

	cfg := config.GetConfig()

	objJwt := jwthandler.JwtHandler{
		Secret: cfg.JWTSecret,
	}

	token, errData = objJwt.GetToken(time.Duration(cfg.SessionEXP)*time.Minute, userData.ID)
	if errData != nil {
		err = errorhandler.ErrorMapping(1002, "terjadi kesalahan - GO-1002", errData)
		return
	}

	sessionData := domain.Session{
		UserID: userData.ID,
		Token:  token,
	}

	errData = obj.repoSession.Create(db, sessionData)
	if errData != nil {
		err = errorhandler.ErrorMapping(1001, "terjadi kesalahan - GO-1001", errData)
		return
	}
	return
}

func (obj *UserService) Logout(db *gorm.DB, userID int) (err errorhandler.ErrorData) {
	errData := obj.repoSession.DeleteByUserID(db, userID)
	if errData != nil {
		err = errorhandler.ErrorMapping(1001, "terjadi kesalahan - GO-1001", errData)
		return
	}
	return
}
