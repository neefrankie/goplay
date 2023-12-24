package stdlib

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMkdir(t *testing.T) {
	err := os.Mkdir("build/testdir", 0750)
	if err != nil && !os.IsExist(err) {
		t.Fatal(err)
	}
}

func TestMkdirAll(t *testing.T) {
	err := os.MkdirAll("build/test/subdir", 0750)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMkdirTemp(t *testing.T) {
	dir, err := os.MkdirTemp("build", "example")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s\n", dir)
	defer os.RemoveAll(dir)
}

func TestMkdirTempSuffix(t *testing.T) {
	logsDir, err := os.MkdirTemp("", "*-logs")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(logsDir)

	globPattern := filepath.Join(os.TempDir(), "*-logs")
	matches, err := filepath.Glob(globPattern)
	if err != nil {
		t.Fatal(err)
	}

	for _, match := range matches {
		if err := os.RemoveAll(match); err != nil {
			t.Logf("Failed to remove %q: %v", match, err)
		}
	}
}

func TestTempDir(t *testing.T) {
	dir := os.TempDir()
	t.Logf("%s\n", dir)
}

func TestReadDir(t *testing.T) {
	files, err := os.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range files {
		t.Logf("%s\n", file.Name())
	}
}

func TestUserDir(t *testing.T) {
	dir, err := os.UserCacheDir()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("User Cache directory: %s\n", dir)

	dir, err = os.UserConfigDir()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("User config directory: %s\n", dir)

	dir, err = os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("User home directory: %s", dir)
}

func TestWorkingDir(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s\n", dir)
}
