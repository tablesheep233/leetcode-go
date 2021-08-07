package Permute_46

func permute(nums []int) [][]int {
	res := [][]int{}
	l := len(nums)
	back(&res, nums, l, []int{}, map[int]bool{})
	return res
}

func back(res *[][]int, nums []int, l int, list []int, user map[int]bool) {
	if l == len(list) {
		temp := make([]int, l)
		copy(temp, list)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < l; i++ {
		if !user[i] {
			user[i] = true
			list = append(list, nums[i])
			back(res, nums, l, list, user)
			user[i] = false
			list = list[:len(list)-1]
		}
	}
}
