package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"techtest/internal/models"
	"techtest/internal/repository"
	"techtest/pkg/fizzbuzz_core"
)

//Sanity check and Error handling before Fizz-Buzz compution
func GetFizzBuzz(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	int1, err := strconv.Atoi(vars.Get("int1"))
	if err != nil {
		http.Error(w, "int1 query parameter has an invalid syntax.", http.StatusBadRequest)
		return
	} else if int1 <= 0 {
		http.Error(w, "int1 can't be less or equal to 0.", http.StatusBadRequest)
		return
	}
	int2, err := strconv.Atoi(vars.Get("int2"))
	if err != nil {
		http.Error(w, "int2 query parameter has an invalid syntax.", http.StatusBadRequest)
		return
	} else if int2 <= 0 {
		http.Error(w, "int2 can't be less or equal to 0.", http.StatusBadRequest)
		return
	}
	limit, err := strconv.Atoi(vars.Get("limit"))
	if err != nil {
		http.Error(w, "limit query parameter has an invalid syntax.", http.StatusBadRequest)
		return
	} else if limit <= 0 {
		http.Error(w, "limit can't be less or equal to 0.", http.StatusBadRequest)
		return
	}
	str1 := vars.Get("str1")
	if str1 == "" {
		http.Error(w, "str1 query parameter is missing or empty.", http.StatusBadRequest)
		return
	}
	str2 := vars.Get("str2")
	if str2 == "" {
		http.Error(w, "str2 query parameter is missing or empty.", http.StatusBadRequest)
		return
	}

	repository.AddQueryEntryToDB(r.URL.RawQuery)
	var response models.Fizzbuzz
	response.Content = fizzbuzz_core.ComputeFizzBuzz(uint8(int1), uint8(int2), uint8(limit), str1, str2)

	json.NewEncoder(w).Encode(response)
}
