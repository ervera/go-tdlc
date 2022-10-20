package domain

import (
	"github.com/dgrijalva/jwt-go"
)

type Claim struct {
	Email string `json:"email"`
	//ID    string `json:"id"`
	UUID string `json:"uuid"`
	jwt.StandardClaims
}
