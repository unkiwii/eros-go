package build

import (
	"fmt"

	"github.com/unkiwii/eros-go/cmd/base"
)

var CmdBuild = &base.Command{
	Run:       runBuild,
	Name:      "build",
	UsageLine: "",
	Short:     "TODO: build.Short",
	Long:      "TODO: build.Long",
}

func runBuild(cmd *base.Command, args []string) {
	fmt.Printf("build %v\n", args)
}
