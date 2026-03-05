package maxproduct

import "sort"

func MaxProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	// Either top 3, or two most-negative × largest positive
	a := nums[n-1] * nums[n-2] * nums[n-3]
	b := nums[0] * nums[1] * nums[n-1]
	if a > b {
		return a
	}
	return b
}
