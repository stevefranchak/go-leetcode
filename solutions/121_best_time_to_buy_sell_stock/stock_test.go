package buy_sell_stock

import "testing"

func TestMaxProfit(t *testing.T) {
	cases := []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{4, 3, 100, 20, 25, 20, 18, 40}, 97},
		{[]int{100, 200, 300, 400, 355, 254, 153}, 300},
		{[]int{1000, 100, 5000, 4900, 10, 20}, 4900},
		{[]int{10000, 9999, 10000}, 1},
		{[]int{1}, 0},
	}

	for _, c := range cases {
		got := maxProfit(c.prices)
		if got != c.want {
			t.Errorf("maxProfit(%v) = %d; want %d", c.prices, got, c.want)
		}
	}
}

func BenchmarkMaxProfit(b *testing.B) {
	testData := []int{100, 200, 300, 400, 355, 254, 153}
	var result int
	for i := 0; i < b.N; i++ {
		result = maxProfit(testData)
	}
	_ = result
}
