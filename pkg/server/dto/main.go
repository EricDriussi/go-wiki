package dto

import "wiki/pkg/page"

type TemplateDTO struct {
	Page *page.Page
	Path string
}

type Multi struct {
	Pages []*page.Page
	Paths map[string]string
}
