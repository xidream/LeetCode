package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/xidream/leetcode/problem/0003_longest_substring_without_repeating_characters"
)

var _ = DescribeTable("LengthOfLongestSubstring",
	func(s string, expected int) {
		Expect(LengthOfLongestSubstring(s)).To(Equal(expected))
	},
	Entry("Example 1", "abcabcbb", 3),
	Entry("Example 2", "bbbbb", 1),
	Entry("Example 3", "pwwkew", 3),
	Entry("Test 1", "abba", 2),
)
