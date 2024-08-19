package content

import (
	"creditcard/data"
	"fmt"
	"math/rand"
	"os"
)

func Issue() {
	issuer, brand := "", ""
	// Looks for a prefix of the chosen issuer
	oki := false
	for key, val := range data.Issuers {
		if val == ChosenIssuer {
			issuer = key
			oki = true
			break
		}
	}
	if !oki {
		fmt.Println("\033[31m" + "No matching issuer found")
		os.Exit(1)
	}
	// Looks for a prefix of the chosen brand
	okb := false
	for key, val := range data.Brands {
		if val == ChosenBrand && key == issuer[:len(key)] {
			brand = key
			okb = true
			break
		}
	}
	if !okb {
		fmt.Println("\033[31m" + "No matching brand found")
		os.Exit(1)
	}
	// if chosen issuer and chosen brand do not have matching prefixes, exits with status 1
	if issuer[:len(brand)] != brand {
		fmt.Println("\033[31m" + "Issuer and brand are not compatible")
		os.Exit(1)
	}

	gen := 16 - len(issuer) - 4
	gened := issuer
	// generates all random numbers after prefix except 4 last ones
	for i := 0; i < gen; i++ {
		gened = gened + string('0'+rand.Intn(10))
	}
	gened = gened + "****"
	// Calls generate feature to generate and print a random valid number
	Generate([]string{gened}, true)
}
