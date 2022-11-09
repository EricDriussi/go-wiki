package server

import (
	"log"
	"net/http"
	"wiki/pkg"
	"wiki/pkg/server/factory"
)

func Run() {
	http.HandleFunc(pkg.ViewRoute, factory.GetViewHandler())
	http.HandleFunc(pkg.EditRoute, factory.GetEditHandler())
	http.HandleFunc(pkg.SaveRoute, factory.GetSaveHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
