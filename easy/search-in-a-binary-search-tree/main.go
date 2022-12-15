package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {

	var t *TreeNode

	var search func(r *TreeNode)
	search = func(r *TreeNode) {
		if r != nil {
			if r.Val == val {
				t = r
				return
			}

			if r.Val > val {
				search(r.Left)
			} else {
				search(r.Right)
			}
		}
	}

	search(root)

	return t
}
