package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func apiBase(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Api is online :)")
}

func JsonTest(w http.ResponseWriter, r *http.Request) {
	problems := ResponsePayload{
		"test",
		[]Question{
			{"This is a test Problem", "This is the answer"},
			{"This is also a test problem", "Also an Asnwer"},
		},
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(problems); err != nil {
		panic(err)
	}
}
