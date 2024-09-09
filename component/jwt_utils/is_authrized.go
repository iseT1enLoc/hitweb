package jwtutils

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func Is_authorized(requestToken string, secretkey string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			print(token)
			return nil, errors.ErrUnsupported
		}
		return []byte(secretkey), nil
	})
	if err != nil {
		print("line 17 is authorized")
		return false, err
	}
	print("Authorized...")
	return true, nil
}
