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

		// Separrating input into command and arrguments
		fields := strings.Fields(strings.TrimRight(input, "\n"))
		command := fields[0]

		switch command {
		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Fprint(os.Stdout, "echo")
		default:
			fmt.Fprint(os.Stdout, input, ": command not found\n")
		}
	}
}
