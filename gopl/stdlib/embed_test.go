package stdlib

import (
	"io"
	"io/fs"
	"testing"
)

func TestEmbedReadDir(t *testing.T) {
	ent, err := templates.ReadDir("templates")
	if err != nil {
		t.Fatal(err)
	}

	for _, e := range ent {
		t.Logf("%s\n", e.Name())
	}
}

func TestEmbedReadFile(t *testing.T) {
	b, err := templates.ReadFile("templates/index.html")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s\n", b)
}

func TestEmbedOpen(t *testing.T) {
	file, err := templates.Open("templates/index.html")
	if err != nil {
		t.Fatal(err)
	}

	b, err := io.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s\n", b)
}

func TestEmbedWalkDir(t *testing.T) {
	err := fs.WalkDir(templates, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		t.Logf("%s | Is dir: %t\n", path, d.IsDir())

		return nil
	})

	t.Fatal(err)
}
