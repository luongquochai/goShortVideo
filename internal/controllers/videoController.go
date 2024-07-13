package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luongquochai/goShortVideo/internal/services"
)

func UploadVideo(c *gin.Context) {
	userID := c.GetUint("userID")

	file, err := c.FormFile("video")
	if err != nil {
		log.Printf("Error getting file from request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}

	description := c.PostForm("description")
	log.Printf("User %d is uploading a file: %s", userID, file.Filename)

	video, err := services.UploadVideo(c, userID, file, description)
	if err != nil {
		log.Printf("Error uploading video: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"video": video})
}

func GetVideos(c *gin.Context) {
	videos, err := services.GetVideos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"videos": videos})
}
