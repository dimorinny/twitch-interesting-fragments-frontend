package main

import (
	"github.com/caarlos0/env"
	"github.com/dimorinny/twitch-interesting-fragments-frontend/configuration"
	"github.com/dimorinny/twitch-interesting-fragments-frontend/data"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2"
	"log"
	"github.com/kataras/go-template/html"
	"net/http"
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
	iris.OptionIsDevelopment(true)
	iris.UseTemplate(html.New(html.Config{
		Layout: "index.html",
	})).Directory(
		"./template",
		".html",
	)

	iris.Get("/", Index)

	iris.Listen(":8080")
}

type IndexPage struct {
	Fragments []data.UploadedFragment
}

func Index(ctx *iris.Context) {
	result, err := storage.GetUploadedImages()
	if err != nil {
		ctx.EmitError(http.StatusInternalServerError)
	}

	ctx.Render(
		"fragments.html",
		IndexPage{
			Fragments: result,
		},
	)
}
