package product_of_array_except_self

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	cases := []struct {
		nums []int
		want []int
	}{
		// Add test cases here
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{[]int{3, -4, 2}, []int{-8, 6, -12}},
		{[]int{-6, -2, 4}, []int{-8, -24, 12}},
		{[]int{}, nil},
		{[]int{1}, nil},
		{[]int{4, 2}, []int{2, 4}},
	}

	for _, c := range cases {
		got := productExceptSelf(c.nums)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("productExceptSelf(%v) = %v; want %v", c.nums, got, c.want)
		}
	}
}

func BenchmarkProductExceptSelf(b *testing.B) {
	testData := []int{1, 2, 3, 4}
	var result []int
	for i := 0; i < b.N; i++ {
		result = productExceptSelf(testData)
	}
	_ = result
}
