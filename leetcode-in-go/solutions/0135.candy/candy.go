package prb0135

import (
	"math"
)

func candy(ratings []int) int {
	if len(ratings) < 2 {
		return len(ratings)
	}
	candies := make([]int, len(ratings))
	for i, rating := range ratings {
		if i == 0 {
			continue
		}
		if rating > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	count := candies[len(ratings)-1]
	for j := len(ratings) - 1; j > 0; j-- {
		if ratings[j] < ratings[j-1] {
			candies[j-1] = int(math.Max(float64(candies[j-1]), float64(candies[j]+1)))
		}
		count += candies[j-1]
	}
	return count + len(ratings)
}
