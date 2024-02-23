package models

import "github.com/golang-jwt/jwt"

type ClientCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RestaurantCredentials struct {
	Id       int    `json:"id"`
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

type ClientClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type RestaurantClaims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}
