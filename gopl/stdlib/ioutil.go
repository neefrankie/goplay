package stdlib

import (
	"io"
	"os"
	"path/filepath"
)

func MkdirParent(filename string) error {
	dir := filepath.Dir(filename)
	return os.MkdirAll(dir, 0750)
}

func SaveFile(filename string, data []byte) error {
	err := MkdirParent(filename)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0666)
	if err != nil {
		return err
	}

	return nil
}

func SaveString(filename string, s string) (int, error) {
	err := MkdirParent(filename)
	if err != nil {
		return 0, err
	}

	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		return 0, err
	}

	return io.WriteString(f, s)
}

func Save(filename string, r io.Reader) (int64, error) {
	err := MkdirParent(filename)
	if err != nil {
		return 0, err
	}

	f, err := os.Create(filename)
	defer f.Close()

	if err != nil {
		return 0, err
	}

	return io.Copy(f, r)
}
