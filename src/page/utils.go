package page

import (
	"fmt"
	"os"
)

var defaultWikiRoute = "wiki_pages/"

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
