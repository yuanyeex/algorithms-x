package prb0135

import "testing"

type problem struct {
	ratings []int
	output  int
}

func Test_Ok(t *testing.T) {
	problems := []problem{
		{[]int{1, 0, 2}, 5},
		{[]int{1, 2, 2}, 4},
		{[]int{1, 2, 87, 87, 87, 2, 1}, 13},
	}

	for _, prb := range problems {
		v := candy(prb.ratings)
		if v != prb.output {
			t.Errorf("%v expect %d but got %d", prb, prb.output, v)
		}
	}
}
