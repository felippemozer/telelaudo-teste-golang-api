package contract

type userGeo struct {
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lng"`
}

type userAddress struct {
	Street   string  `json:"street"`
	Suite    string  `json:"suite"`
	City     string  `json:"city"`
	ZipCode  string  `json:"zipcode"`
	Location userGeo `json:"geo"`
}

type userCompany struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type GetUserByIdRequestDTO struct {
	Id       uint8       `json:"id"`
	Name     string      `json:"name"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Address  userAddress `json:"address"`
	Phone    string      `json:"phone"`
	Website  string      `json:"website"`
	Company  userCompany `json:"company"`
}
