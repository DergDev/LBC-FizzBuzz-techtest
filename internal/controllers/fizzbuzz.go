package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"techtest/pkg/fizzbuzz_core"
)

func GetFizzBuzz(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	int1, err := strconv.Atoi(vars.Get("int1"))
	if err != nil {
		json.NewEncoder(w).Encode("int1 query parameter has an invalid syntax.")
		return
	}
	int2, err := strconv.Atoi(vars.Get("int2"))
	if err != nil {
		json.NewEncoder(w).Encode("int2 query parameter has an invalid syntax.")
		return
	}
	limit, err := strconv.Atoi(vars.Get("limit"))
	if err != nil {
		json.NewEncoder(w).Encode("limit query parameter has an invalid syntax.")
		return
	}

	fizz := vars.Get("str1")
	buzz := vars.Get("str2")

	response := strings.Join(fizzbuzz_core.ComputeFizzBuzz(uint8(int1), uint8(int2), uint8(limit), fizz, buzz), ", ")

	json.NewEncoder(w).Encode(response)
}
