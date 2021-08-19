package MajorityElement_169

//给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 n/2 的元素。
//
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
//Boyer-Moore 摩尔投票算法
// 若干个同政党的候选人投票，一人只能投一票，赞成票与反对票对冲，当票数为0时新的候选人上台竞选，
// 若台上不是同政党的人，则其余政党会投反对票 ，所以最后留在台上的就是当选者(多数元素)
//
func majorityElement(nums []int) int {
	m, i := 0, 0
	for _, n := range nums {
		if m == 0 {
			i = n
		}
		if i == n {
			m = m + 1
		} else {
			m = m - 1
		}
	}
	return i
}
