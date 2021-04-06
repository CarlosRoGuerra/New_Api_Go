package types

import "fmt"

type User struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name, omitempty"`
	Password string `json:"pwsd, omitempty"`
}

func (u User) Show() string {
	return fmt.Sprintf("Test")
}
