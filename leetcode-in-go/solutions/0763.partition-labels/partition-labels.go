package prb0763

func partitionLabels(s string) []int {
	lastIndexMapping := make(map[rune]int)

	for i, v := range s {
		lastIndexMapping[v] = i
	}

	var res []int
	start, last := 0, 0
	for i, v := range s {
		if lastIndexMapping[v] > last {
			last = lastIndexMapping[v]
		}

		if i == last {
			res = append(res, i-start+1)
			start = i + 1
		}
	}

	return res
}
