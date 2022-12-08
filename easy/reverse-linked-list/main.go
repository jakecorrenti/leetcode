package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var next *ListNode = nil
	var prev *ListNode = nil

	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		if next == nil {
			break
		}
		head = next
	}

	return head
}
