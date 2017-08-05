package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		// deel i door 1 en wanneer rest 0 is hebben we een even getal te pakken
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
