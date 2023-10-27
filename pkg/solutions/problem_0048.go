package solutions

import (
	"math/big"
)

func Problem0048() {
	result := big.NewInt(0)

	for i := 1; i <= 100000; i++ {
		bigI := big.NewInt(int64(i))
		result.Add(bigI.Exp(bigI, bigI, nil), result)
	}
	println(len(result.Text(10)))
}
