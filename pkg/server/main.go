package server

import (
	"log"
	"net/http"
	"wiki/pkg/config"
	"wiki/pkg/server/factory"
)

func Run() {
	http.HandleFunc(config.IndexRoute, factory.GetIndexHandler())
	http.HandleFunc(config.ViewRoute, factory.GetViewHandler())
	http.HandleFunc(config.EditRoute, factory.GetEditHandler())
	http.HandleFunc(config.SaveRoute, factory.GetSaveHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
