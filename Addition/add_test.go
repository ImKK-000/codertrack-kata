package add_test

import (
	. "add"
	"testing"
)

func TestAdditionInput48And16ShouldBe64(t *testing.T) {
	expectedNumber := "64"

	actualNumber := Addition("48", "16")

	if expectedNumber != actualNumber {
		t.Errorf("expect %s but it got %s", expectedNumber, actualNumber)
	}
}

func TestAdditionInput995144900824012405238259418456758068524921And879161138176194604342893336019619756243251ShouldBe1874306039000207009581152754476377824768172(t *testing.T) {
	expectedNumber := "1874306039000207009581152754476377824768172"

	actualNumber := Addition("995144900824012405238259418456758068524921", "879161138176194604342893336019619756243251")

	if expectedNumber != actualNumber {
		t.Errorf("expect %s but it got %s", expectedNumber, actualNumber)
	}
}
