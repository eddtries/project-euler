package utils

func Sum(numbers []int) (sum int) {
	for number := range numbers {
		sum += number
	}

	return
}
