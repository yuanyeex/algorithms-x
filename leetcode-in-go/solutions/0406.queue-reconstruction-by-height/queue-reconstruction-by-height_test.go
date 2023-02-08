package prb0406

import (
	"reflect"
	"testing"
)

type problem struct {
	input  [][]int
	queued [][]int
}

func Test_Ok(t *testing.T) {
	problems := []problem{
		{
			input: [][]int{
				{7, 0},
				{4, 4},
				{7, 1},
				{5, 0},
				{6, 1},
				{5, 2},
			},
			// [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
			queued: [][]int{
				{5, 0},
				{7, 0},
				{5, 2},
				{6, 1},
				{4, 4},
				{7, 1},
			},
		},
		{
			input: [][]int{
				{6, 0},
				{5, 0},
				{4, 0},
				{3, 2},
				{2, 2},
				{1, 4},
			},
			// [[4,0],[5,0],[2,2],[3,2],[1,4],[6,0]]
			queued: [][]int{
				{4, 0},
				{5, 0},
				{2, 2},
				{3, 2},
				{1, 4},
				{6, 0},
			},
		},
	}

	for _, p := range problems {
		res := reconstructQueue(p.input)
		if !reflect.DeepEqual(res, p.queued) {
			t.Errorf("%v expected %v, got %v", p, p.queued, res)
		}
	}
}
