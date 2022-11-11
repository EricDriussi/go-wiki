package templateDTO

import "wiki/pkg/page"

type Valid interface {
	AllPaths() []string
	FirstPage() *page.Page
}

type Multi struct {
	Pages []*page.Page
	Paths map[string]string
}

func (this Multi) AllPaths() []string {
	var paths []string
	for _, path := range this.Paths {
		paths = append(paths, path)
	}
	return paths
}

func (this Multi) FirstPage() *page.Page {
	return this.Pages[0]
}

type Single struct {
	Page  *page.Page
	Paths map[string]string
}

func (this Single) AllPaths() []string {
	var paths []string
	for _, path := range this.Paths {
		paths = append(paths, path)
	}
	return paths
}

func (this Single) FirstPage() *page.Page {
	return this.Page
}
