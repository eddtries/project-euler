package solutions

import (
	"bufio"
	"log"
	"math/big"
	"os"
)

func Problem0013() int {
	result := new(big.Int)

	file, _ := os.Open("./pkg/solutions/data/problem_0013.data")
	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		bigNum := new(big.Int)
		bigNum, ok := bigNum.SetString(fscanner.Text(), 10)
		if !ok {
			log.Fatalln("Failed to read fscanner.Text() to big.Int")
		}

		result.Add(result, bigNum)
	}

	println(result.Text(10)[:10])
	return 0
}
