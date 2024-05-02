package handler

import (
	"encoding/json"
	"net/http"

	"github.com/GabrielFMPinheiro/tracing-golang/service_b/domain"
	"github.com/GabrielFMPinheiro/tracing-golang/service_b/integration"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type Body struct {
	Zipcode string `json:"zipcode"`
}

func CalculateTemperature(w http.ResponseWriter, r *http.Request) {
	var body Body

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	marshalled, err := json.Marshal(body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	err = domain.ZipcodeValidate(ctx, w, marshalled)

	defer r.Body.Close()

	if err != nil && err.Error() == "invalid zipcode" {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	zipcodeService := integration.ZipcodeIntegration{}
	addr, err := zipcodeService.GetZipcode(body.Zipcode, ctx)

	if (integration.Address{}) == *addr {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	weatherService := integration.WeatherIntegration{}
	weather, err := weatherService.GetWeather(addr.City, ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(weather)
}
