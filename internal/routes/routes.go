package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luongquochai/goShortVideo/internal/controllers"
	"github.com/luongquochai/goShortVideo/internal/middlewares"
)

func SetupRoutes(r *gin.Engine) {
	// Serve static files
	r.Static("/public", "./frontend/public/")
	r.Static("/icon", "./frontend/src/icon")
	r.Static("/videos", "./videos")

	// Load HTML templates
	r.LoadHTMLGlob("frontend/src/**/*.templ")

	// Serve the main HTML page
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Home.templ", nil)
	})

	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
		api.GET("/videos", controllers.GetVideos)
		// api.POST("/upload", controllers.UploadVideo)

		auth := api.Group("/")
		auth.Use(middlewares.AuthMiddleware())
		{
			auth.POST("/upload", controllers.UploadVideo)

		}
	}
}
