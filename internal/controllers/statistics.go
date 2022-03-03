package controllers

import (
	"encoding/json"
	"net/http"
	"reflect"
	"techtest/internal/repository"
)

//Get the most requested query from the DB
func GetApiStatistics(w http.ResponseWriter, r *http.Request) {
	if len(reflect.ValueOf(r.URL.Query()).MapKeys()) > 0 {
		http.Error(w, "This enpoint doesn't accept query parameters.", http.StatusBadRequest)
	} else {
		response, err := repository.GetQueryEntry()
		if err != nil {
			http.Error(w, "Statistics not found.", http.StatusNotFound)
		} else {
			json.NewEncoder(w).Encode(response)
		}
	}
}
