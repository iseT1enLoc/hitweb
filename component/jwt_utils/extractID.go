package jwtutils

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func ExtractID(requestToken string, secretKey string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrUnsupported
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		return "", err
	}
	print(token)
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", errors.ErrUnsupported
	}

	id := claims["id"].(string)

	idString := string(id)
	fmt.Printf("id String: %d", len(idString))
	return idString, nil

}
