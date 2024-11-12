package int_to_roman

import (
	"testing"
)

func TestIntToRoman(t *testing.T) {
	cases := []struct {
		num  int
		want string
	}{
		{3749, "MMMDCCXLIX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
		{3999, "MMMCMXCIX"},
	}

	for _, c := range cases {
		got := intToRoman(c.num)
		if got != c.want {
			t.Errorf("intToRoman(%d) = %s; want %s", c.num, got, c.want)
		}
	}
}

func BenchmarkIntToRoman(b *testing.B) {
	var result string
	for i := 0; i < b.N; i++ {
		result = intToRoman(3749)
	}
	_ = result
}
