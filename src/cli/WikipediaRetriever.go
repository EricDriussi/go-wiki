package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"wiki/src/page"
)

type wikipediaArticle struct {
	_     bool `json:"batchcomplete"`
	Query struct {
		Pages []struct {
			Extract string `json:"extract"`
			_       int64  `json:"ns"`
			_       int64  `json:"pageid"`
			Title   string `json:"title"`
		} `json:"pages"`
	} `json:"query"`
}

var (
	wombatArticle   = "Wombat"
	platypusArticle = "Platypus"
	templeOSArticle = "TempleOS"
	articles        = []string{"Wombat", "Platypus", "TempleOS"}
)

func DownloadArticles() {
	fmt.Println("Setting up a bunch of pages from wikipedia...")

	for _, article := range articles {
		rawResponse := downloadArticle(article)
		articleTitle, articleExtract := parseWikipediaResponse(rawResponse)

		pageToWrite := page.Page{Title: articleTitle, Body: articleExtract}
		saveErr := pageToWrite.Save()
		if saveErr != nil {
			log.Fatal("[ERROR]: Couldn't save requested page")
		}
	}
}

func downloadArticle(article string) string {
	url := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=query&format=json&prop=extracts&titles=%s&formatversion=2&exintro=1&explaintext=1", article)
	resp, wikipediaError := http.Get(url)
	if wikipediaError != nil {
		fmt.Printf("Couldn't get from wikipedia -> %s", wikipediaError)
	}
	defer resp.Body.Close()
	body, parseError := ioutil.ReadAll(resp.Body)
	if parseError != nil {
		fmt.Printf("Couldn't parse from wikipedia -> %s", parseError)
		return ""
	}
	return string(body)
}

func parseWikipediaResponse(res string) (string, string) {
	var article wikipediaArticle
	jsonError := json.Unmarshal([]byte(res), &article)
	if jsonError != nil {
		fmt.Printf("Couldn't parse json -> %s", jsonError)
		return "", ""
	}
	return article.Query.Pages[0].Title, article.Query.Pages[0].Extract
}
