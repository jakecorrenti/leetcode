package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	vals := []int{}
	if root == nil {
		return vals
	}

	var search func(root *TreeNode)
	search = func(root *TreeNode) {
		if root.Left != nil {
			search(root.Left)
		}

		vals = append(vals, root.Val)

		if root.Right != nil {
			search(root.Right)
		}
	}

	search(root)

	return vals
}
