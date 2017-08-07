package main

import "fmt"

// de funtcie wrapper verwacht een anonieme functie van het type int terug ()
// dat wordt de annonime functie met waarde van x
func wrapper() func() int {
	// x geinitialiseerd naar 0 op blok niveau in het buitenste blok
	// waarom wordt x de tweede keer niet opnieuw geinitialiseerd?
	x := 0
	// de binnenste functie geeft een annonime functie van het type int terug
	return func() int {
		// x wordt opgehoogd met 1
		x++
		// geef de waarde van x als functie terug aan het buitenste blok
		return x
	}
}

func main() {
	// binnen de functie wrapper zit de annonieme functie met de waarde van x
	// die functie wordt eenmalig toegewezen aan de variabele element (dan wordt de variable x op 0 ginitialiseerd)
	// daardoor wordt er een functie element() gemaakt met de waarde van x
	element := wrapper()

	// de annonime functie uit wrapper die toegewezen is aan element wordt aangeroepen en de uitkomst wordt naar het scherm geprint. dat betekend dat wrapper() zelf niet meer wordt aangeroepen maar alleen het binnenste blok (de anonime functie)
	// met andere woorde element() is de annonieme functie uit het binnenblok van wrapper()
	fmt.Println(element())
	fmt.Println(element())
}

//BELANGRIJK!!!
