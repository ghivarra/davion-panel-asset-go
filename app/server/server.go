package server

import (
	"fmt"
	"os"

	"github.com/ghivarra/davion-panel-asset-go/router"
	"github.com/gin-gonic/gin"
)

func Run() {
	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")

	// load gin engine
	ginRouter := gin.Default()

	// load router
	ginRouter = router.Load(ginRouter)

	// serve
	ginRouter.Run(fmt.Sprintf("%s:%s", serverHost, serverPort))
}
