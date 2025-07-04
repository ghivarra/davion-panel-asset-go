package server

import (
	"fmt"

	"github.com/ghivarra/davion-panel-asset-go/environment"
	"github.com/ghivarra/davion-panel-asset-go/router"
	"github.com/gin-gonic/gin"
)

func Run() {
	// load gin engine
	var engine *gin.Engine

	if environment.ENV == "development" {
		engine = gin.Default()
	} else {
		engine = gin.New()
	}

	// load router
	engine = router.Load(engine)

	// use recovery
	engine.Use(gin.Recovery())

	// serve
	config := fmt.Sprintf("%s:%s", environment.SERVER_HOST, environment.SERVER_PORT)
	fmt.Printf("Server is listening to %s", config)
	engine.Run(config)
}
