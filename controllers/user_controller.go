package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gsiddharth/wandr/errors"
	"github.com/gsiddharth/wandr/models"
	"net/http"
)

func Register(rw http.ResponseWriter, r *http.Request) {

	userName := mux.Vars(r)["username"]
	password := mux.Vars(r)["password"]
	name := mux.Vars(r)["name"]
	email := mux.Vars(r)["email"]
	phone := mux.Vars(r)["phone"]
	gender := mux.Vars(r)["gender"]
	user, _ := models.AddUser(userName, password, &models.UserProfile{Name: name, Email: email, Phone: phone, Gender: gender})

	js, _ := json.Marshal(user)
	rw.Write(js)

}

func Login(rw http.ResponseWriter, r *http.Request) {
	userName := mux.Vars(r)["username"]
	password := mux.Vars(r)["password"]

	user, _ := models.AuthenticateUser(userName, password)
	if user != nil {
		js, _ := json.Marshal(user)
		rw.Write(js)
	} else {
		err := errors.UserError{Description: "Username/Password incorrect", Code: errors.USER_LOGIN_PASSWORD_INCORRECT}
		js, _ := json.Marshal(err)
		rw.Write(js)
	}
}

func UpdateProfile(rw http.ResponseWriter, r *http.Request) {

}
