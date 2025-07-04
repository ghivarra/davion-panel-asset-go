package server

import (
	"fmt"

	"github.com/ghivarra/davion-panel-asset-go/environment"
	"github.com/ghivarra/davion-panel-asset-go/router"
	"github.com/gin-gonic/gin"
)

func Run() {
	// load gin engine
	ginRouter := gin.Default()

	// load router
	ginRouter = router.Load(ginRouter)

	// serve
	ginRouter.Run(fmt.Sprintf("%s:%s", environment.SERVER_HOST, environment.SERVER_PORT))
}
