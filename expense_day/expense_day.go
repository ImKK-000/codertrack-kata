package expense_day

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var Date = [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

func CalculateExpense(expenseListPerDay string) (int, int) {
	var sum int
	expenses := strings.Split(expenseListPerDay, " ")
	day, _ := strconv.Atoi(expenses[0])
	expenses = expenses[1:]
	for _, expense := range expenses {
		expenseStringWithTrim := strings.TrimFunc(expense, unicode.IsLetter)
		expenseNumber, _ := strconv.Atoi(expenseStringWithTrim)
		sum += expenseNumber
	}
	return day, sum
}

func SumExpense(expenseList []string) [7]int {
	var sumDays [7]int
	for _, expensePerDay := range expenseList {
		day, sumPerDay := CalculateExpense(expensePerDay)
		sumDays[day%7] += sumPerDay
	}
	return sumDays
}

func PrintSumPerDate(dates [7]int) string {
	var outputString string
	for index, date := range Date {
		outputString += fmt.Sprintln(date, dates[(index+5)%7])
	}
	return outputString
}
