package new

import (
	"errors"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"

	"github.com/arjungandhi/list/pkg/list"
)

// rootCmd is the main command for the list command line tool
// its just holds all the other useful commands
var Cmd = &Z.Cmd{
	Name:    "new",
	Summary: "creates a new list",
	Commands: []*Z.Cmd{
		help.Cmd,
	},
	Call: func(_ *Z.Cmd, args ...string) error {
		/// get list diretory
		listDir := os.Getenv("LISTDIR")
		if listDir == "" {
			return errors.New("LISTDIR not set")
		}

		// get all files
		files, err := list.List()
		if err != nil {
			return err
		}

		// prompt for the name of the list
		listName := ""
		prompt := &survey.Input{
			Message: "Name of the list:",
		}
		survey.AskOne(prompt, &listName)

		// check if the list already exists
		for file, _ := range files {
			if file == listName {
				// if it does, error out
				return fmt.Errorf("list %s already exists", listName)
			}
		}

		// create the list
		path := fmt.Sprintf("%s/%s", listDir, listName+".md")

		// create the file
		_, err = os.Create(path)

		if err != nil {
			return err
		}

		return nil
	},
}
