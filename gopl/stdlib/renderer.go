package stdlib

import (
	"bytes"
	"html/template"
	"io"
	"io/fs"
)

type Renderer struct {
	fsys   fs.FS
	dir    string
	layout *template.Template
	cache  map[string]*template.Template
}

// NewFSRenderer parses tempalte files from a dire in a
// cetain filesystem.
// To use os filesystem, wrap the directory in os.DirFS("dir").
func NewFSRenderer(fsys fs.FS, dir string, layoutDir string) (Renderer, error) {

	var t *template.Template
	var err error
	if layoutDir != "" {
		t, err = ParseFSTemplateDir(fsys, layoutDir)

		if err != nil {
			return Renderer{}, err
		}
	} else {
		t = template.New("__EMPTY_LAYOUT")
	}

	return Renderer{
		fsys:   fsys,
		dir:    dir,
		layout: t,
		cache:  make(map[string]*template.Template),
	}, nil
}

func MustNewFSRenderer(fsys fs.FS, dir string, layoutDir string) Renderer {
	r, err := NewFSRenderer(fsys, dir, layoutDir)
	if err != nil {
		panic(err)
	}

	return r
}

func (r Renderer) RenderString(name string, data any) (string, error) {
	var b bytes.Buffer
	err := r.Render(&b, name, data)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

// Render write outputs to a writer, such as
// file, network stream, etc.
func (r Renderer) Render(wr io.Writer, name string, data any) error {
	if t, ok := r.cache[name]; ok {
		return t.Execute(wr, data)
	}

	layout, err := r.layout.Clone()
	if err != nil {
		return err
	}

	// Do not use filepath.Join here since it is os-specific.
	// Use path.Join or just string concatenation.
	t, err := layout.ParseFS(r.fsys, r.dir+"/"+name)
	if err != nil {
		return err
	}
	r.cache[name] = t
	return t.Execute(wr, data)
}

// ParseFSTemplateDir parses all tempalate files in a directory.
// Please note that templates parsed in sucha a way should not have
// the same name in define appear more than once.
// If the same name is defined multiple times, the last parsed one wins.
// file1.html {{block "parent"}}{{end}}
// file2.html {{define "parent"}}Inherit from parent{{end}}
// file3.html {{define "parent"}}Inherit again{{end}}
// If you parse the three files in this functions, you could only get
// file2.html or file3.html, not both.
// Go's template never acutally implemented Jinjia style inheritance.
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
