package numbers

// Erwartet eine Zahl n und prüft, ob n eine Primzahl ist.
func IsPrime(n int) bool {
	// TODO
	t := 0

	for a := 1; n >= a; a++ {
		if n%a == 0 {
			t++
		}
	}

	if t == 2 {
		return true
	} else {
		return false
	}

}
