package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	UserID        uint
	FilePath      string
	ThumbnailPath string
	Description   string
}
