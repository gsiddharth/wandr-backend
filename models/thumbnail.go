package models

import (
	"github.com/gsiddharth/wandr/errors"
	"github.com/jinzhu/gorm"
)

type Thumbnail struct {
	gorm.Model

	Url         string `sql:"size:255"`
	VideoID     uint
	Width       uint
	Height      uint
	SizeInBytes uint
}

func GetThumbnailById(id uint) (*Thumbnail, error) {

	var thumbnail Thumbnail

	if err := DB.Where("id = ?", id).First(&thumbnail).Error; err != nil {
		return &thumbnail, nil
	} else {
		return nil, &errors.Error{Description: "Thumbnail not found!", Code: errors.DB_GET_ERROR}
	}
}

func GetThumbnailsByVideoID(videoId uint) ([]Thumbnail, error) {
	var thumbnails []Thumbnail

	if err := DB.Where("video_id = ?", videoId).Find(&thumbnails).Error; err != nil {
		return thumbnails, nil
	} else {
		return nil, &errors.Error{Description: "Thumnails not found!", Code: errors.DB_GET_ERROR}
	}
}
