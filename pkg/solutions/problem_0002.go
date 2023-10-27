package solutions

// https://projecteuler.net/problem=2

func calculateNextFibValue(nums []int) int {
	last := nums[len(nums)-1]
	secondLast := nums[len(nums)-2]
	return last + secondLast
}

func Problem0002() int {
	fibValues := []int{1, 1}
	for {
		nextValue := calculateNextFibValue(fibValues)
		if nextValue > 4000000 {
			break
		}

		fibValues = append(fibValues, nextValue)
	}

	sumOfEvenValues := 0

	for _, v := range fibValues {
		if v%2 == 0 {
			sumOfEvenValues += v
		}
	}

	return sumOfEvenValues
}
