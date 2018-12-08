package expense_cat

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type ExpenseCategory struct {
	Food, Game, Movie, Stationery, Transportation, TotalExpense int
}

func (expenseCategory ExpenseCategory) ToString() string {
	return fmt.Sprintln("food", expenseCategory.Food) +
		fmt.Sprintln("game", expenseCategory.Game) +
		fmt.Sprintln("movie", expenseCategory.Movie) +
		fmt.Sprintln("stationery", expenseCategory.Stationery) +
		fmt.Sprintln("transportation", expenseCategory.Transportation) +
		fmt.Sprintln("TOTAL", expenseCategory.TotalExpense)
}

func CalculateExpense(expenseListPerDay string) ExpenseCategory {
	var expenseCategory ExpenseCategory
	expenses := strings.Split(expenseListPerDay, " ")
	expenses = expenses[1:]
	for _, expense := range expenses {
		expenseStringWithTrim := strings.TrimFunc(expense, unicode.IsLetter)
		expenseNumber, _ := strconv.Atoi(expenseStringWithTrim)
		switch expense[0] {
		case 'f':
			expenseCategory.Food += expenseNumber
			break
		case 'g':
			expenseCategory.Game += expenseNumber
			break
		case 'm':
			expenseCategory.Movie += expenseNumber
			break
		case 's':
			expenseCategory.Stationery += expenseNumber
			break
		case 't':
			expenseCategory.Transportation += expenseNumber
			break
		}
		expenseCategory.TotalExpense += expenseNumber
	}
	return expenseCategory
}

func SumExpense(expenseList []string) ExpenseCategory {
	var SumExpenseCategory ExpenseCategory
	for _, expensePerDay := range expenseList {
		expenseCategoryPerDay := CalculateExpense(expensePerDay)
		SumExpenseCategory.Food += expenseCategoryPerDay.Food
		SumExpenseCategory.Game += expenseCategoryPerDay.Game
		SumExpenseCategory.Movie += expenseCategoryPerDay.Movie
		SumExpenseCategory.Stationery += expenseCategoryPerDay.Stationery
		SumExpenseCategory.Transportation += expenseCategoryPerDay.Transportation
		SumExpenseCategory.TotalExpense += expenseCategoryPerDay.TotalExpense
	}
	return SumExpenseCategory
}
