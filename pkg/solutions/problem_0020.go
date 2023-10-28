package solutions

import (
	"math/big"
	"strconv"
)

// https://projecteuler.net/problem=16

func Problem0020() int {
	number := big.NewInt(1)

	for i := 100; i > 0; i-- {
		iAsBigNum := big.NewInt(int64(i))
		number.Mul(number, iAsBigNum)
	}

	numberAsString := number.Text(10)
	result := 0
	for i := 0; i < len(numberAsString); i++ {
		num, _ := strconv.Atoi(string(numberAsString[i]))
		result += num
	}
	return result
}
