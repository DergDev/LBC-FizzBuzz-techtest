package main

import (
	"fmt"
	"log"
	"net/http"
	"techtest/internal/controllers"

	"github.com/gorilla/mux"
)

//Router Setup
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/fb-api/fizz-buzz", controllers.GetFizzBuzz).Methods("GET")
	router.HandleFunc("/fb-api/statistics", controllers.GetApiStatistics).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	fmt.Println("Fizz-buzz API v1.0")
	handleRequests()
}
