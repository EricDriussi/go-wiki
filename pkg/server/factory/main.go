package factory

import (
	"fmt"
	"net/http"
	"regexp"
	"wiki/pkg/config"
	h "wiki/pkg/server/handlers"
)

func GetViewHandler() http.HandlerFunc {
	return HandlerMaker(h.ViewHandler)
}

func GetEditHandler() http.HandlerFunc {
	return HandlerMaker(h.EditHandler)
}

func GetSaveHandler() http.HandlerFunc {
	return HandlerMaker(h.SaveHandler)
}

func HandlerMaker(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	validPath := regexp.MustCompile(
		fmt.Sprintf("^(%s|%s|%s)([a-zA-Z0-9_-]+)$",
			config.ViewRoute,
			config.EditRoute,
			config.SaveRoute,
		),
	)

	return func(res http.ResponseWriter, req *http.Request) {
		match := validPath.FindStringSubmatch(req.URL.Path)
		if match == nil {
			http.NotFound(res, req)
			return
		}
		fn(res, req, match[2])
	}
}
