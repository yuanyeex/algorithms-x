package prb0665

func checkPossibility(nums []int) bool {
	modified := false
	for i, num := range nums {
		if i == 0 {
			continue
		}
		if num < nums[i-1] {
			if modified {
				return false
			}
			if i > 1 && nums[i-2] > nums[i] {
				nums[i] = nums[i-1]
			} else {
				nums[i-1] = nums[i]
			}
			modified = true
		}
	}
	return true
}
