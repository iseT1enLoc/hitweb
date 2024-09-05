package domain

type SignUpReq struct {
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
	Password  string `json:"pass_word"`
}

type ISignUpUseCase interface {
	SignUp(signupReq SignUpReq) (user User, err error)
	GetUserByEmail(email string) error
	CreateAcessToken(expirationHour int, secretKey string, user User) (accesstoken string, err error)
	CreateRefreshToken(expirationHour int, secretKey string, user User) (refreshtoken string, err error)
}
