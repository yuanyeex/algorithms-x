package prb0122

func maxProfit(prices []int) int {
	profit, ending := 0, len(prices)-1
	hold := -1
	for i, v := range prices {
		if hold != -1 && (i == ending || prices[i+1] < v) {
			profit += (v - hold)
			hold = -1
		} else if hold == -1 && (i != ending && v < prices[i+1]) {
			hold = v
		}
	}
	return profit
}
