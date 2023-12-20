package stdlib

import (
	"bytes"
	"html/template"
	"io"
	"io/fs"
	"path/filepath"
)

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
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
	tmpl *template.Template
}

func NewRenderer(dir string) (Renderer, error) {
	t, err := ParseTemplateDir(dir)
	if err != nil {
		return Renderer{}, err
	}

	return Renderer{
		tmpl: t,
	}, nil
}

func (r Renderer) Render(name string, data any) (string, error) {
	var b bytes.Buffer
	err := r.tmpl.ExecuteTemplate(&b, name, data)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

func (r Renderer) RenderTo(wr io.Writer, name string, data any) error {
	return r.tmpl.ExecuteTemplate(wr, name, data)
}
