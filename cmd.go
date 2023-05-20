// zet Command Line tool
package list

import (
	"github.com/arjungandhi/list/pkg/archive"
	"github.com/arjungandhi/list/pkg/edit"
	l "github.com/arjungandhi/list/pkg/list"
	"github.com/arjungandhi/list/pkg/new"
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
		edit.Cmd,
		archive.Cmd,
		new.Cmd,
	},
}
