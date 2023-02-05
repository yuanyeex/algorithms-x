package prb0435

import "testing"

type problem struct {
	input    [][]int
	expected int
}

func Test_Ok(t *testing.T) {
	problems := []problem{
		{
			//Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
			[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
			1,
		},
		{
			// [1,2],[1,2],[1,2]]
			[][]int{{1, 2}, {1, 2}, {1, 2}},
			2,
		},
	}

	for _, prb := range problems {
		solu := eraseOverlapIntervals(prb.input)
		if solu != prb.expected {
			t.Errorf("%v, expected %d but got %d", prb, prb.expected, solu)
		}
	}
}
