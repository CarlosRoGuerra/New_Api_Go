package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CarlosRoGuerra/New_Api_Go/v1/api"
	"github.com/CarlosRoGuerra/New_Api_Go/v1/internal/database"
)

func main() {
	db := database.GetCollection("users")
	fmt.Println(db)
	api := api.New().Router
	log.Fatal(http.ListenAndServe(":8001", api))
}
