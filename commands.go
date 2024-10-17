package main

type CliCommand struct {
	name        string
	description string
	callback    func(...string) error
}

func getCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"generate": {
			name:        "generate",
			description: "generates a new password for a given service",
			callback:    generatePassword,
		},
	}
}
