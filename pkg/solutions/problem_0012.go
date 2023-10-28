package solutions

import "log"

// https://projecteuler.net/problem=12

func getDivisors(x int) (divisors []int) {

	for i := 1; i <= x; i++ {
		if x%i == 0 {
			divisors = append(divisors, i)
		}
	}

	return
}

func getNextTriangleNumber() func() int {
	triangleNumber := 0
	number := 1

	return func() int {
		triangleNumber += number
		number++
		return triangleNumber
	}
}

func Problem0012() int {
	nextTriangleNumber := getNextTriangleNumber()

	for {
		triangleNumber := nextTriangleNumber()
		numOfDivisors := len(getDivisors(triangleNumber))
		log.Printf("%d has %d divisors", triangleNumber, numOfDivisors)
		if numOfDivisors >= 500 {
			return triangleNumber
		}
	}
}
