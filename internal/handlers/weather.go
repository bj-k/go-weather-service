package handlers

import (
	"encoding/json"
	"expvar"
	"net/http"
	"strings"

	"go-weather-service/internal/models"
)

var (
	GetWeatherRequests            = expvar.NewInt("weather_requests_total")
	GetSupportedLocationsRequests = expvar.NewInt("supported_locations_requests_total")
)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	GetWeatherRequests.Add(1)
	location := r.URL.Query().Get("location")
	if location == "" {
		http.Error(w, "Missing location parameter", http.StatusBadRequest)
		return
	}

	data, ok := models.WeatherData[strings.ToLower(location)]
	if !ok {
		http.Error(w, "Location not supported", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GetSupportedLocations(w http.ResponseWriter, r *http.Request) {
	GetSupportedLocationsRequests.Add(1)
	locations := make([]string, 0, len(models.WeatherData))
	for _, data := range models.WeatherData {
		locations = append(locations, data.Location)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(locations)
}

func GetOpenAPI(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "api/openapi.yaml")
}
