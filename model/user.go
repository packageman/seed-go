package model

type User struct {
	Id       int         `json:"id"`
	Name     string      `json:"name"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Address  Address     `json:"address"`
	Phone    string      `json:"phone"`
	Website  string      `json:"website"`
	Company  interface{} `json:"company"`
}

type Address struct {
	Street  string            `json:"street"`
	Suite   string            `json:"suite"`
	City    string            `json:"city"`
	Zipcode string            `json:"zipcode"`
	Geo     map[string]string `json:"geo"`
}
