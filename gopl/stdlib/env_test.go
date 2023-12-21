package stdlib

import (
	"os"
	"testing"
)

func initEnv() {
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")
}

func TestGetEnv(t *testing.T) {
	initEnv()

	t.Logf("%s lives in %s.\n", os.Getenv("NAME"), os.Getenv("BURROW"))
}

func TestUnsetEnv(t *testing.T) {
	os.Setenv("TMPDIR", "/my/tmp")
	defer os.Unsetenv("TMPDIR")
}

func TestExpandEnv(t *testing.T) {
	initEnv()

	t.Logf("%s\n", os.ExpandEnv("$NAME lives in ${BURROW}"))
}

func TestExpand(t *testing.T) {
	mapper := func(placeholderName string) string {
		switch placeholderName {
		case "DAY_PART":
			return "morning"
		case "NAME":
			return "Gopher"
		}

		return ""
	}

	t.Logf("%s\n", os.Expand("Good ${DAY_PART}, $NAME!", mapper))
}

func TestEnviron(t *testing.T) {
	t.Logf("%s\n", os.Environ())
}

func TestLookupEnv(t *testing.T) {
	show := func(key string) {
		val, ok := os.LookupEnv(key)
		if !ok {
			t.Logf("%s not set\n", key)
		} else {
			t.Logf("%s=%s\n", key, val)
		}
	}

	os.Setenv("SOME_KEY", "value")
	os.Setenv("EMPTY_KEY", "")

	show("SOME_KEY")
	show("EMPTY_KEY")
	show("MISSING_KEY")
}
