package page

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
	"wiki/pkg/config"
)

func (this *Page) Save() error {
	err := os.MkdirAll(config.WikiPagesPath, os.ModePerm)
	if err != nil {
		return err
	}
	fileContent := []byte(this.body)
	return os.WriteFile(fullPathTo(this.title), fileContent, 0600)
}

func Load(title string) (*Page, error) {
	body, err := os.ReadFile(fullPathTo(title))
	if err != nil {
		return &Page{title: title, body: "PAGE NOT FOUND"}, err
	}
	return &Page{title: title, body: string(body)}, nil
}

func LoadAll() []*Page {
	pages := []*Page{}
	files, err := readFilesFrom(config.WikiPagesPath)
	if err != nil {
		fmt.Printf("Error reading assets: %s", err)
		return pages
	}
	for _, title := range trimmedFileNamesOf(files) {
		page, loadErr := Load(title)
		if loadErr != nil {
			fmt.Printf("Page %s not found!", title)
		}
		pages = append(pages, page)
	}
	return pages
}

// TODO.ask

func trimmedFileNamesOf(files []fs.FileInfo) []string {
	filenames := []string{}
	for _, file := range files {
		nameWithoutExtension := strings.Split(file.Name(), ".")[0]
		filenames = append(filenames, nameWithoutExtension)
	}
	return filenames
}

func readFilesFrom(dir string) ([]fs.FileInfo, error) {
	fileReader, pathErr := os.Open(dir)
	if pathErr != nil {
		return []fs.FileInfo{}, pathErr
	}
	files, readErr := fileReader.Readdir(0)
	if readErr != nil {
		return []fs.FileInfo{}, pathErr
	}
	return files, nil
}

func fullPathTo(title string) string {
	cwd, _ := os.Getwd()
	return fmt.Sprintf("%s/%s/%s.txt", cwd, config.WikiPagesPath, title)
}
