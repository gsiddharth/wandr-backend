package models

import (
	"crypto/md5"
	"github.com/gsiddharth/wandr/errors"
	"github.com/gsiddharth/wandr/utils"
	"github.com/jinzhu/gorm"
	"io"
	"log"
)

type User struct {
	gorm.Model

	Username      string      `sql:"size:255;unique"`
	Password      string      `sql:"size:255" json:"-"`
	Email         string      `sql:"size:255;unique"`
	UserProfile   UserProfile // One to One relationship
	Videos        []Video     //One to many relationship
	SessionSecret SessionSecret
}

func AddUser(userName string, password string, email string, userProfile *UserProfile) (*User, error) {

	user := &User{Username: userName, Password: utils.MD5Hash(password), Email: email, UserProfile: *userProfile, SessionSecret: *NewSessionSecret()}

	if _, err1 := GetUserByUserName(userName); err1 != nil { // username doesn't exist
		if _, err2 := GetUserByEmail(email); err2 != nil { // email doesn't
			if err3 := DB.Save(user).Error; err3 != nil {
				log.Fatal(err3)
				return nil, &errors.Error{Description: "User not saved!", Code: errors.DB_SAVE_ERROR}
			} else {
				return user, nil
			}
		} else {
			return nil, &errors.Error{Description: "Email already exists!", Code: errors.REGISTER_EMAIL_NOT_UNIQUE}
		}

	} else {
		return nil, &errors.Error{Description: "Username already exists!", Code: errors.REGISTER_USERNAME_NOT_UNIQUE}
	}
}

func GetUserByUserName(userName string) (*User, error) {

	var user User
	if err := DB.Where("username = ?", userName).First(&user).Error; err == nil {
		DB.Model(&user).Related(&user.SessionSecret)
		return &user, nil
	} else {
		return nil, &errors.Error{Description: "User not found!", Code: errors.DB_GET_ERROR}
	}
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	if err := DB.Where("email = ?", email).First(&user).Error; err == nil {
		DB.Model(&user).Related(&user.SessionSecret)
		return &user, nil
	} else {
		return nil, &errors.Error{Description: "User not found!", Code: errors.DB_GET_ERROR}
	}
}

func GetUserById(id uint) (*User, error) {

	var user User

	if err := DB.Where("id = ?", id).First(&user).Error; err == nil {
		DB.Model(&user).Related(&user.SessionSecret)
		return &user, nil
	} else {
		return nil, &errors.Error{Description: "User not found!", Code: errors.DB_GET_ERROR}
	}
}

func AuthenticateUser(userName, password string) (*User, error) {

	pMd5 := md5.New()
	io.WriteString(pMd5, password)

	user := &User{Username: userName, Password: string(pMd5.Sum(nil))}

	if err := DB.Where(user).First(user).Error; err != nil {
		return nil, &errors.Error{Description: "User not found!", Code: errors.DB_GET_ERROR}
	} else {
		DB.Model(&user).Related(&user.SessionSecret)
		return user, nil
	}

}

func ChangePassword(userName, oldpassword, newpassword string) (*User, error) {

	user, _ := AuthenticateUser(userName, oldpassword)

	if user != nil {

		user.Password = newpassword

		if err2 := DB.Save(user).Error; err2 == nil {
			return user, nil
		} else {
			return nil, &errors.Error{Description: "Password Update Error", Code: errors.DB_UPDATE_ERROR}
		}

	} else {
		return nil, &errors.Error{Description: "Password Update Error", Code: errors.DB_UPDATE_ERROR}
	}
}
