package models

import (
	"github.com/gsiddharth/wandr/errors"
	"github.com/jinzhu/gorm"
	"time"
)

type Video struct {
	gorm.Model

	UserId          string `sql:"size:255"`
	Url             string `sql:"size:255"`
	LengthInSeconds float64
	TimeOfVideo     time.Time
	SizeInKbs       uint
	Location        Location    //One to one relationship
	Thumbnails      []Thumbnail //One to many relationship
}

func GetVideoById(id uint) (*Video, *errors.DBError) {

	var video Video

	if err := DB.Where("Id = ?", id).First(&video).Error; err != nil {
		return &video, nil
	} else {
		return nil, &errors.DBError{Description: "Video not found!", Code: errors.DB_GET_ERROR}
	}
}

func AddVideo(video *Video) (*Video, *errors.DBError) {

	if err := DB.Save(video).Error; err == nil {
		return video, nil
	} else {
		return nil, &errors.DBError{Description: "Video add Error!", Code: errors.DB_SAVE_ERROR}
	}

}

func GetVideosOfUser(userId uint) ([]User, *errors.DBError) {
	users := []User{}

	if err := DB.Where("Id = ?", userId).Find(&users).Error; err == nil {
		return users, nil
	} else {
		return nil, &errors.DBError{Description: "Error finding Videos!", Code: errors.DB_GET_ERROR}
	}
}
