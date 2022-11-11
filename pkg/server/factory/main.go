package factory

import (
	"fmt"
	"net/http"
	"regexp"
	"wiki/pkg/config"
	"wiki/pkg/server/handlers"
)

func IndexHandler() http.HandlerFunc {
	return buildValidHandler(handle.Index)
}

func ViewHandler() http.HandlerFunc {
	return buildValidHandler(handle.View)
}

func EditHandler() http.HandlerFunc {
	return buildValidHandler(handle.Edit)
}

func SaveHandler() http.HandlerFunc {
	return buildValidHandler(handle.Save)
}

func buildValidHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	validPath := regexp.MustCompile(
		fmt.Sprintf("^%s$|^(%s|%s|%s)([a-zA-Z0-9_-]+)$",
			config.IndexRoute,
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
