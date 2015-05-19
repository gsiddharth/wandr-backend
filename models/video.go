package models

import (
	"github.com/gsiddharth/wandr/errors"
	"github.com/jinzhu/gorm"
	"time"
)

type Video struct {
	gorm.Model

	UserID          uint
	Url             string `sql:"size:255"`
	LengthInSeconds float64
	TimeOfVideo     time.Time
	SizeInBytes     uint
	Location        Location    //One to one relationship
	Thumbnails      []Thumbnail //One to many relationship
}

func GetVideoById(id uint) (*Video, error) {

	var video Video

	if err := DB.Where("id = ?", id).First(&video).Error; err != nil {
		DB.Model(&video).Related(&video.Thumbnails)
		DB.Model(&video).Related(&video.Location)
		return &video, nil
	} else {
		return nil, &errors.Error{Description: "Video not found!", Code: errors.DB_GET_ERROR}
	}
}

func AddVideo(video *Video) (*Video, error) {

	if err := DB.Save(video).Error; err == nil {
		return video, nil
	} else {
		return nil, &errors.Error{Description: "Video add Error!", Code: errors.DB_SAVE_ERROR}
	}

}

func GetVideosByUserID(userId uint) ([]Video, error) {
	videos := []Video{}

	if err := DB.Where("user_id = ?", userId).Find(&videos).Error; err == nil {
		for _, video := range videos {
			DB.Model(&video).Related(&video.Thumbnails)
			DB.Model(&video).Related(&video.Location)
		}
		return videos, nil
	} else {
		return nil, &errors.Error{Description: "Error finding Videos!", Code: errors.DB_GET_ERROR}
	}
}

func GetVideosOfUser(user *User) ([]Video, error) {
	DB.Model(&user).Related(&user.Videos)

	for _, video := range user.Videos {
		DB.Model(&video).Related(&video.Thumbnails)
		DB.Model(&video).Related(&video.Location)
	}

	return user.Videos, nil
}

func GetNearbyVideos(user *User, location *Location) ([]Video, error) {
	return GetVideosOfUser(user)
}
