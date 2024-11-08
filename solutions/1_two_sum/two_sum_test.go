package two_sum

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{-1, 3, 5}, 4, []int{0, 2}},
		{[]int{-3, -5}, -8, []int{0, 1}},
	}

	for _, c := range cases {
		got := twoSum(c.nums, c.target)
		slices.Sort(got) // the problem states the answer can be returned in any order
		if got[0] != c.want[0] || got[1] != c.want[1] {
			t.Errorf("twoSum(%v, %d) = %v; want %v", c.nums, c.target, got, c.want)
		}
	}
}
