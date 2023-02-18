package prb0633

import "math"

func judgeSquareSum(c int) bool {
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
