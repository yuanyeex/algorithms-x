package prb0167

import (
	"reflect"
	"testing"
)

type problem struct {
	input    []int
	target   int
	expected []int
}

func TestOk(t *testing.T) {
	problems := []problem{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
	}

	for _, p := range problems {
		ret := twoSum(p.input, p.target)
		if !reflect.DeepEqual(ret, p.expected) {
			t.Errorf("%v expected %v, got %v", p, p.expected, ret)
		}
	}
}
