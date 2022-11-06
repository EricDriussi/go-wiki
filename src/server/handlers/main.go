package handlers

import (
	"html/template"
	"net/http"

	config "wiki/src"
	p "wiki/src/page"
)

type templateDTO struct {
	Page *p.Page
	Path string
}

var (
	templatesPath = "src/server/html_templates/"
	templates     = template.Must(
		template.ParseFiles(
			templatesPath+"edit_form.html",
			templatesPath+"view.html",
		),
	)
)

func ViewHandler(res http.ResponseWriter, req *http.Request, title string) {
	page, noPageErr := p.Load(title)
	if noPageErr != nil {
		http.Redirect(res, req, config.EditRoute+title, http.StatusFound)
		return
	}
	dto := templateDTO{Page: page, Path: config.EditRoute}
	renderTemplateCache(res, "view", dto)
}

func EditHandler(res http.ResponseWriter, req *http.Request, title string) {
	page, _ := p.Load(title)
	dto := templateDTO{Page: page, Path: config.SaveRoute}
	renderTemplateCache(res, "edit_form", dto)
}

func SaveHandler(res http.ResponseWriter, req *http.Request, title string) {
	body := req.FormValue("body")
	pageToWrite := p.Page{Title: title, Body: body}
	err := pageToWrite.Save()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(res, req, config.ViewRoute+title, http.StatusFound)
}

func renderTemplateCache(res http.ResponseWriter, templateName string, dto templateDTO) {
	err := templates.ExecuteTemplate(res, templateName+".html", dto)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
