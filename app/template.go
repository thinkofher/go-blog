package app

import (
	"path/filepath"
	"text/template"
)

type BlogTemplate interface {
	TemplateData() *PageData
	Template() (*template.Template, error)
}

func NewBlogTemplate(dir string, data PageData) BlogTemplate {
	return blogTemplate{
		dir:  dir,
		data: data,
	}
}

type blogTemplate struct {
	dir  string
	data PageData
}

func (b blogTemplate) TemplateData() *PageData {
	return &b.data
}

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

const templatesBase string = "base"
const templatesRootDir string = "templates"
const templatesWildcard string = "*.tmpl"
