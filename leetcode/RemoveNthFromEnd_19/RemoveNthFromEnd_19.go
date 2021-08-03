package RemoveNthFromEnd_19

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := &ListNode{Next: head}
	fast, slow := temp, temp
	for ; n > 0; n-- {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return temp.Next
}
