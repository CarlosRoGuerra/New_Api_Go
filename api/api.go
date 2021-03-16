package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CarlosRoGuerra/New_Api_Go/v1/pkg/types"
	"github.com/gorilla/mux"
)

var users []types.User

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
	var tuser types.User
	users = append(users, types.User{Id: "1", Name: "test-tudo", Password: "56983"})
	err := json.NewDecoder(r.Body).Decode(&tuser)
	for index, item := range users {
		if item.Id == tuser.Id {
			users = append(users[:index], users[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&tuser)
			users = append(users, tuser)
			json.NewEncoder(w).Encode(tuser)
		}
		if err != nil {
			fmt.Println(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	if tuser.Name == users[0].Name {
		fmt.Printf("Usuario Atualizado!")
		w.WriteHeader(http.StatusOK)
	} else {
		fmt.Printf("Error")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	var duser types.User
	users = append(users, types.User{Id: "1", Name: "test-tudo", Password: "12456"})
	//params := mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&duser)
	for index, item := range users {
		if item.Id == duser.Id {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(duser)
	if err != nil {
		fmt.Println(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("%v :", users)
	if users == nil {
		fmt.Printf("User Deleted!")
	}
	w.WriteHeader(301)
}

func New() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUser).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/users", deleteUser).Methods("DELETE")

	return router
}
