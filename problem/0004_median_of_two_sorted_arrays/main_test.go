package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/xidream/leetcode/problem/0004_median_of_two_sorted_arrays"
)

var _ = DescribeTable("FindMedianSortedArrays",
	func(nums1 []int, nums2 []int, expected float64) {
		Expect(FindMedianSortedArrays(nums1, nums2)).To(Equal(expected))
	},
	Entry("Example 1", []int{1, 3}, []int{2}, 2.0),
	Entry("Example 2", []int{1, 2}, []int{3, 4}, 2.5),
)
