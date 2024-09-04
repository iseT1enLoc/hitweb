package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/ydb-platform/ydb-go-sdk/v3/log"
	jwtutils "go_practice.com/component/jwt_utils"
	"go_practice.com/domain"
)

type signupUsecase struct {
	userRepo domain.IUserRepository
	timeout  time.Duration
}

// SignUp implements domain.ISignUpUseCase.
func (s signupUsecase) SignUp(signupReq domain.SignUpReq) (username string, err error) {
	encryptedPassword, err := jwtutils.HashPassword(signupReq.Password)
	if err != nil {
		log.Error(err)
		return
	}
	user := domain.User{
		Id:        uuid.NewString(),
		UserName:  signupReq.UserName,
		UserEmail: signupReq.UserEmail,
		Password:  encryptedPassword,
	}
	_, err = s.userRepo.InsertUserToDatabase(user)
	if err != nil {
		log.Error(err)
		return
	}
	return user.UserName, nil

}

func NewSignUpUseCase(user_repo domain.IUserRepository, timeout time.Duration) domain.ISignUpUseCase {
	return &signupUsecase{
		userRepo: user_repo,
		timeout:  timeout,
	}
}
