package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kkn1993/go_webapp/pkg/config"
	"github.com/kkn1993/go_webapp/pkg/handlers"
	"github.com/kkn1993/go_webapp/pkg/render"
)

const portNumber = ":8088"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers((repo))

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Startin application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
