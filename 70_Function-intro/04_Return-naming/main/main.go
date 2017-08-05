package main

import "fmt"

func main() {
	// binnen Println wordt de greet functie aangeroepen met argumenten
	// in plaats van greet("Jane", "Doe") wordt de output van fmt.Sprint(fname, " ", lname) weergegeven
	fmt.Println(greet("Jane ", "Doe"))
}

// de parameters van greet worden gevuld met de argumenten uit main
// ze zijn van het type string
// de return word een variabele s va het type string
func greet(fname, lname string) (s string) {
	// s wordt gelijk gesteld aan de output van Sprint
	s = fmt.Sprint(fname, lname)
	return
}
