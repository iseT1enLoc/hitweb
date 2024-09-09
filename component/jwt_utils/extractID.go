package jwtutils

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func ExtractID(requestToken string, secretKey string) (int, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrUnsupported
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return 0, errors.ErrUnsupported
	}

	id := claims["id"].(float64)

	idInt := int(id)
	return idInt, nil
}
