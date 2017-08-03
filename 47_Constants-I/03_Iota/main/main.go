package main

import "fmt"

//iota
const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
)

func main() {

	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\t\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\t\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\t\n", GB)
}
