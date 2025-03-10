package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/felippemozer/telelaudo-teste-golang-api/internal/contract"
)

type Service interface {
	GetBy(city string) (*Weather, error)
}

func GetBy(city string) (*Weather, error) {
	spplitedCity := strings.Split(city, " ")
	if len(spplitedCity) > 1 {
		city = strings.Join(spplitedCity, "+")
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&lang=pt_br&appid=%s", city, os.Getenv("OPENWEATHERMAP_API_KEY"))
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error on HTTP request: %s", err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode >= http.StatusBadRequest {
		errorBody, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("error on get weather info: %s", string(errorBody))
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error on parse weather data: %s", err.Error())
	}

	var jsonData contract.GetWeatherByCityRequestDTO
	json.Unmarshal(data, &jsonData)

	var weatherInfo []WeatherInfo
	for _, v := range jsonData.Weather {
		weatherInfo = append(weatherInfo, WeatherInfo{
			Id:          v.Id,
			Main:        v.Main,
			Description: v.Description,
			Icon:        v.Icon,
		})
	}

	weather := NewWeather(
		jsonData.Id,
		jsonData.Name,
		jsonData.Coordinates.Latitude,
		jsonData.Coordinates.Longitude,
		weatherInfo,
		jsonData.Base,
		jsonData.MainInfo.Temperature,
		jsonData.MainInfo.FeelsLike,
		jsonData.MainInfo.Minimum,
		jsonData.MainInfo.Maximun,
		jsonData.MainInfo.Pressure,
		jsonData.MainInfo.Humidity,
		jsonData.MainInfo.SeaLevel,
		jsonData.MainInfo.GroundLevel,
		jsonData.Visibility,
		jsonData.Wind.Speed,
		jsonData.Wind.Deg,
		jsonData.Wind.Gust,
		jsonData.Clouds.All,
		jsonData.Rain.OneHour,
		jsonData.Snow.OneHour,
		jsonData.Datetime,
		jsonData.Sys.Type,
		jsonData.Sys.Id,
		jsonData.Sys.Country,
		jsonData.Sys.Sunset,
		jsonData.Sys.Sunrise,
		jsonData.Timezone,
	)
	return weather, nil
}
