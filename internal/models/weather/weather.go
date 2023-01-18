package weather

type Weather struct {
	Coordinate  Coordinate    `json:"coord"`
	WeatherInfo []WeatherInfo `json:"weather"`
	Main        MainInfo      `json:"main"`
	Wind        Wind          `json:"wind"`
	Clouds      Clouds        `json:"clouds"`
	Country     Country       `json:"sys"`
	City        string        `json:"name"`
}

type Coordinate struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

type WeatherInfo struct {
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type MainInfo struct {
	Temperature          float32 `json:"temp"`
	TemperatureFeelsLike float32 `json:"feels_like"` // Температура ощущается как
	TemperatureMin       float32 `json:"temp_min"`
	TemperatureMax       float32 `json:"temp_max"`
	Humidity             int     `json:"humidity"` // Влажность
}

type Wind struct {
	Speed     float64 `json:"speed"`
	Direction int     `json:"deg"`  // Направление ветра, градусы (метеорологические)
	Gust      float64 `json:"gust"` // Порыв ветра. Единица измерения по умолчанию: метр/сек,
}

type Clouds struct {
	CloudCover int `json:"all"`
}

type Country struct {
	Code string `json:"country"`
}
