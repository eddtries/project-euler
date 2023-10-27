package solutions

import (
	"sort"
)

// https://projecteuler.net/problem=3

func getPrimesUpToX(x int) []int {
	var primes []int

	for i := 2; i < x; i++ {
		if isPrime(x) {
			primes = append(primes, x)
		}
	}

	return primes
}

func getPrimeFactors(x int) []int {
	var primeFactors []int

	primes := getPrimesUpToX(x)
	for _, prime := range primes {
		if x%prime == 0 {
			primeFactors = append(primeFactors, prime)
			x /= prime
		}
	}
	return primeFactors
}

func Problem0003() int {
	factors := getPrimeFactors(1000)
	sort.Ints(factors)
	return factors[len(factors)-1]
}
