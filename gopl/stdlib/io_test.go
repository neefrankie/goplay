package stdlib

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
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

func TestStringsJSON(t *testing.T) {
	var b strings.Builder

	enc := json.NewEncoder(&b)
	enc.SetIndent("", "\t")
	err := enc.Encode(map[string]string{
		"hello": "world",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s", b.String())
}

func TestBufferJSON(t *testing.T) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetIndent("", "\t")

	err := enc.Encode(map[string]string{
		"hello": "world",
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s", buf.String())
}

// A BUffer is a variable-sized buffer bytes with Read and Write methods.
// The zero value for Buffer is an empty buffer ready to use.
func TestBufferWriter(t *testing.T) {
	var b bytes.Buffer

	b.Write([]byte("Hello "))
	fmt.Fprint(&b, "world!")

	b.WriteTo(os.Stdout)
}

func TestBufferReader(t *testing.T) {
	buf := bytes.NewBufferString("R29waGVycyBydWxlIQ==")
	dec := base64.NewDecoder(base64.StdEncoding, buf)
	io.Copy(os.Stdout, dec)
}
