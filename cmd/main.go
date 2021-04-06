package main

import (
	"log"
	"net/http"

	"github.com/CarlosRoGuerra/New_Api_Go/v1/api"
)

func main() {
	api := api.New().Router
	log.Fatal(http.ListenAndServe(":8000", api))
}
