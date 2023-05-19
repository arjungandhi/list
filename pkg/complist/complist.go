package complist

import (
	"github.com/rwxrob/bonzai"
	"github.com/rwxrob/fn/filt"

	"github.com/arjungandhi/list/pkg/list"
)

func New() *comp { return new(comp) }

type comp struct{}

func (c *comp) Complete(x bonzai.Command, args ...string) []string {
	// not sure we've completed the command name itself yet
	if len(args) == 0 {
		return []string{x.GetName()}
	}

	// get our list of commands from the List comamdn
	lists, err := list.List()
	if err != nil {
		lists = map[string]string{}
	}

	// build list of visible commands and params
	options := []string{}
	options = append(options, x.GetCommandNames()...)
	for name, _ := range lists {
		options = append(options, name)
	}

	if len(args) == 0 {
		return options
	}

	return filt.HasPrefix(options, args[0])
}
