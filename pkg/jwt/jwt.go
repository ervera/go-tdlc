package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ervera/tdlc-gin/internal/domain"
)

func GenerateJWT(u domain.User) (string, error) {
	miClave := []byte("generate")

	payLoad := jwt.MapClaims{
		"email":            u.Email,
		"nombre":           u.Nombre,
		"apellidos":        u.Apellido,
		"fecha_nacimiento": u.FechaNacimiento,
		"biografia":        u.Biografia,
		"ubicacion":        u.Ubicacion,
		"sitioweb":         u.SitioWeb,
		"_id":              u.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payLoad)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
