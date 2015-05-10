package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/gsiddharth/wandr/controllers"
	"github.com/gsiddharth/wandr/models"
	"net/http"
)

func main() {
	r := InitRouter()
	err := models.Init("sqlite3", "db/wandr.db")
	if err == nil {
		n := negroni.Classic()
		n.UseHandler(r)
		n.Run(":3000")
	} else {
		fmt.Print(err.Error())
	}
}

func InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/login/{username}/{password}", controllers.Login).Methods("GET")
	r.HandleFunc("/register/{username}/{password}/{name}/{email}/{phone}/{gender}", controllers.Register).Methods("GET")
	http.Handle("/", r)

	return r
}
