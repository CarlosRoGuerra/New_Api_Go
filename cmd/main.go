package main

import (
	"log"

	"github.com/CarlosRoGuerra/New_Api_Go/v1/api"
	"github.com/CarlosRoGuerra/New_Api_Go/v1/internal/database"
	"github.com/spf13/viper"
)

func main() {
	client, err := database.NewDefaultMongoClient()
	if err != nil {
		panic(err)
	}
	api := api.NewWithClient(client)
	log.Fatal(api.Listen(":" + viper.GetString("port")))
}
