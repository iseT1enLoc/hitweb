package jwtutils

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go_practice.com/domain"
)

func CreateAcessToken(expirationHour int, secretKey string, user domain.User) (accessToken string, err error) {
	exp := time.Now().Add(time.Duration(expirationHour) * time.Hour)
	customClaim := domain.CustomAccessToken{
		Id:               user.Id,
		User_name:        user.UserName,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: &jwt.NumericDate{exp}},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaim)

	signString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		//log.Fatalf("Access: Error happened while create signString [error]-%v", signString)
		log.Print("Error happened while creating refresh sign string")
		return "", err
	}
	return signString, nil
}
