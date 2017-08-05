package main

import "fmt"

func main() {
	//slice
	data := []float64{43, 56, 87, 12, 45, 57}
	// op deze manier kan de slice in de variabele data een voor een worden overgedragen aan de parameter in de functie avarage
	n := average(data...)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
