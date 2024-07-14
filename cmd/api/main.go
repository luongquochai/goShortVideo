package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luongquochai/goShortVideo/internal/config"
	"github.com/luongquochai/goShortVideo/internal/routes"
)

func main() {
	config.LoadEnv()
	config.InitDB()

	r := gin.Default()

	routes.SetupRoutes(r)
	r.Run() // default to localhost:8080
}
