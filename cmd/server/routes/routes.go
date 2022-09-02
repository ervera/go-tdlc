package routes

import (
	"database/sql"

	"github.com/ervera/tdlc-gin/cmd/server/handler"
	"github.com/ervera/tdlc-gin/internal/localGoogle"
	"github.com/ervera/tdlc-gin/internal/login"
	"github.com/ervera/tdlc-gin/internal/media"
	"github.com/ervera/tdlc-gin/internal/user"
	"github.com/ervera/tdlc-gin/pkg/middleware"
	"github.com/ervera/tdlc-gin/pkg/sendgrid"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func (r *router) MapRoutes() {
	r.setGroup()

	r.buildGoogleRoutes()
	r.buildUserRoutes()
	r.buildLoginRoutes()
	// r.buildTweetRoutes()
	// r.buildImageRoutes()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/")
}

func (r *router) buildUserRoutes() {
	repoUsers := user.NewRepository(r.db)
	serviceMedia := media.NewService()
	serviceSendgrid := sendgrid.NewService()
	serviceUsers := user.NewService(repoUsers, serviceSendgrid)
	user := handler.NewHandlerUser(serviceUsers, serviceMedia)
	r.rg.POST("/user", user.CreateUser())
	r.rg.GET("/user/:id", middleware.TokenAuthMiddleware(), user.GetUserById())
	r.rg.PATCH("/user", middleware.TokenAuthMiddleware(), user.UpdateSelfUser())
	//r.rg.GET("/user/sendmail", user.SendMail())
	r.rg.POST("/user/forgotpassword", user.ForgotPassword())
	r.rg.POST("/user/newpassword", middleware.TokenAuthMiddleware(), user.NewPassword())
	// group := r.rg.Group("/user")
	// group.GET("/:id", middleware.TokenAuthMiddleware(), user.GetUserById())
	// group.POST("/relation/:id", middleware.TokenAuthMiddleware(), user.CreateUserRelation())
}

// func (r *router) buildImageRoutes() {
// 	repoUsers := user.NewRepository(r.db)
// 	serviceUsers := user.NewService(repoUsers)
// 	user := handler.NewImageHandler(serviceUsers)

// 	group := r.rg.Group("/image")

// 	group.POST("/:type", middleware.TokenAuthMiddleware(), user.UploadUserImage())
// 	group.DELETE("/:type", middleware.TokenAuthMiddleware(), user.DeleteUserImage())
// }

// func (r *router) buildTweetRoutes() {
// 	repoTweet := tweet.NewRepository(r.db)
// 	serviceTweet := tweet.NewService(repoTweet)
// 	user := handler.NewHandlerTweet(serviceTweet)
// 	group := r.rg.Group("/tweet")

// 	r.rg.POST("/tweet", middleware.TokenAuthMiddleware(), user.CreateTweet())
// 	group.GET("/:id", middleware.TokenAuthMiddleware(), user.GetTweetsByUserId())
// 	group.DELETE("/:id", middleware.TokenAuthMiddleware(), user.DeleteTweetsById())
// }

func (r *router) buildGoogleRoutes() {
	repoUser := user.NewRepository(r.db)
	serviceSendgrid := sendgrid.NewService()
	serviceUser := user.NewService(repoUser, serviceSendgrid)
	serviceMedia := media.NewService()
	localGoogleService := localGoogle.NewService(repoUser, serviceUser, serviceMedia)
	localGoogle := handler.NewGoogleHandler(localGoogleService)

	group := r.rg.Group("/google")

	group.POST("/login", localGoogle.Login())
}

func (r *router) buildLoginRoutes() {
	repoUsers := user.NewRepository(r.db)
	serviceLogin := login.NewService(repoUsers)
	login := handler.NewLogin(serviceLogin)
	r.rg.POST("/login", login.Login())
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r: r, db: db}
}
