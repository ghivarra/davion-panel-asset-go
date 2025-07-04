package router

import (
	"fmt"

	"github.com/ghivarra/davion-panel-asset-go/common"
	"github.com/ghivarra/davion-panel-asset-go/module/controller/assets/image"
	"github.com/ghivarra/davion-panel-asset-go/module/controller/home"
	corsMiddleware "github.com/ghivarra/davion-panel-asset-go/module/middleware/cors-middleware"
	"github.com/gin-gonic/gin"
)

func Load(router *gin.Engine) *gin.Engine {

	// home router
	router.GET("/", home.Index)

	// favicon
	router.StaticFile("/favicon.ico", fmt.Sprintf("%s/public/favicon.ico", common.ROOTPATH))

	// asset router group
	assetRouterGroup := router.Group("/assets", corsMiddleware.Run)
	assetRouterGroup.GET("/image/*path", image.Get)

	// return router to be served by server
	return router
}
