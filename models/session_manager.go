package models

import (
	"github.com/gorilla/sessions"
	"github.com/gsiddharth/wandr/errors"
	"net/http"
)

var COOKIE_STORE_SECRET = "all that is gold doesn't glitter"

var SESSION_NAME = "wandr"

var STORE = sessions.NewCookieStore([]byte(COOKIE_STORE_SECRET))

const (
	CURRENT_USER_ID_KEY = iota
)

func GetSession(rw http.ResponseWriter, r *http.Request) *sessions.Session {
	session, _ := STORE.Get(r, SESSION_NAME)
	return session
}

func Store(rw http.ResponseWriter, r *http.Request, key interface{}, value interface{}) {
	session := GetSession(rw, r)
	session.Values[key] = value
	session.Save(r, rw)
}

func Get(rw http.ResponseWriter, r *http.Request, key interface{}) interface{} {
	session := GetSession(rw, r)
	return session.Values[key]
}

func StoreString(rw http.ResponseWriter, r *http.Request, key interface{}, value string) {
	Store(rw, r, key, value)
}

func StoreInt(rw http.ResponseWriter, r *http.Request, key interface{}, value int) {
	Store(rw, r, key, value)
}

func StoreUInt(rw http.ResponseWriter, r *http.Request, key interface{}, value uint) {
	Store(rw, r, key, value)
}

func StoreFloat64(rw http.ResponseWriter, r *http.Request, key interface{}, value float64) {
	Store(rw, r, key, value)
}

func GetString(rw http.ResponseWriter, r *http.Request, key interface{}) (string, error) {
	value, found := Get(rw, r, key).(string)
	if found {
		return value, nil
	} else {
		return "", &errors.Error{Description: "Key not found", Code: errors.SESSION_KEY_ERROR}
	}
}

func GetInt(rw http.ResponseWriter, r *http.Request, key interface{}) (int, error) {
	value, found := Get(rw, r, key).(int)
	if found {
		return value, nil
	} else {
		return 0, &errors.Error{Description: "Key not found", Code: errors.SESSION_KEY_ERROR}
	}
}

func GetUInt(rw http.ResponseWriter, r *http.Request, key interface{}) (uint, error) {
	value, found := Get(rw, r, key).(uint)
	if found {
		return value, nil
	} else {
		return 0, &errors.Error{Description: "Key not found", Code: errors.SESSION_KEY_ERROR}
	}
}

func GetFloat64(rw http.ResponseWriter, r *http.Request, key interface{}) (float64, error) {
	value, found := Get(rw, r, key).(float64)
	if found {
		return value, nil
	} else {
		return 0, &errors.Error{Description: "Key not found", Code: errors.SESSION_KEY_ERROR}
	}
}

func SetCurrentUser(rw http.ResponseWriter, r *http.Request, user *User) {
	Store(rw, r, CURRENT_USER_ID_KEY, user.ID)
}

func GetCurrentUser(rw http.ResponseWriter, r *http.Request) (*User, error) {
	id, err0 := GetUInt(rw, r, CURRENT_USER_ID_KEY)

	if err0 == nil {
		user, err1 := GetUserById(id)
		if err1 == nil {
			return user, nil
		} else {
			return nil, &errors.Error{Description: "User not logged In", Code: errors.USER_NOT_LOGGED_IN}
		}
	} else {
		return nil, &errors.Error{Description: "User not logged In", Code: errors.USER_NOT_LOGGED_IN}
	}
}
