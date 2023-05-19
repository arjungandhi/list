package complist_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/arjungandhi/list/pkg/complist"
	mapset "github.com/deckarep/golang-set/v2"
	Z "github.com/rwxrob/bonzai/z"
)

func TestComp(t *testing.T) {
	path, err := filepath.Abs("./testdata")
	if err != nil {
		t.Errorf("filepath.Abs() returned an error: %v", err)
	}
	// first set LISTDIR
	os.Setenv("LISTDIR", path)

	cmd := &Z.Cmd{
		Name:    "edit",
		Summary: "edits a users list",
	}

	completer := complist.New()

	// test the comp command in various situations

	// args:
	// expected: edit
	res := completer.Complete(cmd, []string{}...)
	expected := []string{"edit"}
	if !equal(res, expected) {
		t.Errorf("expected: %v, got: %v", expected, res)
	}

	// args:
	// expected: a, abc, abcde, b
	res = completer.Complete(cmd, []string{""}...)
	expected = []string{"a", "abc", "abcde", "b"}
	if !equal(res, expected) {
		t.Errorf("expected: %v, got: %v", expected, res)
	}

	// args: a
	// expected: a, abc, abcde
	res = completer.Complete(cmd, []string{"a"}...)
	expected = []string{"a", "abc", "abcde"}
	if !equal(res, expected) {
		t.Errorf("expected: %v, got: %v", expected, res)
	}

	// args: ab
	// expected: abc, abcde
	res = completer.Complete(cmd, []string{"ab"}...)
	expected = []string{"abc", "abcde"}
	if !equal(res, expected) {
		t.Errorf("expected: %v, got: %v", expected, res)
	}

}

func equal(a []string, b []string) bool {
	aSet := mapset.NewSet[string]()
	for _, v := range a {
		aSet.Add(v)
	}

	bSet := mapset.NewSet[string]()
	for _, v := range b {
		bSet.Add(v)
	}

	return aSet.Equal(bSet)
}
