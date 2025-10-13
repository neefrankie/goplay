package stdlib

import (
	"os"
	"path"
	"path/filepath"
	"testing"
)

func TestPathBase(t *testing.T) {
	t.Logf("%s\n", path.Base("/a/b"))
	t.Logf("%s\n", path.Base("/"))
	t.Logf("%s\n", path.Base(""))
	t.Logf("%s\n", path.Base("basename.jpg"))
}

func TestPathGlob(t *testing.T) {
	matches, err := filepath.Glob("templates/*")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s\n", matches)

	matches, err = filepath.Glob("templates/**/*")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s\n", matches)
}

func TestPathJoin(t *testing.T) {
	// a/b/c
	p1 := path.Join("a", "b", "c")
	t.Logf("path.Join(): %s\n", p1)

	// a\b\c on Windows.
	p2 := filepath.Join("a", "b", "c")
	t.Logf("filepath.Join(): %s\n", p2)
}

func TestExpandHome(t *testing.T) {
	p := os.ExpandEnv("$HOME/datasource")

	t.Logf("ExpandHome(): %s\n", p)
}
