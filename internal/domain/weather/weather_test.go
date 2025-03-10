package weather

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWeather(t *testing.T) {
	assert := assert.New(t)
	var (
		id                         uint32  = 12564
		name                       string  = "Los Santos"
		latitude                   float32 = 34.014193
		longitude                  float32 = -118.344388
		weatherId                  uint32  = 1
		weatherMain                string  = "Clear"
		weatherDescription         string  = "Tempo aberto, bom pra danar"
		weatherIcon                string  = "c01n"
		base                       string  = "stations"
		temperature                float32 = 20.12
		feelsLike                  float32 = 16.23
		minimum                    float32 = 14.34
		maximum                    float32 = 21.45
		pressure                   uint32  = 1015
		humidity                   uint8   = 55
		seaLevel                   uint32  = 1015
		groundLevel                uint32  = 1015
		visibility                 uint32  = 10000
		speed                      float32 = 10.5
		deg                        int8    = 10
		gust                       float32 = 12.94
		cloudAll                   int8    = 10
		rain1h                     int32   = 0
		snow1h                     int32   = 0
		dt                         int64   = 1741471645
		sysType                    uint8   = 1
		sysId                      int64   = 15
		country                    string  = "US"
		sunset                     int64   = 1741461445
		sunrise                    int64   = 1741414645
		tz                         int64   = -25200
		pascal                     float32 = 1013.2
		kilometer                  uint32  = 1000
		kilometerPerHourMultiplier float32 = 3.6
	)

	w := NewWeather(
		id,
		name,
		latitude,
		longitude,
		[]WeatherInfo{
			{
				Id:          weatherId,
				Main:        weatherMain,
				Description: weatherDescription,
				Icon:        weatherIcon,
			},
		},
		base,
		temperature,
		feelsLike,
		minimum,
		maximum,
		pressure,
		humidity,
		seaLevel,
		groundLevel,
		visibility,
		speed,
		deg,
		gust,
		cloudAll,
		rain1h,
		snow1h,
		dt,
		sysType,
		sysId,
		country,
		sunset,
		sunrise,
		tz,
	)

	assert.NotNil(w)
	assert.Equal(name, w.Name)
	assert.Equal(float32(pressure)/pascal, w.MainInfo.Pressure)
	assert.Equal(visibility/kilometer, w.Visibility)
	assert.Equal(speed*kilometerPerHourMultiplier, w.Wind.Speed)
}
