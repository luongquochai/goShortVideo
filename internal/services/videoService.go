package services

import (
	"fmt"
	"log"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/luongquochai/goShortVideo/internal/config"
	"github.com/luongquochai/goShortVideo/internal/models"
	"github.com/luongquochai/goShortVideo/internal/utils"
)

func UploadVideo(c *gin.Context, userID uint, file *multipart.FileHeader, description string) (*models.Video, error) {
	filePath := fmt.Sprintf("videos/%s-%s", uuid.New().String(), file.Filename)
	thumbnailPath := fmt.Sprintf("thumbnails/%s-thumbnail.png", uuid.New().String())

	log.Printf("Saving file to: %s", filePath)
	if err := utils.SaveFile(c, filePath, file); err != nil {
		log.Printf("Error saving file: %v", err)
		return nil, err
	}

	log.Printf("Generating thumbnail for: %s", filePath)
	if err := utils.GenerateThumbnail(filePath, thumbnailPath); err != nil {
		log.Printf("Error generating thumbnail: %v", err)
		return nil, err
	}

	video := &models.Video{
		UserID:        userID,
		FilePath:      filePath,
		ThumbnailPath: thumbnailPath,
		Description:   description,
	}

	log.Printf("Saving video information to database")
	if err := config.DB.Create(video).Error; err != nil {
		log.Printf("Error saving video to DB: %v", err)
		return nil, err
	}

	return video, nil
}

func GetVideos() ([]models.Video, error) {
	var videos []models.Video
	if err := config.DB.Find(&videos).Error; err != nil {
		return nil, err
	}

	return videos, nil
}
