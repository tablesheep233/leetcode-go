package MoveZeroes_283

func moveZeroes(nums []int) {
	i := 0
	for j, n := range nums {
		if n != 0 {
			t := nums[i]
			nums[i] = n
			nums[j] = t
			i++
		}
	}
}
