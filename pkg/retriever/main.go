package retriever

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"wiki/pkg/page"
)

type wikipediaArticle struct {
	_     bool `json:"batchcomplete"`
	Query struct {
		Pages []struct {
			Extract string `json:"extract"`
			_       int64  `json:"ns"`
			_       int64  `json:"pageid"`
			Title   string `json:"title"`
			Missing bool   `json:"missing"`
		} `json:"pages"`
	} `json:"query"`
}

func DownloadArticles(articles []string) {
	var wg sync.WaitGroup

	for _, article := range articles {
		artCopy := article
		wg.Add(1)
		go func() {
			defer wg.Done()

			rawResponse := downloadSingleArticle(artCopy)
			articleTitle, articleExtract, err := parseWikipediaResponse(rawResponse)

			if err != nil {
				fmt.Printf("[ERROR]: %s\n", err.Error())
			} else {
				pageToWrite := page.New().WithTitle(articleTitle).WithBody(articleExtract)
				saveErr := pageToWrite.Save()
				if saveErr != nil {
					fmt.Printf("[ERROR]: %s\n", saveErr.Error())
				}
			}
		}()
	}
	wg.Wait()
}

func downloadSingleArticle(article string) string {
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

func parseWikipediaResponse(res string) (title string, body string, err error) {
	var article wikipediaArticle
	jsonError := json.Unmarshal([]byte(res), &article)
	if jsonError != nil {
		return "", "", errors.New("Something went wrong when parsing wikipedia's response!")
	}
	if article.Query.Pages[0].Missing {
		return "", "", errors.New("One of your requested articles wasn't found!")
	}
	return article.Query.Pages[0].Title, article.Query.Pages[0].Extract, nil
}
