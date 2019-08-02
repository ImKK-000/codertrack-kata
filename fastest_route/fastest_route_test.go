package fastest_route_test

import (
	. "codertrack-kata/fastest_route"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindShortestPathFromInputTestSet2ShouldBeDistance132(t *testing.T) {
	expectedDistance := int64(132)
	rawInputData, _ := ioutil.ReadFile("input.2")
	inputData := strings.Split(string(rawInputData), "\n")

	actualDistance := FindShortestPathFrom(inputData)

	assert.Equal(t, expectedDistance, actualDistance)
}

func TestFindShortestPathFromInputTestSet1ShouldBeDistance40(t *testing.T) {
	expectedDistance := int64(40)
	rawInputData, _ := ioutil.ReadFile("input.1")
	inputData := strings.Split(string(rawInputData), "\n")

	actualDistance := FindShortestPathFrom(inputData)

	assert.Equal(t, expectedDistance, actualDistance)
}
