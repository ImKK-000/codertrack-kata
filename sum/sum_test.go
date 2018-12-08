package sum_test

import (
	. "codertrack-kata/sum"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Sum_Input_Set_2_Should_Be_635589218(t *testing.T) {
	expectedNumberResult := 635589218
	rawInputData, _ := ioutil.ReadFile("input.2")
	inputData := strings.Split(string(rawInputData), "\n")

	actualNumberResult := Sum(inputData)

	assert.Equal(t, expectedNumberResult, actualNumberResult)
}

func Test_Sum_Input_Set_1_Should_Be_10(t *testing.T) {
	expectedNumberResult := 10
	rawInputData, _ := ioutil.ReadFile("input.1")
	inputData := strings.Split(string(rawInputData), "\n")

	actualNumberResult := Sum(inputData)

	assert.Equal(t, expectedNumberResult, actualNumberResult)
}
