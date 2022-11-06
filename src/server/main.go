package server

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"

	p "wiki/src/page"
)

type templateDTO struct {
	Page *p.Page
	Path string
}

var (
	ViewRoute     = "/wiki/view/"
	EditRoute     = "/wiki/edit/"
	SaveRoute     = "/wiki/save/"
	templatesPath = "src/server/html_templates/"
)

var templates = template.Must(
	template.ParseFiles(
		templatesPath+"edit_form.html",
		templatesPath+"view.html",
	),
)

var validPath = regexp.MustCompile(fmt.Sprintf("^(%s|%s|%s)([a-zA-Z0-9_-]+)$", ViewRoute, EditRoute, SaveRoute))

func HandlerMaker(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		match := validPath.FindStringSubmatch(req.URL.Path)
		if match == nil {
			http.NotFound(res, req)
			return
		}
		fn(res, req, match[2])
	}
}

func ViewHandler(res http.ResponseWriter, req *http.Request, title string) {
	page, noPageErr := p.Load(title)
	if noPageErr != nil {
		http.Redirect(res, req, EditRoute+title, http.StatusFound)
		return
	}
	dto := templateDTO{Page: page, Path: EditRoute}
	renderTemplate(res, "view", dto)
}

func EditHandler(res http.ResponseWriter, req *http.Request, title string) {
	page, _ := p.Load(title)
	dto := templateDTO{Page: page, Path: SaveRoute}
	renderTemplate(res, "edit_form", dto)
}

func SaveHandler(res http.ResponseWriter, req *http.Request, title string) {
	body := req.FormValue("body")
	pageToWrite := p.Page{Title: title, Body: body}
	err := pageToWrite.Save()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(res, req, ViewRoute+title, http.StatusFound)
}

func renderTemplate(res http.ResponseWriter, templateName string, dto templateDTO) {
	err := templates.ExecuteTemplate(res, templateName+".html", dto)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
