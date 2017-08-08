package main

// de import van fmt is in de file scope
import "fmt"

// de scope van x is de package scope en is dus beschipbaar binnen de hele package
var x int = 42

// de functie main is ook in de package scope
func main() {
	fmt.Println(x)
	foo()
	y := 17
	fmt.Println(y)
}

func foo() {
	fmt.Println(x)
	// fmt.Println(y)
}
