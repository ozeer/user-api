package initialize

import (
	"user-api/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()
	ApiGroup := r.Group("/v1")
	router.InitUserRouter(ApiGroup)

	return r
}
