package solutions

func _sum(nums []int) (result int) {
	for num, _ := range nums {
		result += num
	}

	return
}

func getProperDivisors(x int) (divisors []int) {
	for i := 1; i < x; i++ {
		if x%i == 0 {
			divisors = append(divisors, i)
		}
	}

	return
}

func isAbundentNumber(x int) bool {
	divisors := getProperDivisors(x)
	total := sum(divisors)
	return total > x
}

func getAbundentNumbers(limit int) (abundentNumbers []int) {
	for i := 0; i <= limit; i++ {
		if isAbundentNumber(i) {
			abundentNumbers = append(abundentNumbers, i)
		}
	}
	return
}

func Problem0023() int {
	abundentNumbers := getAbundentNumbers(28123)

	var evenNumbersNotSumOfTwoAbundentNumbers []int
	for i := 2; i <= 28123+28123; i = i + 2 {
		for x := 0; x < len(abundentNumbers); x++ {
			complement := i - abundentNumbers[x]
			if complement < 0 {
				break
			}
			for y := 0; y < len(abundentNumbers); y++ {
				if abundentNumbers[y] == complement && x != y {
					evenNumbersNotSumOfTwoAbundentNumbers = append(evenNumbersNotSumOfTwoAbundentNumbers, i)
					break
				}
			}
		}
	}

	return sum(evenNumbersNotSumOfTwoAbundentNumbers)
}
