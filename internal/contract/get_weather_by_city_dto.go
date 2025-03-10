package contract

type weatherCoord struct {
	Longitude float32 `json:"lon"`
	Latitude  float32 `json:"lat"`
}

type weatherInfo struct {
	Id          uint32 `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type weatherMain struct {
	Temperature float32 `json:"temp"`
	FeelsLike   float32 `json:"feels_like"`
	Minimum     float32 `json:"temp_min"`
	Maximun     float32 `json:"temp_max"`
	Pressure    uint32  `json:"pressure"`
	Humidity    uint8   `json:"humidity"`
	SeaLevel    uint32  `json:"sea_level"`
	GroundLevel uint32  `json:"grnd_level"`
}

type weatherWind struct {
	Speed float32 `json:"speed"`
	Deg   int8    `json:"deg"`
	Gust  float32 `json:"gust"`
}

type weatherCloud struct {
	All int8 `json:"all"`
}

type weatherRain struct {
	OneHour int32 `json:"1h"`
}

type weatherSnow struct {
	OneHour int32 `json:"1h"`
}

type weatherSys struct {
	Type    uint8  `json:"type"`
	Id      int64  `json:"id"`
	Country string `json:"country"`
	Sunset  int64  `json:"sunset"`
	Sunrise int64  `json:"sunrise"`
}

type GetWeatherByCityRequestDTO struct {
	Coordinates weatherCoord  `json:"coord"`
	Weather     []weatherInfo `json:"weather"`
	Base        string        `json:"base"`
	MainInfo    weatherMain   `json:"main"`
	Visibility  uint32        `json:"visibility"`
	Wind        weatherWind   `json:"wind"`
	Clouds      weatherCloud  `json:"clouds"`
	Rain        weatherRain   `json:"rain"`
	Snow        weatherSnow   `json:"snow"`
	Datetime    int64         `json:"dt"`
	Sys         weatherSys    `json:"sys"`
	Timezone    int64         `json:"timezone"`
	Id          uint32        `json:"id"`
	Name        string        `json:"name"`
}
