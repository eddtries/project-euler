package solutions

import "fmt"

// https://projecteuler.net/problem=50

func getPrime() func() int {
	currentPrime := 1
	return func() int {
		i := currentPrime + 1

		for {
			if isPrime(i) {
				currentPrime = i
				return currentPrime
			}
			i++
		}
	}
}

func Problem0050() int {
	nextPrime := getPrime()

	var consecutivePrimes []int
	runningTotal := 0

	for {
		prime := nextPrime()

		// Escape if value > 1,000,000
		if runningTotal+prime > 1000000 {
			break
		}

		// Track our running total of consecutive primes
		runningTotal += prime

		// If the total itself is a prime, save it
		if isPrime(runningTotal) {
			consecutivePrimes = append(consecutivePrimes, runningTotal)
		}
	}

	fmt.Println(consecutivePrimes)

	// Return the highest consecutive prime that is also prime
	return consecutivePrimes[len(consecutivePrimes)-1]
}
