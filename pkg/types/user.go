package types

type User struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name, omitempty"`
	Password string `json:"pwsd, omitempty"`
}
