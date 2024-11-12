package contains_duplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	cases := []struct {
		nums []int
		want bool
	}{
		// Add test cases here
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, c := range cases {
		got := containsDuplicate(c.nums)
		if got != c.want {
			t.Errorf("containsDuplicate(%v) = %v; want %v", c.nums, got, c.want)
		}
	}
}

func BenchmarkContainsDuplicate(b *testing.B) {
	testDataLength := 15
	testData := make([]int, testDataLength)
	for i := 0; i < testDataLength; i++ {
		testData = append(testData, i)
	}
	var result bool
	for i := 0; i < b.N; i++ {
		result = containsDuplicate(testData)
	}
	_ = result
}
