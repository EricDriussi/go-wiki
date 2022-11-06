package page

import (
	"fmt"
	"os"
	"path/filepath"
)

type Page struct {
	Title string
	Body  string
	Route string
}

var defaultWikiRoute = "wiki_pages/"

func (page *Page) Save() error {
	path := buildPathToRoute(page.Route)
	fullPath := path + page.Title + ".txt"
	err := createDirs(path)
	if err != nil {
		return err
	}
	fileContent := []byte(page.Body)
	return os.WriteFile(fullPath, fileContent, 0600)
}

func Load(pageRouteWithTitle string) (*Page, error) {
	base := filepath.Base(pageRouteWithTitle)
	fullPath := buildPathToPage(pageRouteWithTitle)
	body, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}
	return &Page{Title: base, Body: string(body)}, nil
}

func createDirs(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func buildPathToRoute(route string) string {
	cwd, _ := os.Getwd()
	return fmt.Sprintf("%s/%s/%s/", cwd, defaultWikiRoute, route)
}

func buildPathToPage(route string) string {
	cwd, _ := os.Getwd()
	return fmt.Sprintf("%s/%s/%s.txt", cwd, defaultWikiRoute, route)
}
