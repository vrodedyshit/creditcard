package content

import (
	"fmt"
	"regexp"
	"strconv"
)

// Applies Luhn's Algorithms to a given string of numbers
func LuhnAlgo(str string) int {
	sum := 0
	for i, s := range str {
		a, _ := strconv.Atoi(string(s))
		if i%2 == 0 {
			a *= 2
			if a > 9 {
				a = a/10 + a%10
			}
		}
		sum += a
	}
	return sum
}

// Validation check
func IsValid(str string) bool {
	// regular expression to check if the argument is only numbers
	re := regexp.MustCompile(`^[0-9]+$`)
	if len(str) < 13 || !re.MatchString(str) {
		return false
	}
	sum := LuhnAlgo(str)
	if sum%10 == 0 {
		return true
	}
	return false
}

// Validate feature's print
func Validate(nums []string) {
	for _, str := range nums {
		if IsValid(str) {
			fmt.Println("OK")
		} else {
			fmt.Println("INCORRECT")
		}
	}
}
