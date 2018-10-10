package sum

func Summation(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}
