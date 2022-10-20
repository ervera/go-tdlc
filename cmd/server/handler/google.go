package handler

// import (
// 	"github.com/ervera/tdlc-gin/internal/localGoogle"
// 	"github.com/ervera/tdlc-gin/pkg/web"
// 	"github.com/gin-gonic/gin"
// )

// type gToken struct {
// 	GoogleToken string `json:"google_token"`
// }

// type GoogleHandler struct {
// 	service localGoogle.Service
// }

// func NewGoogleHandler(p localGoogle.Service) *GoogleHandler {
// 	return &GoogleHandler{
// 		service: p,
// 	}
// }

// func (c *GoogleHandler) Login() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		gToken := gToken{}
// 		ctx.ShouldBindJSON(&gToken)
// 		result, err := c.service.Login(ctx, gToken.GoogleToken)
// 		if err != nil {
// 			web.Error(ctx, 400, err.Error())
// 			return
// 		}
// 		web.Response(ctx, 200, result)
// 	}
// }
