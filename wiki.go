package main

import (
	"fmt"
	"log"
	"net/http"
	"wiki/src/page"
	"wiki/src/server"
)

func main() {
	pageToWrite := page.Page{Title: "test_page", Body: "This is a sample Page."}
	saveErr := pageToWrite.Save()
	if saveErr != nil {
		log.Fatal("[ERROR]: Couldn't save requested page")
	}
	pageToRead, LoadErr := page.Load("test_page")
	if LoadErr != nil {
		log.Fatal("[ERROR]: Couldn't load requested page")
	}
	fmt.Println(string(pageToRead.Body))

	http.HandleFunc("/wiki/", server.ViewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
