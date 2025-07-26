package main

import (
	"github.com/nmartin867/auth-logs/initializers"
	"github.com/nmartin867/auth-logs/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.AuthEntry{})
}
