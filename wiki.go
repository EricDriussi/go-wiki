package main

import (
	"fmt"
	"log"
	"wiki/src/page"
)

func main() {
	pageToWrite := page.Page{Title: "test_page", Route: "wiki_pages", Body: "This is a sample Page."}
	saveErr := pageToWrite.Save()
	if saveErr != nil {
		log.Fatal("[ERROR]: Couldn't save requested page")
	}
	pageToRead, LoadErr := page.Load("test_page", "wiki_pages")
	if LoadErr != nil {
		log.Fatal("[ERROR]: Couldn't load requested page")
	}
	fmt.Println(string(pageToRead.Body))
}
