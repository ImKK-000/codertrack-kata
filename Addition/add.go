package add

import "math/big"

func Addition(numberFirst, numberSecond string) string {
	numberFirstBigInt, _ := new(big.Int).SetString(numberFirst, 10)
	numberSecondBigInt, _ := new(big.Int).SetString(numberSecond, 10)
	accumulator := new(big.Int).Add(numberFirstBigInt, numberSecondBigInt)
	return accumulator.String()
}
