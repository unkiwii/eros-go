package main

import (
	"flag"
	"os"

	"github.com/unkiwii/eros-go/cmd"
)

func init() {
	cmd.Commands = []*cmd.Command{
		cmd.Version,
		cmd.Help,
		cmd.Build,
	}
}

func main() {
	flag.Usage = func() {
		cmd.PrintHelp()
		os.Exit(2)
	}
	flag.Parse()
	cmd.Execute(flag.Args())
}
