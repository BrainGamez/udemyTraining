package main

import (
	"fmt"
)

// x krijgt de waarde van x in main
// daarna wordt de waarde van x op 0 gezet
func zero(z int) {
	z = 0
}

// x krijgt de waarde 5
// de functie zero wordt aangeroepen met de argument 5
// de lokale x wordt naar het scherm geprint (met x in zero wordt eigenlijk niks gedaan)
func main() {
	x := 5
	zero(x)
	fmt.Println(x)

}
