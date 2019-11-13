package app

type PageData struct {
	Title   string
	Flashes []interface{}
}

func NewPageData(title string) *PageData {
	return &PageData{
		Title: title,
	}
}

func (p *PageData) SetFlashes(flashes []interface{}) {
	p.Flashes = flashes
}
