package content

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"regexp"
)

func Generate(args []string, pick bool) {
	res := []string{}
	// Regular expression to check if arguments are a valid input
	re := regexp.MustCompile(`^\d{9,15}\*{1,4}$`)
	for _, str := range args {
		if !re.MatchString(str) {
			fmt.Println("\033[31m" + "Invalid input for a generate feature")
			os.Exit(1)
		}
		ast := 0
		l := len(str)
		for i := l - 1; i >= l-5; i-- {
			if str[i] == '*' {
				ast++
			}
		}
		// Parsing through all possible numbers in ascending oreder and checks if they are valid
		for i := 1; i < int(math.Pow10(ast)); i++ {
			gened := fmt.Sprintf("%0*d", ast, i)
			if IsValid(str[:l-ast] + gened) {
				res = append(res, str[:l-ast]+gened)
			}
		}
	}
	// Prints according to a flag
	if pick {
		fmt.Println(res[rand.Intn(len(res))])
	} else {
		for _, card := range res {
			fmt.Println(card)
		}
	}
}
