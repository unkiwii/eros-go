package version

import (
	"fmt"

	"github.com/unkiwii/eros-go/cmd/base"
)

var CmdVersion = &base.Command{
	Run:       runVersion,
	Name:      "version",
	UsageLine: "",
	Short:     "print eros version",
	Long:      `Version prints the eros version.`,
}

func runVersion(cmd *base.Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
	}

	//TODO: get version, os and arch
	VERSION := "TODO"
	OS := "TODO"
	ARCH := "TODO"

	fmt.Printf("eros version %s %s/%s\n", VERSION, OS, ARCH)
}
