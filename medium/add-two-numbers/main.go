package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// can start at beginning of each, add, then keep track of any carry somehow (if x / 10 > 1, then there's a carry), and when reach the end of the nodes, add the carry node if necessary.
	carry := 0
	var head *ListNode = nil
	tmpHead := head
	for l1 != nil || l2 != nil {
		var sum int

		if l1 == nil && l2 != nil {
			sum = l2.Val
		} else if l1 != nil && l2 == nil {
			sum = l1.Val
		} else {
			sum = l1.Val + l2.Val
		}

		if carry > 0 {
			sum += 1
			carry = 0
		}

		if sum/10 >= 1 {
			carry = 1
			sum -= 10
		}

		node := &ListNode{
			Val:  sum,
			Next: nil,
		}

		if head == nil {
			head = node
			tmpHead = head
		} else {
			tmpHead.Next = node
			tmpHead = node
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		if l1 == nil && l2 == nil && carry > 0 {
			node := &ListNode{
				Val:  carry,
				Next: nil,
			}
			tmpHead.Next = node
			tmpHead = node
		}
	}
	return head
}
