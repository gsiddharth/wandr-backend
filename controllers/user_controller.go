package controllers

import (
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gsiddharth/wandr/errors"
	"github.com/gsiddharth/wandr/models"
	"github.com/gsiddharth/wandr/utils"
	"net/http"
)

func Register(rw http.ResponseWriter, r *http.Request) {

	userName := mux.Vars(r)["username"]
	password := mux.Vars(r)["password"]
	name := mux.Vars(r)["name"]
	email := mux.Vars(r)["email"]
	phone := mux.Vars(r)["phone"]
	gender := mux.Vars(r)["gender"]
	user, err := models.AddUser(userName, password, email, &models.UserProfile{Name: name, Phone: phone, Gender: gender})

	if err == nil {
		models.NewOutput(user, "OK", http.StatusOK).Render(rw)
	} else {
		models.NewErrorOutput(errors.Error{Description: err.Error(), Code: http.StatusConflict}).Render(rw)
	}

}

func Login(rw http.ResponseWriter, r *http.Request) {
	userName := mux.Vars(r)["username"]
	password := mux.Vars(r)["password"]

	user, _ := models.AuthenticateUser(userName, password)
	if user != nil {
		models.NewOutput(user, "OK", http.StatusOK).Render(rw)
	} else {
		err := errors.Error{Description: "Error: Username/Password incorrect", Code: http.StatusNotFound}
		models.NewErrorOutput(err).Render(rw)
	}
}

func Authenticate(rw http.ResponseWriter, r *http.Request) {
	models.NewOutput("", "OK", http.StatusOK).Render(rw)
}

func CurrentUser(rw http.ResponseWriter, r *http.Request) (*models.User, error) {
	if userId := context.Get(r, utils.USER_ID_KEY); userId != nil {
		return models.GetUserById(userId.(uint))
	} else {
		return nil, &errors.Error{Description: "User not found", Code: errors.USER_NOT_FOUND_SESSION_KEY}
	}
}
