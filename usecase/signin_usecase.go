package usecase

import (
	"time"

	jwtutils "go_practice.com/component/jwt_utils"
	"go_practice.com/domain"
)

type signInUsecase struct {
	UserRepository domain.IUserRepository
	timeout        time.Duration
}

// CreateAcessToken implements domain.SignIn.
func (s *signInUsecase) CreateAcessToken(expirationHour int, secretKey string, user domain.User) (accesstoken string, err error) {
	return jwtutils.CreateAcessToken(expirationHour, secretKey, user)
}

// CreateRefreshToken implements domain.SignIn.
func (s *signInUsecase) CreateRefreshToken(expirationHour int, secretKey string, user domain.User) (refreshtoken string, err error) {
	return jwtutils.CreateRefreshToken(expirationHour, secretKey, user)
}

// GetUserByEmail implements domain.SignIn.
func (s *signInUsecase) GetUserByEmail(email string) (domain.User, error) {
	desiredUser, err := s.UserRepository.GetUserByEmail(email)
	if err != nil {
		return domain.User{}, err
	}
	return desiredUser, nil
}

func NewSignInUsecase(user_repository domain.IUserRepository, timeout time.Duration) domain.ISignIn {
	return &signInUsecase{
		UserRepository: user_repository,
		timeout:        timeout,
	}
}
