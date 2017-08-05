package main

import "fmt"

func main() {
	// we kunnen meerdere argumenten gebruiken
	greet("Jane", "Doe")
}

// deze worden door de parameters opgevangen
// func greet(fname string, lname string) {
// dit is een kortere notatie, waarbij beide variabelen van hetzelfde type zijn
func greet(fname, lname string) {
	fmt.Println(fname, lname)
}
