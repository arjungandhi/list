package edit

import (
	"fmt"
	"os"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"

	"github.com/arjungandhi/list/pkg/complist"
	"github.com/arjungandhi/list/pkg/list"
)

// rootCmd is the main command for the list command line tool
// its just holds all the other useful commands
var Cmd = &Z.Cmd{
	Name:    "edit",
	Summary: "edits a users list",
	Commands: []*Z.Cmd{
		help.Cmd,
	},
	Comp:    complist.New(),
	Usage:   "<list name>",
	NumArgs: 1,
	Call: func(_ *Z.Cmd, args ...string) error {

		files, err := list.List()
		if err != nil {
			return err
		}

		if _, ok := files[args[0]]; !ok {
			return fmt.Errorf("list %q does not exist", args[0])
		}

		// start the editor with the list file
		editor, exists := os.LookupEnv("EDITOR")
		if !exists {
			editor = "vi"
		}
		Z.SysExec(editor, files[args[0]])

		return nil
	},
}
