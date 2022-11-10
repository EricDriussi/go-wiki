package page

type Page struct {
	title string
	body  string
}

func New() Page {
	return Page{}
}

func (this Page) WithTitle(title string) Page {
	this.title = title
	return this
}

func (this Page) Title() string {
	return this.title
}

func (this Page) WithBody(body string) Page {
	this.body = body
	return this
}

func (this Page) Body() string {
	return this.body
}
