package prb0001

import (
	"reflect"
	"testing"
)

type input struct {
	data   []int
	target int
}

type answer struct {
	ret []int
}

type question struct {
	in  input
	ans answer
}

func Test_Ok(t *testing.T) {
	questions := []question{
		{
			in:  input{data: []int{2, 7, 11, 15}, target: 9},
			ans: answer{[]int{0, 1}},
		},
		{
			in:  input{data: []int{3, 2, 4, 6}, target: 6},
			ans: answer{[]int{1, 2}},
		},
		{
			in:  input{data: []int{3, 3}, target: 6},
			ans: answer{[]int{0, 1}},
		},
	}

	for _, question := range questions {
		ret := twoSum(question.in.data, question.in.target)
		if !reflect.DeepEqual(ret, question.ans.ret) {
			t.Errorf("expect %v, got %v\n", question.ans.ret, ret)
		}
	}
}
