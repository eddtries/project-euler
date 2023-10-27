package solutions

import (
	"math/big"
)

func Problem0048() {
	result := big.NewInt(0)

	for i := 1; i <= 1000; i++ {
		bigI := big.NewInt(int64(i))
		result.Add(bigI.Exp(bigI, bigI, nil), result)
	}
	println(result.Text(10)[len(result.Text(10))-10:])
}
