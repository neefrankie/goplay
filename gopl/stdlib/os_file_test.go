package stdlib

import (
	"io"
	"os"
	"testing"
)

// os.ReadFile read the file and returns the content as `[]byte`.
// Under the hook it uses os.Open().
// To manipulate files as `[]byte`, use `os.ReadFile“ and `os.WriteFile`
// directly without toucing the File object.
func TestReadFile(t *testing.T) {
	data, err := os.ReadFile("./templates/index.html")
	if err != nil {
		t.Error(err)
	}

	t.Logf("%s\n", data)
}

func TestWriteFile(t *testing.T) {
	err := os.WriteFile("build/hello.txt", []byte("Hello, Gpphers"), 0666)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpen(t *testing.T) {
	f, err := os.Open("./templates/index.html")
	if err != nil {
		t.Fatal(err)
	}

	io.Copy(os.Stdout, f)
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
