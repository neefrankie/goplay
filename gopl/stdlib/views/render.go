package views

import (
	"bytes"
	"io"
	"io/fs"
	"path/filepath"
	"text/template"
)

// ParseTemplateDir parses all templates in one shot.
// To make this function work correctly, do not use `block/define` mechanism.
// to mimic inheritance since it does not work in a simple way.
// Treat the template only has include features, and include header and footer in
// every page.
func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.WalkDir(dir, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

type Renderer struct {
	templates *template.Template
}

// NewRenderer parses template from current filesystem.
func NewRenderer(dir string) (Renderer, error) {
	t, err := ParseTemplateDir(dir)
	if err != nil {
		return Renderer{}, err
	}

	return Renderer{
		templates: t,
	}, nil
}

func MustNewRenderer(dir string) Renderer {
	r, err := NewRenderer(dir)
	if err != nil {
		panic(err)
	}

	return r
}

func (r Renderer) RenderString(name string, data any) (string, error) {
	var b bytes.Buffer
	err := r.templates.ExecuteTemplate(&b, name, data)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

// RenderTo write outputs to a writer, such as
// file, network stream, etc.
func (r Renderer) Render(wr io.Writer, name string, data any) error {
	return r.templates.ExecuteTemplate(wr, name, data)
}
