package user

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/felippemozer/telelaudo-teste-golang-api/internal/contract"
)

type Service interface {
	GetBy(id int8) (*User, error)
}

func GetBy(id int8) (*User, error) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", id)

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error on HTTP request: %s", err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode >= http.StatusBadRequest {
		errorBody, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("error on get user info: %s", string(errorBody))
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error on parse user data: %s", err.Error())
	}

	var jsonData contract.GetUserByIdRequestDTO
	json.Unmarshal(data, &jsonData)

	user := NewUser(
		jsonData.Id,
		jsonData.Name,
		jsonData.Username,
		jsonData.Email,
		jsonData.Address.Street,
		jsonData.Address.Suite,
		jsonData.Address.City,
		jsonData.Address.ZipCode,
		jsonData.Address.Location.Latitude,
		jsonData.Address.Location.Longitude,
		jsonData.Phone,
		jsonData.Website,
		jsonData.Company.Name,
		jsonData.Company.CatchPhrase,
		jsonData.Company.Bs,
	)

	return user, nil
}
