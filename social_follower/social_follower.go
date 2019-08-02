package social_follower

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

func FindPopularFollower(data []string) []string {
	var followerTotalPoint = make(map[string]int)
	var followerSet []string
	var temp, followerName, action string
	var actionCount int
	for _, d := range data {
		fmt.Sscanf(d, "%s %s %s %d", &followerName, &temp, &action, &actionCount)
		if _, ok := followerTotalPoint[followerName]; !ok {
			followerSet = append(followerSet, followerName)
		}
		followerTotalPoint[followerName] += influencePointTable[action] * actionCount
	}
	followerSet = BubbleSort(followerSet, followerTotalPoint)
	for i := 0; i < len(followerSet); i++ {
		followerSet[i] = fmt.Sprintf("%s %d", followerSet[i], followerTotalPoint[followerSet[i]])
	}
	return followerSet[:5]
}
