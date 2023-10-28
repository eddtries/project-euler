package solutions

import "log"

func getCollatzSequenceOfN(n int) []int {
	sequence := []int{n}
	for {
		if n == 1 {
			return sequence
		}
		if n%2 == 0 {
			n /= 2
			sequence = append(sequence, n)
		} else {
			n = 3*n + 1
			sequence = append(sequence, n)
		}
	}
}

func Problem0014() int {
	result := 0
	longestChain := 0
	for i := 1; i < 1000000; i++ {
		lengthOfSequence := len(getCollatzSequenceOfN(i))
		log.Printf("%d created a sequence of length %d", i, lengthOfSequence)
		if lengthOfSequence > longestChain {
			longestChain = lengthOfSequence
			result = i
		}
	}
	return result
}
