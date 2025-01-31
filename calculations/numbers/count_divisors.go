package numbers

// Erwartet eine Zahl n >= 1 und liefert die Anzahl der Teiler dieser Zahl zurück.
func CountDivisors(n int) int {
	// TODO
	t := 0

	for a := 1; n >= a; a++ {
		if n%a == 0 {
			t++
		}
	}
	return t
}
