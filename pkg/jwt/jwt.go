package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ervera/tdlc-gin/internal/domain"
)

func GenerateJWT(u domain.User) (string, error) {
	miClave := []byte("generate")

	payLoad := jwt.MapClaims{
		"email":      u.Email,
		"first_name": u.FirstName,
		"last_name":  u.LastName,
		"created_on": u.CreatedOn,
		"biography":  u.Biography,
		"location":   u.Location,
		"website":    u.Website,
		"id":         u.ID,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payLoad)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
