package main

import "fmt"

// hier begint het programma
func main() {
	// de variabele m wordt geinitialiseerd met de functie make
	// de functie make verwacht een slice of string. wat betekenen 1 en 25?
	m := make([]string, 2, 25)
	// m wordt naar het scherm geprint (leeg?)
	// leeg omdat er nog niks ingezet is. Dat komt pas waneer change me wordt aangeroepen
	fmt.Println(m)
	// changeMe wordt aangeroepen
	changeMe(m)
	// mijn naam wordt naar het scherm geprint
	fmt.Println(m)
}

// de functie changeMe verwacht een slice of string. gaat in de variabele z
func changeMe(z []string) {
	// mijn naam wordt naar de eerste plek in de slice binnen z geschreven
	z[0] = "Bryan"
	z[1] = "Francisca"
	// mijn naam wordt naar het scherm geprint
	fmt.Println(z)
}

// leeg?
// Bryan
// Bryan
