package handlers

import (
	"net/http"
	"wiki/pkg/config"
	"wiki/pkg/page"
	"wiki/pkg/server/dto"
	"wiki/pkg/server/render"
)

func IndexHandler(res http.ResponseWriter, req *http.Request, _ string) {
	paths := map[string]string{
		"ViewPath": config.ViewRoute,
		"EditPath": config.EditRoute,
	}
	peter := dto.Multi{Pages: page.LoadAll(), Paths: paths}
	render.MultiPage(res, "index.html", peter)
}

func ViewHandler(res http.ResponseWriter, req *http.Request, title string) {
	page, noPageErr := page.Load(title)
	if noPageErr != nil {
		http.Redirect(res, req, config.EditRoute+title, http.StatusFound)
		return
	}
	dto := dto.TemplateDTO{Page: page, Path: config.EditRoute}
	render.SinglePage(res, "view.html", dto)
}

func EditHandler(res http.ResponseWriter, req *http.Request, title string) {
	page, _ := page.Load(title)
	dto := dto.TemplateDTO{Page: page, Path: config.SaveRoute}
	render.SinglePage(res, "edit_form.html", dto)
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
