package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex-cli > ")
		scanner.Scan()
		text, err := prepareInput(scanner.Text())
		if err != nil {
			fmt.Println(err)
			continue
		}
		commands := getCliCommands()
		command, ok := commands[text[0]]
		if !ok {
			fmt.Printf("Список доступных команд: %s\n", getAvailableCommandNames(commands))
			continue
		}
		err = command.callback()
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
