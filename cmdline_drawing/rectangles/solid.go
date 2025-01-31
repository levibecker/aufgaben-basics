package rectangles

import "fmt"

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Das Rechteck soll komplett mit `#`-Zeichen gefüllt sein.
func DrawSolidRectangle(height, width int) {

	for h := 0; h < height; h++ {

		for w := 0; w < width; w++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

}
