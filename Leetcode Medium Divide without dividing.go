package kata

import "math"

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	R := dividend
	D := divisor
	Q := 0
	if dividend < 0 && divisor < 0 {
		D = 0 - D
		R = 0 - R
	} else if dividend < 0 {
		R = 0 - R
	} else if divisor < 0 {
		D = 0 - D
	}

	for R >= D {
		sub := D
		add := 1
		for R >= sub<<1 {
			sub <<= 1
			add <<= 1
		}
		R -= sub
		Q += add
	}

	isNegative := (dividend < 0) != (divisor < 0)

	if isNegative {
		return -Q
	}

	return Q
}
