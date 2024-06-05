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

		fmt.Fprint(os.Stdout, strings.TrimRight(input, "\n"), ": command not found\n")
	}
}
