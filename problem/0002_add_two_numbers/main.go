package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(digits []int) *ListNode {
	var (
		rv = &ListNode{
			Val: digits[0],
		}
		node = rv
	)
	for _, digit := range digits[1:] {
		node.Next = &ListNode{
			Val: digit,
		}
		node = node.Next
	}
	return rv
}

/*
AddTwoNumbers
  - Runtime: 0 ms | Beats 100.00%
  - Memory: 6.10 MB | Beats 5.60%
*/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		rv                 = &ListNode{}
		node               = rv
		v1, v2, carry, sum int
	)
	for {
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum = v1 + v2 + carry
		carry = sum / 10
		node.Val = sum % 10
		if l1 == nil && l2 == nil && carry == 0 {
			return rv
		}
		node.Next = &ListNode{}
		node = node.Next
		v1, v2 = 0, 0
	}
}
