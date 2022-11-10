package page

type Page struct {
	Title string
	Body  string
}

func New() Page {
	return Page{}
}

func (this Page) WithTitle(title string) Page {
	this.Title = title
	return this
}

func (this Page) WithBody(body string) Page {
	this.Body = body
	return this
}
