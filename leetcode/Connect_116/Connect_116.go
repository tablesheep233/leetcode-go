package Connect_116

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root != nil {
		connectNext(root, nil, nil)
	}
	return root
}

func connectNext(node *Node, right *Node, parent *Node) {
	if right != nil {
		node.Next = right
	} else if parent != nil {
		if parent.Next != nil {
			node.Next = parent.Next.Left
		}
	}
	connectNext(node.Left, node.Right, nil)
	connectNext(node.Right, nil, node)
}
