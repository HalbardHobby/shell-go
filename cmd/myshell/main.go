package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Naive implementation for REPL cycle
	for {
		// Uncomment this block to pass the first stage
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimRight(input, "\n")

		switch input {
		case "exit 0":
			os.Exit(0)
		default:
			fmt.Fprint(os.Stdout, input, ": command not found\n")
		}
	}
}
