package stdlib

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestMD5String(t *testing.T) {
	h := md5.New()
	io.WriteString(h, "The fog is getting thicker")
	io.WriteString(h, "And Leon's getting laaarger!")

	t.Logf("%x\n", h.Sum(nil))
}

func TestMD5File(t *testing.T) {
	f, err := os.Open("build/pi.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		t.Fatal(err)
	}

	// Arg to h.Sum(b []byte) is the checksum already
	// calcluated. The latest checksum will simply
	// appended to b.
	t.Logf("%x\n", h.Sum(nil))
}

func TestMD5Sum(t *testing.T) {
	data := []byte("These pretzels are making me thirsty.")

	sum := md5.Sum(data)

	t.Logf("%x\n", sum)
	t.Logf("%s\n", hex.EncodeToString(sum[:]))
}

func TestSha256String(t *testing.T) {
	h := sha256.New()
	h.Write([]byte("hello world\n"))

	t.Logf("%x\n", h.Sum(nil))
}

func TestSha256File(t *testing.T) {
	f, err := os.Open("build/pi.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		t.Fatal(err)
	}

	t.Logf("%x\n", h.Sum(nil))
}

func TestSha256Sum(t *testing.T) {
	sum := sha256.Sum256([]byte("hello world\n"))
	fmt.Printf("%x\n", sum)
}
