package main

import (
	"fmt"
)

// z krijgt de waarde van het geheugen adress van x in main
// via *z worde de waarde op dat geheuegen adress (van x) overschreven met 0
func zero(z *int) {
	*z = 0
}

// de waarde van x wordt op 5 geinitialiseerd
// het geheugen adress van x wordt in de func zero gezet
// wanneer de waarde van x nu geprint wordt naar het scherm, zien we de overschreven waarde
func main() {
	x := 5
	zero(&x)
	fmt.Println(x)

}

//
