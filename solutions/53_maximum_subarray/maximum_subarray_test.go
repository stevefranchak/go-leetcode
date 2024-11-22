package maximum_subarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
		{[]int{-5, 21, -6}, 21},
	}

	for _, c := range cases {
		got := maxSubArray(c.nums)
		if got != c.want {
			t.Errorf("maxSubArray(%v) = %d; want %d", c.nums, got, c.want)
		}
	}
}

func BenchmarkMaxSubArray(b *testing.B) {
	testData := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	var result int
	for i := 0; i < b.N; i++ {
		result = maxSubArray(testData)
	}
	_ = result
}
