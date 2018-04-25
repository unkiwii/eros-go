package cmd

import (
	"fmt"
)

var Build = &Command{
	Run:       runBuildCommand,
	Name:      "build",
	UsageLine: "",
	Short:     "TODO: build.Short",
	Long:      "TODO: build.Long",
}

func runBuildCommand(cmd *Command, args []string) {
	fmt.Printf("build %v\n", args)
}
