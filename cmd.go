// zet Command Line tool
package list

import (
	l "github.com/arjungandhi/list/pkg/list"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

// rootCmd is the main command for the list command line tool
// its just holds all the other useful commands
var Cmd = &Z.Cmd{
	Name:    "list",
	Summary: "list is a command line tool for managing lists",
	Commands: []*Z.Cmd{
		help.Cmd,
		l.Cmd,
	},
}