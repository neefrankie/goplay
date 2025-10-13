package thumbnail

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// ImageFile reads an image from infile and writes
// a thumbnail-size version of it in the same directory.
// It returns the generated file name, e.g., "foo.thumb.jpg".
func ImageFile(infile string, outdir string) (string, error) {
	err := os.MkdirAll(outdir, 0750)
	if err != nil {
		return "", err
	}
	outfile := thumbFileName(infile, outdir)

	return outfile, createThumb(outfile, infile)
}

func thumbFileName(infile string, outdir string) string {
	name, ext := SplitExt(filepath.Base(infile))

	return filepath.Join(outdir, name) + ".thumb" + ext
}

func imageFormat(infile string) string {
	ext := filepath.Ext(infile)
	ext = strings.ToLower(ext)
	ext = strings.TrimPrefix(ext, ".")

	if ext == "jpg" || ext == "jpeg" {
		return "jpeg"
	}
	return ext
}

func createThumb(outfile, infile string) error {

	in, err := os.Open(infile)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(outfile)
	if err != nil {
		return err
	}

	format := imageFormat(infile)
	if err := imageStream(out, in, format); err != nil {
		out.Close()
		return fmt.Errorf("scaling %s to %s: %s", infile, outfile, err)
	}

	return out.Close()
}

func imageStream(w io.Writer, r io.Reader, format string) error {
	src, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	dst := resizeImage(src)

	switch format {
	case "jpeg":
		return jpeg.Encode(w, dst, nil)
	case "png":
		return png.Encode(w, dst)
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}
}

func resizeImage(src image.Image) image.Image {
	xs := src.Bounds().Size().X
	ys := src.Bounds().Size().Y

	width, height := 128, 128
	if aspect := float64(xs) / float64(ys); aspect < 1.0 {
		width = int(128 * aspect)
	} else {
		height = int(128 / aspect)
	}

	xscale := float64(xs) / float64(width)
	yscale := float64(ys) / float64(height)

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			srcx := int(float64(x) * xscale)
			srcy := int(float64(y) * yscale)
			dst.Set(x, y, src.At(srcx, srcy))
		}
	}

	return dst
}
