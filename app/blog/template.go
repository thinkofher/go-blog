package blog

import (
	"html/template"
	"path/filepath"
)

// TemplatesBase contains name of filename
// with base template of go-blog application.
const TemplatesBase string = "base"

const templatesRootDir string = "templates"
const templatesWildcard string = "*.tmpl"

// Renderer wraps *html/template.Template to make
// executing templates easier for the sake of this project.
type Renderer interface {
	// Data method returns PageData specific for wrapped templates.
	Data() *Data
	// Render method returns template ready to execute.
	Render() (*template.Template, error)
}

// NewRenderer returns Renderer with templates from
// 'templates/{dir}' folder, where dir is function argument,
// and with given PageData.
func NewRenderer(dir string, data Data) Renderer {
	return renderer{
		dir:  dir,
		data: data,
	}
}

// renderer satisfies Renderer interface.
type renderer struct {
	dir  string
	data Data
}

// Data returns pointer to the PageData,
// containing neccesary data to execute blog template.
func (b renderer) Data() *Data {
	return &b.data
}

// Render method returns pointer to html/template.Template
// ready to execute.
func (b renderer) Render() (*template.Template, error) {
	tmpl, err := b.loadBase()
	if err != nil {
		return nil, err
	}

	tmpl, err = b.loadTemplates(tmpl)
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

func (b renderer) loadBase() (*template.Template, error) {
	return template.ParseGlob(filepath.Join(templatesRootDir, templatesWildcard))
}

func (b renderer) loadTemplates(t *template.Template) (*template.Template, error) {
	return t.ParseGlob(filepath.Join(templatesRootDir, b.dir, templatesWildcard))
}
