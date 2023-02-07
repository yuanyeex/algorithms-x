package prb0763

import (
	"reflect"
	"testing"
)

type problem struct {
	str string
	res []int
}

func Test_Ok(t *testing.T) {
	problems := []problem{
		{
			"ababcbacadefegdehijhklij",
			[]int{9, 7, 8},
		},
		{
			"eccbbbbdec",
			[]int{10},
		},
	}

	for _, p := range problems {
		res := partitionLabels(p.str)
		if !reflect.DeepEqual(res, p.res) {
			t.Errorf("%v expected %v, but get %v", p, p.res, res)
		}
	}
}
