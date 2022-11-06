package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"wiki/src/page"
)

var articlesToDownload = []string{"Wombat", "Platypus", "TempleOS"}

func PrintWikiPage() {
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

func CreateWikiPage() {
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

func SetupWikiPages() {
	fmt.Println("Downloading a bunch of pages from wikipedia...")
	downloadArticles(articlesToDownload)
	fmt.Println("All done!")
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
