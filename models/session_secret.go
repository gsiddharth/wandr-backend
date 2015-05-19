package models

import (
	"github.com/gsiddharth/wandr/errors"
	"github.com/gsiddharth/wandr/utils"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

//TODO : Implement Session Expiry

type SessionSecret struct {
	gorm.Model

	UserID uint   `sql:unique json:"-"`
	Secret string `sql:"size:255"`
	Key    string `sql:"size:255;unique"`
	Expiry time.Time
}

func Authenticate(key string, secret string) (uint, error) {

	var sessionSecret SessionSecret

	if err0 := DB.Where("key = ? and secret = ?", key, secret).First(&sessionSecret).Error; err0 == nil {

		return sessionSecret.UserID, nil

	} else {
		return 0, &errors.Error{Description: "Session Authentication Failed",
			Code: errors.SESSION_AUTHENTICATION_ERROR}
	}
}

func GetSessionSecretByUserID(userId uint) (*SessionSecret, error) {
	var sessionSecret SessionSecret
	if err := DB.Where("user_id = ?", userId).First(&sessionSecret).Error; err == nil {
		return &sessionSecret, nil
	} else {
		return nil, &errors.Error{Description: "Session Secret not found!",
			Code: errors.DB_GET_ERROR}
	}
}

func SetSessionSecretProperties(sessionSecret *SessionSecret) *SessionSecret {

	sessionSecret.Key = utils.RandomString(utils.KEY_LENGTH)
	sessionSecret.Secret = utils.RandomString(utils.SECRET_LENGTH)

	return sessionSecret
}

func NewSessionSecret() *SessionSecret {

	sessionSecret := &SessionSecret{}
	sessionSecret = SetSessionSecretProperties(sessionSecret)

	return sessionSecret
}

func UpdateSessionSecret(userId uint) (*SessionSecret, error) {

	if sessionSecret, err0 := GetSessionSecretByUserID(userId); err0 == nil {
		//found the secret updating now
		SetSessionSecretProperties(sessionSecret)
		if err1 := DB.Model(sessionSecret).UpdateColumns(sessionSecret).Error; err1 != nil {
			log.Fatal(err1)
			return nil, &errors.Error{Description: "Error while updating new session secret",
				Code: errors.SESSION_SECRET_CREATE_ERROR}
		} else {
			return sessionSecret, nil
		}
	} else {
		return nil, &errors.Error{Description: "Error while creating new session secret",
			Code: errors.SESSION_SECRET_CREATE_ERROR}
	}
}
