package rectangles

import "fmt"

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein Quadrat mit dieser Seitenlänge auf der Konsole.
// Das Quadrat soll komplett mit `#`-Zeichen gefüllt sein.
func DrawSolidSquare(length int) {

	for l1 := 0; l1 < length; l1++ {
		for l2 := 0; l2 < length; l2++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
