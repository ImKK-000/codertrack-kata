package grading_avg_test

import (
	. "codertrack-kata/grading_avg"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SumAvg_Input_Set_2_Should_Be_GradeSum_1_Point_412(t *testing.T) {
	expectedGradeSum := 1.412
	rawInputData, _ := ioutil.ReadFile("input.2")
	inputData := strings.Split(string(rawInputData), "\n")

	actualGradeSum := SumAverage(inputData[:len(inputData)-1])

	assert.Equal(t, expectedGradeSum, actualGradeSum)
}

func Test_SumAvg_Input_Set_1_Should_Be_GradeSum_0_Point_7(t *testing.T) {
	expectedGradeSum := 0.7
	rawInputData, _ := ioutil.ReadFile("input.1")
	inputData := strings.Split(string(rawInputData), "\n")

	actualGradeSum := SumAverage(inputData[:len(inputData)-1])

	assert.Equal(t, expectedGradeSum, actualGradeSum)
}

func Test_ConvertScoreToGrade_Input_100_Should_Be_Grade_4(t *testing.T) {
	expectedGrade := 4.0

	actualGrade := ConvertScoreToGrade(89)

	assert.Equal(t, expectedGrade, actualGrade)
}

func Test_ConvertScoreToGrade_Input_84_Should_Be_Grade_3_Point_5(t *testing.T) {
	expectedGrade := 3.5

	actualGrade := ConvertScoreToGrade(84)

	assert.Equal(t, expectedGrade, actualGrade)
}

func Test_ConvertScoreToGrade_Input_79_Should_Be_Grade_3(t *testing.T) {
	expectedGrade := 3.0

	actualGrade := ConvertScoreToGrade(79)

	assert.Equal(t, expectedGrade, actualGrade)
}

func Test_ConvertScoreToGrade_Input_74_Should_Be_Grade_2_Point_5(t *testing.T) {
	expectedGrade := 2.5

	actualGrade := ConvertScoreToGrade(74)

	assert.Equal(t, expectedGrade, actualGrade)
}

func Test_ConvertScoreToGrade_Input_69_Should_Be_Grade_2(t *testing.T) {
	expectedGrade := 2.0

	actualGrade := ConvertScoreToGrade(69)

	assert.Equal(t, expectedGrade, actualGrade)
}

func Test_ConvertScoreToGrade_Input_64_Should_Be_Grade_1_Point_5(t *testing.T) {
	expectedGrade := 1.5

	actualGrade := ConvertScoreToGrade(64)

	assert.Equal(t, expectedGrade, actualGrade)
}

func Test_ConvertScoreToGrade_Input_59_Should_Be_Grade_1(t *testing.T) {
	expectedGrade := 1.0

	actualGrade := ConvertScoreToGrade(59)

	assert.Equal(t, expectedGrade, actualGrade)
}

func Test_ConvertScoreToGrade_Input_54_Should_Be_Grade_0_Point_5(t *testing.T) {
	expectedGrade := 0.5

	actualGrade := ConvertScoreToGrade(54)

	assert.Equal(t, expectedGrade, actualGrade)
}

func Test_ConvertScoreToGrade_Input_0_Should_Be_Grade_0(t *testing.T) {
	expectedGrade := 0.0

	actualGrade := ConvertScoreToGrade(0)

	assert.Equal(t, expectedGrade, actualGrade)
}
