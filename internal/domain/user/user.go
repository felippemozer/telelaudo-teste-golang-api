package user

type Geo struct {
	Latitude  float32
	Longitude float32
}

type Address struct {
	Street   string
	Suite    string
	City     string
	ZipCode  string
	Location Geo
}

type Company struct {
	Name        string
	CatchPhrase string
	Bs          string
}

type User struct {
	Id       uint8
	Name     string
	Username string
	Email    string
	Address  Address
	Phone    string
	Website  string
	Company  Company
}

func NewUser(
	id uint8,
	name string,
	username string,
	email string,
	street string,
	suite string,
	city string,
	zipCode string,
	latitude float32,
	longitude float32,
	phone string,
	website string,
	companyName string,
	catchPhrase string,
	companyBs string,
) *User {
	var u User

	u.Id = id
	u.Name = name
	u.Username = username
	u.Email = email
	u.Address.Street = street
	u.Address.Suite = suite
	u.Address.City = city
	u.Address.ZipCode = zipCode
	u.Address.Location.Latitude = latitude
	u.Address.Location.Longitude = longitude
	u.Phone = phone
	u.Website = website
	u.Company.Name = companyName
	u.Company.CatchPhrase = catchPhrase
	u.Company.Bs = companyBs

	return &u
}
