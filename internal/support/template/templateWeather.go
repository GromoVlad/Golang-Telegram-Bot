package template

import (
	"fmt"
	"golang_telegram_bot/internal/models/weather"
	"math"
	"strconv"
	"strings"
)

func GetWeatherIcon(code string) string {
	switch {
	case code == "01d" || code == "01n":
		return "☀"
	case code == "02d" || code == "02n":
		return "⛅"
	case code == "03d" || code == "03n" || code == "04d" || code == "04n":
		return "☁"
	case code == "09d" || code == "09n" || code == "10d" || code == "10n":
		return "🌧"
	case code == "11d" || code == "11n":
		return "⛈"
	case code == "13d" || code == "13n":
		return "❄"
	case code == "50d" || code == "50n":
		return "🌫"
	default:
		return ""
	}
}

func GetWindDirection(direction int) string {
	switch {
	case direction >= 0 && direction < 25:
		return "C"
	case direction >= 25 && direction < 65:
		return "СЗ"
	case direction >= 65 && direction < 115:
		return "З"
	case direction >= 115 && direction < 155:
		return "ЮЗ"
	case direction >= 155 && direction < 205:
		return "Ю"
	case direction >= 205 && direction < 245:
		return "ЮВ"
	case direction >= 245 && direction < 295:
		return "В"
	case direction >= 295 && direction < 335:
		return "СВ"
	case direction >= 335 && direction <= 360:
		return "C"
	default:
		return "Неизвестно"
	}
}

func BuildWeatherTemplate(weather *weather.Weather) string {
	result := "🌄 Погода в " + weather.City + " [" + weather.Country.Code + "]\n\n"
	for _, info := range weather.WeatherInfo {
		result += strings.Title(info.Description) + " " + GetWeatherIcon(info.Icon) + "\n\n"
	}
	result += "🌡Температура: " + fmt.Sprintf("%.0f", math.Ceil(float64(weather.Main.Temperature))) + "°C\n" +
		"🌡Ощущается как: " + fmt.Sprintf("%.0f", math.Ceil(float64(weather.Main.TemperatureFeelsLike))) + "°C\n\n" +
		"🌬Скорость ветра: " + fmt.Sprintf("%.2f", weather.Wind.Speed) + "м/с\n" +
		"🧭Направление ветра: " + GetWindDirection(weather.Wind.Direction) + "\n\n" +
		"💧Влажность: " + strconv.Itoa(weather.Main.Humidity) + "%\n" +
		"☁Облачность: " + strconv.Itoa(weather.Clouds.CloudCover) + "%\n"

	return result
}
