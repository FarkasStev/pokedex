package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommandRegistry()
	configuration := config{}
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
