package domain

type Wallet struct {
	ID      int     `json:"id" gorm:"column:id"`
	UserID  int     `json:"" gorm:"column:userid"`
	Balance float64 `json:"" gorm:"column:balance"`
}

type TopUp struct {
	Nominal float64 `json:"nominal"`
}
