package main

import "fmt"

// functie visit verwacht een variabele (numbers) van het type slice van het type int en een variabele (calback) van het type annonieme functie met een parameter van het type int
func visit(numbers []int, callback func(int)) {
	// range loopt voor elke waarde in de variabele numbers (1, 2, 3, 4) en geeft elke keer de waarde terug
	// in n zit de waarde van numbers per keer.
	// dan wordt callback elke keer aangeroepen met de waarde van n
	for _, n := range numbers {
		// hier wordt de Println uitgevoerd
		callback(n)
	}
}

// hier begint het programma
func main() {
	// de functie visit wordt aangeroepen met de argumenten (1)slice van het type int met waardes en (2)een annonieme functie met de variable n van het type int.
	// LET OP!! de Println zit in het 2e argument
	visit([]int{1, 2, 3, 4}, func(n int) {
		// de verschillende waardes worden 'om de beurt' (doordat range loopt) naar het scherm geprint.
		fmt.Println(n)
	})
}
