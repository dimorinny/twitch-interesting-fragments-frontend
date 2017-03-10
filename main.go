package main

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/dimorinny/twitch-interesting-fragments-frontend/configuration"
	"github.com/dimorinny/twitch-interesting-fragments-frontend/data"
	"gopkg.in/mgo.v2"
	"log"
)

var (
	config  configuration.Configuration
	storage data.Storage
)

func initConfiguration() {
	config = configuration.Configuration{}

	err := env.Parse(&config)
	if err != nil {
		log.Fatal(err)
	}
}

func initMongoStorage() {
	session, err := mgo.Dial(config.StorageHost)
	if err != nil {
		log.Fatal(err)
	}

	storage = data.NewMongoStorage(
		session,
	)
}

func init() {
	initConfiguration()
	initMongoStorage()
}

func main() {
	result, err := storage.GetUploadedImages()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(result))
	fmt.Println(result)
}
