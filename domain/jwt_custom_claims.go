package domain

import "github.com/golang-jwt/jwt/v5"

type CustomAccessToken struct {
	Id        string `json:"id"`
	User_name string `json:"user_name"`
	jwt.RegisteredClaims
}
type CustomRefreshClaim struct {
	Id string `json:"id"`
	jwt.RegisteredClaims
}
