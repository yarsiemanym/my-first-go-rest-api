package main

import (
	"fmt"
	"my-first-go-rest-api/app"
	"my-first-go-rest-api/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET")

	router.Use(app.JwtAuthentication)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Print(err)
	}
}
