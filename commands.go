package main

type CliCommand struct {
	name        string
	description string
	callback    func(...string) error
	flags       []string
}

func getCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"generate": {
			name:        "generate",
			description: "Generates a new password for a given service and username",
			callback:    commandGeneratePassword,
			flags:       []string{"-s", "-u", "-l", "-db"},
		},
		"get": {
			name:        "get",
			description: "Displays password of given service and username",
			callback:    commandGetPassword,
			flags:       []string{"-s", "-u", "-db"},
		},
		"help": {
			name:        "help",
			description: "displays a help message",
			callback:    commandHelp,
			flags:       []string{},
		},
	}
}

var flags map[string]map[string]string = map[string]map[string]string{
	"-s": {
		"description": "service name e.g; example.come",
		"required":    "yes",
	},
	"-u": {
		"description": "user name e.g; test@example.come",
		"required":    "yes",
	},
	"-l": {
		"description": "password length must be in the range of [8, 20] inclusive, defaults to 8 if not length is provided",
		"required":    "no",
	},
	"-db": {
		"description": "file path where the password is stored",
		"required":    "yes",
	},
}
