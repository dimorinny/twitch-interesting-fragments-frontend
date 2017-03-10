package main

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/dimorinny/twitch-interesting-fragments-frontend/configuration"
	"github.com/dimorinny/twitch-interesting-fragments-frontend/data"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/view"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
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
	app := iris.New()
	app.Adapt(
		iris.DevLogger(),
		view.HTML("./template", ".html"),
		httprouter.New(),
	)


	app.Get("/", Index)

	app.Listen(fmt.Sprintf(
		"%s:%d",
		config.Host,
		config.Port,
	))
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
