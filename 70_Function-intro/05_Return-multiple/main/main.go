package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe "))
}

// de return parameters zijn gescheiden door een komma
func greet(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
