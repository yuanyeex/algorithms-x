package prb0455

import "testing"

type question struct {
	children []int
	cookies  []int
	answer   int
}

func Test_Ok(t *testing.T) {
	questions := []question{
		{
			children: []int{1, 2, 3},
			cookies:  []int{1, 1},
			answer:   1,
		},
		{
			[]int{1, 2},
			[]int{1, 2, 3},
			2,
		},
		{
			[]int{10, 9, 8, 7},
			[]int{5, 6, 7, 8},
			2,
		},
	}

	for _, q := range questions {
		solution := findContentChildren(q.children, q.cookies)
		if solution != q.answer {
			t.Errorf("%v: expect %d, got %d", q, q.answer, solution)
		}
	}
}
