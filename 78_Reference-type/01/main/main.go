package main

import "fmt"

// hier begint het programma
func main() {
	// de variabele m wordt geinitialiseerd met de functie make
	// de argument van make is een map van het type int met daar een string in
	m := make(map[string]int)
	// de functie changeme wordt aangeroepen met de functie in m als een argument
	changeMe(m)
	// de waarde van de string m wordt weergegeven (42)
	fmt.Println(m["Bryan"])
}

// de functie changeme verwacht map van het type int met daar een string in
// deze waarde wordt in de variabele z gestopt
func changeMe(z map[string]int) {
	// de waarde van de string Bryan in de map z wordt ingevuld met de int 42
	z["Bryan"] = 42
}
