package render

import (
	"html/template"
	"net/http"
	"wiki/pkg/config"
	"wiki/pkg/server/dto"
)

func SinglePage(res http.ResponseWriter, templateName string, dto dto.TemplateDTO) {
	templates := template.Must(template.ParseFiles(
		config.TemplatesPath+"edit_form.html",
		config.TemplatesPath+"view.html",
	),
	)

	err := templates.ExecuteTemplate(res, templateName, dto)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func MultiPage(res http.ResponseWriter, templateName string, dtos dto.Multi) {
	funcMap := template.FuncMap{"extract": firstFewLines}
	templates := template.Must(template.New(templateName).Funcs(funcMap).ParseFiles(
		config.TemplatesPath + "index.html",
	),
	)

	err := templates.ExecuteTemplate(res, templateName, dtos)
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
