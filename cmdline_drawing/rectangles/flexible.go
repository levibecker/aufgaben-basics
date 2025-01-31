package rectangles

import "fmt"

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Die Zeichen für Rand und Füllung des Rechtecks werden als Parameter erwartet.
func DrawRectangle(height, width int, inner, outer string) {
	// TODO

	hi := height - 2
	wi := width - 2

	for w1 := 0; w1 < width; w1++ {
		fmt.Print(outer)
	}
	fmt.Println()

	for h1 := 0; h1 < hi; h1++ {
		fmt.Print(outer)
		for w := 0; w < wi; w++ {
			fmt.Print(inner)
		}
		fmt.Print(outer)
		fmt.Println()

	}
	for w1 := 0; w1 < width; w1++ {
		fmt.Print(outer)
	}
	fmt.Println()
}

// REMARKS
// - Wenn Sie diese Aufgabe gelöst haben, können Sie die Aufgaben
//   für das Zeichnen von leeren und gefüllten Rechtecken sehr viel einfacher lösen.
