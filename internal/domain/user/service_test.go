package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// {
//     "Id": 1,
//     "Name": "Leanne Graham",
//     "Username": "Bret",
//     "Email": "Sincere@april.biz",
//     "Address": {
//         "Street": "Kulas Light",
//         "Suite": "Apt. 556",
//         "City": "Gwenborough",
//         "ZipCode": "92998-3874",
//         "Location": {
//             "Latitude": 0,
//             "Longitude": 0
//         }
//     },
//     "Phone": "1-770-736-8031 x56442",
//     "Website": "hildegard.org",
//     "Company": {
//         "Name": "Romaguera-Crona",
//         "CatchPhrase": "Multi-layered client-server neural-net",
//         "Bs": "harness real-time e-markets"
//     }
// }

var (
	u = NewUser(
		1,
		"Leanne Graham",
		"Bret",
		"Sincere@april.biz",
		"Kulas Light",
		"Apt. 556",
		"Gwenborough",
		"92998-3874",
		0,
		0,
		"1-770-736-8031 x56442",
		"hildegard.org",
		"Romaguera-Crona",
		"Multi-layered client-server neural-net",
		"harness real-time e-markets",
	)
)

func TestGetBy(t *testing.T) {
	assert := assert.New(t)
	var idTest int8 = 1

	userTest, err := GetBy(idTest)

	assert.NoError(err)
	assert.Equal(u.Name, userTest.Name)
	assert.Equal(u.Username, userTest.Username)
	assert.Equal(u.Email, userTest.Email)
	assert.Equal(u.Address.Street, userTest.Address.Street)
	assert.Equal(u.Website, userTest.Website)
	assert.Equal(u.Company.Name, userTest.Company.Name)
}
