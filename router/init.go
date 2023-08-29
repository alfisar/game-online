package router

import (
	_userControll "game-online/application/user/controller"
	_userRepo "game-online/application/user/repository"
	_userServ "game-online/application/user/service"
	"game-online/internal/jwthandler"
	"game-online/internal/middleware"
	"os"

	_walletControll "game-online/application/wallet/controller"
	_walletRepo "game-online/application/wallet/repository"
	_walletServ "game-online/application/wallet/service"

	_bankControll "game-online/application/bank_account/controller"
	_bankRepo "game-online/application/bank_account/repository"
	_bankServ "game-online/application/bank_account/service"

	_redisRepo "game-online/application/redis/repository"
	_walletLogRepo "game-online/application/wallet_log/repository"
	"game-online/database"

	"github.com/joho/godotenv"
)

var (
	_         = godotenv.Load(".env")
	db        = database.NewDatabaseMySql()
	redis     = database.NewDatabaseRedis()
	repoRedis = _redisRepo.NewRedisRepository(redis)
)

func setUserController() *_userControll.UserController {
	repo := _userRepo.NewUserRepository()
	repoWallet := _walletRepo.NewWalletRepository()
	serv := _userServ.NewUserService(repo, repoWallet, repoRedis)
	controller := _userControll.NewUserController(serv, db)
	return &controller
}

func setWalletController() *_walletControll.WalletController {
	repoWallet := _walletRepo.NewWalletRepository()
	repoWalletLog := _walletLogRepo.NewWalletLogRepository()
	serv := _walletServ.NewWalletSevice(repoWallet, repoWalletLog)
	controller := _walletControll.NewWalletController(serv, db)
	return &controller
}

func setBankController() *_bankControll.BankAccountController {
	repo := _bankRepo.NewBankAccountRepository()
	serv := _bankServ.NewBankAccountService(repo)
	controller := _bankControll.NewBankAccountController(serv, db)
	return &controller
}

func setMiddleware() *middleware.AuthenticateMiddleware {
	repo := _userRepo.NewUserRepository()
	repoWallet := _walletRepo.NewWalletRepository()
	serv := _userServ.NewUserService(repo, repoWallet, repoRedis)
	jwtData := jwthandler.GetJwt()
	jwtData.Secret = os.Getenv("JWT_SECRET")
	middleWR := middleware.NewAuthenticateMiddleware(serv, jwtData, db)
	return middleWR
}
