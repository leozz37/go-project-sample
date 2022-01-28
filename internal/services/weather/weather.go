package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type CityWeather struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity float64 `json:"humidity"`
	}
}

func GetWeatherByCityName(cityName string) (float64, error) {
	owaClientID := os.Getenv("OWA_CLIENT_ID")
	owaURL := os.Getenv("OWA_URL") + fmt.Sprintf("q=%s&appid=%s&units=metric", cityName, owaClientID)

	resp, err := http.Get(owaURL)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var weather CityWeather
	json.NewDecoder(resp.Body).Decode(&weather)

	return weather.Main.Temp, nil
}
