package expense_day_test

import (
	. "codertrack-kata/expense_day"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PrintSumPerDate_Input_Set_2_Should_Be_Output_Set_2(t *testing.T) {
	expectedSumExpenseString, _ := ioutil.ReadFile("output.2")
	rawInputData, _ := ioutil.ReadFile("input.2")
	inputData := strings.Split(string(rawInputData), "\n")

	sumExpense := SumExpense(inputData)
	actualSumExpenseString := PrintSumPerDate(sumExpense)

	assert.Equal(t, string(expectedSumExpenseString), actualSumExpenseString)
}

func Test_PrintSumPerDate_Input_Set_1_Should_Be_Output_Set_1(t *testing.T) {
	expectedSumExpenseString, _ := ioutil.ReadFile("output.1")
	rawInputData, _ := ioutil.ReadFile("input.1")
	inputData := strings.Split(string(rawInputData), "\n")

	sumExpense := SumExpense(inputData)
	actualSumExpenseString := PrintSumPerDate(sumExpense)

	assert.Equal(t, string(expectedSumExpenseString), actualSumExpenseString)
}

func Test_SumExpense_Input_Set_1_Should_Be_1300(t *testing.T) {
	expectedSumExpense := [7]int{71, 290, 607, 0, 0, 95, 237}
	rawInputData, _ := ioutil.ReadFile("input.1")
	inputData := strings.Split(string(rawInputData), "\n")

	actualSumExpense := SumExpense(inputData)

	assert.Equal(t, expectedSumExpense, actualSumExpense)
}

func Test_CalculateExpense_Input_3_Should_Be_0(t *testing.T) {
	expectedExpenseResult := 0
	inputData := "3"

	_, actualExpenseResult := CalculateExpense(inputData)

	assert.Equal(t, expectedExpenseResult, actualExpenseResult)
}

func Test_CalculateExpense_Input_1_f20_t15_f40_m200_t15_Should_Be_290(t *testing.T) {
	expectedExpenseResult := 290
	inputData := "1 f20 t15 f40 m200 t15"

	_, actualExpenseResult := CalculateExpense(inputData)

	assert.Equal(t, expectedExpenseResult, actualExpenseResult)
}
