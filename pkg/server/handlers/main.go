package handlers

import (
	"html/template"
	"net/http"
	"wiki/pkg"

	p "wiki/pkg/page"
)

type templateDTO struct {
	Page *p.Page
	Path string
}

func ViewHandler(res http.ResponseWriter, req *http.Request, title string) {
	page, noPageErr := p.Load(title)
	if noPageErr != nil {
		http.Redirect(res, req, pkg.EditRoute+title, http.StatusFound)
		return
	}
	dto := templateDTO{Page: page, Path: pkg.EditRoute}
	renderTemplateCache(res, "view", dto)
}

func EditHandler(res http.ResponseWriter, req *http.Request, title string) {
	page, _ := p.Load(title)
	dto := templateDTO{Page: page, Path: pkg.SaveRoute}
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
	http.Redirect(res, req, pkg.ViewRoute+title, http.StatusFound)
}

func renderTemplateCache(res http.ResponseWriter, templateName string, dto templateDTO) {
	templates := template.Must(template.ParseFiles(
		pkg.TemplatesPath+"edit_form.html",
		pkg.TemplatesPath+"view.html",
	),
	)

	err := templates.ExecuteTemplate(res, templateName+".html", dto)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
