package page

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  string
}

var defaultWikiRoute = "wiki_pages/"

func (page *Page) Save() error {
	path := buildPathToRoute()
	fullPath := path + page.Title + ".txt"
	err := createDirs(path)
	if err != nil {
		return err
	}
	fileContent := []byte(page.Body)
	return os.WriteFile(fullPath, fileContent, 0600)
}

func Load(title string) (*Page, error) {
	fullPath := buildPathToPage(title)
	body, err := os.ReadFile(fullPath)
	if err != nil {
		return &Page{Title: title}, err
	}
	return &Page{Title: title, Body: string(body)}, nil
}

func createDirs(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func buildPathToRoute() string {
	cwd, _ := os.Getwd()
	return fmt.Sprintf("%s/%s/", cwd, defaultWikiRoute)
}

func buildPathToPage(title string) string {
	cwd, _ := os.Getwd()
	return fmt.Sprintf("%s/%s/%s.txt", cwd, defaultWikiRoute, title)
}
