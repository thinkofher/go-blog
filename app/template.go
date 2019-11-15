package app

import (
	"html/template"
	"path/filepath"
)

const templatesBase string = "base"
const templatesRootDir string = "templates"
const templatesWildcard string = "*.tmpl"

// BlogTemplate wraps *html/template.Template to make
// executing templates easier for the sake of this project.
type BlogTemplate interface {
	// TemplateData method returns PageData specific for wrapped templates.
	TemplateData() *PageData
	// Template method returns template ready to execute.
	Template() (*template.Template, error)
}

// NewBlogTemplate returns BlogTemplate with templates from
// 'templates/{dir}' folder, where dir is function argument,
// and with given PageData.
func NewBlogTemplate(dir string, data PageData) BlogTemplate {
	return blogTemplate{
		dir:  dir,
		data: data,
	}
}

// blogTemplate satisfies BlogTemplate interface.
type blogTemplate struct {
	dir  string
	data PageData
}

// TemplateData returns pointer to the PageData,
// containing neccesary data to execute blog template.
func (b blogTemplate) TemplateData() *PageData {
	return &b.data
}

// Template method returns pointer to html/template.Template
// ready to execute.
func (b blogTemplate) Template() (*template.Template, error) {
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

func (b blogTemplate) loadBase() (*template.Template, error) {
	return template.ParseGlob(filepath.Join(templatesRootDir, templatesWildcard))
}

func (b blogTemplate) loadTemplates(t *template.Template) (*template.Template, error) {
	return t.ParseGlob(filepath.Join(templatesRootDir, b.dir, templatesWildcard))
}
