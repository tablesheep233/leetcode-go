package Rotate_189

func rotate(nums []int, k int) {
	l := len(nums)
	arr := make([]int, l)
	for i, n := range nums {
		arr[(i+k)%l] = n
	}
	for i, n := range arr {
		nums[i] = n
	}
}
