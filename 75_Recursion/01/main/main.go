package main

import "fmt"

// de functie factorial verwacht een variabele x van het type int
// de functie geeft een int terug
func factorial(x int) int {

	//wanneer x gelijk is aan 0
	if x == 0 {
		//geef dan een 1 terug
		return 1
	}
	// geef in alle andere gevallen de uit komst van x maal de functie factorial
	// waarbij factorial dit keer aangeroepen wordt met x - 1 (en dus dezelfde waarde)
	// in andere woorden ik snap er niks van!!!
	// 4 * 3 * 2 * 1 * 1 = 24
	return x * factorial(x-1)
}

// Hier begint het programma
func main() {
	// print de output van factorial naar het scherm wanneer deze aangeroepen word met 4 als argument
	// de output zal van het type int zijn
	fmt.Println(factorial(4))
}
