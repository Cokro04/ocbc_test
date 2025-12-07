package main

import (
	"bufio"
	"fmt"
	"os"
)

func minYearsToSurpass(aliceInitial, bobBonus, aliceRate, bobRate float64) string {
	// Jika mustahil Bob menyalip
	if bobRate <= aliceRate && bobBonus <= aliceInitial {
		return "bob tidak akan pernah menyalip alice"
	}

	alice := aliceInitial
	bob := bobBonus
	years := 0

	// Batasi loop aman
	for years <= 10_000_000 {
		if bob > alice {
			return fmt.Sprintf("bob menyalip di tahun ke %d", years)
		}

		// Tahun berikutnya
		alice = alice * (1 + aliceRate)
		bob = bob * (1 + bobRate)
		years++
	}

	return "bob tidak akan pernah menyalip alice"
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var aliceInitial, bobBonus, aliceRate, bobRate float64

	// Input format:
	// aliceInitial bobBonus aliceRate bobRate
	fmt.Fscan(reader, &aliceInitial, &bobBonus, &aliceRate, &bobRate)

	result := minYearsToSurpass(aliceInitial, bobBonus, aliceRate, bobRate)
	fmt.Println(result)
}
