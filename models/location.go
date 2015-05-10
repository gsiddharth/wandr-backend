package models

import (
	"github.com/gsiddharth/wandr/errors"
	"github.com/jinzhu/gorm"
)

type Location struct {
	gorm.Model

	City      string `sql:"size:255"`
	Longitude float64
	Latitude  float64
}

func GetLocationById(id uint) (*Location, *errors.DBError) {

	var location Location

	if err := DB.Where("Id = ?", id).First(&location).Error; err == nil {
		return &location, nil
	} else {
		return nil, &errors.DBError{Description: "Location not found!", Code: errors.DB_GET_ERROR}
	}
}

func AddNewLocation(location *Location) (*Location, *errors.DBError) {

	if err := DB.Save(location).Error; err == nil {
		return location, nil
	} else {
		return nil, &errors.DBError{Description: "Location update error!", Code: errors.DB_SAVE_ERROR}
	}
}
