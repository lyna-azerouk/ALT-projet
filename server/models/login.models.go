package models

import "github.com/golang-jwt/jwt"

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Restaurant struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
type Restaurant_Menu struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int
}

type Response struct {
	Success int    `json:"success"`
	Message string `json:"message,omitempty"`
}

type UserClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
