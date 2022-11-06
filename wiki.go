package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"wiki/src/page"
	"wiki/src/server"
)

func main() {
	for _, a := range os.Args[1:] {
		if a == "-r" {
			fmt.Println("Sample Read/Write:")
			readWrite()
		} else if a == "-s" {
			fmt.Println("Launching server")
			serve()
		} else {
			fmt.Println("NOPE")
		}
	}
}

func readWrite() {
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
}

func serve() {
	http.HandleFunc("/wiki/", server.ViewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
