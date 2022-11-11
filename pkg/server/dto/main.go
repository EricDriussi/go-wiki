package dto

import "wiki/pkg/page"

type Multi struct {
	Pages []*page.Page
	Paths map[string]string
}

type Single struct {
	Page  *page.Page
	Paths map[string]string
}
