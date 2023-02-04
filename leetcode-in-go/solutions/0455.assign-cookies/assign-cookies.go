package prb0455

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	childIndex := 0

	for _, cookie := range s {
		if childIndex >= len(g) {
			break
		}

		if g[childIndex] <= cookie {
			childIndex++
		}
	}

	return childIndex
}
