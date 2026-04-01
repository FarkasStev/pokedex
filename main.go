package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommandRegistry()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		command, exists := commands[input[0]]
		if !exists {
			fmt.Println("Unknown command")
		} else {
			command.callback()
		}

	}

}
