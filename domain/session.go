package domain

type Session struct {
	ID     int    `json:"id" gorm:"column:id"`
	UserID int    `json:"userId" gorm:"column:userId"`
	Token  string `json:"token" gorm:"column:token"`
}
