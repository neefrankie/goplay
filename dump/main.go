package main

import (
	_ "embed"
	"os"
	"path/filepath"

	"example.com/dump/config"
	"example.com/dump/db"
	"example.com/dump/repo"
)

//go:embed build/api.toml
var tomlConfig string

var (
	baseDir = "code/story_v2"
)

func main() {
	config.MustSetupViper([]byte(tomlConfig))

	myDBs := db.MustNewMyDBs()

	storyRepo := repo.New(myDBs.Read)

	ids, err := repo.ReadIDs(mustGetIDFile())
	if err != nil {
		panic(err)
	}

	dataDir := mustDataDir()

	storyRepo.StartDump(ids, dataDir)
}

func mustGetIDFile() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return filepath.Join(homeDir, "code/ids.txt")
}

func mustDataDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	dir := filepath.Join(homeDir, baseDir)
	os.MkdirAll(dir, 0740)

	return dir
}
