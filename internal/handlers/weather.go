package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/brunoofgod/goexpert-lesson-4/internal/services"
)

type WeatherRequest struct {
	CEP string `json:"cep"`
}

type WeatherResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

// GetWeather processa a requisição do usuário
// @Summary Obtém a temperatura de uma cidade a partir do CEP
// @Description Retorna a temperatura em Celsius, Fahrenheit e Kelvin
// @Tags Clima
// @Accept json
// @Produce json
// @Param request body WeatherRequest true "CEP para consulta"
// @Success 200 {object} WeatherResponse
// @Failure 422 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /weather [post]
func GetWeather(w http.ResponseWriter, r *http.Request) {
	var req WeatherRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if len(req.CEP) != 8 {
		http.Error(w, `{"message": "invalid zipcode"}`, http.StatusUnprocessableEntity)
		return
	}

	city, err := services.GetCityByZip(req.CEP)
	if err != nil {
		http.Error(w, `{"message": "can not find zipcode "}`+err.Error(), http.StatusNotFound)
		return
	}

	tempC, err := services.GetWeatherByCity(http.DefaultClient, city)
	if err != nil {
		http.Error(w, `{"message": "error fetching weather"}`, http.StatusInternalServerError)
		return
	}

	response := WeatherResponse{
		TempC: tempC,
		TempF: tempC*1.8 + 32,
		TempK: tempC + 273,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
