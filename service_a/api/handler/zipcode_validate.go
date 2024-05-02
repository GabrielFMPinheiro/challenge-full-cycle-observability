package handler

import (
	"encoding/json"
	"net/http"

	"gopkg.in/Nhanderu/brdoc.v1"
)

type RequestData struct {
	Zipcode string `json:"zipcode"`
}

func ZipCodeValidate(w http.ResponseWriter, r *http.Request) {
	var body RequestData
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !brdoc.IsCEP(body.Zipcode) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}
}
