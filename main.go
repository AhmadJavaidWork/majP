package main

import (
	"errors"
	"log"
	"os"
)

const (
	DefaultPassLen = 8
	DefaultDBPath  = "data"
)

func main() {
	cmd, err := parseCmd(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.callback(os.Args[2:]...)
	if err != nil {
		log.Fatal("Error: ", err)
	}
}

func parseCmd(args []string) (CliCommand, error) {
	cmdName := ""
	for i, arg := range args {
		if arg == "-c" && i+1 < len(args) {
			cmdName = args[i+1]
		}
	}

	if cmdName == "" {
		return CliCommand{}, errors.New("please provide a command after -c flag")
	}

	cmd, ok := getCommands()[cmdName]

	if !ok {
		return CliCommand{}, errors.New("command not found")
	}
	return cmd, nil
}
