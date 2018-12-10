package couple_test

import (
	. "codertrack-kata/couple"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetPersonSingle_Input_Set_2_Should_Be_Output_Set_2(t *testing.T) {
	rawOutputData, _ := ioutil.ReadFile("output.2")
	expectedPersonSingle := strings.Split(string(rawOutputData), "\n")
	rawInputData, _ := ioutil.ReadFile("input.2")
	inputData := strings.Split(string(rawInputData), "\n")

	actualPersonSingle := GetPersonSingle(inputData)

	assert.Equal(t, expectedPersonSingle, actualPersonSingle)
}

func Test_GetPersonSingle_Input_Set_1_Should_Be_Output_Set_1(t *testing.T) {
	rawOutputData, _ := ioutil.ReadFile("output.1")
	expectedPersonSingle := strings.Split(string(rawOutputData), "\n")
	rawInputData, _ := ioutil.ReadFile("input.1")
	inputData := strings.Split(string(rawInputData), "\n")

	actualPersonSingle := GetPersonSingle(inputData)

	assert.Equal(t, expectedPersonSingle, actualPersonSingle)
}

func Test_GetPerson_Input_Boy_E5_Should_Be_Name_Boy_Destination_E5(t *testing.T) {
	expectedPersonName := "Boy"
	expectedPersonDestination := "E5"
	inputData := "Boy E5"

	actualPersonName, actualPersonDestination := GetPerson(inputData)

	assert.Equal(t, expectedPersonName, actualPersonName)
	assert.Equal(t, expectedPersonDestination, actualPersonDestination)
}

func Test_SortName_Input_Set_1_Should_Be_Input_Set_1_Is_Sort(t *testing.T) {
	expectedSortName := []string{
		"BOY E5",
		"JAN X2",
		"JIB B6",
		"KIM A1",
		"NOK A2",
		"TU A1",
	}
	rawInputData, _ := ioutil.ReadFile("input.1")
	inputData := strings.Split(string(rawInputData), "\n")

	actualSortName := SortName(inputData)

	assert.Equal(t, expectedSortName, actualSortName)
}
