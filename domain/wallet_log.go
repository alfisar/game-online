package domain

type WalletLog struct {
	ID        int     `json:"id" gorm:"column:id"`
	UserID    int     `json:"userId" gorm:"column:userid"`
	WalletID  int     `json:"walletId" gorm:"column:walletid"`
	WalletIn  float64 `json:"walletIn" gorm:"column:walletin"`
	WalletOut float64 `json:"walletOut" gorm:"column:walletout"`
}
