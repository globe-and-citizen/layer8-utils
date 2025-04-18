package utilities

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	UserName string `json:"username"`
	UserID   uint   `json:"user_id"`
	jwt.RegisteredClaims
}

type ClientClaims struct {
	UserName string `json:"username"`
	ClientID string `json:"user_id"`
	jwt.RegisteredClaims
}
