package content

import (
	"bufio"
	"creditcard/data"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	BrandsPath   string
	IssuersPath  string
	ChosenBrand  string
	ChosenIssuer string
)

func ParseFlags(args []string) bool {
	// If args length is smaller than 2 thene there is no feature that can be supported
	if len(args) < 2 {
		fmt.Println("\033[31m" + "Not enough arguments")
		return false
	}
	// Flags declaration
	feature := args[0]
	stdin := flag.Bool("stdin", false, "Support standard input")
	pick := flag.Bool("pick", false, "Support random pick")
	flag.StringVar(&BrandsPath, "brands", "", "Filepath to brands.txt")
	flag.StringVar(&IssuersPath, "issuers", "", "Filepath to issuers.txt")
	flag.StringVar(&ChosenBrand, "brand", "", "Chosen brand")
	flag.StringVar(&ChosenIssuer, "issuer", "", "Chosen issuers")
	// Custom flag parsing to aboid errors caused by feature flags
	flag.CommandLine.Parse(args[1:])

	var cardNumbers []string
	if BrandsPath != "" {
		args = args[1:]
		data.ReadBrands(BrandsPath)
	}
	if IssuersPath != "" {
		args = args[1:]
		data.ReadIssuers(IssuersPath)
	}
	nonFlagArgs := flag.Args()
	// Standard input handling
	if *stdin {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			cardNumbers = append(cardNumbers, strings.Fields(line)...)
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
			return false
		}
	} else {
		cardNumbers = nonFlagArgs
	}
	// Start of an interface
	switch feature {
	case "validate":
		Validate(cardNumbers)
	case "generate":
		Generate(cardNumbers, *pick)
	case "information":
		Info(cardNumbers)
	case "issue":
		Issue()
	default:
		fmt.Println("\033[31m" + "Flag is not supported")
		return false
	}
	return true
}
