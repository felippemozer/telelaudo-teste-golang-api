package weather

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var (
	w = NewWeather(
		2643743,
		"London",
		51.5085,
		-0.1257,
		[]WeatherInfo{
			{
				Id:          804,
				Main:        "Clouds",
				Description: "overcast clouds",
				Icon:        "04n",
			},
		},
		"stations",
		7.65,
		5.05,
		6.98,
		8.4,
		1007,
		83,
		1007,
		1002,
		6000,
		4.12,
		10,
		0,
		100,
		0,
		0,
		1741645279,
		2,
		2091269,
		"GB",
		1741587952,
		1741629353,
		0,
	)
)

func TestGetBy(t *testing.T) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		t.Fatalf("Error loading .env file: %s", err.Error())
	}
	assert := assert.New(t)
	city := "London"

	res, err := GetBy(city)

	assert.NoError(err)
	assert.Equal(w.Name, res.Name)
	assert.Equal(w.Sys.Country, res.Sys.Country)
}
