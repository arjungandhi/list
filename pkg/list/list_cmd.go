package list

import (
	"fmt"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

// rootCmd is the main command for the list command line tool
// its just holds all the other useful commands
var Cmd = &Z.Cmd{
	Name:    "list",
	Summary: "list lists all of the users lists",
	Commands: []*Z.Cmd{
		help.Cmd,
	},
	Call: func(_ *Z.Cmd, _ ...string) error {
		files, err := List()
		if err != nil {
			return err
		}

		for name, _ := range files {
			fmt.Println(name)
		}

		return nil
	},
}
