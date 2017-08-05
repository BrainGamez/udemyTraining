package main

import "fmt"

func main() {
	var myName string
	fmt.Print("Typ naam hier: ")
	fmt.Scan(&myName)
	fmt.Println("Hello", myName)
}
