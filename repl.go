package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	splitArgs := strings.Split(strings.ToLower(text), " ")
	cleanedInput := make([]string, 0)
	for _, arg := range splitArgs {
		fmt.Println("arg:" + arg)
		trimmedArg := strings.TrimSpace(arg)
		if len(trimmedArg) > 0 {
			fmt.Printf("trimmedArg:%v len(trimmedArg): %v\n", trimmedArg, len(trimmedArg))
			cleanedInput = append(cleanedInput, trimmedArg)
		}
	}
	fmt.Printf("cleanedInput:%v len(cleanedInput): %v\n", cleanedInput, len(cleanedInput))
	return cleanedInput
}
