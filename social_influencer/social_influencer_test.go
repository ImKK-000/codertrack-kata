package social_influencer_test

import (
	. "codertrack-kata/social_influencer"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPopularInfluencerInputSet2ShouldBeList(t *testing.T) {
	expectedPopularInfluencer := []string{
		"gradientpub",
		"am_i_tom",
		"Peerapat_JK",
		"TacticalGrace",
		"stuarthalloway",
	}
	rawInputData, _ := ioutil.ReadFile("input.2")
	inputData := strings.Split(string(rawInputData), "\n")

	actualPopularInfluencer := FindPopularInfluencer(inputData)

	assert.Equal(t, expectedPopularInfluencer, actualPopularInfluencer)
}

func TestFindPopularInfluencerInputSet1ShouldBeList(t *testing.T) {
	expectedPopularInfluencer := []string{
		"DeepLearningHub",
		"martinfowler",
		"KentBeck",
		"AndrewYNg",
		"SFMachineLearn",
	}
	rawInputData, _ := ioutil.ReadFile("input.1")
	inputData := strings.Split(string(rawInputData), "\n")

	actualPopularInfluencer := FindPopularInfluencer(inputData)

	assert.Equal(t, expectedPopularInfluencer, actualPopularInfluencer)
}
