package solutions

func Problem0005() int {
	result := 0
	numbers := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for {
		result += 20
		isDivisible := true
		for _, number := range numbers {
			if result%number != 0 {
				isDivisible = false
			}
		}

		if isDivisible {
			return result
		}
	}
}
