package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func commandHelp(args ...string) error {
	fmt.Println("usage: majP -c <command_name> [args]")

	fmt.Println()
	w := tabwriter.NewWriter(os.Stdout, 4, 4, 1, ' ', 0)
	fmt.Fprintln(w, "Command\t\tDescription")
	for _, cmd := range getCommands() {
		if cmd.name == "help" {
			continue
		}
		fmt.Fprintf(w, "%s\t\t%s\n", cmd.name, cmd.description)

	}
	w.Flush()

	for _, cmd := range getCommands() {
		if cmd.name == "help" {
			continue
		}
		fmt.Println()
		fmt.Printf("%s command flags:\n", cmd.name)
		w := tabwriter.NewWriter(os.Stdout, 4, 4, 1, ' ', 0)
		fmt.Fprintln(w, "Flag\t\tDescription\t\tRequired")
		for _, flag := range getCommands()[cmd.name].flags {
			fmt.Fprintf(w, "%s\t\t%s\t\t%s\n", flag, flags[flag]["description"], flags[flag]["required"])
		}
		w.Flush()
	}
	return nil
}
