package social_influencer

import "fmt"

var influencePointTable = map[string]int{
	"like":    1,
	"reply":   2,
	"retweet": 3,
}

func BubbleSort(set []string, totalPoint map[string]int) []string {
	setLength := len(set)
	for i := 0; i < setLength-1; i++ {
		for j := 0; j < setLength-i-1; j++ {
			if totalPoint[set[j]] < totalPoint[set[j+1]] {
				set[j], set[j+1] = set[j+1], set[j]
			}
		}
	}
	return set
}

func FindPopularInfluencer(data []string) []string {
	var influencerTotalPoint = make(map[string]int)
	var influencerSet []string
	var temp, influencerName, action string
	var actionCount int
	for _, d := range data {
		fmt.Sscanf(d, "%s %s %s %d", &temp, &influencerName, &action, &actionCount)
		if _, ok := influencerTotalPoint[influencerName]; !ok {
			influencerSet = append(influencerSet, influencerName)
		}
		influencerTotalPoint[influencerName] += influencePointTable[action] * actionCount
	}
	return BubbleSort(influencerSet, influencerTotalPoint)[:5]
}
