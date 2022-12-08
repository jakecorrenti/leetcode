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
	// check to make sure that neither l1 and l2 are nil
	for l1 != nil || l2 != nil {
		var sum int

		// if only l2 is a valid node, then that is going to be the sum
		if l1 == nil && l2 != nil {
			sum = l2.Val
		} else if l1 != nil && l2 == nil { // if l1 is the only valid node, then that is going to be the sum
			sum = l1.Val
		} else {
			sum = l1.Val + l2.Val
		}

		if carry > 0 {
			sum += 1
			carry = 0
		}

		// if the sum is greater than 9, then there is a carry
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

        // if at the end of both lists there is still a carry, add an additional node at the end for the carry
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
