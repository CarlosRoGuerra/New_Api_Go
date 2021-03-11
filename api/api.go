package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CarlosRoGuerra/New_Api_Go/v1/pkg/types"
	"github.com/gorilla/mux"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(&types.User{})

}
func createUser(w http.ResponseWriter, r *http.Request) {
	var user types.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("%v", user)

	w.WriteHeader(http.StatusCreated)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	upduser := []struct {
		name     string
		user     types.User
		expected error
	}{
		{
			user: types.User{
				Name:     "test-update",
				Password: "123",
			},
		},
	}
	params := mux.Vars(r)
	for index, item := range upduser {
		if item.user.Name == params["name"] {
			upduser = append(upduser[:index], upduser[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(upduser)
	}
	var updateUser types.User
	_ = json.NewDecoder(r.Body).Decode(&updateUser)
	updateUser.Name = params["name"]
	upduser = append(upduser)
	json.NewEncoder(w).Encode(upduser)
	w.WriteHeader(301)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	deluser := []struct {
		name     string
		user     types.User
		expected error
	}{
		{
			user: types.User{
				Name:     "test-delete",
				Password: "123",
			},
		},
	}
	params := mux.Vars(r)
	for index, item := range deluser {
		if item.user.Name == params["name"] {
			deluser = append(deluser[:index], deluser[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(deluser)
	}
	//fmt.Printf("%v", deleteUser)
	w.WriteHeader(http.StatusMovedPermanently)
}

func New() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users", getUser).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users", updateUser).Methods("PUT")
	router.HandleFunc("/users", deleteUser).Methods("DELETE")

	return router
}
