package routes

import (
	"github.com/ervera/tdlc-gin/cmd/server/handler"
	"github.com/ervera/tdlc-gin/internal/user"
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
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/")
}

func (r *router) buildUserRoutes() {
	repoUsers := user.NewRepository(r.db)
	serviceusers := user.NewService(repoUsers)
	user := handler.NewHandlerUser(serviceusers)
	group := r.rg.Group("/user")
	group.POST("/", user.CreateUser())
}

func NewRouter(r *gin.Engine, db *mongo.Client) Router {
	return &router{r: r, db: db}
}
