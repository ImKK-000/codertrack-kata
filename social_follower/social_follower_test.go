package social_follower_test

import (
	. "codertrack-kata/social_follower"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPopularFollowerInputSet2ShouldBeList(t *testing.T) {
	expectedPopularFollower := []string{
		"elmstuff 5574",
		"dan_abramov 5485",
		"math3ma 5471",
		"kmett 5437",
		"martinfowler 5423",
	}
	rawInputData, _ := ioutil.ReadFile("input.2")
	inputData := strings.Split(string(rawInputData), "\n")

	actualPopularFollower := FindPopularFollower(inputData)

	assert.Equal(t, expectedPopularFollower, actualPopularFollower)
}

func TestFindPopularFollowerInputSet1ShouldBeList(t *testing.T) {
	expectedPopularFollower := []string{
		"AndrewYNg 489",
		"BartoszMilewski 408",
		"unclebobmartin 365",
		"martinfowler 268",
		"KentBeck 251",
	}
	rawInputData, _ := ioutil.ReadFile("input.1")
	inputData := strings.Split(string(rawInputData), "\n")

	actualPopularFollower := FindPopularFollower(inputData)

	assert.Equal(t, expectedPopularFollower, actualPopularFollower)
}
