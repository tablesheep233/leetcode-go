package RemoveDuplicates_26

func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	i, j := 1, 1
	for ; j < len(nums); j++ {
		if nums[j] != nums[j-1] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
