package models

import (
	"crypto/md5"
	"github.com/gsiddharth/wandr/errors"
	"github.com/jinzhu/gorm"
	"io"
)

type User struct {
	gorm.Model

	Username    string      `sql:"size:255;unique"`
	Password    string      `sql:"size:255" json:"-"`
	UserProfile UserProfile // One to One relationship
	Videos      []Video     //One to many relationship
}

func AddUser(userName string, password string, userProfile *UserProfile) (*User, *errors.DBError) {

	p_md5 := md5.New()
	io.WriteString(p_md5, password)

	user := &User{Username: userName, Password: string(p_md5.Sum(nil)), UserProfile: *userProfile}

	if err := DB.Save(user).Error; err != nil {
		return nil, &errors.DBError{Description: "User not saved!", Code: errors.DB_SAVE_ERROR}
	} else {
		return user, nil
	}
}

func GetUserByUserName(userName string) (*User, *errors.DBError) {

	var user User
	if err := DB.Where("username = ?", userName).First(&user).Error; err == nil {
		return &user, nil
	} else {
		return nil, &errors.DBError{Description: "User not found!", Code: errors.DB_GET_ERROR}
	}
}

func GetUserById(id uint) (*User, *errors.DBError) {

	var user User

	if err := DB.Where("Id = ?", id).First(&user).Error; err == nil {
		return &user, nil
	} else {
		return nil, &errors.DBError{Description: "User not found!", Code: errors.DB_GET_ERROR}
	}
}

func AuthenticateUser(userName, password string) (*User, *errors.DBError) {

	pMd5 := md5.New()
	io.WriteString(pMd5, password)

	user := &User{Username: userName, Password: string(pMd5.Sum(nil))}

	if err := DB.Where(user).First(user).Error; err != nil {
		return nil, &errors.DBError{Description: "User not found!", Code: errors.DB_GET_ERROR}
	} else {
		return user, nil
	}

}

func ChangePassword(userName, oldpassword, newpassword string) (*User, *errors.DBError) {

	user, _ := AuthenticateUser(userName, oldpassword)

	if user != nil {

		user.Password = newpassword

		if err2 := DB.Save(user).Error; err2 == nil {
			return user, nil
		} else {
			return nil, &errors.DBError{Description: "Password Update Error", Code: errors.DB_UPDATE_ERROR}
		}

	} else {
		return nil, &errors.DBError{Description: "Password Update Error", Code: errors.DB_UPDATE_ERROR}
	}
}
