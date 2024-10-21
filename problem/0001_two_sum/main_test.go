package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/xidream/leetcode/problem/0001_two_sum"
)

var _ = DescribeTable("TwoSum",
	func(nums []int, target int, expected []int) {
		Expect(TwoSum(nums, target)).To(Equal(expected))
	},
	Entry("Example 1", []int{2, 7, 11, 15}, 9, []int{0, 1}),
	Entry("Example 2", []int{3, 2, 4}, 6, []int{1, 2}),
	Entry("Example 3", []int{3, 3}, 6, []int{0, 1}),
)
