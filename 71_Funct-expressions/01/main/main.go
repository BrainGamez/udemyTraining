package main

import "fmt"

// func greeting() {
// 	fmt.Println("Hello World!")
// }

func main() {
	//anonieme functie (een functie toegewezen aan een variabele)
	greeting := func() {
		fmt.Println("Hello World!")
	}

	//de variabele aangeroepen als functie?
	greeting()

	fmt.Printf("%T\n", greeting)
}
