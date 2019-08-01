package grading_avg

import (
	"fmt"
	"strconv"
)

func ConvertScoreToGrade(score int) float64 {
	switch {
	case score >= 85 && score <= 100:
		return 4
	case score >= 80 && score <= 84:
		return 3.5
	case score >= 75 && score <= 79:
		return 3
	case score >= 70 && score <= 74:
		return 2.5
	case score >= 65 && score <= 69:
		return 2
	case score >= 60 && score <= 64:
		return 1.5
	case score >= 55 && score <= 59:
		return 1
	case score >= 50 && score <= 54:
		return 0.5
	}
	return 0
}

func SumAverage(scores []string) float64 {
	var sum float64
	for _, score := range scores {
		scoreInt, _ := strconv.Atoi(score)
		sum += ConvertScoreToGrade(scoreInt)
		fmt.Println(scoreInt)
	}
	return sum / float64(len(scores))
}
