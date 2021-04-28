package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CarlosRoGuerra/New_Api_Go/v1/internal/database"
	"github.com/CarlosRoGuerra/New_Api_Go/v1/pkg/types"
	"github.com/gorilla/mux"
)

func (a *Api) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := a.Client.GetUsers("users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
	w.WriteHeader(http.StatusOK)
}

func (a *Api) createUser(w http.ResponseWriter, r *http.Request) {
	var user types.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	returnedUser, err := a.Client.CreateUser("users", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(returnedUser)
}

func (a *Api) updateUser(w http.ResponseWriter, r *http.Request) {
	users, err := a.Client.GetUsers("users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var toUpdate types.User
	err = json.NewDecoder(r.Body).Decode(&toUpdate)
	if err != nil {
		fmt.Println(w, err.Error(), http.StatusBadRequest)
	}
	for _, user := range users {
		if user.Id == toUpdate.Id {
			updateUser, err := a.Client.UpdateUser("users", toUpdate)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updateUser)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func (a *Api) deleteUser(w http.ResponseWriter, r *http.Request) {
	users, err := a.Client.GetUsers("users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var toDelete types.User
	err = json.NewDecoder(r.Body).Decode(&toDelete)
	if err != nil {
		fmt.Println(w, err.Error(), http.StatusBadRequest)
	}
	for _, user := range users {
		if user.Id == toDelete.Id {
			err := a.Client.DeleteUser("users", user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			break
		}
	}

	http.Error(w, "user not found", http.StatusNotFound)
}

type Api struct {
	Router *mux.Router
	Client database.DatabaseClient
}

func NewWithClient(client database.DatabaseClient) Api {
	var a Api
	a.Client = client
	a.buildRouter()
	return a
}

func (a *Api) buildRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/users", a.getUsers).Methods("GET")
	router.HandleFunc("/users", a.createUser).Methods("POST")
	router.HandleFunc("/users/{id}", a.updateUser).Methods("PUT")
	router.HandleFunc("/users", a.deleteUser).Methods("DELETE")
	router.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "WORKING\n")
	}).Methods("GET")

	a.Router = router
}

func (a *Api) Listen(port string) error {
	return http.ListenAndServe(port, a.Router)
}
