package config

import (
	"os"
	"path/filepath"
)

func MustUserHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return home
}

type DataPath struct {
	Root    string
	Meta    string
	Content string
}

func NewDataPath() *DataPath {
	home := MustUserHomeDir()

	rootDir := filepath.Join(
		home,
		"codes",
		"youkre-editor",
	)
	metaDir := filepath.Join(rootDir, "data")
	contentDir := filepath.Join(rootDir, "content")

	return &DataPath{
		Root:    rootDir,
		Meta:    metaDir,
		Content: contentDir,
	}
}
