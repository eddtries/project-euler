package utils

func CollatzSequence(n int) []int {
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
