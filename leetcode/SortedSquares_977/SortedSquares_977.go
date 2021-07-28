package SortedSquares_977

func sortedSquares(nums []int) []int {
	l, r := 0, len(nums)-1
	result := make([]int, len(nums))
	for i := r; i >= 0; i-- {
		if nums[l]*nums[l] > nums[r]*nums[r] {
			result[i] = nums[l] * nums[l]
			l++
		} else {
			result[i] = nums[r] * nums[r]
			r--
		}
	}
	return result
}
