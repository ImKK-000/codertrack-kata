package expense_avg_test

import (
	. "codertrack-kata/expense_avg"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CalculateExpenseAverage_Input_Set_2_Should_Be_192Dot81(t *testing.T) {
	expectedExpenseAverage := "192.81"
	rawInputData, _ := ioutil.ReadFile("input.2")
	inputData := strings.Split(string(rawInputData), "\n")
	day := len(inputData)

	actualExpenseAverage := CalculateExpenseAverage(inputData, day)

	assert.Equal(t, expectedExpenseAverage, actualExpenseAverage)
}

func Test_CalculateExpenseAverage_Input_Set_1_Should_Be_185Dot71(t *testing.T) {
	expectedExpenseAverage := "185.71"
	rawInputData, _ := ioutil.ReadFile("input.1")
	inputData := strings.Split(string(rawInputData), "\n")
	day := len(inputData)

	actualExpenseAverage := CalculateExpenseAverage(inputData, day)

	assert.Equal(t, expectedExpenseAverage, actualExpenseAverage)
}

func Test_SumExpense_Input_Set_1_Should_Be_1300(t *testing.T) {
	expectedSumExpense := 1300
	rawInputData, _ := ioutil.ReadFile("input.1")
	inputData := strings.Split(string(rawInputData), "\n")

	actualSumExpense := SumExpense(inputData)

	assert.Equal(t, expectedSumExpense, actualSumExpense)
}
func Test_CalculateExpense_Input_3_Should_Be_0(t *testing.T) {
	expectedExpenseResult := 0
	inputData := "3"

	actualExpenseResult := CalculateExpense(inputData)

	assert.Equal(t, expectedExpenseResult, actualExpenseResult)
}

func Test_CalculateExpense_Input_1_f20_t15_f40_m200_t15_Should_Be_290(t *testing.T) {
	expectedExpenseResult := 290
	inputData := "1 f20 t15 f40 m200 t15"

	actualExpenseResult := CalculateExpense(inputData)

	assert.Equal(t, expectedExpenseResult, actualExpenseResult)
}
