package page

import (
	"os"
)

type Page struct {
	Title string
	Body  string
}

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
