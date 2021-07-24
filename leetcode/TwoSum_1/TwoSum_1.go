package TwoSum_1

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		if n, ok := m[target-num]; ok {
			return []int{i, n}
		}
		m[num] = i
	}
	return nil
}
