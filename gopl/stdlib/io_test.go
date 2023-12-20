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

func TestReadFile(t *testing.T) {
	data, err := os.ReadFile("./templates/pipeline_var.text")
	if err != nil {
		t.Error(err)
	}

	t.Logf("%s\n", data)
}
