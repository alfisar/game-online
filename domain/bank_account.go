package domain

type BankAccount struct {
	ID            int    `json:"id" gorm:"column:id"`
	UserId        int    `json:"userId" gorm:"column:userid"`
	AccountNumber string `json:"accountNumber" gorm:"column:accountnumber"`
	AccountName   string `json:"accountName" gorm:"column:accountname"`
	BankName      string `json:"bankName" gorm:"column:bankname"`
}
