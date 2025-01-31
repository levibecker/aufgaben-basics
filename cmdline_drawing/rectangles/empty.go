package rectangles

import "fmt"

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Der Rand des Rechtecks soll aus `#`-Zeichen bestehen, der Innenraum soll leer sein.
func DrawEmptyRectangle(height, width int) {
	hi := height - 2
	wi := width - 2

	for w := 0; w < width; w++ {
		fmt.Print("#")
	}

	fmt.Println()

	for h1 := 0; h1 < hi; h1++ {
		fmt.Print("#")
		for w2 := 0; w2 < wi; w2++ {
			fmt.Print(" ")
		}
		fmt.Print("#")
		fmt.Println()
	}

	for w := 0; w < width; w++ {
		fmt.Print("#")
	}

	fmt.Println()
}

// REMARKS
// - Diese Aufgabe ist eine etwas komplexere Variante der Aufgabe "Gefülltes Rechteck zeichnen".
//   Man geht dabei ähnlich vor, muss allerdings zusätzlich prüfen, ob man sich am Rand des Rechtecks befindet.
