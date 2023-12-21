package stdlib

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReadAll(t *testing.T) {
	r := strings.NewReader("Go is a general-purpose language designed with system promgramming in mind")

	b, err := io.ReadAll(r)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%s", b)
}

func TestReadFull(t *testing.T) {
	r := strings.NewReader("some io.Reader stream to be read\n")

	buf := make([]byte, 4)
	if _, err := io.ReadFull(r, buf); err != nil {
		t.Error(err)
	}

	t.Logf("%s\n", buf)

	longBuf := make([]byte, 64)
	if _, err := io.ReadFull(r, longBuf); err != nil {
		t.Log("error: ", err)
	}
}

func TestWriteString(t *testing.T) {
	_, err := io.WriteString(os.Stdout, "Hello World")

	if err != nil {
		t.Fatal(err)
	}
}

func TestReadFile(t *testing.T) {
	data, err := os.ReadFile("./templates/index.html")
	if err != nil {
		t.Error(err)
	}

	t.Logf("%s\n", data)
}

func TestOpenFile(t *testing.T) {
	f, err := os.OpenFile("build/notes.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestWriteFile(t *testing.T) {
	err := os.WriteFile("build/hello", []byte("Hello, Gpphers"), 0666)
	if err != nil {
		t.Fatal(err)
	}
}

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

func TestReadDir(t *testing.T) {
	files, err := os.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range files {
		t.Logf("%s\n", file.Name())
	}
}
