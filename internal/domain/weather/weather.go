package weather

import (
	"time"
)

type Coord struct {
	Longitude float32
	Latitude  float32
}

type WeatherInfo struct {
	Id          uint32
	Main        string
	Description string
	Icon        string
}

type Main struct {
	Temperature float32
	FeelsLike   float32
	Minimum     float32
	Maximum     float32
	Pressure    float32
	Humidity    uint8
	SeaLevel    float32
	GroundLevel float32
}

type Wind struct {
	Speed float32
	Deg   int8
	Gust  float32
}

type Cloud struct {
	All int8
}

type Rain struct {
	OneHour int32
}

type Snow struct {
	OneHour int32
}

type Sys struct {
	Type    uint8
	Id      int64
	Country string
	Sunset  time.Time
	Sunrise time.Time
}

type Weather struct {
	Coordinates Coord
	Weather     []WeatherInfo
	Base        string
	MainInfo    Main
	Visibility  uint32
	Wind        Wind
	Clouds      Cloud
	Rain        Rain
	Snow        Snow
	Datetime    time.Time
	Sys         Sys
	Timezone    int8
	Id          uint32
	Name        string
}

func NewWeather(
	id uint32,
	name string,
	latitude float32,
	longitude float32,
	weather []WeatherInfo,
	base string,
	temperature float32,
	feelsLike float32,
	minimum float32,
	maximun float32,
	pressure uint32,
	humidity uint8,
	seaLevel uint32,
	groundLevel uint32,
	visibility uint32,
	speed float32,
	deg int8,
	gust float32,
	cloudAll int8,
	rain1h int32,
	snow1h int32,
	dt int64,
	sysType uint8,
	sysId int64,
	country string,
	sunset int64,
	sunrise int64,
	tz int64,
) *Weather {
	var w Weather

	w.Id = id
	w.Name = name
	w.Coordinates.Latitude = latitude
	w.Coordinates.Longitude = longitude
	w.Weather = weather
	w.Base = base
	w.MainInfo.Temperature = temperature
	w.MainInfo.FeelsLike = feelsLike
	w.MainInfo.Minimum = minimum
	w.MainInfo.Maximum = maximun
	w.MainInfo.Humidity = humidity
	w.Wind.Deg = deg
	w.Clouds.All = cloudAll
	w.Rain.OneHour = rain1h
	w.Snow.OneHour = snow1h
	w.Sys.Type = sysType
	w.Sys.Id = sysId
	w.Sys.Country = country

	// Treat pressure
	var pascal float32 = 1013.2
	w.MainInfo.Pressure = float32(pressure) / pascal
	w.MainInfo.SeaLevel = float32(seaLevel) / pascal
	w.MainInfo.GroundLevel = float32(groundLevel) / pascal

	// Treat visibility
	var kilometer uint32 = 1000
	w.Visibility = visibility / kilometer

	// Treat wind
	var kilometerPerHourMultiplier float32 = 3.6
	w.Wind.Speed = speed * kilometerPerHourMultiplier
	w.Wind.Gust = gust * kilometerPerHourMultiplier

	// Treat datetime
	w.Datetime = time.Unix(dt, 0).UTC()
	w.Sys.Sunset = time.Unix(sunset, 0).UTC()
	w.Sys.Sunrise = time.Unix(sunrise, 0).UTC()

	// Treat Timezone
	var totalSecondsByHour int64 = 3600
	w.Timezone = int8(tz / totalSecondsByHour)

	return &w
}
