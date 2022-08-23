package domain

import (
	"github.com/dgrijalva/jwt-go"
)

type Claim struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	jwt.StandardClaims
}
