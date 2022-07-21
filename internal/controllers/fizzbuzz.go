package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"techtest/internal/models"
	"techtest/internal/repository"
	"techtest/pkg/fizzbuzz_core"
)

const stringErrorMessage = "One of the string parameter was missing or empty, please check your query."
const intErrorMessage = "One of the numerical parameter was below or equal to 0, please check your query."
const parameterErrorMessage = "Something went wrong while getting your numerical parameters, please check your query, you might have assigned a wrong type to an int."

//Sanity check and Error handling before Fizz-Buzz compution
func GetFizzBuzz(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	int1, errInt1 := strconv.Atoi(vars.Get("int1"))
	int2, errInt2 := strconv.Atoi(vars.Get("int2"))
	limit, errLimit := strconv.Atoi(vars.Get("limit"))
	str1 := vars.Get("str1")
	str2 := vars.Get("str2")
	if errInt1 != nil || errInt2 != nil || errLimit != nil {
		http.Error(w, parameterErrorMessage, http.StatusBadRequest)
		return
	}
	if int1 <= 0 || int2 <= 0 || limit <= 0 {
		http.Error(w, intErrorMessage, http.StatusBadRequest)
		return
	}
	if str1 == "" || str2 == "" {
		http.Error(w, stringErrorMessage, http.StatusBadRequest)
		return
	}

	repository.AddQueryEntryToDB(r.URL.RawQuery)
	var response models.Fizzbuzz
	response.Content = fizzbuzz_core.ComputeFizzBuzz(int1, int2, limit, str1, str2)

	json.NewEncoder(w).Encode(response)
}
