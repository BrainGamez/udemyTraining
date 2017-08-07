package main

import "fmt"

func hello() {
	fmt.Print("Hello ")
}

func world() {
	fmt.Println("World")
}

func main() {
	// functies met defer ervoor wordt pas vlak voor de exit van een functie uitegvoerd. Als laatste dus.
	defer world()
	hello()
}
