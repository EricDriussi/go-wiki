package server

import (
	"log"
	"net/http"
	config "wiki/src"
	factory "wiki/src/server/factory"
)

func Run() {
	http.HandleFunc(config.ViewRoute, factory.GetViewHandler())
	http.HandleFunc(config.EditRoute, factory.GetEditHandler())
	http.HandleFunc(config.SaveRoute, factory.GetSaveHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
