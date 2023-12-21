package stdlib

import (
	"path"
	"path/filepath"
	"testing"
)

func TestPathBase(t *testing.T) {
	t.Logf("%s\n", path.Base("/a/b"))
	t.Logf("%s\n", path.Base("/"))
	t.Logf("%s\n", path.Base(""))
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
