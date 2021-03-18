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
		fmt.Println(w, err.Error(), http.StatusBadRequest)
	}
	for _, item := range users {
		if item.Id == user.Id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)

}

// func createUser(w http.ResponseWriter, r *http.Request) {
// 	users := []types.User{
// 		{Id: "123", Name: "seila", Password: "456"},
// 		{Id: "456", Name: "test", Password: "678"},
// 	}
// 	var user types.User
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		fmt.Println(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	fmt.Printf("%v", user)

// 	w.WriteHeader(http.StatusCreated)
// }

// func updateUser(w http.ResponseWriter, r *http.Request) {
// 	users := []types.User{
// 		{Id: "123", Name: "seila", Password: "456"},
// 		{Id: "456", Name: "test", Password: "678"},
// 	}
// 	var tuser types.User
// 	err := json.NewDecoder(r.Body).Decode(&tuser)
// 	for index, item := range users {
// 		if item.Id == tuser.Id {
// 			users = append(users[:index], users[index+1:]...)
// 			_ = json.NewDecoder(r.Body).Decode(&tuser)
// 			users = append(users, tuser)
// 		}
// 		if err != nil {
// 			fmt.Println(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}
// 	}
// 	if tuser.Name == item.Name {
// 		fmt.Printf("Usuario Atualizado!")
// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(tuser)
// 		return
// 	} else {
// 		fmt.Printf("Error")
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// }

// func deleteUser(w http.ResponseWriter, r *http.Request) {
// 	var duser types.User
// 	users = append(users, types.User{Id: "1", Name: "test-tudo", Password: "12456"})
// 	//params := mux.Vars(r)
// 	err := json.NewDecoder(r.Body).Decode(&duser)
// 	for index, item := range users {
// 		if item.Id == duser.Id {
// 			users = append(users[:index], users[index+1:]...)
// 			break
// 		}
// 	}
// 	json.NewEncoder(w).Encode(duser)
// 	if err != nil {
// 		fmt.Println(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	fmt.Printf("%v :", users)
// 	if users == nil {
// 		fmt.Printf("User Deleted!")
// 	}
// 	w.WriteHeader(301)
// }

func New() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUser).Methods("GET")
	//router.HandleFunc("/users", createUser).Methods("POST")
	//router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	//router.HandleFunc("/users", deleteUser).Methods("DELETE")

	return router
}
