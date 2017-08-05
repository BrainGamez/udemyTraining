package main

import "fmt"

// binnen functie main wordt de functie average aangeroepen en de uit komst in n gezet
// daarna wordt n naar het scherm geprint
func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

// de functie average heeft de parameter sf als insert welke van het type variadic float64 is
// dit betekend dat we zoveel float64 argumenten kunnen invoeren als we willen.
// het lijkt erop alsof we een array vullen
// de return is van het type float64
func average(sf ...float64) float64 {
	// de variabele total wordt geinitialiseerd op 0.0
	total := 0.0
	// dan draait er een loop
	// die begint met een ongedefinieerde lege waarde
	// dit stuk snap ik niet... (Worden alle argumenten bij elkaar opgeteld?)
	for _, v := range sf {
		total += v
	}
	// de functie avarage geeft als return het totaal gedeeld door de lengte van sf (het totaal aanal argumenten)
	return total / float64(len(sf))
}
