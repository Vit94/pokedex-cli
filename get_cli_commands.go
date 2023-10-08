package main

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var conf = config{}

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
		"map": {
			name:        "map",
			description: "shows next 20 pokemon location areas",
			callback: func() error {
				Map(&conf)
				return nil
			},
		},
		"mapb": {
			name:        "mapb",
			description: "shows previous 20 pokemon location areas",
			callback: func() error {
				Mapb(&conf)
				return nil
			},
		},
	}
}
