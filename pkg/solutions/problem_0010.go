package solutions

// https://projecteuler.net/problem=10

func allPrimesUnderX(x int) []int {
	var primes []int

	i := 2

	for {
		if i >= x {
			break
		}

		if isPrime(i) {
			primes = append(primes, i)
		}

		i++
	}

	return primes
}

func Problem0010() int {
	var result int

	primes := allPrimesUnderX(2000000)
	for _, prime := range primes {
		result += prime
	}

	return result
}
