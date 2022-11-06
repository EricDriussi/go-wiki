package page

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  string
	Route string
}

func (page *Page) Save() error {
	filename := page.Title + ".txt"
	path := buildPath(page.Route)
	fullPath := path + filename
	err := createDirs(path)
	if err != nil {
		return err
	}
	fileContent := []byte(page.Body)
	return os.WriteFile(fullPath, fileContent, 0600)
}

func Load(title string, route string) (*Page, error) {
	filename := title + ".txt"
	path := buildPath(route)
	fullPath := path + filename
	body, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: string(body)}, nil
}

func createDirs(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func buildPath(route string) string {
	cwd, _ := os.Getwd()
	return fmt.Sprintf("%v/%v/", cwd, route)
}
