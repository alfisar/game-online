package domain

type BankAccount struct {
	ID            int    `json:"id" gorm:"column:id"`
	UserId        int    `json:"userId" gorm:"column:userId"`
	AccountNumber string `json:"accountNumber" gorm:"column:accountNumber"`
	AccountName   string `json:"accountName" gorm:"column:accountName"`
	BankName      string `json:"bankName" gorm:"column:bankName"`
}
