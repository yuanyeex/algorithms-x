package prb0088

import (
	"reflect"
	"testing"
)

type problem struct {
	nums1  []int
	m      int
	nums2  []int
	n      int
	output []int
}

func TestOk(t *testing.T) {
	problems := []problem{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1}, 1, []int{}, 0, []int{1}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
	}

	for _, p := range problems {
		merge(p.nums1, p.m, p.nums2, p.n)
		if !reflect.DeepEqual(p.output, p.nums1) {
			t.Errorf("Expected %v, got %v", p.output, p.nums1)
		}
	}
}
