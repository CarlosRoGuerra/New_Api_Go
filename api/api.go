package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

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
	sync.Mutex
	Server          *http.Server
	Client          database.DatabaseClient
	shutdown        chan os.Signal
	shutdownTimeout time.Duration
}

func (a *Api) buildServer() {
	router := a.buildRouter()
	a.Server = &http.Server{
		Handler: router,
	}
	a.shutdownTimeout = time.Second * 10
}

func NewWithClient(client database.DatabaseClient) *Api {
	var a Api
	a.Client = client
	a.buildServer()
	return &a
}

func (a *Api) buildRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", a.getUsers).Methods("GET")
	router.HandleFunc("/users", a.createUser).Methods("POST")
	router.HandleFunc("/users/{id}", a.updateUser).Methods("PUT")
	router.HandleFunc("/users", a.deleteUser).Methods("DELETE")
	router.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "WORKING\n")
	}).Methods("GET")

	return router
}

func (a *Api) handleSignals() {
	a.shutdown = make(chan os.Signal)
	signal.Notify(a.shutdown, os.Interrupt, syscall.SIGTERM)
	<-a.shutdown
	close(a.shutdown)
	ctx, cancel := context.WithTimeout(context.Background(), a.shutdownTimeout)
	defer cancel()
	a.Server.Shutdown(ctx)
}

func (a *Api) Listen(port string) error {
	a.Server.Addr = port
	fmt.Printf("Server starting at address: %s\n", a.Server.Addr)
	go a.handleSignals()
	return a.Server.ListenAndServe()
}
