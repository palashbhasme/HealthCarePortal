package utils

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/palashbhasme/healthcare-portal/internal/domain/models"
)

func ParseToken(tokenString, jwtSecret string) (*models.UserClaims, error) {
	claims := &models.UserClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil

}
