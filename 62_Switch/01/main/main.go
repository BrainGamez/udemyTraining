package main

import "fmt"

func main() {
	var input string
	fmt.Println("Typ naam: ")
	fmt.Scan(&input)

	switch input {
	case "Bryan":
		fmt.Println("Hi, Bryan!")

	case "Piet":
		fmt.Println("Hi, Piet!")

	case "Klaas":
		fmt.Println("Hi, Klaas!")

	default:
		fmt.Println("Rare naam")
	}
}
