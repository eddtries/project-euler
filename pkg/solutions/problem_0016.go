package solutions

import (
	"math/big"
	"strconv"
)

// https://projecteuler.net/problem=16]

func Problem0016() int {
	number := new(big.Int)
	number.Exp(big.NewInt(2), big.NewInt(1000), nil)
	numberAsString := number.Text(10)
	result := 0
	for i := 0; i < len(numberAsString); i++ {
		num, _ := strconv.Atoi(string(numberAsString[i]))
		result += num
	}
	return result
}
