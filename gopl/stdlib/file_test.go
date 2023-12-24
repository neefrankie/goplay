package stdlib

import (
	"io"
	"os"
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
