package cache

import (
	"os"
	"path/filepath"
)

func CacheDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".ghscaffold", "repos")
}

func ClearReposCache() error {
	return os.RemoveAll(CacheDir())
}

// ...functions to cache and retrieve README files...
