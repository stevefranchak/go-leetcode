package int_to_roman

import (
	"strings"
)

func dividesInto(dividend int, divisor int) (int, bool) {
	if divisor == 0 {
		return 0, false
	}
	quotient := dividend / divisor
	return quotient, quotient > 0
}

func intToRoman(num int) string {
	var b strings.Builder

	if q, ok := dividesInto(num, 1000); ok {
		b.WriteString(strings.Repeat("M", q))
		num -= 1000 * q
	}
	if _, ok := dividesInto(num, 900); ok {
		b.WriteString("CM")
		num -= 900
	} else {
		if _, ok := dividesInto(num, 500); ok {
			b.WriteByte('D')
			num -= 500
		} else {
			if _, ok := dividesInto(num, 400); ok {
				b.WriteString("CD")
				num -= 400
			}
		}
	}
	if q, ok := dividesInto(num, 100); ok {
		b.WriteString(strings.Repeat("C", q))
		num -= 100 * q
	}
	if _, ok := dividesInto(num, 90); ok {
		b.WriteString("XC")
		num -= 90
	} else {
		if _, ok := dividesInto(num, 50); ok {
			b.WriteByte('L')
			num -= 50
		} else {
			if _, ok := dividesInto(num, 40); ok {
				b.WriteString("XL")
				num -= 40
			}
		}
	}
	if q, ok := dividesInto(num, 10); ok {
		b.WriteString(strings.Repeat("X", q))
		num -= 10 * q
	}
	if _, ok := dividesInto(num, 9); ok {
		b.WriteString("IX")
		num -= 9
	} else {
		if _, ok := dividesInto(num, 5); ok {
			b.WriteByte('V')
			num -= 5
		} else {
			if _, ok := dividesInto(num, 4); ok {
				b.WriteString("IV")
				num -= 4
			}
		}
	}
	if q, ok := dividesInto(num, 1); ok {
		b.WriteString(strings.Repeat("I", q))
	}

	return b.String()
}
