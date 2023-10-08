package main

import "fmt"

func help() error {
	fmt.Println("Welcome to Pokedex!\nAvailable commands:")
	for _, command := range getCliCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
