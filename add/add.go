package add

import (
	"math/big"
)

func Add(firstNumber, secondNumber string) string {
	firstNumberBigInt := new(big.Int)
	firstNumberBigInt.SetString(firstNumber, 10)
	secondNumberBigInt := new(big.Int)
	secondNumberBigInt.SetString(secondNumber, 10)
	resultNumberBigInt := new(big.Int)
	resultNumberBigInt.Add(firstNumberBigInt, secondNumberBigInt)
	return resultNumberBigInt.String()
}

func convertToNumber(number string) []int {
	var numberInt []int
	for _, number := range number {
		digitInt := int(number - '0')
		digitIntSlice := []int{digitInt}
		numberInt = append(digitIntSlice, numberInt...)
	}
	return numberInt
}
