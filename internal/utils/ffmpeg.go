package utils

import (
	"bytes"
	"log"
	"mime/multipart"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func SaveFile(c *gin.Context, filePath string, file *multipart.FileHeader) error {
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		log.Printf("Error saving uploaded file: %v", err)
		return err
	}
	return nil
}

func GenerateThumbnail(videoPath, thumbnailPath string) error {
	ffmpegPath := "ffmpeg" // Ensure ffmpeg is in your PATH
	cmd := exec.Command(ffmpegPath, "-i", videoPath, "-ss", "00:00:01.000", "-vframes", "1", thumbnailPath)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	log.Printf("Running FFmpeg command: %s", cmd.String())
	err := cmd.Run()
	if err != nil {
		log.Printf("FFmpeg output: %s", out.String())
		log.Printf("FFmpeg error: %s", stderr.String())
		log.Printf("Error running FFmpeg command: %v", err)
		return err
	}
	log.Printf("FFmpeg output: %s", out.String())
	return nil
}
