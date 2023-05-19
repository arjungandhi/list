package list_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/arjungandhi/list/pkg/list"
	"github.com/golang-collections/collections/set"
)

func TestList(t *testing.T) {
	path, err := filepath.Abs("./testdata")
	if err != nil {
		t.Errorf("filepath.Abs() returned an error: %v", err)
	}
	// first set LISTDIR
	os.Setenv("LISTDIR", path)
	// setup the test cases
	answers := set.New()
	answers.Insert("1.md")
	answers.Insert("2.md")
	answers.Insert("5.md")

	// then run test
	files, err := list.List()
	if err != nil {
		t.Errorf("list.List() returned an error: %v", err)
	}

	// convert files into a set
	filesSet := set.New()
	for _, file := range files {
		filesSet.Insert(filepath.Base(file))
	}

	// compare the two sets
	diff := answers.Difference(filesSet)
	if diff.Len() != 0 {
		t.Errorf("list.List() returned the wrong files: expected %v got %v", answers, filesSet)
	}

}

func TestListNames(t *testing.T) {
	path, err := filepath.Abs("./testdata")
	if err != nil {
		t.Errorf("filepath.Abs() returned an error: %v", err)
	}
	// first set LISTDIR
	os.Setenv("LISTDIR", path)
	// setup the test cases
	answers := set.New()
	answers.Insert("1")
	answers.Insert("2")
	answers.Insert("5")

	// then run test
	files, err := list.ListNames()
	if err != nil {
		t.Errorf("list.ListNames() returned an error: %v", err)
	}

	// convert files into a set
	filesSet := set.New()
	for _, file := range files {
		filesSet.Insert(file)
	}

	// compare the two sets
	diff := answers.Difference(filesSet)
	if diff.Len() != 0 {
		t.Errorf("list.ListNames() returned the wrong files: expected %v got %v", answers, filesSet)
	}

}
