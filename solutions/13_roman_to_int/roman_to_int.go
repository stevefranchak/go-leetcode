package roman_to_int

var lookup map[rune]int

func init() {
	lookup = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
}

func romanToInt(s string) int {
	var prev int
	sum := 0
	for _, c := range s {
		curr := lookup[c]
		if prev != 0 && curr > prev {
			sum = (sum - prev) + (curr - prev)
		} else {
			sum += curr
		}
		prev = curr
	}

	return sum
}
