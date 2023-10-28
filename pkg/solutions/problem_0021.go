package solutions

// https://projecteuler.net/problem=21

func sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func properDivisors(n int) []int {
	var divisors []int
	for i := 1; i < n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

func isAmicableNumber(n int) bool {
	x := sum(properDivisors(n))
	y := sum(properDivisors(x))

	// Avoid duplicates where n == y but n == x
	if n == x {
		return false
	}

	return n == y
}

func Problem0021() int {
	sum := 0

	for i := 0; i < 10000; i++ {
		if isAmicableNumber(i) {
			sum += i
		}
	}

	return sum
}
