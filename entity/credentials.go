package entity

import "github.com/golang-jwt/jwt"

type Credentials struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CustomClaims struct {
	Phone string `json:"phone"`
	Type  string `json:"type"`
	jwt.StandardClaims
}

var SecretKey = []byte("your-secret-key")
