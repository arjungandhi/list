package archive

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"

	"github.com/arjungandhi/list/pkg/complist"
	"github.com/arjungandhi/list/pkg/list"
)

// rootCmd is the main command for the list command line tool
// its just holds all the other useful commands
var Cmd = &Z.Cmd{
	Name:    "archive",
	Summary: "archives a users list",
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
		// confirm that user wantes to archvie the list
		archive := false
		prompt := &survey.Confirm{
			Message: fmt.Sprintf("Are you sure you want to archive %s?", args[0]),
		}
		survey.AskOne(prompt, &archive)

		if !archive {
			return nil
		}
		// start the editor with the list file
		directory := filepath.Dir(files[args[0]])
		base := filepath.Base(files[args[0]])

		// make a new archive directory if it does not exist
		archiveDir := filepath.Join(directory, "archive")
		if _, err := os.Stat(archiveDir); os.IsNotExist(err) {
			err := os.Mkdir(archiveDir, 0755)
			if err != nil {
				return err
			}
		}

		// move the list file to the archive directory
		newPath := filepath.Join(archiveDir, base)
		err = os.Rename(files[args[0]], newPath)
		if err != nil {
			return err
		}

		return nil
	},
}
