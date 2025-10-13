package thumbnail_test

import (
	"gopl/ch8/thumbnail"
	"path/filepath"
	"testing"
)

func TestNameExt(t *testing.T) {
	name, ext := thumbnail.SplitExt("/Users/a/a.jpg")

	t.Logf("Name: %s, ext: %s\n", name, ext)

	basename := filepath.Base("/Users/a/b.png")
	t.Logf("basename: %s\n", basename)
}
