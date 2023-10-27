package solutions

/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10001st prime number?
*/

func isPrime(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}

	return true
}

func Problem0007() int {
	primes := []int{2}

	current_int := 3

	for {
		if len(primes) == 10001 {
			break
		}

		if isPrime(current_int) {
			primes = append(primes, current_int)
		}

		current_int += 1
	}

	return primes[len(primes)-1]
}
