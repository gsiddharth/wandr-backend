package models

import (
	"github.com/gsiddharth/wandr/errors"
	"github.com/jinzhu/gorm"
)

type UserProfile struct {
	gorm.Model

	UserId   uint
	Name     string   `sql:"size:255"`
	Email    string   `sql:"size:255"`
	Phone    string   `sql:"size:20"`
	Gender   string   `sql:"size:2"`
	Location Location // One to One relationship
}

func AddUserProfile(user *User, name string, email string,
	phone string, location *Location) (*UserProfile, *errors.DBError) {

	userProfile := &UserProfile{UserId: user.ID, Phone: phone, Location: *location}

	if err := DB.Save(userProfile).Error; err == nil {
		return userProfile, nil
	} else {
		return nil, &errors.DBError{Description: "Profile Save Error!", Code: errors.DB_SAVE_ERROR}
	}
}

func UpdateUserProfile(newProfile *UserProfile) (*UserProfile, *errors.DBError) {

	oldProfile, err0 := GetUserProfileByUserId(newProfile.UserId)

	if oldProfile != nil && newProfile != nil && err0 == nil {
		if newProfile.Name != "" {
			oldProfile.Name = newProfile.Name
		}

		if newProfile.Phone != "" {
			oldProfile.Phone = newProfile.Phone
		}

		if newProfile.Gender != "" {
			oldProfile.Gender = newProfile.Gender
		}

		if newProfile.Location.ID != oldProfile.Location.ID {
			var l *Location
			l, err1 := AddNewLocation(&newProfile.Location)

			if err1 != nil {
				return nil, &errors.DBError{Description: "Location add error!", Code: errors.DB_SAVE_ERROR}
			} else {
				oldProfile.Location = *l
			}
		}

		if err2 := DB.Save(oldProfile).Error; err2 == nil {
			return oldProfile, nil
		} else {
			return nil, &errors.DBError{Description: "Profile update error!", Code: errors.DB_UPDATE_ERROR}
		}
	} else {
		return nil, &errors.DBError{Description: "Profile update error!", Code: errors.DB_UPDATE_ERROR}
	}
}

func GetUserProfileById(id uint) (*UserProfile, *errors.DBError) {

	var userProfile UserProfile

	if err := DB.Where("Id = ?", id).First(&userProfile).Error; err == nil {
		return &userProfile, nil
	} else {
		return nil, &errors.DBError{Description: "Profile not found!", Code: errors.DB_GET_ERROR}
	}
}

func GetUserProfileByUserId(userId uint) (*UserProfile, *errors.DBError) {

	userProfile := &UserProfile{UserId: userId}

	if err := DB.Where(userProfile).First(userProfile).Error; err == nil {
		return userProfile, nil
	} else {
		return nil, &errors.DBError{Description: "Profile not found!", Code: errors.DB_GET_ERROR}
	}
}

func GetUserProfileByUsername(userName string) (*UserProfile, *errors.DBError) {
	user, _ := GetUserByUserName(userName)
	return GetUserProfileByUserId(user.ID)
}
