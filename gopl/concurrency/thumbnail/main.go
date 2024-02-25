package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

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

func imageStream(w io.Writer, r io.Reader) error {
	src, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	dst := resizeImage(src)

	return jpeg.Encode(w, dst, nil)
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

	if err := imageStream(out, in); err != nil {
		out.Close()
		return fmt.Errorf("scaling %s to %s: %s", infile, outfile, err)
	}

	return out.Close()
}

func imageFile(infile string) (string, error) {
	outfile := thumFileName(infile)

	return outfile, createThumb(outfile, infile)
}

func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go imageFile(f)
	}
}

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			imageFile(f)
		}(f)
	}

	for range filenames {
		<-ch
	}
}

func makeThumbnails4(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		go func(f string) {
			_, err := imageFile(f)

			errors <- err
		}(f)
	}

	for range filenames {
		// When it enctouners the first non-nil error, it returns the error
		// to the caller, leaving no goroutine draining the errors channel.
		// Each reamaining goroutine will block forever when it tries to send
		// a value on that channel, and will never terminate.
		if err := <-errors; err != nil {
			return err // Goroutine leaks
		}
	}

	return nil
}

func makeThumnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = imageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

func makeThumnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup

	for f := range filenames {
		// Add must be called before the worker goroutine starts
		wg.Add(1)

		go func(f string) {
			defer wg.Done()
			thumb, err := imageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	// wait and close mut be concurrent with the loop over sizes.
	//
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}

	return total
}

func main() {
	filenames, err := listFiles("build/images")
	if err != nil {
		log.Fatal(err)
	}

	makeThumbnails3(filenames)
}

func listFiles(dir string) ([]string, error) {
	var paths []string

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			paths = append(paths, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return paths, nil
}

func thumFileName(infile string) string {
	ext := path.Ext(infile)
	base := strings.TrimSuffix(infile, ext)

	return base + ".thumb" + ext
}
