package main

import "fmt"

func main() {
	// b wordt geinitialiseerd met 'waar'
	b := true

	// als food geinitialiseerd met Chocolade waar (b) is dan wordt food naar het scherm geprint
	if food := "Choclate"; b {
		fmt.Println(food)
	}
}
