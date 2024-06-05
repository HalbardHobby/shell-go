package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
		case "type":
			switch args[0] {
			case "exit", "echo", "type":
				fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", args[0])
			default:
				path, found := lookupExecutable(args[0])
				if found {
					fmt.Fprintf(os.Stdout, "%s is %s\n", args[0], path+"/"+args[0])
				} else {
					fmt.Fprintf(os.Stdout, "%s: not found\n", args[0])
				}
			}
		default:
			err := executeProgram(command, args...)
			if err != nil {
				fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
			}
		}
	}
}

func executeProgram(command string, args ...string) (err error) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	return
}

func lookupExecutable(cmd string) (string, bool) {
	path_val := os.Getenv("PATH")

	// Iterate through all provided paths
	for _, path := range strings.Split(path_val, ":") {

		// Iterate through all entries in given path
		entries, _ := os.ReadDir(path)
		for _, entry := range entries {

			// check name and if it is a directory
			if entry.Name() == cmd && !entry.IsDir() {
				return path, true
			}
		}
	}

	// return empty path and false by default
	return "", false
}
