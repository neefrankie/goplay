package store

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func mkDir(dir string) error {
	return os.MkdirAll(filepath.Dir(dir), 0755)
}

type Dir struct {
	rootDir string
}

func NewDir(rootDir string) *Dir {
	return &Dir{rootDir: rootDir}
}

func (b *Dir) GetRootDir() string {
	return b.rootDir
}

func (b *Dir) BuildPath(path ...string) string {
	return filepath.Join(append([]string{b.rootDir}, path...)...)
}

func (b *Dir) WriteText(name string, content string) (string, error) {
	p := b.BuildPath(name)

	if err := mkDir(filepath.Dir(p)); err != nil {
		return "", err
	}

	err := os.WriteFile(p, []byte(content), 0666)
	if err != nil {
		return "", err
	}

	return p, nil
}

func (b *Dir) ReadText(name string) (string, error) {
	p := b.BuildPath(name)
	content, err := os.ReadFile(p)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (b *Dir) WriteJSON(name string, v interface{}) (string, error) {
	p := b.BuildPath(name)

	if err := mkDir(filepath.Dir(p)); err != nil {
		return "", err
	}

	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	err = os.WriteFile(p, data, 0666)
	if err != nil {
		return "", err
	}

	return p, nil
}

func (b *Dir) ReadJSON(name string, v interface{}) error {
	p := b.BuildPath(name)

	data, err := os.ReadFile(p)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}
