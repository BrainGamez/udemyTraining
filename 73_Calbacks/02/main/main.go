package main

import "fmt"

// functie filter heeft 2 paremeters en geeft een slice int terug
// de eerste paremeter is een variabele numbers van het type slice int
// de tweede paraneter is een functie genaamd callback met als parameter een int en geeft een boolean terug (true of false)
func filter(numbers []int, callback func(int) bool) []int {
	// binnen de scope van filter wordt xs geinitialiseerd met een lege slice int
	xs := []int{}
	// loop: voor elke item in de variabele numbers
	for _, n := range numbers {
		// als callback waar is (functie in het tweede argument van de aanroep uit main)
		if callback(n) {
			// xs krijgt de waarde in xs en n aan elkaar vast? wat doet append?//
			xs = append(xs, n)
		}
	}
	return xs
}

// hier begint het programma
func main() {
	// binnen de scope van main wordt de variabele xs geinitialiseerd met het reultaat van de aanroep van functie filter met 2 argumenten:
	// het eerste argument is een slice met de waardes 1 - 4
	// het tweede argument is een functie met de variabele n van het type int als parameter. het geeft een boolean terug (true of false)
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})
	// alle waardes van de lokale xs worden in een keer naar het scherm geprint
	fmt.Println(xs)
}
