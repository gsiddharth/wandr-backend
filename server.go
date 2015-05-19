package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/gsiddharth/wandr/controllers"
	"github.com/gsiddharth/wandr/models"
	"github.com/gsiddharth/wandr/security"
	"github.com/gsiddharth/wandr/utils"
	"net/http"
)

func main() {

	app := negroni.Classic()

	err := models.Init("sqlite3", "db/wandr.db")

	if err == nil {
		app.UseHandler(NewRouter())
		app.Run(":3000")

	} else {
		fmt.Print(err.Error())
	}
}

func NewRouter() *mux.Router {

	rStart := NewUnAuthenticatedRouter()
	rSecure := NewAuthenticatedRouter()

	nStart := negroni.New()
	nStart.UseHandler(rStart)

	nSecure := negroni.New()
	nSecure.Use(security.New("X-Auth-Key", "X-Auth-Secret", utils.USER_ID_KEY))
	nSecure.UseHandler(rSecure)

	mainRouter := mux.NewRouter().StrictSlash(true)
	mainRouter.Handle("/login/{username}/{password}", nStart)
	mainRouter.Handle("/register/{username}/{password}/{name}/{email}/{phone}/{gender}", nStart)
	mainRouter.Handle("/register/{username}/{password}/{name}/{email}/{phone}/", nStart)
	mainRouter.Handle("/register/{username}/{password}/{name}/{email}/", nStart)

	mainRouter.Handle("/video", nSecure)
	mainRouter.Handle("/authenticate", nSecure)
	mainRouter.Handle("/recommendedvideos/{city}/{longitude}/{latitude}", nSecure)

	http.Handle("/", mainRouter)
	return mainRouter

}

func NewUnAuthenticatedRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/login/{username}/{password}", controllers.Login).Methods("GET", "POST")
	r.HandleFunc("/register/{username}/{password}/{name}/{email}/{phone}/{gender}", controllers.Register).Methods("GET", "POST")
	r.HandleFunc("/register/{username}/{password}/{name}/{email}/{phone}/", controllers.Register).Methods("GET", "POST")
	r.HandleFunc("/register/{username}/{password}/{name}/{email}/", controllers.Register).Methods("GET", "POST")

	return r
}

func NewAuthenticatedRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/authenticate", controllers.Authenticate).Methods("GET", "POST")
	r.HandleFunc("/video", controllers.SaveVideoAndThumbnail).Methods("GET", "POST")
	r.HandleFunc("/recommendedvideos/{city}/{longitude}/{latitude}", controllers.Authenticate).Methods("GET", "POST")
	return r
}
