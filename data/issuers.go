package data

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Map to store data about Issuers prefix to issuer
var Issuers = make(map[string]string)

func ReadIssuers(path string) {
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
	// Regular expression to check input format from the file
	re := regexp.MustCompile(`^[A-Za-z\s]+:\d+$`)
	for _, line := range fileLines {
		if !re.MatchString(line) {
			fmt.Println("\033[31m" + "Invalid input in issuers file")
			os.Exit(1)
		}
		issuer := strings.Split(line, ":")
		Issuers[issuer[1]] = issuer[0]
	}
}
