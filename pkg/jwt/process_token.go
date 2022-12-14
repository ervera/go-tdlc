package jwt

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/ervera/tdlc-gin/internal/domain"
	"github.com/google/uuid"
)

var UserID uuid.UUID
var Email string

func ProcessToken(tk string) (*domain.Claim, bool, string, error) {
	ID := ""

	myPass := []byte("generate")
	claims := &domain.Claim{}
	splitToken := strings.Split(tk, "Bearer ")
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("error token split")
	}
	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myPass, nil
	})
	if err == nil {
		/***** SIRVE PARA VALIDAR BUSCANDO EN LA BD ****/
		/*repoLogin := user.NewRepository(db)
		finded := repoLogin.Exist(ctx, claims.Username, claims.Email)
		if finded {
			Email = claims.Email
			Username = claims.Username
		}

		return claims, finded, ID, nil*/
		/***** SIRVE PARA VALIDAR BUSCANDO EN LA BD ****/
		Email = claims.Email
		UserID = uuid.Must(uuid.Parse(claims.ID))
		//guid := uuid.Must(uuid.Parse(ctx.Param("id")))
		return claims, true, ID, nil
	}
	if !tkn.Valid {
		return claims, false, "", errors.New("token invalido")
	}
	return claims, false, "", err
}
