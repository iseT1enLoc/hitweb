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

// CreateAcessToken implements domain.ISignUpUseCase.
func (s *signupUsecase) CreateAcessToken(expirationHour int, secretKey string, user domain.User) (accesstoken string, err error) {
	return jwtutils.CreateAcessToken(expirationHour, secretKey, user)
}

// CreateRefreshToken implements domain.ISignUpUseCase.
func (s *signupUsecase) CreateRefreshToken(expirationHour int, secretKey string, user domain.User) (refreshtoken string, err error) {
	return jwtutils.CreateRefreshToken(expirationHour, secretKey, user)
}

// GetUserByEmail implements domain.ISignUpUseCase.
func (s *signupUsecase) GetUserByEmail(email string) error {
	_, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return err
	}
	return nil
}

// SignUp implements domain.ISignUpUseCase.
func (s signupUsecase) SignUp(signupReq domain.SignUpReq) (insertedUser domain.User, err error) {
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
	iuser, err := s.userRepo.InsertUserToDatabase(user)
	if err != nil {
		log.Error(err)
		return
	}
	return iuser, nil

}

func NewSignUpUseCase(user_repo domain.IUserRepository, timeout time.Duration) domain.ISignUpUseCase {
	return &signupUsecase{
		userRepo: user_repo,
		timeout:  timeout,
	}
}
