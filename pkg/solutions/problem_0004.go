package solutions

import (
	"sort"
	"strconv"
)

// https://projecteuler.net/problem=4

func reverseString(input string) string {
	output := ""
	for _, c := range input {
		output = string(c) + output
	}
	return output
}

func Problem0004() int {
	var palindromeNumbers []int
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			product := i * j
			productAsString := strconv.Itoa(product)
			if productAsString == reverseString(productAsString) {
				palindromeNumbers = append(palindromeNumbers, product)
			}
		}
	}

	sort.Ints(palindromeNumbers)
	return palindromeNumbers[len(palindromeNumbers)-1]
}
