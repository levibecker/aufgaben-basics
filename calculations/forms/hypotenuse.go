package forms

import "math"

// Erwartet die Längen der Katheten eines rechtwinkligen Dreiecks.
// Liefert die Länge der Hypotenuse.
func Hypotenuse(a, b float64) float64 {
	// TODO
	c := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	return c
}
