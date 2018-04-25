// Package base defines shared basic pieces of the eros commands,
// in particular logging and the Command structure.
package cmd

import (
	"flag"
	"fmt"
	"os"
)

// A Command is an implementation of a eros command
// like 'eros build' or 'eros version'.
type Command struct {
	// Run runs the command.
	// The args are the arguments after the command name.
	Run func(cmd *Command, args []string)

	// Name is the name of the comand
	Name string

	// UsageLine is the one-line usage message.
	UsageLine string

	// Short is the short description shown in the 'eros help' output.
	Short string

	// Long is the long message shown in the 'eros help <this-command>' output.
	Long string

	// Flag is a set of flags specific to this command.
	Flag flag.FlagSet
}

// Commands lists the available commands and help topics.
// The order here is the order in which they are printed by 'eros help'.
var Commands []*Command

func (c *Command) Usage() {
	fmt.Fprintf(os.Stderr, "usage: %s %s\n", c.Name, c.UsageLine)
	fmt.Fprintf(os.Stderr, "Run 'eros help %s' for details.\n", c.Name)
	os.Exit(2)
}

func Execute(args []string) {
	if len(args) < 1 {
		Repl()
		return
	}

	cmdName := args[0]
	for _, cmd := range Commands {
		if cmd.Name == cmdName {
			cmd.Flag.Usage = func() { cmd.Usage() }
			cmd.Flag.Parse(args[1:])
			args = cmd.Flag.Args()
			cmd.Run(cmd, args)
			return
		}
	}

	UnknownCommand(cmdName)
}
