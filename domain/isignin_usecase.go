package domain

type SignInReq struct {
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"pass_word"`
}
type SignInResponse struct {
	AccessToken  string
	RefreshToken string
}

type ISignIn interface {
	GetUserByEmail(email string) (user User, err error)
	CreateAcessToken(expirationHour int, secretKey string, user User) (accesstoken string, err error)
	CreateRefreshToken(expirationHour int, secretKey string, user User) (refreshtoken string, err error)
}
