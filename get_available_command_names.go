package main

import "strings"

func getAvailableCommandNames(commands map[string]cliCommand) string {
	var availableCommandNames []string
	for _, availableCommand := range commands {
		availableCommandNames = append(availableCommandNames, availableCommand.name)
	}
	return strings.Join(availableCommandNames, ", ")
}
