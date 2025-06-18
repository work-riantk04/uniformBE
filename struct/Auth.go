package structure

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	Appname        string    `json:"name"`
	jwt.StandardClaims
}