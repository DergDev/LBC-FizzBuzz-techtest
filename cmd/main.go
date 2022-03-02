package main

import (
	"fmt"
	"log"
	"net/http"
	"techtest/internal/controllers"

	"github.com/gorilla/mux"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/fizz-buzz", controllers.GetFizzBuzz).Methods("GET")
	//router.HandleFunc("/all", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
