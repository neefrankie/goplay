package stdlib

import (
	_ "embed"
	"io/fs"
	"os"
	"testing"
)

//go:embed templates/index.html
var s string

//go:embed templates/index.html
var b []byte

func TestEbmedFileToString(t *testing.T) {
	print(s)
}

func TestEmbedFileToBytes(t *testing.T) {
	print(string(b))
}

func TestEmbedFS(t *testing.T) {
	fs.WalkDir(templates, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("%s | Is dir: %t", path, d.IsDir())

		return nil
	})
}

func TestFsWalkDir(t *testing.T) {
	root := "./templates"
	filesystem := os.DirFS(root)

	// fs.WalkDir always uses slash separated paths.
	// filepath.WalkDir uses the seaprator character appropriate for the operating system.
	fs.WalkDir(filesystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%s | Is dir: %t\n", path, d.IsDir())
		return nil
	})
}
