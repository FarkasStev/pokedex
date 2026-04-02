package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/farkasstev/pokedex/internal/pokeapi"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommandRegistry()
	configuration := config{
		Client: pokeapi.NewClient(5 * time.Second),
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		command, exists := commands[input[0]]
		if !exists {
			fmt.Println("Unknown command")
		} else {
			if err := command.callback(&configuration); err != nil {
				fmt.Printf("Encountered Error: %v\n", err)
			}
		}

	}

}
