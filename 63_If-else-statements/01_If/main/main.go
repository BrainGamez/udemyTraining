package main

import "fmt"

func main() {
	// Wanneer 'waar' wordt de prinln uitgevoerd (standaard is true)
	if !true {
		fmt.Println("This ran")
	}
	// Wanneer 'vals' wordt de prinln uitgevoerd
	if !false {
		fmt.Println("This did not run")
	}
}
