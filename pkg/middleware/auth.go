package middleware

import (
	"fmt"

	"github.com/ervera/tdlc-gin/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, _, _, err := jwt.ProcessToken(ctx.GetHeader("Authorization"))
		if err != nil {
			fmt.Println(err)
			return
		}
		ctx.Next()
	}
}
