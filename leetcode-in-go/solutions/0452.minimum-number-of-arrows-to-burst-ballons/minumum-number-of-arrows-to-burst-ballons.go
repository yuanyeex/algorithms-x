package prb0452

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	count := 1
	currArrow := points[0][1]

	for _, point := range points[1:] {
		if point[0] <= currArrow {
			continue
		}
		count++
		currArrow = point[1]
	}

	return count
}
