package stdlib

import (
	"path"
	"testing"
)

func TestPathBase(t *testing.T) {
	t.Logf("%s\n", path.Base("/a/b"))
	t.Logf("%s\n", path.Base("/"))
	t.Logf("%s\n", path.Base(""))
}
