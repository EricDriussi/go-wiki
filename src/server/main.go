package server

import (
	"log"
	"net/http"
	"wiki/src"
	"wiki/src/server/factory"
)

func Run() {
	http.HandleFunc(src.ViewRoute, factory.GetViewHandler())
	http.HandleFunc(src.EditRoute, factory.GetEditHandler())
	http.HandleFunc(src.SaveRoute, factory.GetSaveHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
