package cmd

import (
	"fmt"
	"os"
	"text/template"
)

const helpTemplate = `eros is a tool for managing eros source code.

Usage:

	eros [command] [arguments]

The commands are:
{{range .}}
	{{.Name | printf "%-12s"}} {{.Short}}{{end}}

If no command is specified an interactive session is started.

Use "eros help [command]" for more information about a command.
`

const usageTemplate = `usage: eros {{.Name}} {{.UsageLine}}

{{.Long}}
`

var Help = &Command{
	Run:       runHelpCommand,
	Name:      "help",
	UsageLine: "<command>",
	Short:     "shows help for a given command",
	Long:      `shows help for a given command`,
}

func runHelpCommand(helpCommand *Command, args []string) {
	if len(args) == 0 {
		PrintHelp()
		return
	}

	if len(args) != 1 {
		printTemplate(usageTemplate, helpCommand)
		os.Exit(2)
	}

	arg := args[0]

	for _, cmd := range Commands {
		if cmd.Name == arg {
			printTemplate(usageTemplate, cmd)
			return
		}
	}

	UnknownCommand(arg)
}

func PrintHelp() {
	printTemplate(helpTemplate, Commands)
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
