package prb0142

import "testing"

type problem struct {
	values  []int
	pos     int
	header  *ListNode
	cycleAt *ListNode
}

func TestOk(t *testing.T) {
	var problems []problem
	problems = append(problems, newProblem([]int{3, 2, 0, -4}, 1))
	problems = append(problems, newProblem([]int{1, 2}, 0))
	problems = append(problems, newProblem([]int{1}, -1))
	//problems = append(problems, newProblem([]int{-21, 10, 17, 8, 4, 26, 5, 35, 33, -7, -16, 27, -12, 6, 29, -12, 5, 9, 20, 14, 14, 2, 13, -24, 21, 23, -21, 5}, 24))

	for _, p := range problems {
		cycle := detectCycle(p.header)
		if cycle != p.cycleAt {
			t.Errorf("%v got %v", p, *cycle)
		}
	}
}

func newProblem(values []int, pos int) problem {
	var header *ListNode
	var cycle *ListNode
	var prev *ListNode = nil
	for i, value := range values {
		node := &ListNode{
			Val: value,
		}
		if i != 0 {
			prev.Next = node
		} else {
			header = node
		}
		prev = node
		if i == pos {
			cycle = node
		}
	}
	if cycle != nil {
		prev.Next = cycle
	}

	return problem{
		values:  values,
		pos:     pos,
		header:  header,
		cycleAt: cycle,
	}
}
