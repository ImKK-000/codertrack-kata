package expense_cat_test

import (
	. "codertrack-kata/expense_cat"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExpenseCategory_ToString_Input_Set_2_Should_Be_Output_Set_2(t *testing.T) {
	expectedExpenseResult, _ := ioutil.ReadFile("output.2")
	rawInputData, _ := ioutil.ReadFile("input.2")
	inputData := strings.Split(string(rawInputData), "\n")

	expenseResult := SumExpense(inputData)
	actualExpenseString := expenseResult.ToString()

	assert.Equal(t, string(expectedExpenseResult), actualExpenseString)
}

func Test_ExpenseCategory_ToString_Input_Set_1_Should_Be_Output_Set_1(t *testing.T) {
	expectedExpenseResult, _ := ioutil.ReadFile("output.1")
	rawInputData, _ := ioutil.ReadFile("input.1")
	inputData := strings.Split(string(rawInputData), "\n")

	expenseResult := SumExpense(inputData)
	actualExpenseString := expenseResult.ToString()

	assert.Equal(t, string(expectedExpenseResult), actualExpenseString)
}

func Test_SumExpense_Input_Set_1_Should_Be_Food_833_Game_140_Movie_200_Stationery_0_Transportation_127_TotalExpense_1300(t *testing.T) {
	expectedExpenseResult := ExpenseCategory{
		Food:           833,
		Game:           140,
		Movie:          200,
		Stationery:     0,
		Transportation: 127,
		TotalExpense:   1300,
	}
	rawInputData, _ := ioutil.ReadFile("input.1")
	inputData := strings.Split(string(rawInputData), "\n")

	actualExpenseResult := SumExpense(inputData)

	assert.Equal(t, expectedExpenseResult, actualExpenseResult)
}

func Test_CalculateExpense_Input_3_Should_Be_0(t *testing.T) {
	expectedExpenseResult := ExpenseCategory{}
	inputData := "3"

	actualExpenseResult := CalculateExpense(inputData)

	assert.Equal(t, expectedExpenseResult, actualExpenseResult)
}

func Test_CalculateExpense_Input_1_f20_t15_f40_m200_t15_Should_Be_290(t *testing.T) {
	expectedExpenseResult := ExpenseCategory{
		Food:           60,
		Game:           0,
		Movie:          200,
		Stationery:     0,
		Transportation: 30,
		TotalExpense:   290,
	}
	inputData := "1 f20 t15 f40 m200 t15"

	actualExpenseResult := CalculateExpense(inputData)

	assert.Equal(t, expectedExpenseResult, actualExpenseResult)
}
