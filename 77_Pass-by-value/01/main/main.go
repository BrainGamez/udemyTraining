package main

import "fmt"

//hier begint het programma
// functie main wordt aangeroepen
func main() {
	// de variabele age wordt geinitialiseerd op 44
	age := 44
	// het geheugenadres van age wordt naar het scherm geprint
	fmt.Println(&age)

	// de functie change me wordt aangeroepen met het geheugenadres van age als argument
	changeMe(&age)

	// het geheugenadres van age wordt naar het scherm geprint
	fmt.Println(&age)

	// de waarde van age wordt naar het scherm geprint (24)
	fmt.Println(age)
}

// de functie change me verwacht een 'pointer' naar een int die dan opgenomen wordt in de variabele z
func changeMe(z *int) {
	// de variabele z wordt naar het scherm geprint. (het geheugenadres van age)
	fmt.Println(z)
	// de waarde op dat geheugenadres wordt naar het scherm geprint (44)
	fmt.Println(*z)

	// de waarde op het geheugen adress wordt overschreven met 24
	*z = 24

	// de variabele z wordt naar het scherm geprint. (het geheugenadres van age)
	fmt.Println(z)
	// de waarde op dat geheugenadres wordt naar het scherm geprint (24)
	fmt.Println(*z)
}

//geheugenadres
//geheugenadres
//44
//geheugenadres
//24
//geheugenadres
//24
