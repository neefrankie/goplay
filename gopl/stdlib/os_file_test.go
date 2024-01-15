package stdlib

import (
	"os"
	"testing"
)

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
