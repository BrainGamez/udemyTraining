package main

import "fmt"

func main() {
	var h int
	var l int
	fmt.Print("Kies een groot getal: ")
	fmt.Scan(&h)
	fmt.Print("Kies een klein getal: ")
	fmt.Scan(&l)

	fmt.Println(h, "delen door", l, "is", h/l, ",rest", h%l)
}
