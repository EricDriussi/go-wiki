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
		fmt.Println("Some sample pages are already provided")
		fmt.Println("Go to http://localhost:8080/wiki/view/Wombat to check it out")
		cli.DownloadArticlesInParallel()
		server.Run()
	}
	for _, arg := range os.Args[1:] {
		if arg == "-r" || arg == "--read" {
			cli.PrintWikiPage()
		} else if arg == "-w" || arg == "--write" {
			cli.CreateWikiPage()
		} else if arg == "-s" || arg == "--setup" {
			cli.DownloadArticlesInParallel()
		} else {
			fmt.Println("Invalid flag")
		}
	}
}
