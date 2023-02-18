package prb0633

import "testing"

type problem struct {
	input  int
	output bool
}

func Test_Ok(t *testing.T) {
	problems := []problem{
		{5, true},
		{3, false},
		{1, true},
		{999999999, false},
	}

	for _, prb := range problems {
		ret := judgeSquareSum_time_exceeded(prb.input)
		if ret != prb.output {
			t.Errorf("prb %v expect %v, but got %v", prb.input, prb.output, ret)
		}
	}
}

func Test_long_time_one(t *testing.T) {
	problems := []problem{
		{5, true},
		{3, false},
		{1, true},
	}

	for _, prb := range problems {
		ret := judgeSquareSum_time_exceeded(prb.input)
		if ret != prb.output {
			t.Errorf("prb %v expect %v, but got %v", prb.input, prb.output, ret)
		}
	}
}
