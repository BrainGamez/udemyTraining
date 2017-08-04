package main

import "fmt"

func main() {
	i := 0

	for {
		i++
		// de volgende twee loops staan naast elkaar
		// Wanneer rest 0 is dan gaat het programma door naar de volgende loop
		// Wanneer rest 1 is dan gaat het programma een bovenliggend niveau en begint dan opnieuw
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		// Zodra i groter of gelijk is aan 50 wordt er een break uitgevoerd (naar het eind van het programma)
		if i >= 50 {
			break
		}
	}
}
