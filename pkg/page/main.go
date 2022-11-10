package page

// TODO.refactor

import (
	"fmt"
	"os"
	"strings"
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

func LoadAll() []*Page {
	res := []*Page{}
	for _, filename := range allWikiFileNames() {
		page, _ := Load(filename)
		res = append(res, page)
	}
	return res
}

func allWikiFileNames() []string {
	fileReader, err := os.Open(config.WikiPagesPath)
	if err != nil {
		fmt.Println(err)
		return []string{}
	}
	files, err := fileReader.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return []string{}
	}
	filenames := []string{}
	for _, file := range files {
		nameWithoutExtension := strings.Split(file.Name(), ".")[0]
		filenames = append(filenames, nameWithoutExtension)
	}
	return filenames
}

func buildPathToPage(title string) string {
	cwd, _ := os.Getwd()
	return fmt.Sprintf("%s/%s/%s.txt", cwd, config.WikiPagesPath, title)
}
