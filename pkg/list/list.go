package list

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func List() (map[string]string, error) {
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

	filesMap := make(map[string]string)
	for _, file := range files {
		fileName := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
		filesMap[fileName] = file
	}

	return filesMap, nil
}
