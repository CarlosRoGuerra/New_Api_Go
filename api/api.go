package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CarlosRoGuerra/New_Api_Go/v1/internal/database"
	"github.com/CarlosRoGuerra/New_Api_Go/v1/pkg/types"
	"github.com/gorilla/mux"
)

func (a *Api) getUser(w http.ResponseWriter, r *http.Request) {
	users, err := a.Client.GetUsers("users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var user types.User
	err = json.NewDecoder(r.Body).Decode(&user)
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

func (a *Api) createUser(w http.ResponseWriter, r *http.Request) {
	user, err := a.Client.CreateUser("users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//var user types.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusCreated)
}

func (a *Api) updateUser(w http.ResponseWriter, r *http.Request) {
	users, err := a.Client.GetUsers("users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var user types.User
	err = json.NewDecoder(r.Body).Decode(&user)
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

func (a *Api) deleteUser(w http.ResponseWriter, r *http.Request) {
	users, err := a.Client.GetUsers("users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var user types.User
	err = json.NewDecoder(r.Body).Decode(&user)
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

type Api struct {
	Router *mux.Router
	Client database.DatabaseClient
}

func New() Api {
	var a Api
	router := mux.NewRouter()
	a = Api{Router: router}
	router.HandleFunc("/users", a.getUser).Methods("GET")
	router.HandleFunc("/users", a.createUser).Methods("POST")
	router.HandleFunc("/users/{id}", a.updateUser).Methods("PUT")
	router.HandleFunc("/users", a.deleteUser).Methods("DELETE")
	router.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "WORKING")
	}).Methods("GET")
	return a
}
