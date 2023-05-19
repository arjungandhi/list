package list

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func List() ([]string, error) {
	// get LISTDIR from environment
	listDir := os.Getenv("LISTDIR")
	if listDir == "" {
		return nil, errors.New("LISTDIR not set")
	}

	// get all .md files in LISTDIR
	files, err := filepath.Glob(filepath.Join(listDir, "*.md"))
	if err != nil {
		return nil, err
	}

	return files, nil
}

func ListNames() ([]string, error) {
	files, err := List()
	if err != nil {
		return nil, err
	}

	names := make([]string, len(files))
	// return files without ext
	for i, file := range files {
		names[i] = strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
	}

	return names, nil
}
