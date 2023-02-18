package prb0633

import "math"

/**
 * worst o(sqrt(n))
 */
func judgeSquareSum(c int) bool {
	for m := 0; m*m <= c; m++ {
		b := math.Sqrt(float64(c - m*m))
		if b == math.Floor(b) {
			return true
		}
	}
	return false
}

/**
 * O(sqrt(n)): always
 */
func judgeSquareSum_2(c int) bool {
	m, n := 0, int(math.Sqrt(float64(c)))
	for m <= n {
		tmp := m*m + n*n
		if tmp == c {
			return true
		} else if tmp > c {
			n--
		} else {
			m++
		}
	}
	return false
}

/**
 * O(n)
 */
func judgeSquareSum_time_exceeded(c int) bool {
	m, n := 0, c
	for m <= n {
		tmp := m*m + n*n
		if tmp == c {
			return true
		} else if tmp > c {
			n--
		} else {
			m++
		}
	}
	return false
}
