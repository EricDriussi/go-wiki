package server

import (
	"html/template"
	"net/http"

	p "wiki/src/page"
)

type templateDTO struct {
	Page p.Page
	Path string
}

var (
	ViewPath = "/wiki/view/"
	EditPath = "/wiki/edit/"
	SavePath = "/wiki/save/"
)

func ViewHandler(res http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len(ViewPath):]
	page := tryLoadPage(title)
	dto := templateDTO{Page: page, Path: EditPath}
	renderTemplate(res, "view", dto)
}

func EditHandler(res http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len(EditPath):]
	page := tryLoadPage(title)
	dto := templateDTO{Page: page, Path: SavePath}
	renderTemplate(res, "edit_form", dto)
}

func SaveHandler(res http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len(SavePath):]
	body := req.FormValue("body")
	pageToWrite := p.Page{Title: title, Body: body}
	pageToWrite.Save()
	http.Redirect(res, req, ViewPath+title, http.StatusFound)
}

func renderTemplate(res http.ResponseWriter, templateName string, dto templateDTO) {
	editTemplate, _ := template.ParseFiles("src/server/html_templates/" + templateName + ".html")
	editTemplate.Execute(res, dto)
}

func tryLoadPage(title string) p.Page {
	page, err := p.Load(title)
	if err != nil {
		page.Title = "Nothing much to load..."
	}
	return *page
}
