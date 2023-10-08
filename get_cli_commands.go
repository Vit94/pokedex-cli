package main

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Shows usage details",
			callback:    help,
		},
		"exit": {
			name:        "exit",
			description: "Stops the program",
			callback:    exit,
		},
	}
}
