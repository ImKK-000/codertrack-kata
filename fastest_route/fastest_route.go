package fastest_route

import (
	"fmt"

	"github.com/RyanCarrier/dijkstra"
)

func FindShortestPathFrom(data []string) int64 {
	graph := dijkstra.NewGraph()
	var sourcePath, destinationPath string
	var distance int64
	for _, path := range data {
		fmt.Sscanf(path, "%s %s %d", &sourcePath, &destinationPath, &distance)
		graph.AddMappedArc(sourcePath, destinationPath, distance)
	}
	homeID, _ := graph.GetMapping("HOME")
	destID, _ := graph.GetMapping("DEST")
	bestPath, _ := graph.Shortest(homeID, destID)
	return bestPath.Distance
}
