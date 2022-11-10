package page

import (
	"fmt"
	"os"
	"wiki/pkg/config"
)

func (this *Page) Save() error {
	err := os.MkdirAll(config.WikiPagesPath, os.ModePerm)
	if err != nil {
		return err
	}
	fullPath := buildPathToPage(this.title)
	fileContent := []byte(this.body)
	return os.WriteFile(fullPath, fileContent, 0600)
}

func Load(title string) (*Page, error) {
	fullPath := buildPathToPage(title)
	body, err := os.ReadFile(fullPath)
	if err != nil {
		return &Page{title: title}, err
	}
	return &Page{title: title, body: string(body)}, nil
}

func buildPathToPage(title string) string {
	cwd, _ := os.Getwd()
	return fmt.Sprintf("%s/%s/%s.txt", cwd, config.WikiPagesPath, title)
}
