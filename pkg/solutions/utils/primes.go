package utils

func IsPrime(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}

	return true
}

func GetPrimesUpToLimit(limit int) (primes []int) {
	for i := 2; i < limit; i++ {
		if IsPrime(limit) {
			primes = append(primes, limit)
		}
	}

	return
}

func PrimesGenerator() func() int {
	currentPrime := 1
	return func() int {
		i := currentPrime + 1

		for {
			if IsPrime(i) {
				currentPrime = i
				return currentPrime
			}
			i++
		}
	}
}
