package main

import (
	"fmt"
	"os"
	"wiki/src/cli"
	"wiki/src/server"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Launching server on port 8080")
		fmt.Println("Go to http://localhost:8080/wiki/view/[WIKI_PAGE_TITLE] to check it out")
		server.Run()
	}
	for _, a := range os.Args[1:] {
		if a == "-r" {
			cli.PrintWikiPage()
		} else if a == "-w" {
			cli.CreateWikiPage()
		} else if a == "-s" {
			// TODO
			fmt.Println("Setup a bunch of pages (wikipedia)")
		} else {
			fmt.Println("Invalid flag")
		}
	}
}
