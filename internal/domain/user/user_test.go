package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	assert := assert.New(t)
	var (
		id          uint8   = 1
		name        string  = "Carl Johnson"
		username    string  = "CJ"
		email       string  = "cj@grovefamily.com"
		street      string  = "Grove Street"
		suite       string  = "Ganton"
		city        string  = "Los Santos"
		zipCode     string  = "90008"
		latitude    float32 = 34.014193
		longitude   float32 = -118.344388
		phone       string  = "1-800-GROVE-CJ"
		website     string  = "grovefamily.com"
		companyName string  = "Grove Street Families"
		catchPhrase string  = "In Ballas we shot"
		companyBs   string  = "Private protection and sales"
	)

	u := NewUser(
		id,
		name,
		username,
		email,
		street,
		suite,
		city,
		zipCode,
		latitude,
		longitude,
		phone,
		website,
		companyName,
		catchPhrase,
		companyBs,
	)

	assert.NotNil(u)
	assert.Equal(id, u.Id)
	assert.Equal(name, u.Name)
	assert.Equal(username, u.Username)
	assert.Equal(email, u.Email)
	assert.Equal(street, u.Address.Street)
	assert.Equal(suite, u.Address.Suite)
	assert.Equal(city, u.Address.City)
	assert.Equal(zipCode, u.Address.ZipCode)
	assert.Equal(latitude, u.Address.Location.Latitude)
	assert.Equal(longitude, u.Address.Location.Longitude)
	assert.Equal(phone, u.Phone)
	assert.Equal(website, u.Website)
	assert.Equal(companyName, u.Company.Name)
	assert.Equal(catchPhrase, u.Company.CatchPhrase)
	assert.Equal(companyBs, u.Company.Bs)
}
