package services

import (
	"serveur/server/models"
)

import (
	"github.com/golang-jwt/jwt"
	_ "github.com/golang-jwt/jwt"
	_ "os"
)

func NewAccessToken(claims models.UserClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return accessToken.SignedString([]byte("SECRET"))
}

func ParseAccessToken(accessToken string) *models.UserClaims {
	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET"), nil
	})
	return parsedAccessToken.Claims.(*models.UserClaims)
}
