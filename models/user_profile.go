package models

import (
	"github.com/gsiddharth/wandr/errors"
	"github.com/jinzhu/gorm"
)

type UserProfile struct {
	gorm.Model

	UserID   uint
	Name     string   `sql:"size:255"`
	Phone    string   `sql:"size:20"`
	Gender   string   `sql:"size:2"`
	Location Location // One to One relationship
}

func AddUserProfile(user *User, name string, phone string,
	location *Location) (*UserProfile, error) {

	userProfile := &UserProfile{UserID: user.ID, Phone: phone, Location: *location}

	if err := DB.Save(userProfile).Error; err == nil {
		return userProfile, nil
	} else {
		return nil, &errors.Error{Description: "Profile Save Error!", Code: errors.DB_SAVE_ERROR}
	}
}

func UpdateUserProfile(newProfile *UserProfile) (*UserProfile, error) {

	oldProfile, err0 := GetUserProfileByUserId(newProfile.UserID)

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
				return nil, &errors.Error{Description: "Location add error!", Code: errors.DB_SAVE_ERROR}
			} else {
				oldProfile.Location = *l
			}
		}

		if err2 := DB.Save(oldProfile).Error; err2 == nil {
			return oldProfile, nil
		} else {
			return nil, &errors.Error{Description: "Profile update error!", Code: errors.DB_UPDATE_ERROR}
		}
	} else {
		return nil, &errors.Error{Description: "Profile update error!", Code: errors.DB_UPDATE_ERROR}
	}
}

func GetUserProfileById(id uint) (*UserProfile, error) {

	var userProfile UserProfile

	if err := DB.Where("Id = ?", id).First(&userProfile).Error; err == nil {
		return &userProfile, nil
	} else {
		return nil, &errors.Error{Description: "Profile not found!", Code: errors.DB_GET_ERROR}
	}
}

func GetUserProfileByUserId(userId uint) (*UserProfile, error) {

	userProfile := &UserProfile{UserID: userId}

	if err := DB.Where(userProfile).First(userProfile).Error; err == nil {
		return userProfile, nil
	} else {
		return nil, &errors.Error{Description: "Profile not found!", Code: errors.DB_GET_ERROR}
	}
}

func GetUserProfileByUsername(userName string) (*UserProfile, error) {
	user, _ := GetUserByUserName(userName)
	return GetUserProfileByUserId(user.ID)
}
