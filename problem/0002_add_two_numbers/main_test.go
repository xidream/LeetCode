package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/xidream/leetcode/problem/0002_add_two_numbers"
)

var _ = DescribeTable("NewListNode",
	func(digits []int, expected *ListNode) {
		Expect(NewListNode(digits)).To(Equal(expected))
	},
	Entry("[0]", []int{0}, &ListNode{Val: 0}),
	Entry(
		"[2,4,3]",
		[]int{2, 4, 3},
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	),
)

var _ = DescribeTable("AddTwoNumbers",
	func(l1 *ListNode, l2 *ListNode, expected *ListNode) {
		Expect(AddTwoNumbers(l1, l2)).To(Equal(expected))
	},
	Entry(
		"Example 1",
		NewListNode([]int{2, 4, 3}),
		NewListNode([]int{5, 6, 4}),
		NewListNode([]int{7, 0, 8}),
	),
	Entry(
		"Example 2",
		NewListNode([]int{0}),
		NewListNode([]int{0}),
		NewListNode([]int{0}),
	),
	Entry(
		"Example 3",
		NewListNode([]int{9, 9, 9, 9, 9, 9, 9}),
		NewListNode([]int{9, 9, 9, 9}),
		NewListNode([]int{8, 9, 9, 9, 0, 0, 0, 1}),
	),
)
