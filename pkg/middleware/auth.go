package middleware

import (
	"github.com/ervera/tdlc-gin/pkg/jwt"
	"github.com/ervera/tdlc-gin/pkg/web"
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, _, _, err := jwt.ProcessToken(ctx.GetHeader("Authorization"))
		if err != nil {
			web.Error(ctx, 400, "token vencido"+err.Error())
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
