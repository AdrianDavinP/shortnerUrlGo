package Struct

import "github.com/dgrijalva/jwt-go"

type Token struct {
	Date    string
	Company string
	*jwt.StandardClaims
}
