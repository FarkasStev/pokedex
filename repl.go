package main

import (
	"fmt"
	"math/rand"
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
	callback    func(*config, []string) error
}

type config struct {
	Next     string
	Previous string
	Client   pokeapi.Client
	Pokedex  map[string]pokeapi.PokemonInfo
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
		"explore": {
			name:        "explore",
			description: "Displays a list of all the Pokémon located in an area. Takes the name of a location area as an argument.",
			callback:    explore,
		},
		"catch": {
			name:        "catch",
			description: "It takes the name of a Pokemon as an argument and adds them to the users Pokedex",
			callback:    catch,
		},
	}
}

func commandExit(configuration *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(configuration *config, args []string) error {
	fmt.Print(`Welcome to the Pokedex!
Usage:

`)
	registry := getCommandRegistry()
	for _, value := range registry {
		fmt.Printf("%v: %v\n", value.name, value.description)
	}
	return nil
}

func commandMap(configuration *config, args []string) error {

	url := "https://pokeapi.co/api/v2/location-area/"

	if configuration.Next != "" {
		url = configuration.Next
	}
	return executeMap(url, configuration)
}

func commandMapb(configuration *config, args []string) error {

	url := "https://pokeapi.co/api/v2/location-area/"

	if configuration.Previous != "" {
		url = configuration.Previous
	}

	return executeMap(url, configuration)
}

func executeMap(url string, configuration *config) error {
	res, err := configuration.Client.GetLocationAreas(url)
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

func explore(configuration *config, args []string) error {
	res, err := configuration.Client.GetLocationArea(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", args[0])
	for i := range res.PokemonEncounters {
		fmt.Printf(" - %v\n", res.PokemonEncounters[i].Pokemon.Name)
	}

	return nil
}

func catch(configuration *config, args []string) error {
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %v...\n", name)
	res, err := configuration.Client.GetPokemonInfo(name)
	if err != nil {
		return err
	}
	catchNum := (100.0 / float64(res.BaseExperience)) * rand.Float64()
	caught := catchNum > .50
	if caught {
		fmt.Printf("%v was caught!\n", name)
		//add to pokedex
		configuration.Pokedex[name] = res
	} else {
		fmt.Printf("%v escaped!\n", name)
	}
	return nil
}
