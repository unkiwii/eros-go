package help

import (
	"fmt"
	"os"
	"text/template"

	"github.com/unkiwii/eros-go/cmd/base"
)

var usageTemplate = `eros is a tool for managing eros source code.

Usage:

	eros [command] [arguments]

The commands are:
{{range .}}
	{{.Name | printf "%-12s"}} {{.Short}}{{end}}

If no command is specified an interactive session is started.

Use "eros help [command]" for more information about a command.
`

var helpTemplate = `usage: eros {{.Name}} {{.UsageLine}}

{{.Long}}
`

func Help(args []string) {
	if len(args) == 0 {
		PrintUsage()
		return
	}

	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "usage: eros help <command>\n\nToo many arguments given.\n")
		os.Exit(2)
	}

	arg := args[0]

	for _, cmd := range base.Commands {
		if cmd.Name == arg {
			printTemplate(helpTemplate, cmd)
			return
		}
	}

	UnknownCommand(arg)
}

func PrintUsage() {
	printTemplate(usageTemplate, base.Commands)
}

func UnknownCommand(cmd string) {
	fmt.Fprintf(os.Stderr, "Unknown command %#q. Run 'eros help' for usage.\n", cmd)
	os.Exit(2)
}

func printTemplate(text string, data interface{}) {
	t := template.New("top")
	template.Must(t.Parse(text))
	err := t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
