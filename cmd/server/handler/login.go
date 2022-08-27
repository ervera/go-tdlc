package handler

import (
	"github.com/ervera/tdlc-gin/internal/domain"
	"github.com/ervera/tdlc-gin/internal/login"
	"github.com/ervera/tdlc-gin/pkg/jwt"
	"github.com/ervera/tdlc-gin/pkg/web"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	service login.Service
}

func (c *LoginHandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//ctx.Writer.Header().Add("content-type", "application/json")
		var u domain.User
		err := ctx.ShouldBindJSON(&u)
		if err != nil {
			web.Error(ctx, 400, "usuario y/o contraseña incorrecta"+err.Error())
			return
		}
		if len(u.Email) == 0 {
			web.Error(ctx, 400, "el email del usuario es requereido")
			return
		}
		user, err := c.service.Login(ctx, u.Email, u.Password)
		if err != nil {
			web.Error(ctx, 400, "usuario y/o contraseña invalido")
			return
		}
		jwtKey, err := jwt.GenerateJWT(user)
		if err != nil {
			web.Error(ctx, 400, "Ocurrio un error al intentar generar el token"+err.Error())
			return
		}

		user.Token = jwtKey
		web.Response(ctx, 200, user)
	}
}

func NewLogin(p login.Service) *LoginHandler {
	return &LoginHandler{
		service: p,
	}
}
