package server

import (
	"fmt"
	"net/http"
	p "wiki/src/page"
)

func ViewHandler(res http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len("/wiki_pages/"):]
	page, _ := p.Load(title, "wiki_pages")
	fmt.Fprintf(res, "<h1>%s</h1><div>%s</div>", page.Title, page.Body)
}
