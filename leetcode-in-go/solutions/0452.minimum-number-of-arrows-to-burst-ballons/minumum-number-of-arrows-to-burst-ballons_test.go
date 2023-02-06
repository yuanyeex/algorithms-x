package prb0452

import "testing"

type Problem struct {
	points [][]int
	output int
}

func Test_ok(t *testing.T) {
	tests := []Problem{
		{
			[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}},
			2,
		},
		{
			[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			4,
		},
		{
			[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
			2,
		},
	}

	for _, test := range tests {
		ret := findMinArrowShots(test.points)
		if ret != test.output {
			t.Errorf("Test %v expected %v, got %v", test, test.output, ret)
		}
	}
}
