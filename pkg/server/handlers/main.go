package handlers

// TODO.refactor

import (
	"html/template"
	"net/http"
	"wiki/pkg/config"
	"wiki/pkg/page"
)

type templateDTO struct {
	Page *page.Page
	Path string
}

func IndexHandler(res http.ResponseWriter, req *http.Request, _ string) {
	allPages := page.LoadAll()
	dtos := []templateDTO{}
	for _, singlePage := range allPages {
		dtos = append(dtos, templateDTO{Page: singlePage, Path: config.ViewRoute})
	}
	renderTemplateCacheIndex(res, "index", dtos)
}

func ViewHandler(res http.ResponseWriter, req *http.Request, title string) {
	page, noPageErr := page.Load(title)
	if noPageErr != nil {
		http.Redirect(res, req, config.EditRoute+title, http.StatusFound)
		return
	}
	dto := templateDTO{Page: page, Path: config.EditRoute}
	renderTemplateCache(res, "view", dto)
}

func EditHandler(res http.ResponseWriter, req *http.Request, title string) {
	page, _ := page.Load(title)
	dto := templateDTO{Page: page, Path: config.SaveRoute}
	renderTemplateCache(res, "edit_form", dto)
}

func SaveHandler(res http.ResponseWriter, req *http.Request, title string) {
	body := req.FormValue("body")
	pageToWrite := page.New().WithTitle(title).WithBody(body)
	err := pageToWrite.Save()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(res, req, config.ViewRoute+title, http.StatusFound)
}

func renderTemplateCache(res http.ResponseWriter, templateName string, dto templateDTO) {
	templates := template.Must(template.ParseFiles(
		config.TemplatesPath+"edit_form.html",
		config.TemplatesPath+"view.html",
	),
	)

	err := templates.ExecuteTemplate(res, templateName+".html", dto)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func renderTemplateCacheIndex(res http.ResponseWriter, templateName string, dtos []templateDTO) {
	funcMap := template.FuncMap{"extract": firstFewLines}
	templates := template.Must(template.New(templateName).Funcs(funcMap).ParseFiles(
		config.TemplatesPath + "index.html",
	),
	)

	err := templates.ExecuteTemplate(res, templateName+".html", dtos)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func firstFewLines(body string) string {
	if len(body) < 501 {
		return body
	}
	return body[0:500] + "..."
}
