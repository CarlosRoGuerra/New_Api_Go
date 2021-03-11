package types

type User struct {
	Name        string `json:"name, omitempty"`
	Password    string `json:"pwsd, omitempty"`
	UserAddress Address
}

type Address struct {
	Address string
}
