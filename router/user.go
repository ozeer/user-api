package router

import (
	"user-api/api"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserGroup := Router.Group("user")
	{
		UserGroup.GET("list", api.GetUserList)
		UserGroup.POST("login", api.PasswordLogin)
	}
}
