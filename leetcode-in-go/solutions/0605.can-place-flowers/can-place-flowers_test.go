package prb0605

import "testing"

type problem struct {
	flowered []int
	n        int
	ret      bool
}

func Test_Ok(t *testing.T) {
	problems := []problem{
		{
			[]int{1, 0, 0, 0, 1},
			1,
			true,
		},
		{
			[]int{1, 0, 0, 0, 1},
			2,
			false,
		},
		{
			[]int{1, 0, 0, 0, 0, 1},
			2,
			false,
		},
		{
			[]int{1, 0, 1, 0, 1, 0, 1},
			0,
			true,
		},
	}

	for _, prb := range problems {
		ret := canPlaceFlowers(prb.flowered, prb.n)
		if ret != prb.ret {
			t.Errorf("%T expect %v but got %v", prb, prb.ret, ret)
		}
	}
}
