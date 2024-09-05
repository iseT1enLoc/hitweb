package jwtutils

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go_practice.com/domain"
)

func CreateRefreshToken(expirationHour int, secretKey string, user domain.User) (refreshToken string, err error) {
	exp := time.Now().Add(time.Duration(expirationHour) * time.Hour)
	customeClaim := domain.CustomRefreshClaim{
		Id:               user.Id,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: &jwt.NumericDate{exp}},
	}
	reToken := jwt.NewWithClaims(jwt.SigningMethodHS256, customeClaim)

	signString, err := reToken.SignedString([]byte(secretKey))
	if err != nil {
		//log.Fatalf("Refresh: Error happened while creating sign string [err]-%v", err)
		log.Print("Error happened while create refresh sign string")
		return "", err
	}
	return signString, nil
}
