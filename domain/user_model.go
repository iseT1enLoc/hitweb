package domain

type User struct {
	Id        string `gorm:"primaryKey" json:"id"`
	UserName  string `gorm:"column: user_name" json:"user_name"`
	UserEmail string `gorm:"column: user_email" json:"user_email"`
	Password  string `gorm:"column: pass_word" json:"pass_word"`
}
