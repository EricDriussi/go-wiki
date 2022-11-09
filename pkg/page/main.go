package page

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  string
}

var defaultWikiRoute = "assets/"

func (page *Page) Save() error {
	err := os.MkdirAll(defaultWikiRoute, os.ModePerm)
	if err != nil {
		return err
	}
	fullPath := buildPathToPage(page.Title)
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

func buildPathToPage(title string) string {
	cwd, _ := os.Getwd()
	return fmt.Sprintf("%s/%s/%s.txt", cwd, defaultWikiRoute, title)
}
