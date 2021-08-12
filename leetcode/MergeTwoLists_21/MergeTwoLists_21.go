package MergeTwoLists_21

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	n := &ListNode{}
	l := n
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			m := &ListNode{Val: l1.Val}
			l.Next = m
			l = l.Next
			l1 = l1.Next
		} else {
			m := &ListNode{Val: l2.Val}
			l.Next = m
			l = l.Next
			l2 = l2.Next
		}
	}
	if l1 == nil {
		l.Next = l2
	} else {
		l.Next = l1
	}
	return n.Next
}
