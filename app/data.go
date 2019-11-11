package app

type PageData struct {
	Title string
}

func NewPageData(title string) *PageData {
	return &PageData{
		Title: title,
	}
}
