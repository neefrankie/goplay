package stdlib

import (
	"image/png"
	"os"
	"testing"
)

func Test_createImage(t *testing.T) {
	img := createImage()

	f, err := os.Create("build/myimage.png")
	defer f.Close()

	if err != nil {
		t.Fatal(err)
	}

	err = png.Encode(f, img)
	if err != nil {
		t.Fatal(err)
	}
}
