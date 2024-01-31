package stdlib

import (
	"bytes"
	"html/template"
	"io"
	"io/fs"
	"path/filepath"
)

type Renderer struct {
	templates *template.Template
}

// NewFSRenderer parses tempalte files from a dire in a
// cetain filesystem.
func NewFSRenderer(fsys fs.FS, dir string) (Renderer, error) {
	t, err := ParseFSTemplateDir(fsys, dir)

	if err != nil {
		return Renderer{}, err
	}

	return Renderer{
		templates: t,
	}, nil
}

func MustNewFSRenderer(fsys fs.FS, dir string) Renderer {
	r, err := NewFSRenderer(fsys, dir)
	if err != nil {
		panic(err)
	}

	return r
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

func (r Renderer) Render(name string, data any) (string, error) {
	var b bytes.Buffer
	err := r.templates.ExecuteTemplate(&b, name, data)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

// RenderTo write outputs to a writer, such as
// file, network stream, etc.
func (r Renderer) RenderTo(wr io.Writer, name string, data any) error {
	return r.templates.ExecuteTemplate(wr, name, data)
}

func ParseFSTemplateDir(fsys fs.FS, dir string) (*template.Template, error) {
	var paths []string
	// The problem with this approach is that inheritance does not work.
	// The same name in define will override each other. The define name must be unique.
	err := fs.WalkDir(fsys, dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			paths = append(paths, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFS(fsys, paths...)
}

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
