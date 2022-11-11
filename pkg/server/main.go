package server

import (
	"log"
	"net/http"
	"wiki/pkg/config"
	"wiki/pkg/server/factory"
)

func Run() {
	http.HandleFunc(config.IndexRoute, factory.IndexHandler())
	http.HandleFunc(config.ViewRoute, factory.ViewHandler())
	http.HandleFunc(config.EditRoute, factory.EditHandler())
	http.HandleFunc(config.SaveRoute, factory.SaveHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
