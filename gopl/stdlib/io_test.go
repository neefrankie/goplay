package stdlib

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestIOReadAll(t *testing.T) {
	r := strings.NewReader("Go is a general-purpose language designed with system promgramming in mind")

	b, err := io.ReadAll(r)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%s", b)
}

func TestIOReadFull(t *testing.T) {
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

func TestIOWriteString(t *testing.T) {
	_, err := io.WriteString(os.Stdout, "Hello World")

	if err != nil {
		t.Fatal(err)
	}
}
