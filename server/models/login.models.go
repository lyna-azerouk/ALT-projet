package models

import (
	"github.com/golang-jwt/jwt"
)

type ClientCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RestaurantCredentials struct {
	Id       uint64 `json:"id"`
	Password string `json:"password"`
}

type Response struct {
	Success int    `json:"success"`
	Message string `json:"message,omitempty"`
}

type ClientClaims struct {
	Email string `json:"email"`
	Id    uint64 `json:"id"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type RestaurantClaims struct {
	Id uint64 `json:"id"`
	jwt.StandardClaims
}
