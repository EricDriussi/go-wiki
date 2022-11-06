package server

import (
	"fmt"
	"net/http"
	p "wiki/src/page"
)

var (
	ViewPath = "/wiki/view/"
	EditPath = "/wiki/edit/"
	SavePath = "/wiki/save/"
)

func ViewHandler(res http.ResponseWriter, req *http.Request) {
	pageRouteWithTitle := req.URL.Path[len(ViewPath):]
	page, _ := p.Load(pageRouteWithTitle)
	fmt.Fprintf(res, "<h1>%s</h1><div>%s</div>", page.Title, page.Body)
}

func EditHandler(res http.ResponseWriter, req *http.Request) {
	pageRouteWithTitle := req.URL.Path[len(EditPath):]
	page, err := p.Load(pageRouteWithTitle)
	if err != nil {
		page.Title = "Nothing much really..."
	}
	fmt.Fprintf(res, "<h1>Editing %s</h1>"+
		"<form action=\"/wiki/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>", page.Title, page.Title, page.Body)
}

func SaveHandler(res http.ResponseWriter, req *http.Request) {
	pageRouteWithTitle := req.URL.Path[len(SavePath):]
	body := req.FormValue("body")
	pageToWrite := p.Page{Title: pageRouteWithTitle, Body: body}
	pageToWrite.Save()
	http.Redirect(res, req, "/wiki/view/"+pageRouteWithTitle, http.StatusFound)
}
