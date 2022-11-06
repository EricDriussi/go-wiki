package server

import (
	"fmt"
	"html/template"
	"net/http"

	p "wiki/src/page"
)

var (
	ViewPath = "/wiki/view/"
	EditPath = "/wiki/edit/"
	SavePath = "/wiki/save/"
)

func ViewHandler(res http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len(ViewPath):]
	page := tryLoadPage(title)
	fmt.Fprintf(res, "<h1>%s</h1><div>%s</div>", page.Title, page.Body)
}

func EditHandler(res http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len(EditPath):]
	page := tryLoadPage(title)
	editTemplate, _ := template.ParseFiles("src/server/html_templates/edit_form.html")
	editTemplate.Execute(res, page)
}

func SaveHandler(res http.ResponseWriter, req *http.Request) {
	pageRouteWithTitle := req.URL.Path[len(SavePath):]
	body := req.FormValue("body")
	pageToWrite := p.Page{Title: pageRouteWithTitle, Body: body}
	pageToWrite.Save()
	http.Redirect(res, req, "/wiki/view/"+pageRouteWithTitle, http.StatusFound)
}

func tryLoadPage(title string) p.Page {
	page, err := p.Load(title)
	if err != nil {
		page.Title = "Nothing much to load..."
	}
	return *page
}
