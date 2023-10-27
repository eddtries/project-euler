package solutions

import "math"

// https://projecteuler.net/problem=9

func Problem0009() int {
	// Search for 1000 to ensure the entire search space is checked on
	// the very unlikely answer of 998 + 1 + 1 or something
	for a := 1; a <= 1000; a++ {
		for b := 1; b <= 1000; b++ {
			cSquared := (a * a) + (b * b)
			c := math.Sqrt(float64(cSquared))
			// Hacky way to check that c is a whole number
			if c == math.Trunc(c) {
				if a+b+int(c) == 1000 {
					return a * b * int(c)
				}
			}
		}
	}

	return -1
}
