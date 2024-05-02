package main

import (
	"net/http"

	"github.com/GabrielFMPinheiro/tracing-golang/service_a/api/handler"
)

type RequestData struct {
	Zipcode string `json:"zipcode"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /zipcode", handler.ZipCodeValidate)
	http.ListenAndServe(":8080", mux)
}
