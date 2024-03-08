package models

import (
	"github.com/golang-jwt/jwt"
)

type ClientCredentials struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type RestaurantCredentials struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type Response struct {
	Success int    `json:"success"`
	Message string `json:"message,omitempty"`
}

type ClientClaims struct {
	Email string `json:"email"`
	Id    string `json:"id"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type RestaurantClaims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}
