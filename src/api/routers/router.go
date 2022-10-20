package routers

import "github.com/gin-gonic/gin"

func CreateRouter(m Mapper) *gin.Engine {

	router := gin.Default()
	m.configureMappings(router)

	return router
}
