package solutions

func Problem0006() int {
	sumOfSquares := 0
	squareOfSums := 0

	for i := 1; i <= 100; i++ {
		sumOfSquares += i * i
	}

	for i := 1; i <= 100; i++ {
		squareOfSums += i
	}

	squareOfSums *= squareOfSums

	return squareOfSums - sumOfSquares
}
