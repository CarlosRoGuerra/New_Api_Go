package main

import (
	"fmt"
	"log"
	"time"

	"github.com/CarlosRoGuerra/New_Api_Go/v1/api"
	"github.com/CarlosRoGuerra/New_Api_Go/v1/internal/database"
	"github.com/spf13/viper"
)

func main() {
	var client database.DatabaseClient
	var err error
	go func() {
		for {
			client, err = database.NewDefaultMongoClient()
			if err == nil {
				break
			}
			fmt.Printf("error creating database client\n")
			time.Sleep(time.Second * 5)
		}
	}()
	api := api.NewWithClient(client)
	log.Fatal(api.Listen(":" + viper.GetString("port")))
}
