package TwoSum_167

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		n := numbers[l] + numbers[r]
		if n == target {
			return []int{l + 1, r + 1}
		} else if n > target {
			r--
		} else {
			l++
		}
	}
	return nil
}
