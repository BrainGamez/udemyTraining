package main

import "fmt"

// deze functie wordt gedeclareerd
// geen reciever (tussen func en greet)
// de naam (identifier) van de functie is greet
// de parameter is de variable name, welke van het type string is
// geen return (tussen haakje sluiten en de open accolade)
// er wordt een variable ontvangen en die wordt toegewezen aan name

//Is de eerste functie de reciever?
// of is de return een anonieme functie van het type string?
func makeGreeter() func() string {
	// de return wordt de annonieme functie
	return func() string {
		// de return van de binnenste functie is "Hello Word!"
		return "Hello World!"
	}

}

func main() {
	greet := makeGreeter()
	fmt.Println(greet())
	fmt.Printf("%T\n", greet)
}
