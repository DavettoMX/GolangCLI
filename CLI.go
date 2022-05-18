package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Golang Terminal")
	fmt.Println("------------------")
	fmt.Print("$ ")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		userInput := scanner.Text()

		if strings.Contains(userInput, "exit") {
			os.Exit(0)
		}

		// Say Hello
		if strings.Contains(userInput, "salute") {
			fmt.Println("Hello! I am your cloud CLI")
		}

		if strings.Contains(userInput, "clear") {
			// Clear the screen
			fmt.Print("\033[H\033[2J")
		}

		// List of files
		if strings.Contains(userInput, "ls") {
			dir, err := os.Getwd() // Get the current working directory

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(dir)
		}

		// Change directory
		if strings.Contains(userInput, "cd") {
			dir := strings.Split(userInput, " ")
			err := os.Chdir(dir[1])

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(dir[1])
		}

		// Help Command
		if strings.Contains(userInput, "help") {
			fmt.Println("Available commands:")
			fmt.Println("\tls - List files in current directory")
			fmt.Println("\tcd - Change directory")
			fmt.Println("\texit - Exit the terminal")
			fmt.Println("\tsalute - Say hello")
			fmt.Println("\tclear - Clear the screen")
		}

		fmt.Print("$ ")
	}
}

func execPINGCommand(cmdParam string) string {
	return "Error! Command not found"
}
