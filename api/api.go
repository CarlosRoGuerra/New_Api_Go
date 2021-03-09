package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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

func newRout() {
	routes := mux.NewRouter()

	routes.HandleFunc("/users", GetUser).Methods("GET")
	routes.HandleFunc("/users/{name}", CreateUser).Methods("POST")

}
