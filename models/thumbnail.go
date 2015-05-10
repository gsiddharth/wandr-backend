package models

import (
	"github.com/gsiddharth/wandr/errors"
	"github.com/jinzhu/gorm"
)

type Thumbnail struct {
	gorm.Model

	Url       string `sql:"size:255"`
	VideoId   uint
	Width     uint
	Height    uint
	SizeInKbs uint
}

func GetThumbnailById(id uint) (*Thumbnail, *errors.DBError) {

	var thumbnail Thumbnail

	if err := DB.Where("Id = ?", id).First(&thumbnail).Error; err != nil {
		return &thumbnail, nil
	} else {
		return nil, &errors.DBError{Description: "User not found!", Code: errors.DB_GET_ERROR}
	}
}
