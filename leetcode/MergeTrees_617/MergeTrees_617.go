package MergeTrees_617

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	root := &TreeNode{}
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root.Val = root1.Val + root2.Val
	root.Right = mergeTrees(root1.Right, root2.Right)
	root.Left = mergeTrees(root1.Left, root2.Left)
	return root
}
