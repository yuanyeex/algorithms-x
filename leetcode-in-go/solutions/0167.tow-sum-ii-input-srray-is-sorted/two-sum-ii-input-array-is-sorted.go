package prb0167

func twoSum(numbers []int, target int) []int {
	for l, r := 0, len(numbers)-1; l < r; {
		switch {
		case numbers[l]+numbers[r] < target:
			l++
		case numbers[l]+numbers[r] > target:
			r--
		default:
			return []int{l + 1, r + 1}
		}
	}
	return nil
}
