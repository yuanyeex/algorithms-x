package prb0665

import "testing"

type problem struct {
	nums []int
	sol  bool
}

func TestOk(t *testing.T) {
	tests := [...]problem{
		{
			nums: []int{4, 2, 3},
			sol:  true,
		},
		{
			nums: []int{4, 2, 1},
			sol:  false,
		},
		{
			nums: []int{2, 3, 3, 2, 2},
			sol:  false,
		},
		{
			nums: []int{1, 1, 1},
			sol:  true,
		},
		{
			nums: []int{3, 4, 2, 3},
			sol:  false,
		},
		{
			nums: []int{5, 7, 1, 8},
			sol:  true,
		},
		{
			nums: []int{1, 4, 1, 2},
			sol:  true,
		},
	}

	for _, prb := range tests {
		copied := []int{}
		copied = append(copied, prb.nums...)
		result := checkPossibility(prb.nums)
		if result != prb.sol {
			t.Errorf("case %v expected %v, got %v", copied, prb.sol, result)
		}
	}
}
