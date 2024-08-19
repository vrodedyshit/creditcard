package content

import (
	"creditcard/data"
	"fmt"
	"os"
	"regexp"
)

func Info(args []string) {
	// Regular expression to check if arguments are only numbers
	re := regexp.MustCompile(`^[0-9]+$`)
	for _, str := range args {
		if !re.MatchString(str) {
			fmt.Println("\033[31m" + "Invalid Argument")
			os.Exit(1)
		}
		fmt.Println(str)
		fmt.Print("Correct: ")
		// if a valid card number brand and issuer are found
		if IsValid(str) {
			fmt.Print("yes\n")
			fmt.Print("Card Brand: ")
			found := false
			for key, val := range data.Brands {
				if key == str[:len(key)] {
					fmt.Print(val)
					fmt.Println()
					found = true
					break
				}
			}
			if !found {
				fmt.Print("-\n")
			}
			fmt.Print("Card Issuer: ")
			found = false
			for key, val := range data.Issuers {
				if key == str[:len(key)] {
					fmt.Print(val)
					fmt.Println()
					found = true
					break
				}
			}
			if !found {
				fmt.Print("-\n")
			}
		} else {
			fmt.Print("no\n")
			fmt.Print("Card Brand: -\n")
			fmt.Print("Card Issuer: -\n")
		}
	}
}
