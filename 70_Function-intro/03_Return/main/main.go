package main

import "fmt"

func main() {
	// binnen Println wordt de greet functie aangeroepen met argumenten
	// in plaats van greet("Jane", "Doe") wordt de output van fmt.Sprint(fname, " ", lname) weergegeven
	fmt.Println(greet("Jane ", "Doe"))
}

// de parameters van greet worden gevuld met de argumenten uit main
// ze zijn van het type string
// waarom staat er string op de plaats van de return? -- is de return waarde ook een string?
func greet(fname, lname string) string {
	// wat is Sprint? -- is dit string print?
	return fmt.Sprint(fname, lname)
}
