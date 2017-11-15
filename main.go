package main

import (
	"flag"
	"os"

	"github.com/unkiwii/eros-go/cmd/base"
	"github.com/unkiwii/eros-go/cmd/build"
	"github.com/unkiwii/eros-go/cmd/help"
	"github.com/unkiwii/eros-go/cmd/repl"
	"github.com/unkiwii/eros-go/cmd/version"
)

func init() {
	base.Commands = []*base.Command{
		version.CmdVersion,
		build.CmdBuild,
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		repl.Repl()
		return
	}

	arg := args[0]

	if arg == "help" {
		help.Help(args[1:])
		return
	}

	for _, cmd := range base.Commands {
		if cmd.Name == arg {
			cmd.Flag.Usage = func() { cmd.Usage() }
			cmd.Flag.Parse(args[1:])
			args = cmd.Flag.Args()
			cmd.Run(cmd, args)
			return
		}
	}

	help.UnknownCommand(arg)
}

func usage() {
	help.PrintUsage()
	os.Exit(2)
}
