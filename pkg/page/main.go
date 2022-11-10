package page

import (
	"fmt"
	"os"
	"wiki/pkg/config"
)

type Page struct {
	Title string
	Body  string
}

func New() Page {
	return Page{}
}

func (this Page) WithTitle(title string) Page {
	this.Title = title
	return this
}

func (this Page) WithBody(body string) Page {
	this.Body = body
	return this
}

func (this *Page) Save() error {
	err := os.MkdirAll(config.WikiPagesPath, os.ModePerm)
	if err != nil {
		return err
	}
	fullPath := buildPathToPage(this.Title)
	fileContent := []byte(this.Body)
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
	return fmt.Sprintf("%s/%s/%s.txt", cwd, config.WikiPagesPath, title)
}
