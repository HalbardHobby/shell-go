package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Naive implementation for REPL cycle
	for {
		// Uncomment this block to pass the first stage
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		// Separrating input into command and arguments
		fields := strings.Fields(strings.TrimRight(input, "\n"))
		command := fields[0]
		args := fields[1:]

		switch command {
		case "exit":
			code, _ := strconv.Atoi(args[0])
			os.Exit(code)
		case "echo":
			fmt.Fprint(os.Stdout, strings.Join(args, " "), "\n")
		default:
			fmt.Fprint(os.Stdout, command, ": command not found\n")
		}
	}
}
