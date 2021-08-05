package Combine_77

func combine(n int, k int) [][]int {
	res := [][]int{}
	back(n, k, 1, []int{}, &res)
	return res
}

func back(n, k, start int, list []int, res *[][]int) {
	if len(list) == k {
		temp := make([]int, k)
		copy(temp, list)
		*res = append(*res, temp)
		return
	}
	for i := start; i <= n; i++ {
		list = append(list, i)
		back(n, k, i+1, list, res)
		list = list[:len(list)-1]
	}
}
