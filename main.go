package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Nome        string `json:"nome, omitempty"`
	Password    string `json:"pwsd, omitempty"`
	UserAddress Address
}

type Address struct {
	Address string
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(&User{})

}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(w, err.Error(), http.StatusBadRequest)
	}

}
func main() {
	routes := mux.NewRouter()

	routes.HandleFunc("/users", GetUser).Methods("GET")
	routes.HandleFunc("/users/{id}", CreateUser).Methods("POST")
}
