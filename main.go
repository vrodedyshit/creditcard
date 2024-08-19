package main

import (
	"creditcard/content"
	"fmt"
	"os"
)

func main() {
	//  Reads command line arguments
	args := os.Args
	if len(args[1:]) != 0 {
		// Starts the main interface of the program
		if !content.ParseFlags([]string(args[1:])) {
			os.Exit(1)
		}
	} else {
		// Exits with status 1 if no flags were passed
		fmt.Println("\033[31m" + "No flags were chosen")
		os.Exit(1)
	}
}
