package domain

type Wallet struct {
	ID      int     `json:"id" gorm:"column:id"`
	UserID  int     `json:"" gorm:"column:userId"`
	Balance float64 `json:"" gorm:"column:balance"`
}
