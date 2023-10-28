package solutions

import (
	"fmt"
	"math/big"
)

// https://projecteuler.net/problem=25

func numberOfDigits(num big.Int) int {
	return len(num.Text(10))
}

func getFib() func() big.Int {
	x := big.NewInt(1)
	y := big.NewInt(1)

	return func() big.Int {
		nextValue := new(big.Int)
		nextValue.Add(x, y)

		if x.Cmp(y) == 1 {
			y = nextValue
		} else {
			x = nextValue
		}

		return *nextValue
	}
}

func Problem0025() int {
	fibSeq := getFib()
	counter := 3 // Because we start with 1, 1 in the closure and have just calculated the 3rd value

	for {
		nextFibValue := fibSeq()
		fmt.Printf("Fib value at %d is %d digits long\n", counter, numberOfDigits(nextFibValue))
		if numberOfDigits(nextFibValue) >= 1000 {
			return counter
		}

		counter++
	}
}
