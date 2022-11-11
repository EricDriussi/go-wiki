package handle

import (
	"net/http"
	"wiki/pkg/config"
	"wiki/pkg/page"
	"wiki/pkg/server/dtos"
	"wiki/pkg/server/renderers"
)

func Index(res http.ResponseWriter, req *http.Request, _ string) {
	paths := map[string]string{
		"ViewPath": config.ViewRoute,
		"EditPath": config.EditRoute,
	}
	peter := templateDTO.Multi{Pages: page.LoadAll(), Paths: paths}
	render.MultiPage(res, "index.html", peter)
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
	render.SinglePage(res, "view.html", dto)
}

func Edit(res http.ResponseWriter, req *http.Request, title string) {
	paths := map[string]string{
		"SavePath": config.SaveRoute,
		"BackPath": config.ViewRoute,
	}
	page, _ := page.Load(title)
	dto := templateDTO.Single{Page: page, Paths: paths}
	render.SinglePage(res, "edit_form.html", dto)
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
