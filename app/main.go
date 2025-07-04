package main

import (
	"os"

	"github.com/ghivarra/davion-panel-asset-go/common"
	"github.com/ghivarra/davion-panel-asset-go/environment"
	"github.com/ghivarra/davion-panel-asset-go/server"
	"github.com/joho/godotenv"
)

func main() {

	// load dotenv
	err := godotenv.Load()
	if err != nil {
		panic("failed to locate .env file")
	}

	// save dotenv
	environment.Save()

	// set rootpath
	var errPath error
	common.ROOTPATH, errPath = os.Getwd()
	if errPath != nil {
		panic("failed to locate the root working directory")
	}

	// run server
	server.Run()
}
