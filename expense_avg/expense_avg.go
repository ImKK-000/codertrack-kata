package expense_avg

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func CalculateExpense(expenseListPerDay string) int {
	var sum int
	expenses := strings.Split(expenseListPerDay, " ")
	expenses = expenses[1:]
	for _, expense := range expenses {
		expenseStringWithTrim := strings.TrimFunc(expense, unicode.IsLetter)
		expenseNumber, _ := strconv.Atoi(expenseStringWithTrim)
		sum += expenseNumber
	}
	return sum
}

func SumExpense(expenseList []string) int {
	var sum int
	for _, expensePerDay := range expenseList {
		sum += CalculateExpense(expensePerDay)
	}
	return sum
}

func CalculateExpenseAverage(expenseList []string, day int) string {
	sumExpense := SumExpense(expenseList)
	return fmt.Sprintf("%.2f", float64(sumExpense)/float64(day))
}
