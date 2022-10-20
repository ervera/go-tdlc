package routers

import (
	"github.com/ervera/tdlc-gin/pkg/middleware"
	"github.com/ervera/tdlc-gin/src/api/handlers"
	"github.com/gin-gonic/gin"
)

// Mapper handles router mappings
type Mapper struct {
	userHandler   *handlers.UserHandler
	loginHandler  *handlers.LoginHandler
	googleHandler *handlers.GoogleHandler
	teamHandler   *handlers.TeamHandler
}

// NewMapper returns a new instance of a Mapper
func NewMapper(userHandler *handlers.UserHandler,
	loginHandler *handlers.LoginHandler,
	googleHandler *handlers.GoogleHandler,
	teamHandler *handlers.TeamHandler,
) Mapper {
	return Mapper{
		userHandler,
		loginHandler,
		googleHandler,
		teamHandler,
	}
}

func (m Mapper) configureMappings(router *gin.Engine) {
	baseGroup := router.Group("/")

	userGroup := router.Group("/user")
	userGroup.POST("", m.userHandler.CreateUser())
	userGroup.GET("/:id", middleware.TokenAuthMiddleware(), m.userHandler.GetUserById())
	userGroup.PATCH("", middleware.TokenAuthMiddleware(), m.userHandler.UpdateSelfUser())
	//r.rg.GET("/user/sendmail", user.SendMail())
	userGroup.POST("/forgotpassword", m.userHandler.ForgotPassword())
	userGroup.POST("/newpassword", middleware.TokenAuthMiddleware(), m.userHandler.NewPassword())
	userGroup.PATCH("/media/:type", middleware.TokenAuthMiddleware(), m.userHandler.UpdateMedia())
	userGroup.DELETE("/media/:type", middleware.TokenAuthMiddleware(), m.userHandler.DeleteMedia())

	baseGroup.POST("/login", m.loginHandler.Login())

	baseGroup.POST("/google/login", m.googleHandler.Login())

	teamGroup := router.Group("/team")
	teamGroup.POST("", middleware.TokenAuthMiddleware(), m.teamHandler.CreateTeam())
	teamGroup.GET("", middleware.TokenAuthMiddleware(), m.teamHandler.GetUserTeam())
}
