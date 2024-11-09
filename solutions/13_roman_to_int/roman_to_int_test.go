package roman_to_int

import "testing"

func TestRomanToInt(t *testing.T) {
	cases := []struct {
		s    string
		want int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"IV", 4},
	}

	for _, c := range cases {
		got := romanToInt(c.s)
		if got != c.want {
			t.Errorf("romanToInt(%s) = %d; want %d", c.s, got, c.want)
		}
	}
}

func BenchmarkRomanToInt(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = romanToInt("MCMXCIV")
	}
	_ = result
}
