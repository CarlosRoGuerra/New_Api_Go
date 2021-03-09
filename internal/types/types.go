package internal

type User struct {
	Nome        string `json:"nome, omitempty"`
	Password    string `json:"pwsd, omitempty"`
	UserAddress Address
}

type Address struct {
	Address string
}
