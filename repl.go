package main

import (
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	splitArgs := strings.Split(strings.ToLower(text), " ")
	cleanedInput := make([]string, 0)
	for _, arg := range splitArgs {
		trimmedArg := strings.TrimSpace(arg)
		if len(trimmedArg) > 0 {
			cleanedInput = append(cleanedInput, trimmedArg)
		}
	}
	return cleanedInput
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommandRegistry() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print(`Welcome to the Pokedex!
Usage:

`)
	registry := getCommandRegistry()
	for _, value := range registry {
		fmt.Printf("%v: %v\n", value.name, value.description)
	}
	return nil
}
