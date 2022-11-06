package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	config "wiki/src"
	"wiki/src/page"
	factory "wiki/src/server"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Launching server on port 8080")
		fmt.Println("Go to http://localhost:8080/wiki/view/[WIKI_PAGE_TITLE] to check it out")
		runServer()
	}
	for _, a := range os.Args[1:] {
		if a == "-r" {
			runCliRead()
		} else if a == "-w" {
			runCliWrite()
		} else if a == "-s" {
			// TODO
			fmt.Println("Setup a bunch of pages (wikipedia)")
		} else {
			fmt.Println("Invalid flag")
		}
	}
}

func runCliRead() {
	fmt.Println("Write the title of the desired wiki page (no spaces allowed):")
	var title string
	fmt.Scanln(&title)

	pageToRead, LoadErr := page.Load(title)
	if LoadErr != nil {
		log.Fatal("[ERROR]: Couldn't load requested page")
	}
	fmt.Println("Page found:")
	fmt.Println(string(pageToRead.Body))
}

func runCliWrite() {
	fmt.Println("Write a title for the new wiki page (no spaces allowed):")
	var title string
	fmt.Scanln(&title)

	fmt.Println("Write the content of the new wiki page:")
	body := basicReader()

	pageToWrite := page.Page{Title: title, Body: body}
	saveErr := pageToWrite.Save()
	if saveErr != nil {
		log.Fatal("[ERROR]: Couldn't save requested page")
	}
}

func basicReader() string {
	reader := bufio.NewReader(os.Stdin)
	body, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input", err)
		return ""
	}
	return strings.TrimSuffix(body, "\n")
}

func runServer() {
	http.HandleFunc(config.ViewRoute, factory.GetViewHandler())
	http.HandleFunc(config.EditRoute, factory.GetEditHandler())
	http.HandleFunc(config.SaveRoute, factory.GetSaveHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
