package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
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
		case "pwd":
			getPwd()
		case "cd":
			cd(args[0])
		case "type":
			typeCommand(args[0])
		default:
			err := executeProgram(command, args...)
			if err != nil {
				fmt.Fprintf(os.Stdout, "%s: command not found\n", command)
			}
		}
	}
}

func getPwd() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error printing current directtory: %s\n", err)
	} else {
		fmt.Fprintf(os.Stdout, "%s\n", pwd)
	}
}

func cd(target string) {
	if target == "~" {
		target = os.Getenv("HOME")
	}
	target = path.Clean(target)
	if !path.IsAbs(target) {
		pwd, _ := os.Getwd()
		target = filepath.Join(pwd, target)
	}
	err := os.Chdir(target)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: No such file or directory\n", target)
	}
}

func typeCommand(command string) {
	switch command {
	case "exit", "echo", "pwd", "cd", "type":
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", command)
	default:
		path, found := lookupExecutable(command)
		if found {
			fmt.Fprintf(os.Stdout, "%s is %s\n", command, path+"/"+command)
		} else {
			fmt.Fprintf(os.Stdout, "%s: not found\n", command)
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
