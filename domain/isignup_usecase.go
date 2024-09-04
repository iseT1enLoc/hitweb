package domain

type SignUpReq struct {
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
	Password  string `json:"pass_word"`
}

type ISignUpUseCase interface {
	SignUp(signupReq SignUpReq) (username string, err error)
}
