package main

import "fmt"

// funct zero krijgt het geheugenadres van x in main (welke de waarde 5 heeft)
// de waarde van x wordt naar het scherm geprint
// dan wordt de waarde op dat geheugenadres overschreven met 0
func zero(x *int) {
	fmt.Println("Waarde van get geheugenadres in zero: ", x)
	*x = 0

}

// x wordt geinitialiseerd op 5
// het geheugenadres van x wordt naar het scherm geprint
// func zero wordt aangeroepen met het geheugenadres van x als argument
// de waarde van de lokale x (welke via het geheugenadres overschreven is in func zero) wordt naar het scherm geprint.
func main() {
	x := 5
	fmt.Println("Waarde van get geheugenadres in main: ", &x)
	zero(&x)
	fmt.Println(x)
}

// de Println van zero wordt pas geprint wanneer zero aangeroepen wordt.
// de Println van main wordt dus eerder geprint
// de waarde van x binnen main wordt als laatste geprint.
