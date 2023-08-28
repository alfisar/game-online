package domain

type User struct {
	ID        int    `json:"id" gorm:"column:id"`
	Name      string `json:"name" gorm:"column:name"`
	Password  string `json:"password" gorm:"column:password"`
	Email     string `json:"email" gorm:"column:email"`
	Telephone string `json:"telephone" gorm:"column:telephone"`
}
