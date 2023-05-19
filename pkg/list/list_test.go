package list_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/arjungandhi/list/pkg/list"
)

func TestList(t *testing.T) {
	path, err := filepath.Abs("./testdata")
	if err != nil {
		t.Errorf("filepath.Abs() returned an error: %v", err)
	}
	// first set LISTDIR
	os.Setenv("LISTDIR", path)
	// setup the test cases
	answers := map[string]string{
		"1": filepath.Join(path, "1.md"),
		"2": filepath.Join(path, "2.md"),
		"5": filepath.Join(path, "5.md"),
	}

	// then run test
	files, err := list.List()
	if err != nil {
		t.Errorf("list.List() returned an error: %v", err)
	}

	// convert both maps to json and compare
	filesJSON, err := json.Marshal(files)
	if err != nil {
		t.Errorf("json.Marshal(files) returned an error: %v", err)
	}

	answersJSON, err := json.Marshal(answers)
	if err != nil {
		t.Errorf("json.Marshal(answers) returned an error: %v", err)
	}

	if string(filesJSON) != string(answersJSON) {
		t.Errorf("Expected %v, got %v", answers, files)
	}
}
