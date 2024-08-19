package data

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Map to store data about Brands prefix to brand
var Brands = make(map[string]string)

func ReadBrands(path string) {
	readFile, err := os.Open(path)
	if err != nil {
		fmt.Println("\033[31m" + "Error reading a file")
		os.Exit(1)
	}

	// Reading and splitting a file by lines into a string slice
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
	// Regular expression to chekc alidity of the input in the file
	re := regexp.MustCompile(`^[A-Za-z\s]+:\d+$`)
	for _, line := range fileLines {
		if !re.MatchString(line) {
			fmt.Println("\033[31m" + "Invalid input in brands file")
			os.Exit(1)
		}
		brand := strings.Split(line, ":")
		Brands[brand[1]] = brand[0]
	}
}
