package weather

import (
	"encoding/json"
	"fmt"
	"golang_telegram_bot/internal/models/weather"
	"io"
	"net/http"
	"os"
)

const WEATHER_URL = "https://api.openweathermap.org/data/2.5/weather"

func GetWeatherByGeo(latitude float64, longitude float64) *weather.Weather {
	response, err := http.Get(
		WEATHER_URL + "?lat=" + fmt.Sprintf("%f", latitude) + "&lon=" +
			fmt.Sprintf("%f", longitude) + "&units=metric&appid=" + os.Getenv("OPEN_WEATHER_MAP_TOKEN") +
			"&lang=ru",
	)
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	weather := new(weather.Weather)
	json.Unmarshal(body, &weather)

	return weather
}
