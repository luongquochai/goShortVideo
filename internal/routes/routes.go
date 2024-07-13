package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/luongquochai/goShortVideo/internal/controllers"
	"github.com/luongquochai/goShortVideo/internal/middlewares"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		auth := api.Group("/")
		auth.Use(middlewares.AuthMiddleware())
		{
			auth.POST("/upload", controllers.UploadVideo)
			auth.GET("/videos", controllers.GetVideos)
		}
	}
}
