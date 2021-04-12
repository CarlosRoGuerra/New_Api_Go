package main

import (
	"log"

	"github.com/CarlosRoGuerra/New_Api_Go/v1/api"
	"github.com/CarlosRoGuerra/New_Api_Go/v1/internal/database"
)

func main() {
	client, err := database.NewDefaultMongoClient()
	if err != nil {
		panic(err)
	}
	api := api.NewWithClient(client)
	log.Fatal(api.Listen(":8888"))
}
