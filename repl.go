package main

import (
	"fmt"
	"os"
	"strings"

	pokeapi "github.com/farkasstev/pokedex/internal/pokeapi"
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
	callback    func(*config) error
}

type config struct {
	Next     string
	Previous string
	Client   pokeapi.Client
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
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in Pokemon world. Each subsequent call will display the next 20",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of 20 location areas in Pokemon world. Each subsequent call will display the previous 20",
			callback:    commandMapb,
		},
	}
}

func commandExit(configuration *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(configuration *config) error {
	fmt.Print(`Welcome to the Pokedex!
Usage:

`)
	registry := getCommandRegistry()
	for _, value := range registry {
		fmt.Printf("%v: %v\n", value.name, value.description)
	}
	return nil
}

func commandMap(configuration *config) error {

	url := "https://pokeapi.co/api/v2/location-area/"

	if configuration.Next != "" {
		url = configuration.Next
	}
	return executeMap(url, configuration)
}

func commandMapb(configuration *config) error {

	url := "https://pokeapi.co/api/v2/location-area/"

	if configuration.Previous != "" {
		url = configuration.Previous
	}

	return executeMap(url, configuration)
}

func executeMap(url string, configuration *config) error {
	res, err := configuration.Client.GetLocationArea(url)
	if err != nil {
		return err
	}
	configuration.Next = res.Next
	configuration.Previous = res.Previous

	results := res.Results

	for i := range results {
		fmt.Println(results[i].Name)
	}

	return nil
}
