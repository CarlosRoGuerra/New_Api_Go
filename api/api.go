package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CarlosRoGuerra/New_Api_Go/v1/pkg/types"
	"github.com/gorilla/mux"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	users := []types.User{
		{Id: "123", Name: "seila", Password: "456"},
		{Id: "456", Name: "test", Password: "678"},
	}
	var user types.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, item := range users {
		if item.Id == user.Id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)

}

func createUser(w http.ResponseWriter, r *http.Request) {
	// users := []types.User{
	// 	{Id: "123", Name: "seila", Password: "456"},
	// 	{Id: "456", Name: "test", Password: "678"},
	// }
	var user types.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusCreated)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	users := []types.User{
		{Id: "123", Name: "seila", Password: "456"},
		{Id: "984", Name: "seila", Password: "456"},
	}
	var user types.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(w, err.Error(), http.StatusBadRequest)
	}
	for index, item := range users {
		if item.Id == user.Id {
			users = append(users[:index], users[index+1:]...)
			json.NewEncoder(w).Encode(item)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	users := []types.User{
		{Id: "123", Name: "seila", Password: "456"},
	}
	var user types.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(w, err.Error(), http.StatusBadRequest)
	}
	for index, item := range users {
		if item.Id == user.Id {
			users = append(users[:index], users[index+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	http.Error(w, "user not found", http.StatusNotFound)
}

func New() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUser).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/users", deleteUser).Methods("DELETE")

	return router
}
