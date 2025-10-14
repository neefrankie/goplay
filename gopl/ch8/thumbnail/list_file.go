package thumbnail

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// ExpandPath expands paths starting with ~/ to the user's home directory.
// It does not expand environment variables like $HOME.
func ExpandPath(path string) (string, error) {
	if len(path) == 0 {
		return path, nil
	}

	if path[0] != '~' {
		return path, nil
	}

	// Only handle ~/ or ~<separator>
	if len(path) == 1 || path[1] == '/' || path[1] == '\\' {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, path[2:]), nil
	}

	// Do not expand ~username or other forms
	return path, nil
}

func SplitExt(path string) (string, string) {
	ext := filepath.Ext(path)

	return strings.TrimSuffix(path, ext), ext
}

func WithSuffix(path, suffix string) string {
	p, _ := SplitExt(path)
	return p + suffix
}

func ListFiles(dir string) ([]string, error) {
	dir, err := ExpandPath(dir)
	if err != nil {
		return nil, err
	}

	var paths []string

	err = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if strings.HasPrefix(d.Name(), ".") {
			return nil
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

func listFilesAsync(dir string) <-chan string {

	ch := make(chan string)

	go func() {
		absDir, err := ExpandPath(dir)
		if err != nil {
			log.Printf("error: %s", err)
			return
		}

		err = filepath.WalkDir(absDir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if strings.HasPrefix(d.Name(), ".") {
				return nil
			}

			if !d.IsDir() {
				ch <- path
			}

			return nil
		})
		if err != nil {
			log.Printf("error: %", err)
		}
	}()

	return ch
}
