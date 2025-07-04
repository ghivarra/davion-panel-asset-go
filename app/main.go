package main

import (
	"github.com/ghivarra/davion-panel-asset-go/server"
	"github.com/joho/godotenv"
)

func main() {

	// load dotenv
	err := godotenv.Load()
	if err != nil {
		panic("failed to locate .env file")
	}

	// run server
	server.Run()
}
