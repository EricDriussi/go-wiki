package server

import (
	"fmt"
	"net/http"
	p "wiki/src/page"
)

func ViewHandler(res http.ResponseWriter, req *http.Request) {
	pageRouteWithTitle := req.URL.Path[len("/wiki/"):]
	page, _ := p.Load(pageRouteWithTitle)
	fmt.Fprintf(res, "<h1>%s</h1><div>%s</div>", page.Title, page.Body)
}
