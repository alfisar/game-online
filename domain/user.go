package domain

type User struct {
	ID        int    `json:"id" gorm:"column:id"`
	Name      string `json:"name" gorm:"column:username"`
	Password  string `json:"password" gorm:"column:password"`
	Email     string `json:"email" gorm:"column:email"`
	Telephone string `json:"telephone" gorm:"column:telephone"`
}

type UserDetail struct {
	ID            int     `json:"id" gorm:"column:id"`
	Name          string  `json:"name" gorm:"column:username"`
	Password      string  `json:"password" gorm:"column:password"`
	Email         string  `json:"email" gorm:"column:email"`
	Telephone     string  `json:"telephone" gorm:"column:telephone"`
	WalletID      int     `json:"-" gorm:"column:walletid"`
	WalletBalance float64 `json:"balance" gorm:"column:balance"`
	AccountName   string  `json:"accountName" gorm:"column:accountname"`
	AccountNumber string  `json:"accountNumber" gorm:"column:accountnumber"`
	BankName      string  `json:"bankName" gorm:"column:bankname"`
}
