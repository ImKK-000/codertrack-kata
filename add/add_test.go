package add_test

import (
	. "codertrack-kata/add"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add_Input_995144900824012405238259418456758068524921_And_879161138176194604342893336019619756243251_Should_Be_XXX(t *testing.T) {
	expectedNumberResult := "1874306039000207009581152754476377824768172"

	actualNumberResult := Add("995144900824012405238259418456758068524921", "879161138176194604342893336019619756243251")

	assert.Equal(t, expectedNumberResult, actualNumberResult)
}

func Test_Add_Input_48_And_16_Should_Be_64(t *testing.T) {
	expectedNumberResult := "64"

	actualNumberResult := Add("48", "16")

	assert.Equal(t, expectedNumberResult, actualNumberResult)
}
