package system

import (
	"en-lang-bk/api"

	"github.com/gin-gonic/gin"
)

type WordRouter struct{}

func (w *WordRouter) InitWordRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	wordRouter := Router.Group("word")
	wordApi := api.ApiGroupApp.SystemApiGroup.WordApi
	{
		wordRouter.POST("create", wordApi.Create)
	}
	return wordRouter
}
