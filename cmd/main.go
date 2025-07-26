package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nmartin867/auth-logs/initializers"
	"github.com/nmartin867/auth-logs/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// AuthEntry Routes
	routes.AuthEntryRoutes(r)

	r.Run()
}
