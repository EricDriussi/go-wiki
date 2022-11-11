package handle

import (
	"html/template"
	"net/http"
	"wiki/pkg/config"
	"wiki/pkg/page"
	templateDTO "wiki/pkg/server/dtos"
)

func Index(res http.ResponseWriter, req *http.Request, _ string) {
	paths := map[string]string{
		"ViewPath": config.ViewRoute,
		"EditPath": config.EditRoute,
	}
	peter := templateDTO.Multi{Pages: page.LoadAll(), Paths: paths}
	render(res, "index.html", peter)
}

func View(res http.ResponseWriter, req *http.Request, title string) {
	paths := map[string]string{
		"EditPath": config.EditRoute,
		"BackPath": config.IndexRoute,
	}
	page, noPageErr := page.Load(title)
	if noPageErr != nil {
		http.Redirect(res, req, config.EditRoute+title, http.StatusFound)
		return
	}
	dto := templateDTO.Single{Page: page, Paths: paths}
	render(res, "view.html", dto)
}

func Edit(res http.ResponseWriter, req *http.Request, title string) {
	paths := map[string]string{
		"SavePath": config.SaveRoute,
		"BackPath": config.ViewRoute,
	}
	page, _ := page.Load(title)
	dto := templateDTO.Single{Page: page, Paths: paths}
	render(res, "edit_form.html", dto)
}

func Save(res http.ResponseWriter, req *http.Request, title string) {
	body := req.FormValue("body")
	pageToWrite := page.New().WithTitle(title).WithBody(body)
	err := pageToWrite.Save()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(res, req, config.ViewRoute+title, http.StatusFound)
}

func render(res http.ResponseWriter, templateName string, dto interface{}) {
	funcMap := template.FuncMap{"extract": firstFewLines}
	templates := template.Must(
		template.
			New(templateName).
			Funcs(funcMap).
			ParseFiles(
				config.TemplatesPath+templateName,
				config.TemplatesPath+templateName,
				config.TemplatesPath+templateName,
			),
	)

	err := templates.ExecuteTemplate(res, templateName, dto)
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
