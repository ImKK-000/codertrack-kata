package transpose_test

import (
	. "codertrack-kata/transpose"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TransposeArray_Input_Set_2_Should_Be_Output_Set_2(t *testing.T) {
	expectedChords, _ := ioutil.ReadFile("output.2")
	rawInputData, _ := ioutil.ReadFile("input.2")
	inputData := strings.Split(string(rawInputData), "\n")

	actualChords := TransposeArray(inputData)

	assert.Equal(t, string(expectedChords), actualChords)
}

func Test_TransposeArray_Input_Set_1_Should_Be_Output_Set_1(t *testing.T) {
	expectedChords, _ := ioutil.ReadFile("output.1")
	rawInputData, _ := ioutil.ReadFile("input.1")
	inputData := strings.Split(string(rawInputData), "\n")

	actualChords := TransposeArray(inputData)

	assert.Equal(t, string(expectedChords), actualChords)
}

func Test_Transpose_Input_Am_C_F_G_Should_Be_Bm_D_G_A(t *testing.T) {
	expectedChords := "Bm D G A"
	inputData := "Am C F G"

	actualChords := Transpose(inputData)

	assert.Equal(t, expectedChords, actualChords)
}

func Test_Transpose_Input_C_Am_Dm_G_Should_Be_D_Bm_Em_A(t *testing.T) {
	expectedChords := "D Bm Em A"
	inputData := "C Am Dm G"

	actualChords := Transpose(inputData)

	assert.Equal(t, expectedChords, actualChords)
}

func Test_ConvertChord_Input_Chord_Am_Should_Be_Bm(t *testing.T) {
	expectedChord := "Bm"

	actualChord := ConvertChord("Am")

	assert.Equal(t, expectedChord, actualChord)
}

func Test_ConvertChord_Input_Chord_G_Should_Be_A(t *testing.T) {
	expectedChord := "A"

	actualChord := ConvertChord("G")

	assert.Equal(t, expectedChord, actualChord)
}

func Test_ConvertChord_Input_Chord_F_Should_Be_G(t *testing.T) {
	expectedChord := "G"

	actualChord := ConvertChord("F")

	assert.Equal(t, expectedChord, actualChord)
}

func Test_ConvertChord_Input_Chord_Em_Should_Be_FSharpm(t *testing.T) {
	expectedChord := "F#m"

	actualChord := ConvertChord("Em")

	assert.Equal(t, expectedChord, actualChord)
}

func Test_ConvertChord_Input_Chord_Dm_Should_Be_Em(t *testing.T) {
	expectedChord := "Em"

	actualChord := ConvertChord("Dm")

	assert.Equal(t, expectedChord, actualChord)
}

func Test_ConvertChord_Input_Chord_C_Should_Be_D(t *testing.T) {
	expectedChord := "D"

	actualChord := ConvertChord("C")

	assert.Equal(t, expectedChord, actualChord)
}
