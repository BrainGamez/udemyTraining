package main

import "fmt"

func main() {
	// wanneer we een fuctie aanroepen gebruiken we argumenten
	greet("Jane")
	greet("John")
}

// deze functie wordt gedeclareerd
// geen reciever (tussen func en greet)
// de naam (identifier) van de functie is greet
// de parameter is de variable name, welke van het type string is
// geen return (tussen haakje sluiten en de open accolade)
// er wordt een variable ontvangen en die wordt toegewezen aan name
func greet(name string) {
	// de variabele name kan binnen het blok worden gebruikt
	fmt.Println(name)
}

// argumenten gaan dus in de parameter van een functie
