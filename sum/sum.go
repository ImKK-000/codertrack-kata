package sum

import "strconv"

func Sum(numberStrings []string) int {
	var sum int
	for _, numberString := range numberStrings {
		number, _ := strconv.Atoi(numberString)
		sum += number
	}
	return sum
}
