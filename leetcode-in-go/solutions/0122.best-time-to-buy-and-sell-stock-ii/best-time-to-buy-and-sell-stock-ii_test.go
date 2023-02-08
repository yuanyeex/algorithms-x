package prb0122

import "testing"

type problem struct {
	prices   []int
	solution int
}

func Test_Ok(t *testing.T) {
	problems := []problem{
		{
			[]int{7, 1, 5, 3, 6, 4},
			7,
		},
		{
			[]int{1, 2, 3, 4, 5},
			4,
		},
		{
			[]int{7, 6, 4, 3, 1},
			0,
		},
	}

	for _, p := range problems {
		res := maxProfit(p.prices)
		if res != p.solution {
			t.Errorf("%v expected %v, got %v", p, p.solution, res)
		}
	}
}
