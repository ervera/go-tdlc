package routes

import (
	"github.com/ervera/tdlc-gin/cmd/server/handler"
	"github.com/ervera/tdlc-gin/internal/login"
	"github.com/ervera/tdlc-gin/internal/tweet"
	"github.com/ervera/tdlc-gin/internal/user"
	"github.com/ervera/tdlc-gin/pkg/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *mongo.Client
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.buildUserRoutes()
	r.buildLoginRoutes()
	r.buildTweetRoutes()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/")
}

func (r *router) buildUserRoutes() {
	repoUsers := user.NewRepository(r.db)
	serviceUsers := user.NewService(repoUsers)
	user := handler.NewHandlerUser(serviceUsers)
	group := r.rg.Group("/user")

	group.POST("/", user.CreateUser())
	group.GET("/:id", middleware.TokenAuthMiddleware(), user.GetUserById())
	group.PATCH("/", middleware.TokenAuthMiddleware(), user.UpdateSelfUser())
	//group.GET("/", user.GetUser())
}

func (r *router) buildTweetRoutes() {
	repoUsers := tweet.NewRepository(r.db)
	serviceUsers := tweet.NewService(repoUsers)
	user := handler.NewHandlerTweet(serviceUsers)
	group := r.rg.Group("/tweet")

	group.POST("/", middleware.TokenAuthMiddleware(), user.CreateTweet())
	group.GET("/:id", middleware.TokenAuthMiddleware(), user.GetTweetsByUserId())
	group.DELETE("/:id", middleware.TokenAuthMiddleware(), user.DeleteTweetsById())
}

func (r *router) buildLoginRoutes() {
	repoUsers := user.NewRepository(r.db)
	serviceLogin := login.NewService(repoUsers)
	login := handler.NewLogin(serviceLogin)
	group := r.rg.Group("/login")

	group.POST("/", login.Login())
}

func NewRouter(r *gin.Engine, db *mongo.Client) Router {
	return &router{r: r, db: db}
}
