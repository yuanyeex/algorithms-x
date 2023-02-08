package prb0406

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		} else {
			return people[i][0] > people[j][0]
		}
	})

	for i, person := range people {
		ind := person[1]
		copy(people[ind+1:i+1], people[ind:i])
		people[ind] = person
	}

	return people
}
