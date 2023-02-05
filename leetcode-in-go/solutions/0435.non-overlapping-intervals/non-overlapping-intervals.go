package prb0435

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	// 1. sort by the ending interval
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	// 2. iterate and select the
	// > min interval ending
	// > not overlapping with the prev intervals
	removed := 0
	prev := intervals[0][1]
	for _, interval := range intervals[1:] {
		if interval[0] < prev {
			removed++
		} else {
			prev = interval[1]
		}
	}

	return removed
}
